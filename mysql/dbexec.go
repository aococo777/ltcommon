package wymysql

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"github.com/aococo777/ltcommon/commonfunc"
	"github.com/aococo777/ltcommon/commonstruct"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	WyMysql        *gorm.DB // 公司数据DB
	WyMysql_read   *gorm.DB // 公司数据DB-只读副本
	WyMysql_ltdata *gorm.DB // 彩票数据DB
	WyMysql_cokey  *gorm.DB // 公司的keyinfo
)

func init() {
	{
		var err error
		mysqlinfo := new(commonstruct.MysqlDBInfo)
		mysqlinfo.Hostsip = beego.AppConfig.String("mysqlurls")
		mysqlinfo.Username = beego.AppConfig.String("mysqluser")
		mysqlinfo.Password = beego.AppConfig.String("mysqlpass")
		mysqlinfo.DBName = beego.AppConfig.String("mysqldb")
		WyMysql, err = InitDB(mysqlinfo)
		if err != nil {
			beego.Error("Init mysqlurls failed! ", err)
		}
	}

	{
		var err error
		mysqlinfo := new(commonstruct.MysqlDBInfo)
		mysqlinfo.Hostsip = beego.AppConfig.String("mysqlurlsread")
		mysqlinfo.Username = beego.AppConfig.String("mysqluserread")
		mysqlinfo.Password = beego.AppConfig.String("mysqlpassread")
		mysqlinfo.DBName = beego.AppConfig.String("mysqldbread")
		WyMysql_read, err = InitDB(mysqlinfo)
		if err != nil {
			beego.Error("Init mysqlurlsread failed! ", err)
		}
	}

	{
		// 推送开奖信息
		var err error
		mysqlinfo := new(commonstruct.MysqlDBInfo)
		mysqlinfo.Hostsip = beego.AppConfig.String("mysqlurlsscj")
		mysqlinfo.Username = beego.AppConfig.String("mysqluserscj")
		mysqlinfo.Password = beego.AppConfig.String("mysqlpassscj")
		mysqlinfo.DBName = beego.AppConfig.String("mysqldbscj")
		WyMysql_ltdata, err = InitDB(mysqlinfo)
		if err != nil {
			beego.Error("Init mysqlurlsscj failed! %v", err)
		}
	}

	{
		// 推送开奖信息
		var err error
		mysqlinfo := new(commonstruct.MysqlDBInfo)
		mysqlinfo.Hostsip = beego.AppConfig.String("mysqlurlscokey")
		mysqlinfo.Username = beego.AppConfig.String("mysqlusercokey")
		mysqlinfo.Password = beego.AppConfig.String("mysqlpasscokey")
		mysqlinfo.DBName = beego.AppConfig.String("mysqldbcokey")
		WyMysql_cokey, err = InitDB(mysqlinfo)
		if err != nil {
			beego.Error("Init mysqlurlscokey failed! %v", err)
		}
	}
}

// 初始化数据库连接
func InitDB(dbinfo *commonstruct.MysqlDBInfo) (*gorm.DB, error) {
	if dbinfo.Username == "" || dbinfo.Password == "" || dbinfo.Hostsip == "" || dbinfo.DBName == "" {
		return nil, errors.New("dbconfig is null")
	}

	Dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", dbinfo.Username, dbinfo.Password, dbinfo.Hostsip, dbinfo.DBName)
	DB, err := gorm.Open("mysql", Dsn)
	if err != nil {
		return nil, err
	}

	// DB.LogMode(true)

	file, err := os.Create("DB.log")
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	DB.SetLogger(log.New(file, "", log.LstdFlags|log.Llongfile))

	DB.DB().SetMaxIdleConns(10)
	return DB, nil
}

// 获取所有彩种信息
func GetRoominfoS() ([]commonstruct.CRoomInfo, error) {
	var RoomInfoS []commonstruct.CRoomInfo
	nowtime := commonfunc.GetNowtime()
	err := WyMysql.Table(commonstruct.WY_gm_config_room).
		Where("in_valid_time > ? ", nowtime).Order("id").Find(&RoomInfoS).Error
	if err != nil {
		beego.Error("GetRoominfoS err ", err)
	}
	return RoomInfoS, err
}

// read 连接保存
func KeepConnRead() ([]commonstruct.CRoomInfo, error) {
	var RoomInfoS []commonstruct.CRoomInfo
	err := WyMysql_read.Table(commonstruct.WY_gm_config_room).Find(&RoomInfoS).Error
	if err != nil {
		beego.Error("KeepConnRead err ", err)
	}
	return RoomInfoS, err
}

// ltdata 连接保存
func KeepConnCokey() ([]commonstruct.KeyJumpurl, error) {
	var RoomInfoS []commonstruct.KeyJumpurl
	err := WyMysql_cokey.Table(commonstruct.Daohang_key_url).Find(&RoomInfoS).Error
	if err != nil {
		beego.Error("KeepConnCokey err ", err)
	}
	return RoomInfoS, err
}

// read 连接保存
func KeepConnLtdata() ([]commonstruct.CRoomInfo, error) {
	var RoomInfoS []commonstruct.CRoomInfo
	err := WyMysql_ltdata.Table(commonstruct.WY_gm_config_room).Find(&RoomInfoS).Error
	if err != nil {
		beego.Error("KeepConnLtdata err ", err)
	}
	return RoomInfoS, err
}

// 通过房间名获取房间信息
func GetRoominfoByName(NameEng string) (commonstruct.CRoomInfo, error) {
	var roominfo commonstruct.CRoomInfo
	err := WyMysql.Table(commonstruct.WY_gm_config_room).Where("name_eng = ?", NameEng).Find(&roominfo).Error
	if err != nil {
		beego.Error("GetRoominfoByName %v", err)
	}
	return roominfo, err
}

// 通过房间名获取房间信息
func GetRoominfoByID(roomid int64) (commonstruct.CRoomInfo, error) {
	var roominfo commonstruct.CRoomInfo
	err := WyMysql.Table(commonstruct.WY_gm_config_room).Where("id = ?", roomid).Find(&roominfo).Error
	if err != nil {
		beego.Error("GetRoominfoByID %v", err)
	}
	return roominfo, err
}

func UpdateRoomcfg(roomid int64, column string, value interface{}) error {
	var updateValues map[string]interface{}

	updateValues = map[string]interface{}{
		column: value,
	}
	if err := WyMysql.Table(commonstruct.WY_gm_config_room).Where("id = ? ", roomid).Update(updateValues).Error; err != nil {
		beego.Error("UpdateRoomcfg err ", roomid, column, value, err)
		return err
	}
	return nil
}

func AddUserOplog(oplog commonstruct.UserOpLog) error {
	err := WyMysql.Table(commonstruct.WY_tmp_log_userop).Create(&oplog).Error
	if err != nil {
		beego.Error("SaveUserOp err %v", err)
	}
	return nil
}

// 用户钱包转换记录
func AddUserTransferlog(uuid int64, optype commonstruct.MoneyUpdateType, OpGold float64, OldGold float64, NewGold float64, Desc string) {
	var uptlog commonstruct.MoneyUpdateLog
	uptlog.Uuid = uuid
	uptlog.Time = commonfunc.GetNowtime()
	uptlog.OpType = optype
	uptlog.OpGold = OpGold
	uptlog.OldGold = OldGold
	uptlog.NewGold = NewGold
	uptlog.Expinfo = Desc
	if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Create(&uptlog).Error; err != nil {
		beego.Error("AddUserTransferlog err ", uptlog, err)
	}
}

// 用户资金变动记录
func AddUserMoneylog(uuid int64, walletname string, optype commonstruct.MoneyUpdateType, OpGold float64, OldGold float64, NewGold float64, Desc string) {
	var uptlog commonstruct.MoneyUpdateLog
	uptlog.Uuid = uuid
	uptlog.Time = commonfunc.GetNowtime()
	uptlog.OpType = optype
	uptlog.WalletName = walletname
	uptlog.OpGold = OpGold
	uptlog.OldGold = OldGold
	uptlog.NewGold = NewGold
	uptlog.Opinfo = Desc
	if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Create(&uptlog).Error; err != nil {
		beego.Error("AddUserMoneylog err ", uptlog, err)
	}
}

func SetDataresExpinfo(uuid int64, expinfo string) {
	if err := WyMysql.Table(commonstruct.WY_tmp_user_datares).Where("uuid = ?", uuid).Update("expinfo", expinfo).Error; err != nil {
		beego.Error("SetDataresExpinfo err %v", err)
	}
}

func GetDataresinfo(uuid int64) commonstruct.Datares {
	var ret commonstruct.Datares
	if err := WyMysql.Table(commonstruct.WY_tmp_user_datares).Where("uuid = ?", uuid).Find(&ret).Error; err != nil {
		beego.Error("GetDataresinfo err", uuid, err)
	}
	return ret
}

// 通过游戏简写获取最新一期开奖信息
func GetRecentWininfo(tablesuf string) commonstruct.Cli_Wininfo {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_num_, tablesuf)
	var RecentWininfo commonstruct.Cli_Wininfo
	if err := WyMysql_ltdata.Table(tblname).Order("opentimestamp desc").Limit(1).Find(&RecentWininfo).Error; err != nil {
		beego.Error("GetRecentWininfo err", tablesuf, err)
		return RecentWininfo
	}
	return RecentWininfo
}

// 获取彩种某期开奖信息
func GetExpectWininfo(tablesuf string, expect string) (commonstruct.Cli_Wininfo, error) {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_num_, tablesuf)
	var RecentWininfo commonstruct.Cli_Wininfo
	err := WyMysql_ltdata.Table(tblname).Where("expect = ?", expect).Find(&RecentWininfo).Error
	// if err != nil {
	// 	 beego.Error("GetRecentWininfo err", tablesuf, err)
	// }
	return RecentWininfo, err
}

func IntToTime(itime int64) time.Time {
	// t, _ := time.Parse("20060102150405", fmt.Sprintf("%d", itime))
	t := commonfunc.StrToBjTime_YL(fmt.Sprintf("%d", itime))
	return t
}

func GetNextTimeinfo_back(tablesuf string, frequency string, now int64) commonstruct.LotteryOpentime {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_lotterytime_, tablesuf)
	nowtime := IntToTime(now)

	var laststr string
	switch frequency {
	case "day":
		laststr = commonfunc.GetBjTime20060102150405(nowtime.AddDate(0, 0, 10))
	default:
		laststr = commonfunc.GetBjTime20060102150405(nowtime.AddDate(0, 0, 2))
	}

	var Opentime commonstruct.LotteryOpentime
	if err := WyMysql_ltdata.Table(tblname).Select("id,expect,opentimestamp").
		Where("opentimestamp > ? and opentimestamp between ? and ?", now, now, laststr).Order("opentimestamp").Limit(1).Find(&Opentime).Error; err != nil {

		//		beego.Error("GetNewExpectinfo err ", tablesuf, now, err)
	}

	return Opentime
}

func GetPreTimeinfo_back(tablesuf string, frequency string, now int64) commonstruct.LotteryOpentime {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_lotterytime_, tablesuf)
	nowtime := IntToTime(now)

	var laststr string
	switch frequency {
	case "day":
		laststr = commonfunc.GetBjTime20060102150405(nowtime.AddDate(0, 0, -10))
	default:
		laststr = commonfunc.GetBjTime20060102150405(nowtime.AddDate(0, 0, -2))
	}

	var Opentime commonstruct.LotteryOpentime
	if err := WyMysql_ltdata.Table(tblname).Select("id,expect,opentimestamp").
		Where("opentimestamp < ? and opentimestamp between ? and ?", now, laststr, now).Order("opentimestamp desc").Limit(1).Find(&Opentime).Error; err != nil {
		return Opentime
	}
	return Opentime
}

func GetNextTimeinfo(tablesuf string, now int64, offset int) commonstruct.LotteryOpentime {
	var Opentime commonstruct.LotteryOpentime
	if tablesuf == "" {
		beego.Error("为空调用===", GetFuncName(3))
		return Opentime
	}

	tblname := fmt.Sprintf("%s%s", commonstruct.WY_lotterytime_, tablesuf)

	nowtime := time.Unix(now, 0)
	nowstr := commonfunc.GetBjTime20060102150405(nowtime)
	var laststr string
	switch tablesuf {
	case "hk6", "fc3d", "pl3":
		laststr = commonfunc.GetBjTime20060102150405(nowtime.AddDate(0, 0, 10))
	default:
		laststr = commonfunc.GetBjTime20060102150405(nowtime.AddDate(0, 0, 2))
	}

	if err := WyMysql_ltdata.Table(tblname).Select("id,expect,opentimestamp").
		Where("opentimestamp > ? and opentimestamp between ? and ?", nowstr, nowstr, laststr).Order("opentimestamp").Limit(1).Offset(offset).Find(&Opentime).Error; err != nil {

		beego.Error("GetNewExpectinfo err ", err, tablesuf, now, offset)
	}
	return Opentime
}

func GetRoomReadyinfo(tablesuf string, now int64) commonstruct.LotteryOpentime {
	var Opentime commonstruct.LotteryOpentime
	if tablesuf == "" {
		beego.Error("为空调用===", GetFuncName(3))
		return Opentime
	}

	tblname := fmt.Sprintf("%s%s", commonstruct.WY_lotterytime_, tablesuf)

	nowtime := time.Unix(now, 0)
	nowstr := commonfunc.GetBjTime2006_01_02_15_04_05(nowtime)

	if err := WyMysql_ltdata.Table(tblname).Select("count(*) as id").
		Where("opentimestamp > ?", nowstr).Find(&Opentime).Error; err != nil {
		beego.Error("GetRoomReadyinfo err ", err, tablesuf, now)
	}
	return Opentime
}

func GetSlotsRecentWininfo(tablesuf string, roomid int64) commonstruct.WininfoFQZS {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_num_, tablesuf)
	// 中奖列表
	var RecentWininfo commonstruct.WininfoFQZS
	if err := WyMysql.Table(tblname).Select("room_id,expect,opencode,opentime").
		Where("room_id = ?", roomid).Order("opentime desc").Limit(1).Find(&RecentWininfo).Error; err != nil {
		beego.Error("GetSlotsRecentWininfo err", tablesuf, err)
		return RecentWininfo
	}
	return RecentWininfo
}

/*************************
* 获取Port的下注金额统计
 ************************/
func GetSumamountGroupbyPortid(uuid int64, roomid int64, expect string) []commonstruct.LotdataUseritem {
	var datas []commonstruct.LotdataUseritem
	err := WyMysql.Table(commonstruct.WY_tmp_user_lotdata_item).
		Select("uuid,room_id,room_eng,room_cn,expect,port_id,sum(order_amount) as order_amount,sum(shizhanhuoliang) as shizhanhuoliang").
		Where("uuid = ? and room_id = ? and expect = ?", uuid, roomid, expect).
		Group("uuid,room_id,room_eng,room_cn,expect,port_id").
		Find(&datas).Error
	if err != nil {
		beego.Error("GetSumamountGroupbyPortid err", err)
	}
	return datas
}

func UpsertWarninglog(newlog commonstruct.WarningLog) error {
	var oldlog commonstruct.WarningLog

	if retinfo := WyMysql.Table(commonstruct.WY_tmp_log_warning).
		Where("uuid = ? and room_id = ? and expect = ? and item_id = ?", newlog.Uuid, newlog.RoomID, newlog.Expect, newlog.ItemID).
		Find(&oldlog); retinfo.Error != nil {
		if retinfo.RecordNotFound() {

			if err := WyMysql.Table(commonstruct.WY_tmp_log_warning).Create(&newlog).Error; err != nil {
				beego.Error("create err ", newlog, err)
				return err
			}
		} else {
			beego.Error("UpsertInoutStatistic err ", newlog, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	updateValues := map[string]interface{}{
		"amount": newlog.Amount,
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_log_warning).
		Where("uuid = ? and room_id = ? and expect = ? and item_id = ?", newlog.Uuid, newlog.RoomID, newlog.Expect, newlog.ItemID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsertWarninglog err ", newlog, err)
	}
	return nil
}

func GetWarninginfoS(uuid int64, roomid int64, expect string) ([]commonstruct.WarningLog, error) {
	var orders []commonstruct.WarningLog
	err := WyMysql.Table(commonstruct.WY_tmp_log_warning).Where("uuid = ? and expect = ?", uuid, expect).Order("warning_time desc").Limit(50).Find(&orders).Error

	return orders, err
}

func GetPortclassSByPlatform(platform string) ([]commonstruct.PortClass, error) {
	var info []commonstruct.PortClass
	err := WyMysql.Table(commonstruct.WY_gm_config_portclass).
		Where("lottery_dalei = ?", platform).Find(&info).Error
	return info, err
}

func GetWalletinfoS() ([]commonstruct.Walletinfo, error) {
	var info []commonstruct.Walletinfo
	err := WyMysql.Table(commonstruct.WY_gm_config_wallet).Find(&info).Error
	return info, err
}

func GetWalletinfo(walletid int64) (commonstruct.Walletinfo, error) {
	var info commonstruct.Walletinfo
	err := WyMysql.Table(commonstruct.WY_gm_config_wallet).Where("id = ?", walletid).Find(&info).Error
	return info, err
}

// 获取公司的黑名单列表
func GetCompanyBlacklist(uuid int64) ([]commonstruct.LimitUserinfo, error) {
	var list []commonstruct.LimitUserinfo
	if err := WyMysql.Table(commonstruct.WY_company_limitinfo).Where("company_id = ? and limit_flag = 1", uuid).Order("limit_time desc").Find(&list).Error; err != nil {
		beego.Error("GetCompanyBlacklist err", err)
		return list, err
	}
	return list, nil
}

func GetUserLimitinfo(uuid int64) commonstruct.LimitUserinfo {
	var user commonstruct.LimitUserinfo
	if err := WyMysql.Table(commonstruct.WY_company_limitinfo).Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		return user
	}
	return user
}

/************************************
* 查询系统公告
************************************/
func QuerySystemnotice(gameplatform string) ([]commonstruct.CNoticeInfo, error) {
	var selectarg string
	selectarg = fmt.Sprintf("company_id = 10000 and game_type = '%v'", gameplatform)

	var newinfo []commonstruct.CNoticeInfo
	err := WyMysql.Table(commonstruct.WY_company_notice).Where(selectarg).Find(&newinfo).Error
	return newinfo, err
}

/*
 获取用户的所有后台目录
*/
func GetUsernaviS(uuid int64) ([]commonstruct.UserNavi, error) {
	var infos []commonstruct.UserNavi
	usernavitablename := fmt.Sprintf("%v%v", commonstruct.WY_user_navi_, uuid/500)
	if err := WyMysql.Table(usernavitablename).Where("uuid = ? and value > 0", uuid).Order("group_id,navi_id").Find(&infos).Error; err != nil {
		beego.Error("GetUsernaviS err", uuid, err)
		return infos, err
	}
	return infos, nil
}

func GetPortclassByID(portid int64) (commonstruct.PortClass, error) {
	var info commonstruct.PortClass
	err := WyMysql.Table(commonstruct.WY_gm_config_portclass).Where("id = ?", portid).Find(&info).Error
	return info, err
}

func GetPreTimeinfo(tablesuf string, Frequency string, now int64) commonstruct.LotteryOpentime {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_lotterytime_, tablesuf)
	nowtime := IntToTime(now)

	var laststr string
	switch Frequency {
	case "day":
		laststr = commonfunc.GetBjTime20060102150405(nowtime.AddDate(0, 0, -10))
	default:
		laststr = commonfunc.GetBjTime20060102150405(nowtime.AddDate(0, 0, -2))
	}

	var Opentime commonstruct.LotteryOpentime
	if err := WyMysql_ltdata.Table(tblname).Select("id,expect,opentimestamp").
		Where("opentimestamp < ? and opentimestamp between ? and ?", now, laststr, now).Order("opentimestamp desc").Limit(1).Find(&Opentime).Error; err != nil {
		return Opentime
	}
	return Opentime
}

// func GetAgentoddsByUuid(uuid int64) commonstruct.AgentOdds {
// 	var oddsinfo commonstruct.AgentOdds
// 	if err := WyMysql.Table(commonstruct.WY_user_odds).Where("uuid = ?", uuid).Find(&oddsinfo).Error; err != nil {
// 		beego.Error("GetAgentoddsByUuid err", uuid, err)
// 	}
// 	return oddsinfo
// }

// 获取公司的玩法设置信息
func GetCompanyPortoddsS(uuid int64, roomid int64) ([]commonstruct.CompanyPortinfo, error) {
	var oddsinfoS []commonstruct.CompanyPortinfo
	err := WyMysql.Table(commonstruct.WY_company_portclass).
		Where("company_id = ? and room_id = ?", uuid, roomid).Find(&oddsinfoS).Error
	return oddsinfoS, err
}

// 获取公司的玩法设置信息
func GetCompanyPortodds(uuid int64, roomid int64, portid int64) (commonstruct.CompanyPortinfo, error) {
	var oddsinfo commonstruct.CompanyPortinfo
	err := WyMysql.Table(commonstruct.WY_company_portclass).
		Where("company_id = ?  and room_id = ? and port_id = ?", uuid, roomid, portid).Find(&oddsinfo).Error
	return oddsinfo, err
}

// 获取公司的玩法设置信息
func GetRoomPortoddsS(roomid int64) ([]commonstruct.CompanyPortinfo, error) {
	var oddsinfoS []commonstruct.CompanyPortinfo
	err := WyMysql.Table(commonstruct.WY_company_portclass).
		Where("room_id = ?", roomid).Find(&oddsinfoS).Error
	return oddsinfoS, err
}

// 获取加时玩法设置信息
func GetJiashiPortclassS(roomid int64) ([]commonstruct.CompanyJiashiPortinfo, error) {
	var oddsinfoS []commonstruct.CompanyJiashiPortinfo
	err := WyMysql.Table(commonstruct.WY_company_portclass_jiashi).
		Where("room_id = ?", roomid).Find(&oddsinfoS).Error
	return oddsinfoS, err
}

func GetCoRechargewayinfo(masterid int64, wayid string) (commonstruct.CoRechargeWay, error) {
	var data commonstruct.CoRechargeWay
	err := WyMysql.Table(commonstruct.WY_company_rechargeway).Where("company_id = ? and way_id = ?", masterid, wayid).Find(&data).Error
	return data, err
}

func DateTime_Str2int64(datetime string) int64 {
	timeUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
	return timeUnix.Unix()
}

func Date_Str2int64(date string) int64 {
	timeUnix, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return timeUnix.Unix()
}

func GetResultidExpinfo(lottery_dalei string, game_type string) []commonstruct.CGameItem {
	var OrderS []commonstruct.CGameItem
	if err := WyMysql.Table(commonstruct.WY_gm_config_item).Where("lottery_dalei = ? and game_type = ? and islianma = 0", lottery_dalei, game_type).
		Find(&OrderS).Error; err != nil {
		beego.Error("GetResultidExpinfo err ", err)
		return nil
	}
	return OrderS
}

func SetNumExpinfo(tablesuf string, expect string, expinfo string) {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_num_, tablesuf)
	if err := WyMysql_ltdata.Table(tblname).Where("expect = ?", expect).Update("expinfo", expinfo).Error; err != nil {
		beego.Error("SetNumExpinfo err ", err)
	}
}

// // 获取开房信息
// func GetAuthority(uuid int64) commonstruct.Authority {
// 	var authority commonstruct.Authority
// 	if err := WyMysql.Table(commonstruct.WY_user_authority).Where("uuid = ? ", uuid).Find(&authority).Error; err != nil {
// 		beego.Error("GetAuthority err ", err)
// 		return authority
// 	}
// 	return authority
// }

// 设置开奖信息
func SetOpencode(tablesuf string, expect string, opencode string) error {
	tblname := fmt.Sprintf("%v%v", commonstruct.WY_num_, tablesuf)
	// 插入数据
	info := &commonstruct.Cli_Wininfo{
		Expect:        expect,
		Opencode:      opencode,
		Opentime:      commonfunc.GetBjTime2006_01_02_15_04_05(time.Now()),
		Opentimestamp: time.Now().Unix(),
	}

	if err := WyMysql_ltdata.Table(tblname).Create(&info).Error; err != nil {
		beego.Error("Insert count err ", *info, tblname, err)
		return err
	}
	return nil
}

// 获取开房信息
func GetRoomgame(gameid int64) commonstruct.CRoomGame {
	var gameinfo commonstruct.CRoomGame
	if gameid <= 0 {
		return gameinfo
	}
	if err := WyMysql.Table(commonstruct.WY_gm_config_roomgame).Where("id = ? ", gameid).Find(&gameinfo).Error; err != nil {
		beego.Error("GetRoomgame err ", gameid, err)
	}
	return gameinfo
}

func GetCompanygame(preid int64, roomid int64) (commonstruct.CompanyGame, error) {
	var value commonstruct.CompanyGame
	err := WyMysql.Table(commonstruct.WY_company_game).Where("company_id = ? and room_id = ?", preid, roomid).Find(&value).Error
	if err != nil {
		// beego.Error("GetClosevalue err ", preid, roomid, err.Error())
	}
	return value, err
}

func GetAllCompanygameS() ([]commonstruct.CompanyGame, error) {
	var value []commonstruct.CompanyGame
	err := WyMysql.Table(commonstruct.WY_company_game).Find(&value).Error
	if err != nil {
		// beego.Error("GetClosevalue err ", preid, roomid, err.Error())
	}
	return value, err
}

// 获取公司玩法设置
func GetAllCompanyoddsS() ([]commonstruct.CompanyPortinfo, error) {
	var portinfoS []commonstruct.CompanyPortinfo

	var logcount commonstruct.CompanyPortinfo
	if err := WyMysql_read.Table(commonstruct.WY_company_portclass).Select("count(*) as company_id").
		Find(&logcount).Error; err != nil {
		beego.Error("GetAllCompanyoddsS count err", err)
		return nil, err
	}

	count := int(math.Ceil(float64(logcount.CompanyID) / dancibishu)) // 需要执行次数
	for i := 0; i < count; i++ {

		var danciorders []commonstruct.CompanyPortinfo
		if err := WyMysql_read.Table(commonstruct.WY_company_portclass).
			Limit(int(dancibishu)).Offset(i * int(dancibishu)).
			Find(&danciorders).Error; err != nil {
			return portinfoS, err
		} else {
			for _, order := range danciorders {
				portinfoS = append(portinfoS, order)
			}
		}
	}

	// err := WyMysql.Table(commonstruct.WY_company_portclass).Find(&portinfoS).Error
	// if err != nil {
	// 	beego.Error("GetAllCompanyoddsS err ", err)
	// }
	return portinfoS, nil
}

func InsertWininfo_baijia(tablesuf string, wininfo commonstruct.Wininfo_actual) {
	//	tblname := fmt.Sprintf("%v%v", commonstruct.WY_num_actual_, tablesuf)

	//	var oldinfo commonstruct.Wininfo_actual
	//	if retinfo := WyMysql.Table(tblname).Where("expect = ?", wininfo.Expect).Find(&oldinfo); retinfo.Error != nil {
	//		if retinfo.RecordNotFound() {
	//			if err := WyMysql.Table(tblname).Create(&wininfo).Error; err != nil {
	//				beego.Error("InsertWininfo err ", tablesuf, wininfo, err)
	//			}
	//		} else {
	//			beego.Error("InsertWininfo_baijia err ", tablesuf, wininfo.Expect, retinfo.Error.Error())
	//		}
	//	}
}

func IsOpened(tablesuf string, expect string) bool {
	//	tblname := fmt.Sprintf("%v%v", commonstruct.WY_num_actual_, tablesuf)
	//	var wininfo commonstruct.Wininfo_actual
	//	if err := WyMysql.Table(tblname).Where("expect = ? ", expect).Find(&wininfo).Error; err != nil {
	//		beego.Error("IsOpened err ", tablesuf, err)
	//		return false
	//	}
	//	if wininfo.Opencode == "" {
	//		return false
	//	}
	return true
}

// func GetMoneylogPageinfo(uuid int64, optype int64, begindate int64, enddate int64) (int, int) {
// 	// 查询订单记录
// 	var logs commonstruct.MoneyUpdateLog
// 	if optype > 0 {
// 		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Select("count(*) as id").
// 			Where("uuid = ? and op_type = ? and time > ? and time < ? ", uuid, optype, GetBegintime(begindate), GetEndtime(enddate)).
// 			Find(&logs).Error; err != nil {
// 			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
// 		}
// 	} else {
// 		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Select("count(*) as id").
// 			Where("uuid = ? and time > ? and time < ? ", uuid, GetBegintime(begindate), GetEndtime(enddate)).
// 			Find(&logs).Error; err != nil {
// 			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
// 		}
// 	}

// 	return int(logs.ID), int(math.Ceil(float64(logs.ID) / 20))
// }

func GetMoneylogPageinfo(uuid int64, optype int64, begindate int64, enddate int64, pagecount int) (int64, int64) {
	timeMin := begindate * 1000000
	timeMax := enddate*1000000 + 235959

	// 查询订单记录
	var logs commonstruct.MoneyUpdateLog
	if optype > 0 {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Select("count(*) as id").
			Where("uuid = ? and op_type = ? and time between ? and ?  ", uuid, optype, timeMin, timeMax).
			Find(&logs).Error; err != nil {
			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
		}
	} else {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Select("count(*) as id").
			Where("uuid = ? and time between ? and ?  ", uuid, timeMin, timeMax).
			Find(&logs).Error; err != nil {
			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
		}
	}

	return int64(logs.ID), int64(math.Ceil(float64(logs.ID) / float64(pagecount)))
}

/***********
* 分页获取资金变动日志
 ***********/
func GetMoneylog(uuid int64, optype int64, offset int, begindate int64, enddate int64, pagecount int) []commonstruct.MoneyUpdateLog {
	timeMin := begindate * 1000000
	timeMax := enddate*1000000 + 235959

	// 查询订单记录
	var logs []commonstruct.MoneyUpdateLog
	if optype > 0 {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Where("uuid = ? and op_type = ? and time between ? and ? ", uuid, optype, timeMin, timeMax).
			Order("id desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
		}
	} else {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Where("uuid = ? and time between ? and ? ", uuid, timeMin, timeMax).
			Order("id desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
		}

	}
	return logs
}

func BKGetMoneylogPageinfo(uuid int64, optype int64, begindate int64, enddate int64) (int64, int64) {
	// 查询订单记录
	var logs commonstruct.MoneyUpdateLog
	if optype > 0 {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Select("count(*) as id").
			Where("uuid = ? and op_type = ? and time > ? and time < ? ", uuid, optype, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate)).
			Find(&logs).Error; err != nil {
			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
		}
	} else {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Select("count(*) as id").
			Where("uuid = ? and time > ? and time < ? ", uuid, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate)).
			Find(&logs).Error; err != nil {
			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
		}
	}

	return int64(logs.ID), int64(math.Ceil(float64(logs.ID) / 20))
}

// func GetMoneylog(uuid int64, optype int64, offset int, begindate int64, enddate int64) []commonstruct.MoneyUpdateLog {
// 	// 查询订单记录
// 	var logs []commonstruct.MoneyUpdateLog
// 	if optype > 0 {
// 		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Where("uuid = ? and op_type = ? and time > ? and time < ? ", uuid, optype, GetBegintime(begindate), GetEndtime(enddate)).
// 			Order("id desc").Limit(20).Offset((offset - 1) * 20).Find(&logs).Error; err != nil {
// 			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
// 		}
// 	} else {
// 		if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Where("uuid = ? and time > ? and time < ? ", uuid, GetBegintime(begindate), GetEndtime(enddate)).
// 			Order("id desc").Limit(20).Offset((offset - 1) * 20).Find(&logs).Error; err != nil {
// 			beego.Error("GetMoneylog err", uuid, optype, begindate, enddate, err)
// 		}

// 	}
// 	return logs
// }

func GetWithdrawByDate(uuid int64, begindate int64, enddate int64) []commonstruct.MoneyInOut {
	// 查询订单记录
	var logs []commonstruct.MoneyInOut
	if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Where("uuid = ? and req_time between ? and ?", uuid, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate)).
		Order("req_time desc").Find(&logs).Error; err != nil {
		beego.Error("GetWithdrawByDate err", uuid, begindate, enddate, err)
		return logs
	}
	return logs
}

func GetRechargeOrderinfo(orderid int64) commonstruct.ReCharge {
	var info commonstruct.ReCharge
	if err := WyMysql.Table(commonstruct.WY_tmp_user_recharge).Where("order_id = ?", orderid).Find(&info).Error; err != nil {
		beego.Error("GetRechargeOrderinfo err", err)
		return info
	}
	return info
}

func SetRechargeChecked(orderid int64, Amount float64, oldgold float64, newgold float64, checked int64) {
	updateValues := map[string]interface{}{
		"checked": checked,
		"amount":  Amount,
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_user_recharge).Where("order_id = ?", orderid).Update(updateValues).Error; err != nil {
		beego.Error("SetRechargeChecked err", err)
	}
}

func UpdateRechargeinfo(orderid int64, amount float64, oldgold float64, newgold float64, state int64, expinfo string) error {
	updateValues1 := map[string]interface{}{
		"checked":  1,
		"state":    state,
		"amount":   amount,
		"old_gold": oldgold,
		"new_gold": newgold,
		"res_time": commonfunc.GetNowtime(),
		"expinfo":  gorm.Expr("CONCAT_WS('|',expinfo,?)", expinfo),
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Update(updateValues1).Error; err != nil {
		beego.Error("UpdateUnWithdraw err", err)
		return err
	}
	return nil
}

func UpdateUnWithdraw(uuid int64, orderid int64, op_result int64, expinfo string) error {
	//	updateValues := map[string]interface{}{
	//		"checked":   1,
	//		"op_result": op_result,
	//		"expinfo":   expinfo,
	//	}
	//	if err := WyMysql.Table(commonstruct.WY_gm_user_withdraw).Where("order_id = ?", orderid).Update(updateValues).Error; err != nil {
	//		beego.Error("UpdateUnWithdraw err", err)
	//		return err
	//	}

	updateValues1 := map[string]interface{}{
		"checked":   1,
		"op_result": op_result,
		"expinfo":   expinfo,
		"res_time":  commonfunc.GetNowtime(),
		"state":     1,
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Update(updateValues1).Error; err != nil {
		beego.Error("UpdateUnWithdraw err", err)
		return err
	}
	return nil
}

func SetQRCodeurl(orderid int64, Amount float64, QRCodeUrl string) {
	updateValues := map[string]interface{}{
		"qr_code_url": QRCodeUrl,
		"amount":      Amount,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Update(updateValues).Error; err != nil {
		beego.Error("SetQRCodeurl err", err)
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_recharge).Where("order_id = ?", orderid).Update(updateValues).Error; err != nil {
		beego.Error("SetQRCodeurl err", err)
	}
}

func GetUserStatistic(uuid int64) (commonstruct.OpStatistic, bool) {
	var ret commonstruct.OpStatistic
	if err := WyMysql.Table(commonstruct.WY_tmp_user_op_statistic).Where("uuid = ?", uuid).Find(&ret).Error; err != nil {
		beego.Error("GetUserOpStatistic err", uuid, err)
		return ret, false
	}
	return ret, true
}

func InitUserStatistic(uuid int64) error {
	var statisticinfo commonstruct.OpStatistic
	statisticinfo.Uuid = uuid
	if err := WyMysql.Table(commonstruct.WY_tmp_user_op_statistic).Create(&statisticinfo).Error; err != nil {
		beego.Error("InitUserStatistic err ", uuid, err)
		return err
	}
	return nil
}

type IP_Count struct {
	IP    string
	Count int64
}

type DateCount struct {
	Date  string
	Count int64
}

type GameAmount struct {
	Game   string
	Amount float64
}

func SetUserLogin(uuid int64, ip string) {
	statisticinfo, ok := GetUserStatistic(uuid)
	if !ok {
		InitUserStatistic(uuid)
	}
	statisticinfo.Uuid = uuid

	bHave := false
	var iplist []IP_Count
	if statisticinfo.IPList != "" {
		json.Unmarshal([]byte(statisticinfo.IPList), &iplist)
		for k, v := range iplist {
			if v.IP == ip {
				v.Count = v.Count + 1
				bHave = true
				iplist[k] = v
				break
			}
		}

		if len(iplist) > 20 {
			iplist = iplist[len(iplist)-15 : len(iplist)]
		}

		iplistbyte, _ := json.Marshal(iplist)
		statisticinfo.IPList = string(iplistbyte)
	}
	if !bHave {
		newip := &IP_Count{
			IP:    ip,
			Count: 1,
		}
		iplist = append(iplist, *newip)
		iplistbyte, _ := json.Marshal(iplist)
		statisticinfo.IPList = string(iplistbyte)
	}

	if statisticinfo.Count <= 0 {
		statisticinfo.Count = 1
	} else {
		statisticinfo.Count = statisticinfo.Count + 1
	}

	var datecount []DateCount
	date := fmt.Sprintf("%04d-%02d-%02d", commonfunc.BeijingTime().Year(), commonfunc.BeijingTime().Month(), commonfunc.BeijingTime().Day())
	if statisticinfo.DateCount != "" {
		json.Unmarshal([]byte(statisticinfo.DateCount), &datecount)

		if len(datecount) > 50 {
			datecount = datecount[len(datecount)-40 : len(datecount)]
		}

		lastinfo := datecount[len(datecount)-1]

		if lastinfo.Date == date {
			lastinfo.Count = lastinfo.Count + 1
			datecount[len(datecount)-1] = lastinfo
		} else {
			lastinfo.Date = date
			lastinfo.Count = 1
			datecount = append(datecount, lastinfo)
		}
		datecountbyte, _ := json.Marshal(datecount)
		statisticinfo.DateCount = string(datecountbyte)
	} else {
		newdate := &DateCount{
			Date:  date,
			Count: 1,
		}
		datecount = append(datecount, *newdate)
		datecountbyte, _ := json.Marshal(datecount)
		statisticinfo.DateCount = string(datecountbyte)
	}

	updateValues := map[string]interface{}{
		"ip_list":    statisticinfo.IPList,
		"count":      statisticinfo.Count,
		"date_count": statisticinfo.DateCount,
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_user_op_statistic).Where("uuid = ?", statisticinfo.Uuid).Update(updateValues).Error; err != nil {
		beego.Error("SetUserLogin update err", err)
	}
}

func SetUserOrder(uuid int64, game_eng string, amount float64) {
	statisticinfo, ok := GetUserStatistic(uuid)
	if !ok {
		InitUserStatistic(uuid)
	}

	bHave := false
	var gameamountlist []GameAmount
	if statisticinfo.GameAmount != "" {
		json.Unmarshal([]byte(statisticinfo.GameAmount), &gameamountlist)
		for k, v := range gameamountlist {
			if v.Game == game_eng {
				v.Amount = v.Amount + amount
				bHave = true
				gameamountlist[k] = v
				break
			}
		}
		listbyte, _ := json.Marshal(gameamountlist)
		statisticinfo.GameAmount = string(listbyte)
	}

	if !bHave {
		newgameamount := &GameAmount{
			Game:   game_eng,
			Amount: amount,
		}
		gameamountlist = append(gameamountlist, *newgameamount)
		listbyte, _ := json.Marshal(gameamountlist)
		statisticinfo.GameAmount = string(listbyte)
	}

	updateValues := map[string]interface{}{
		"amount":      statisticinfo.Amount + amount,
		"game_amount": statisticinfo.GameAmount,
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_user_op_statistic).Where("uuid = ?", uuid).Update(updateValues).Error; err != nil {
		beego.Error("SetUserOrder err", err)
	}
}

// 获取需要客户端显示的长龙信息 (num >= 2 )
func GetChanglongShowinfo(GameDalei string, GameName string, key string) []commonstruct.ChanglongInfo {
	var ret []commonstruct.ChanglongInfo
	var num int

	switch key {
	case "inc":
		if GameName == "pcdd" {
			num = 1
		} else {
			num = 2
		}
		if err := WyMysql.Table(commonstruct.WY_tmp_tool_changlong).Where("game_dalei = ? and game_name = ? and num >= ?", GameDalei, GameName, num).Order("num desc,id").Find(&ret).Error; err != nil {
			beego.Error("GetChanglongInfo err %v\n", err)
			return ret
		}
	case "dec":
		if GameName == "pcdd" {
			num = -1
		} else {
			num = -2
		}
		if err := WyMysql.Table(commonstruct.WY_tmp_tool_changlong).Where("game_dalei = ? and game_name = ? and num <= ?", GameDalei, GameName, num).Order("num,id").Find(&ret).Error; err != nil {
			beego.Error("GetChanglongInfo err %v\n", err)
			return ret
		}
	default:
		beego.Error("GetChanglongInfo err key %v\n", key)
	}
	return ret
}

func UpsertChanglonginfo(GameDalei string, GameName string, items []int64, upttype int) error {
	var err error
	switch upttype {
	case 0: // 连开
		err = WyMysql.Table(commonstruct.WY_tmp_tool_changlong).Where("game_dalei = ? and game_name =? and item_id in (?) ", GameDalei, GameName, items).Update("num", gorm.Expr("num + 1")).Error
	case 1: // 连开中断
		err = WyMysql.Table(commonstruct.WY_tmp_tool_changlong).Where("game_dalei = ? and game_name =? and item_id in (?) ", GameDalei, GameName, items).Update("num", 0).Error
	case 2: // 连不开
		err = WyMysql.Table(commonstruct.WY_tmp_tool_changlong).Where("game_dalei = ? and game_name =? and item_id in (?) ", GameDalei, GameName, items).Update("num", gorm.Expr("num - 1")).Error
	case 3: // 连不开中断
		err = WyMysql.Table(commonstruct.WY_tmp_tool_changlong).Where("game_dalei = ? and game_name =? and item_id in (?) ", GameDalei, GameName, items).Update("num", 1).Error
	}
	return err
}

// 获取游戏的所有长龙信息
func GetChanglongTotalinfo(GameDalei string, GameName string) ([]commonstruct.ChanglongInfo, error) {
	var infos []commonstruct.ChanglongInfo
	err := WyMysql.Table(commonstruct.WY_tmp_tool_changlong).Where("game_dalei = ? and game_name = ?", GameDalei, GameName).Find(&infos).Error
	return infos, err
}

// 初始化item的长龙信息
func InitChanglong(GameDalei string, GameName string, itemid int64, winflag int64, info string) error {
	var newinfo commonstruct.ChanglongInfo
	newinfo.GameDalei = GameDalei
	newinfo.GameName = GameName
	newinfo.ItemID = itemid
	if winflag > 0 {
		newinfo.Num = 1
	} else {
		newinfo.Num = 0
	}
	newinfo.Info = info
	err := WyMysql.Table(commonstruct.WY_tmp_tool_changlong).Create(&newinfo).Error
	if err != nil {
		beego.Error("initChanglong err ", GameDalei, GameName, itemid, err)
	}
	return err
}

func GetPortOddsinfo(uuid int64, roomid int64, port string, portid int64) (commonstruct.CompanyPortinfo, error) {
	var oddsinfo commonstruct.CompanyPortinfo
	err := WyMysql.Table(commonstruct.WY_company_portclass).
		Where("company_id = ? and port = ? and room_id = ? and port_id = ?", uuid, port, roomid, portid).Find(&oddsinfo).Error
	if err != nil {
		beego.Error("GetPortOdds err %v\n", uuid, roomid, port, err)
	}
	return oddsinfo, err
}

// 获取公司开通的游戏列表
func GetGamelistByUuid(uuid int64) []commonstruct.CompanyGame {
	var list []commonstruct.CompanyGame
	if err := WyMysql.Table(commonstruct.WY_company_game).
		Where("company_id = ? and valid_time >= ? and pre_kaiguan = 1", uuid, commonfunc.GetNowtime()).
		Order("sort_id").Find(&list).Error; err != nil {
		beego.Error("GetGamelistByUuid err", err)
		return list
	}
	return list
}

// 获取公司开通的游戏列表
func GetPlayerGamelist(uuid int64) []commonstruct.CompanyGame {
	var list []commonstruct.CompanyGame
	if err := WyMysql.Table(commonstruct.WY_company_game).
		Where("company_id = ? and valid_time >= ? and pre_kaiguan = 1 and in_use in (1,2)", uuid, commonfunc.GetNowtime()).
		Order("sort_id").Find(&list).Error; err != nil {
		beego.Error("GetGamelistByUuid err", err)
		return list
	}
	return list
}

// 获取结算分类S,后端赔率阈值使用
func GetGametypes(platform string) ([]commonstruct.CRoomInfo, error) {
	var info []commonstruct.CRoomInfo
	err := WyMysql.Table(commonstruct.WY_gm_config_room).Select("distinct settle_type,settle_type_cn,settle_type_sort").
		Where("game_dalei = ? and in_valid_time > 0", platform).Order("settle_type_sort").Find(&info).Error
	return info, err
}

func GetAgentOddsinfo(uuid int64, port string) commonstruct.AgentOdds {
	var list commonstruct.AgentOdds
	if err := WyMysql.Table(commonstruct.WY_user_odds).Where("uuid = ? and port = ? and in_use = 1", uuid, port).Find(&list).Error; err != nil {
		beego.Error("GetAgentOddsinfo err", err)
	}
	return list
}

func GetTeamMoneyInoutByDate(sufids []int64, begindate int64, enddate int64) []commonstruct.InoutStatistic {
	var list []commonstruct.InoutStatistic
	if len(sufids) <= 0 {
		return list
	}

	selectinfo := `date,
	sum(in_count) as in_count,
	sum(in_amount) as in_count,
	sum(out_count) as out_count,
	sum(out_amount) as out_count`

	if err := WyMysql.Table(commonstruct.WY_tmp_inout_statistic).Select(selectinfo).
		Where("uuid in (?) and date between ? and ? ", sufids, begindate, enddate).
		Find(&list).Error; err != nil {
		beego.Error("GetTeamMoneyInoutByDate err", begindate, enddate, err)
	}
	return list
}

func GetSufAgentoddslist(idlist []int64, pan string) []commonstruct.AgentOdds {
	var list []commonstruct.AgentOdds
	if err := WyMysql.Table(commonstruct.WY_user_odds).Where("uuid in (?) and port = ?", idlist, pan).Find(&list).Error; err != nil {
		beego.Error("GetSufAgentoddslist err", idlist, err)
	}
	return list
}

func GetSysDefaultodds(roomlist []int64) []commonstruct.DefaultOdds {
	var list []commonstruct.DefaultOdds
	if err := WyMysql.Table(commonstruct.WY_gm_config_defaultodds).Where("room_id in (?)", roomlist).Find(&list).Error; err != nil {
		beego.Error("GetSysDefaultodds err ", err)
	}
	return list
}

// 修改代理的水赔比
func UpdateOddspercent(masterid int64, pan string, odds float64, shui float64, shuitype int64) error {
	var updateValues map[string]interface{}
	switch shuitype {
	case 0:
		updateValues = map[string]interface{}{
			"odds": gorm.Expr("odds + ?", odds),
			// "shui_lottery": gorm.Expr("shui_lottery + ?", shui),
		}
	}

	if err := WyMysql.Table(commonstruct.WY_user_odds).Where("masterid = ? and port = ?", masterid, pan).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateOddspercent err ", err)
		return err
	}
	return nil
}

func GetLastVercode() commonstruct.ReCharge {
	var retinfo commonstruct.ReCharge
	if err := WyMysql.Table(commonstruct.WY_tmp_user_recharge).Where("pay_type in (18,19)").Order("time desc").Limit(1).Find(&retinfo).Error; err != nil { // 注单失败
		beego.Error("GetLastVercode err", err)
	}
	return retinfo
}

/**********************************
* 获取账变分类配置信息
**********************************/
func GetMoneychangetype(uuid int64, key string) ([]commonstruct.MoneychangType, error) {
	var infos []commonstruct.MoneychangType
	var selectarg string
	switch key {
	case "system":
		selectarg = "uuid = 10000"
	case "company":
		selectarg = fmt.Sprintf("uuid = %v", uuid)
	case "all":
		selectarg = fmt.Sprintf("uuid in (10000,%v)", uuid)
	default:
		return infos, errors.New("错误的Key")
	}

	if err := WyMysql.Table(commonstruct.WY_gm_config_moneytype).Where(selectarg).Find(&infos).Error; err != nil {
		beego.Error("GetMoneychangetype err", err)
		return infos, err
	}
	return infos, nil
}

func GetCoVercode(waytype int64) commonstruct.ReCharge {
	var retinfo commonstruct.ReCharge
	if err := WyMysql.Table(commonstruct.WY_tmp_user_recharge).Where("pay_type = ?", waytype).Order("time desc").Limit(1).Find(&retinfo).Error; err != nil {
		beego.Error("GetCoVercode err", err)
	}
	return retinfo
}

func GetWXRechargeinfo(vercode string) commonstruct.ReCharge {
	var retinfo commonstruct.ReCharge

	mintime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().AddDate(0, 0, -1)), 10, 64)
	nowtime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now()), 10, 64)

	if err := WyMysql.Table(commonstruct.WY_tmp_user_recharge).Where("order_title = ? and time between ? and ?", vercode, mintime, nowtime).Find(&retinfo).Error; err != nil { // 注单失败
		beego.Error("newrecharge err", retinfo, mintime, nowtime, err)

	}
	return retinfo
}

func GetUnrecharge(begintime int64, endtime int64) []commonstruct.ReCharge {
	var list []commonstruct.ReCharge
	if err := WyMysql.Table(commonstruct.WY_tmp_user_recharge).Where("pay_type = ? and time between ? and ? and checked = 0", 19, begintime, endtime).Find(&list).Error; err != nil { // 注单失败
		beego.Error("GetUnrecharge err", err)
	}
	return list
}

func NewRecharge(uuid int64, orderid int64, payway string, ordertitle string, orderinfo string, CallbackUrl string, amount float64,
	PartnerID string, SignType string, md5_key string, checked int64, enabletime int64) (*commonstruct.ReCharge, error) {
	newrecharge := &commonstruct.ReCharge{
		OrderID:    orderid,
		Time:       fmt.Sprintf("%d", commonfunc.GetNowtime()),
		Uuid:       uuid,
		PayType:    payway,
		OrderTitle: ordertitle,
		OrderInfo:  orderinfo,
		Callback:   CallbackUrl,
		Amount:     amount,
		PartnerID:  PartnerID,
		SignType:   SignType,
		SignData:   md5_key,
		Checked:    checked,
		EnableTime: enabletime,
	}
	retdb := WyMysql.Table(commonstruct.WY_tmp_user_recharge).Create(newrecharge)
	if err := retdb.Error; err != nil { // 注单失败
		beego.Error("newrecharge err", newrecharge, err)
		return nil, err
	}
	retrecharge := retdb.Value.(*commonstruct.ReCharge)
	return retrecharge, nil
}

func AddMoneyInoutLog(uuid int64, account string, masterid int64, operator string, reqtime int64, restime int64,
	reqtype int64, way int64, cashtype string, amount float64,
	ProfitCO float64, ProfitSaleIn float64, ProfitSaleThird float64, optip float64,
	oldgold float64, newgold float64, ip string, info string, expinfo string) (*commonstruct.MoneyInOut, error) {

	var oplog commonstruct.MoneyInOut
	idstr := fmt.Sprintf("%v%05d", commonfunc.BeijingTime().Format("060102150405"), uuid%100000)
	oplog.OrderID, _ = strconv.ParseInt(idstr, 10, 64)
	beego.Error("OrderID === ", idstr, oplog.OrderID)

	oplog.Uuid = uuid
	oplog.Account = account
	oplog.MasterID = masterid
	oplog.ReqTime = reqtime
	oplog.ResTime = restime
	oplog.ReqType = reqtype
	oplog.State = commonstruct.State_Success
	oplog.Checked = 1
	oplog.OpResult = commonstruct.State_Success
	oplog.Operator = operator
	oplog.Way = way
	oplog.CashType = cashtype
	oplog.Amount = amount
	oplog.ProfitCO = ProfitCO
	oplog.ProfitSaleIn = ProfitSaleIn
	oplog.ProfitSaleThird = ProfitSaleThird
	oplog.OpTip = optip
	oplog.OldGold = oldgold
	oplog.NewGold = newgold
	oplog.IP = ip
	oplog.IPPlace = info
	oplog.Expinfo = expinfo

	retdb := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Create(&oplog)
	if err := retdb.Error; err != nil {
		beego.Error("AddMoneyInoutLog err", oplog, err)
		return nil, err
	}
	retinfo := retdb.Value.(*commonstruct.MoneyInOut)
	return retinfo, nil
}

func BKAddMoneyInoutLog(uuid int64, account string, masterid int64, reqtime int64, restime int64,
	reqtype int64, way int64, cashtype string, amount float64,
	ProfitCO float64, ProfitSaleIn float64, ProfitSaleThird float64, optip float64,
	oldgold float64, newgold float64, ip string, info string) (*commonstruct.MoneyInOut, error) {

	var oplog commonstruct.MoneyInOut
	oplog.Uuid = uuid
	oplog.Account = account
	oplog.MasterID = masterid
	oplog.ReqTime = reqtime
	oplog.ResTime = restime
	oplog.ReqType = reqtype
	oplog.State = commonstruct.State_Success
	oplog.Checked = 1
	oplog.OpResult = commonstruct.State_Success
	oplog.Operator = account
	oplog.Way = way
	oplog.CashType = cashtype
	oplog.Amount = amount
	oplog.ProfitCO = ProfitCO
	oplog.ProfitSaleIn = ProfitSaleIn
	oplog.ProfitSaleThird = ProfitSaleThird
	oplog.OpTip = optip
	oplog.OldGold = oldgold
	oplog.NewGold = newgold
	oplog.IP = ip
	oplog.IPPlace = info

	retdb := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Create(&oplog)
	if err := retdb.Error; err != nil {
		beego.Error("AddMoneyInoutLog err", oplog, err)
		return nil, err
	}
	retinfo := retdb.Value.(*commonstruct.MoneyInOut)
	return retinfo, nil
}

const (
	PageCount = 50
)

func GetWalletTransferLog(uuids int64, begindate int64, enddate int64, offset int) []commonstruct.WalletTransfer {
	// 查询订单记录
	var logs []commonstruct.WalletTransfer

	//	var sqlArg string
	if err := WyMysql.Table(commonstruct.WY_tmp_user_wallettransfer).
		Where("uuid = ? and time between ? and ?", uuids, begindate, enddate).
		Order("time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
		beego.Error("GetWalletTransferLog err", uuids, err)
	}
	return logs
}

func GetWalletTransferPageinfo(uuids int64, begindate int64, enddate int64, offset int) (int64, int64) {
	// 查询订单记录
	var logs commonstruct.WalletTransfer

	//	var sqlArg string
	if err := WyMysql.Table(commonstruct.WY_tmp_user_wallettransfer).Select("count(*) as id").
		Where("uuid = ? and time between ? and ?", uuids, begindate, enddate).
		Find(&logs).Error; err != nil {
		beego.Error("GetWalletTransferPageinfo err", uuids, err)
	}
	return logs.ID, int64(math.Ceil(float64(logs.ID) / 20))
}

func GetWalletTransferPageinfoWithwallet(uuids int64, begindate int64, enddate int64, offset int, reswallet string, destwallet string) (int64, int64) {
	// 查询订单记录
	var logs commonstruct.WalletTransfer

	//	var sqlArg string
	if err := WyMysql.Table(commonstruct.WY_tmp_user_wallettransfer).Select("count(*) as id").
		Where("uuid = ? and time between ? and ? and res_wallet = ? and dest_wallet = ?", uuids, begindate, enddate, reswallet, destwallet).
		Find(&logs).Error; err != nil {
		beego.Error("GetWalletTransferPageinfoWithwallet err", uuids, err)
	}
	return logs.ID, int64(math.Ceil(float64(logs.ID) / 20))
}

func GetWalletTransferLogWithwallet(uuids int64, begindate int64, enddate int64, offset int, reswallet string, destwallet string) []commonstruct.WalletTransfer {
	// 查询订单记录
	var logs []commonstruct.WalletTransfer

	//	var sqlArg string
	if err := WyMysql.Table(commonstruct.WY_tmp_user_wallettransfer).
		Where("uuid = ? and time between ? and ? and res_wallet = ? and dest_wallet = ?", uuids, begindate, enddate, reswallet, destwallet).
		Order("time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
		beego.Error("GetWalletTransferLogWithwallet err", uuids, err)
	}
	return logs
}

func AddTransferWallet(tranferid int64, uuid int64, reswallet string, destwallet string, amount float64) error {
	var newtransfer commonstruct.WalletTransfer
	newtransfer.ID = tranferid
	newtransfer.Uuid = uuid
	newtransfer.Time = commonfunc.GetNowtime()
	newtransfer.ResWallet = reswallet
	newtransfer.DestWallet = destwallet
	newtransfer.Amount = amount

	if err := WyMysql.Table(commonstruct.WY_tmp_user_wallettransfer).Create(newtransfer).Error; err != nil { // 增加转换钱包记录
		beego.Error("AddTransferWallet err", newtransfer, err)
		return err
	}
	return nil
}

func UpdateTransferIninfo(orderid int64, action string, state int64, expinfo string) error {
	var updateValues map[string]interface{}
	switch action {
	case "IN":
		updateValues = map[string]interface{}{
			"dest_state": state,
			"in_time":    commonfunc.GetNowtime(),
			"expinfo":    gorm.Expr("CONCAT_WS('|',expinfo,?)", expinfo),
		}
	case "OUT":
		updateValues = map[string]interface{}{
			"res_state": state,
			"out_time":  commonfunc.GetNowtime(),
			"expinfo":   gorm.Expr("CONCAT_WS('|',expinfo,?)", expinfo),
		}
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_wallettransfer).Where("id = ?", orderid).Update(updateValues).Error; err != nil {
		beego.Error("UpdateTransferIninfo err", updateValues, err)
		return err
	}
	return nil
}

var (
	InoutType = map[int64]string{
		1: "In",
		2: "Out",
	}
)

func UpsertInoutStatistic(uuid int64, inouttype int64, amount float64) error {
	var info commonstruct.InoutStatistic
	date := commonfunc.GetNowdate()

	if retinfo := WyMysql.Table(commonstruct.WY_tmp_inout_statistic).Where("uuid = ? and date = ?", uuid, date).
		Find(&info); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			info.Uuid = uuid
			info.Date = date

			switch InoutType[inouttype] {
			case "In":
				info.InCount = 1
				info.InAmount = amount
			case "Out":
				info.OutCount = 1
				info.OutAmount = amount
			default:

			}
			if err := WyMysql.Table(commonstruct.WY_tmp_inout_statistic).Create(&info).Error; err != nil {
				beego.Error("create err ", info, err)
				return err
			}
		} else {
			beego.Error("UpsertInoutStatistic err ", uuid, inouttype, amount, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	var updateValues map[string]interface{}
	switch InoutType[inouttype] {
	case "In":
		updateValues = map[string]interface{}{
			"in_count":  gorm.Expr("in_count + 1"),
			"in_amount": gorm.Expr("in_amount + ?", amount),
		}
	case "Out":
		updateValues = map[string]interface{}{
			"out_count":  gorm.Expr("out_count + 1"),
			"out_amount": gorm.Expr("out_amount + ?", amount),
		}
	default:

	}
	if err := WyMysql.Table(commonstruct.WY_tmp_inout_statistic).
		Where("uuid = ? and date = ? ", uuid, date).Update(updateValues).Error; err != nil {
		beego.Error("UpsertInoutStatistic err ", uuid, amount, err)
	}

	return nil
}

// 获取服务器的特殊配置
func GetServerset(key string) (commonstruct.ServerSet, error) {
	var info commonstruct.ServerSet
	err := WyMysql.Table(commonstruct.WY_gm_config_serverset).Where("set_key = ?", key).Find(&info).Error
	if err != nil {
		beego.Error("GetServerset err ", key, err)
	}
	return info, err
}

func GetUsermoney(uuid int64) commonstruct.UserMoney {
	var Usermoney commonstruct.UserMoney
	if uuid == 0 {
		return Usermoney
	}

	if err := WyMysql.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&Usermoney).Error; err != nil {
		beego.Error("Select Usermoney err ", uuid, GetFuncName(3), err)
	}

	return Usermoney
}

func GetCreditusermoneyinfoS() ([]commonstruct.UserMoney, error) {
	var retinfos []commonstruct.UserMoney

	err := WyMysql.Table(commonstruct.WY_user_money).Where("uuid in (select uuid from wy_user_base where wallet_type in ('credit','credit,cash'))").Find(&retinfos).Error
	if err != nil {
		beego.Error("GetCreditusermoneyinfoS err ", err)
	}
	return retinfos, err
}

func GetUsersmoney(uuids []int64) commonstruct.UserMoney {
	selectinfo := `sum(cash) as cash,
		sum(cash_limit) as cash_limit,
		sum(tcp3) as tcp3,
		sum(tcp3_limit) as tcp3_limit,
		sum(fc3d) as fc3d,
		sum(fc3d_limit) as fc3d_limit,
		sum(taiwanliuhe) as taiwanliuhe,
		sum(taiwanliuhe_limit) as taiwanliuhe_limit,
		sum(taiwandaletou) as taiwandaletou,
		sum(taiwandaletou_limit) as taiwandaletou_limit,
		sum(xianggangliuhe) as xianggangliuhe,
		sum(xianggangliuhe_limit) as xianggangliuhe_limit,
		sum(aomenliuhe) as aomenliuhe,
		sum(aomenliuhe_limit) as aomenliuhe_limit`

	var Usermoney commonstruct.UserMoney
	if err := WyMysql.Table(commonstruct.WY_user_money).Select(selectinfo).Where("uuid in (?)", uuids).Find(&Usermoney).Error; err != nil {
		beego.Error("GetUsersmoney err ", uuids, err)
	}
	return Usermoney
}

func GetTeamAmount(sufids []int64) commonstruct.UserMoney {
	var moneyinfo commonstruct.UserMoney
	selectinfo := `sum(cash) as cash`
	if err := WyMysql.Table(commonstruct.WY_user_money).Select(selectinfo).Where("uuid in (?)", sufids).
		Find(&moneyinfo).Error; err != nil {
		beego.Error("GetTeamAmount err", err)
	}
	return moneyinfo
}

func GetSufmoneyS(idlist []int64) []commonstruct.UserMoney {
	var list []commonstruct.UserMoney
	if err := WyMysql.Table(commonstruct.WY_user_money).Where("uuid in (?)", idlist).Find(&list).Error; err != nil {
		beego.Error("GetSufmoneyS err", idlist, err)
	}
	return list
}

func UpdateBalance(uuid int64,
	optype commonstruct.MoneyUpdateType, //账变类型
	walletname string, // 钱包名称
	balance float64, // 账变金额
	expinfo string,
	Desc string) (float64, float64, error) { // 账变备注

	var usermoney commonstruct.UserMoney
	var newgold float64
	tx := WyMysql.Begin()
	switch walletname {
	case "cash":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("cash", gorm.Expr("cash + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Cash < 0 {
			tx.Rollback()
			return 0, 0, errors.New("现金钱包余额不足")
		}
		newgold = usermoney.Cash
	case "gaopin":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("gaopin", gorm.Expr("gaopin + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Gaopin < 0 {
			tx.Rollback()
			return 0, 0, errors.New("高频游戏余额不足")
		}
		newgold = usermoney.Gaopin
	case "gaopinlimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("gaopin_limit", gorm.Expr("gaopin_limit + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.GaopinLimit < 0 {
			tx.Rollback()
			return 0, 0, errors.New("高频游戏额度不足")
		}
		newgold = usermoney.GaopinLimit
	case "tcp3":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("tcp3", gorm.Expr("tcp3 + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Tcp3 < 0 {
			tx.Rollback()
			return 0, 0, errors.New("体彩P3余额不足")
		}
		newgold = usermoney.Tcp3
	case "tcp3limit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("tcp3_limit", gorm.Expr("tcp3_limit + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Tcp3Limit < 0 {
			tx.Rollback()
			return 0, 0, errors.New("体彩P3额度不足")
		}
		newgold = usermoney.Tcp3Limit
	case "fc3d":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("fc3d", gorm.Expr("fc3d + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Fc3d < 0 {
			tx.Rollback()
			return 0, 0, errors.New("福彩3D余额不足")
		}
		newgold = usermoney.Fc3d
	case "fc3dlimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("fc3d_limit", gorm.Expr("fc3d_limit + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Fc3dLimit < 0 {
			tx.Rollback()
			return 0, 0, errors.New("福彩3D额度不足")
		}
		newgold = usermoney.Fc3dLimit
	case "xianggangliuhe":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("xianggangliuhe", gorm.Expr("xianggangliuhe + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Xianggangliuhe < 0 {
			tx.Rollback()
			return 0, 0, errors.New("香港六合余额不足")
		}
		newgold = usermoney.Xianggangliuhe
	case "xianggangliuhelimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("xianggangliuhe_limit", gorm.Expr("xianggangliuhe_limit + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.XianggangliuheLimit < 0 {
			tx.Rollback()
			return 0, 0, errors.New("香港六合额度不足")
		}
		newgold = usermoney.XianggangliuheLimit
	case "taiwandaletou":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwandaletou", gorm.Expr("taiwandaletou + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Taiwandaletou < 0 {
			tx.Rollback()
			return 0, 0, errors.New("台湾大乐透余额不足")
		}
		newgold = usermoney.Taiwandaletou
	case "taiwandaletoulimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwandaletou_limit", gorm.Expr("taiwandaletou_limit + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.TaiwandaletouLimit < 0 {
			tx.Rollback()
			return 0, 0, errors.New("台湾大乐透额度不足")
		}
		newgold = usermoney.TaiwandaletouLimit
	case "aomenliuhe":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("aomenliuhe", gorm.Expr("aomenliuhe + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Aomenliuhe < 0 {
			tx.Rollback()
			return 0, 0, errors.New("澳门六合余额不足")
		}
		newgold = usermoney.Aomenliuhe
	case "aomenliuhelimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("aomenliuhe_limit", gorm.Expr("aomenliuhe_limit + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.AomenliuheLimit < 0 {
			tx.Rollback()
			return 0, 0, errors.New("澳门六合额度不足")
		}
		newgold = usermoney.AomenliuheLimit
	case "taiwanliuhe":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwanliuhe", gorm.Expr("taiwanliuhe + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.Taiwanliuhe < 0 {
			tx.Rollback()
			return 0, 0, errors.New("台湾六合余额不足")
		}
		newgold = usermoney.Taiwanliuhe
	case "taiwanliuhelimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwanliuhe_limit", gorm.Expr("taiwanliuhe_limit + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if usermoney.TaiwanliuheLimit < 0 {
			tx.Rollback()
			return 0, 0, errors.New("台湾六合额度不足")
		}
		newgold = usermoney.TaiwanliuheLimit
	default:
		tx.Rollback()
		return 0, 0, errors.New(fmt.Sprintf("错误的walletname类型[%v]", walletname))
	}

	var uptlog commonstruct.MoneyUpdateLog
	uptlog.Uuid = uuid
	uptlog.Time = commonfunc.GetNowtime()
	uptlog.OpType = optype
	uptlog.WalletName = walletname
	uptlog.OpGold = balance
	uptlog.OldGold = newgold - balance
	uptlog.NewGold = newgold
	uptlog.Expinfo = expinfo
	uptlog.Opinfo = Desc

	if err := tx.Table(commonstruct.WY_tmp_log_money).Create(&uptlog).Error; err != nil {
		beego.Error("AddUserMoneylog err ", uptlog, err)
		tx.Rollback()
		return 0, 0, err
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("UpdateBalance commit err ", err)
		tx.Rollback()
		return 0, 0, err
	}
	return newgold - balance, newgold, nil
}

func RecoverUsercreditamount(uuid int64,
	walletname string, // 钱包名称
	oldbalance float64,
	balance float64) error { // 账变备注

	updateValues := map[string]interface{}{
		walletname: balance,
	}

	tx := WyMysql.Begin()

	if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update(updateValues).Error; err != nil {
		tx.Rollback()
		return err
	}

	var uptlog commonstruct.MoneyUpdateLog
	uptlog.Uuid = uuid
	uptlog.Time = commonfunc.GetNowtime()
	uptlog.OpType = commonstruct.UptType_Recoveramount
	uptlog.WalletName = walletname
	uptlog.OpGold = balance
	uptlog.OldGold = oldbalance
	uptlog.NewGold = balance
	uptlog.Opinfo = fmt.Sprintf("每日信用额度恢复")

	if err := tx.Table(commonstruct.WY_tmp_log_money).Create(&uptlog).Error; err != nil {
		beego.Error("AddUserMoneylog err ", uptlog, err)
		tx.Rollback()
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("RecoverUsercreditamount commit err ", err)
		tx.Rollback()
		return err
	}
	return nil
}

func ReclaimBalance(uuid int64,
	optype commonstruct.MoneyUpdateType, //账变类型
	walletname string, // 钱包名称
	balance float64, // 账变金额
	expinfo string,
	Desc string) (float64, error) { // 账变备注

	var usermoney commonstruct.UserMoney
	var oldgold float64
	tx := WyMysql.Begin()
	switch walletname {
	case "cash":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("cash", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Cash
	case "gaopin":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("gaopin", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Gaopin
	case "gaopinlimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("gaopin_limit", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.GaopinLimit
	case "tcp3":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("tcp3", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Tcp3
	case "tcp3limit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("tcp3_limit", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Tcp3Limit
	case "fc3d":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("fc3d", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Fc3d
	case "fc3dlimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("fc3d_limit", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Fc3dLimit
	case "xianggangliuhe":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("xianggangliuhe", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Xianggangliuhe
	case "xianggangliuhelimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("xianggangliuhe_limit", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.XianggangliuheLimit
	case "taiwandaletou":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwandaletou", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Taiwandaletou
	case "taiwandaletoulimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwandaletou_limit", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.TaiwandaletouLimit
	case "aomenliuhe":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("aomenliuhe", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Aomenliuhe
	case "aomenliuhelimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("aomenliuhe_limit", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.AomenliuheLimit
	case "taiwanliuhe":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwanliuhe", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.Taiwanliuhe
	case "taiwanliuhelimit":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwanliueh_limit", 0).Error; err != nil {
			tx.Rollback()
			return oldgold, err
		}
		oldgold = usermoney.TaiwanliuheLimit
	default:
		tx.Rollback()
		return oldgold, errors.New(fmt.Sprintf("错误的walletname类型[%v]", walletname))
	}

	var uptlog commonstruct.MoneyUpdateLog
	uptlog.Uuid = uuid
	uptlog.Time = commonfunc.GetNowtime()
	uptlog.OpType = optype
	uptlog.WalletName = walletname
	uptlog.OpGold = balance
	uptlog.OldGold = oldgold
	uptlog.NewGold = 0
	uptlog.Expinfo = expinfo
	uptlog.Opinfo = Desc

	if err := tx.Table(commonstruct.WY_tmp_log_money).Create(&uptlog).Error; err != nil {
		beego.Error("AddUserMoneylog err ", uptlog, err)
		tx.Rollback()
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("ReclaimBalance commit err ", err)
		tx.Rollback()
		return oldgold, err
	}
	return oldgold, nil
}

func TXUpdateBalance(tx *gorm.DB, uuid int64, walletname string, balance float64, bcheckzero bool) (float64, error) {
	var usermoney commonstruct.UserMoney
	var newgold float64

	switch walletname {
	case "cash":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("cash", gorm.Expr("cash + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if bcheckzero {
			if usermoney.Cash < 0 {
				tx.Rollback()
				return 0, errors.New("现金额度不足")
			}
		}
		newgold = usermoney.Cash
	case "gaopin":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("gaopin", gorm.Expr("gaopin + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if bcheckzero {
			if usermoney.Gaopin < 0 {
				tx.Rollback()
				return 0, errors.New("高频额度不足")
			}
		}
		newgold = usermoney.Gaopin
	case "tcp3":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("tcp3", gorm.Expr("tcp3 + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if bcheckzero {
			if usermoney.Tcp3 < 0 {
				tx.Rollback()
				return 0, errors.New("P3额度不足")
			}
		}
		newgold = usermoney.Tcp3
	case "fc3d":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("fc3d", gorm.Expr("fc3d + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		if bcheckzero {
			if usermoney.Fc3d < 0 {
				tx.Rollback()
				return 0, errors.New("3D额度不足")
			}
		}
		newgold = usermoney.Fc3d
	case "xianggangliuhe":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("xianggangliuhe", gorm.Expr("xianggangliuhe + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		if bcheckzero {
			if usermoney.Xianggangliuhe < 0 {
				tx.Rollback()
				return 0, errors.New("香港额度不足")
			}
		}
		newgold = usermoney.Xianggangliuhe
	case "taiwandaletou":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwandaletou", gorm.Expr("taiwandaletou + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		if bcheckzero {
			if usermoney.Taiwandaletou < 0 {
				tx.Rollback()
				return 0, errors.New("台湾额度不足")
			}
		}
		newgold = usermoney.Taiwandaletou
	case "aomenliuhe":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("aomenliuhe", gorm.Expr("aomenliuhe + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		if bcheckzero {
			if usermoney.Aomenliuhe < 0 {
				tx.Rollback()
				return 0, errors.New("澳门额度不足")
			}
		}
		newgold = usermoney.Aomenliuhe
	case "taiwanliuhe":
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Update("taiwanliuhe", gorm.Expr("taiwanliuhe + ?", balance)).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Table(commonstruct.WY_user_money).Where("uuid = ?", uuid).Find(&usermoney).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		if bcheckzero {
			if usermoney.Taiwanliuhe < 0 {
				tx.Rollback()
				return 0, errors.New("台湾额度不足")
			}
		}
		newgold = usermoney.Taiwanliuhe
	default:
		tx.Rollback()
		return 0, errors.New(fmt.Sprintf("错误的walletname类型[%v]", walletname))
	}

	return newgold, nil
}

/***********
* 设置公司外接游戏限额
 ***********/
func UpsertExplimit(companyid int64, expname string, limit_amount float64) error {

	// var limit commonstruct.ExpLimit
	// tblname := commonstruct.WY_exp_limit
	// if retinfo := WyMysql.Table(tblname).Find(&limit, "company_id = ? and exp_name = ?", companyid, expname); retinfo.Error != nil {
	// 	if retinfo.RecordNotFound() {
	// 		limit.CompanyID = companyid
	// 		limit.ExpName = expname
	// 		limit.LimitAmount = limit_amount
	// 		if err := WyMysql.Table(tblname).Create(&limit).Error; err != nil {
	// 			beego.Error("create err ", limit, err)
	// 			return err
	// 		}
	// 	} else {
	// 		beego.Error("UpsertExplimit err ", companyid, expname, retinfo.Error.Error())
	// 		return retinfo.Error
	// 	}
	// 	return nil
	// }

	// updateValues := map[string]interface{}{
	// 	"limit_amount": limit_amount,
	// }
	// if err := WyMysql.Table(tblname).
	// 	Where("company_id = ? and exp_name = ?", companyid, expname).Update(updateValues).Error; err != nil {
	// 	beego.Error("UpsertExplimit err ", companyid, expname, err)
	// 	return err
	// }
	return nil
}

func SetLastgame(uuid int64, lastgame string) error {
	if err := WyMysql.Table(commonstruct.WY_user_money).Where("uuid = ? ", uuid).Update("lastgame", lastgame).Error; err != nil {
		beego.Error("SetLastgame err ", uuid, lastgame, err.Error())
		return err
	}
	return nil
}

func SetTag(resid int64, destid int64, tag string) error {

	var oldtag commonstruct.UserTag

	tblname := commonstruct.WY_tmp_user_tag
	if retinfo := WyMysql.Table(tblname).
		Where("res_id = ? and dest_id = ?", resid, destid).Find(&oldtag); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			oldtag.ResID = resid
			oldtag.DestID = destid
			oldtag.Tag = tag
			if err := WyMysql.Table(tblname).Create(&oldtag).Error; err != nil {
				beego.Error("create err ", oldtag, err)
				return err
			}
		} else {
			beego.Error("SetTag err ", resid, destid, tag, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	updateValues := map[string]interface{}{
		"tag": tag,
	}
	if err := WyMysql.Table(tblname).
		Where("res_id = ? and dest_id = ?", resid, destid).Update(updateValues).Error; err != nil {
		beego.Error("SetTag err ", resid, destid, tag, err)
		return err
	}
	return nil
}

func GetTag(resid int64, destid int64) commonstruct.UserTag {
	var oldtag commonstruct.UserTag
	tblname := commonstruct.WY_tmp_user_tag
	if err := WyMysql.Table(tblname).Where("res_id = ? and dest_id = ?", resid, destid).Find(&oldtag).Error; err != nil {
		//		beego.Error("SetTag err ", resid, destid, err.Error())
	}
	return oldtag
}

// 修改代理的水赔比
func UpdateAgentOdds(uuid int64, pan string, odds float64, shui_lottery float64, shui_actual float64, shui_electric float64, shui_card float64, shui_sport float64) error {
	updateValues := map[string]interface{}{
		"odds":          odds,
		"shui_lottery":  shui_lottery,
		"shui_actual":   shui_actual,
		"shui_electric": shui_electric,
		"shui_card":     shui_card,
		"shui_sport":    shui_sport,
	}
	if err := WyMysql.Table(commonstruct.WY_user_odds).Where("uuid = ? and port = ?", uuid, pan).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateOddspercent err ", err)
		return err
	}
	return nil
}

func GetRechargeNavilist() []string {
	var navilist []string

	// 支付方式的详细内容
	var list []commonstruct.RechargeWayinfo
	if err := WyMysql.Table(commonstruct.WY_gm_config_recharge_way).
		Select("distinct navi_type as navi_type").Where("navi_type != ''").
		Find(&list).Error; err != nil {
		return navilist
	}

	for _, v := range list {
		navilist = append(navilist, v.NaviType)
	}
	return navilist
}

func GetRechargeNavidetail(navi string) []commonstruct.RechargeWayinfo {
	// 支付方式的详细内容
	var list []commonstruct.RechargeWayinfo
	if err := WyMysql.Table(commonstruct.WY_gm_config_recharge_way).
		Where("navi_type = ?", navi).
		Find(&list).Error; err != nil {
		beego.Error("GetRechargeNavidetail err", err.Error())
	}

	return list
}

// 查询未结算期号信息
func GetUnreadmsgStats(uuid int64) commonstruct.FeedbackMsg {
	var info commonstruct.FeedbackMsg
	if err := WyMysql.Table(commonstruct.WY_tmp_user_retmsg).Select("Count(*) as id").Where("uuid = ? and flag = 0", uuid).Find(&info).Error; err != nil {
		beego.Error("GetUnreadmsgStats err ", info, err)
	}
	return info
}

// 查询未结算期号信息
func GetUnreadmsgInfos(uuid int64) []commonstruct.FeedbackMsg {
	var infos []commonstruct.FeedbackMsg
	if err := WyMysql.Table(commonstruct.WY_tmp_user_retmsg).Where("uuid = ?", uuid).Order("time desc").Find(&infos).Error; err != nil {
		beego.Error("GetUnreadmsgInfos err ", infos, err)
	}
	return infos
}

// 修改已读标志位
func UpdateReadflag(id int64) error {
	updateValues := map[string]interface{}{
		"flag": 1,
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_user_retmsg).Where("id = ?", id).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateReadflag err ", err)
		return err
	}
	return nil
}

// 删除单条信息
func DeleteFeedbackmsgByid(id int64) error {

	sql := fmt.Sprintf("delete from %v where id = %v;", commonstruct.WY_tmp_user_retmsg, id)
	if err := WyMysql.Exec(sql).Error; err != nil {
		beego.Error("DeleteFeedbackmsgByid err ", err)
	}
	return nil
}

// 删除已读信息
func DeleteFeedbackmsgByflag(uuid int64, flag int64) error {
	sql := fmt.Sprintf("delete from %v where flag = %v and uuid = %v;", commonstruct.WY_tmp_user_retmsg, flag, uuid)
	if err := WyMysql.Exec(sql).Error; err != nil {
		beego.Error("DeleteFeedbackmsgByid err ", err)
	}
	return nil
}

func GetCoSelfrechargeway(uuid int64) ([]commonstruct.CoSelfRechargeway, error) {
	tblname := commonstruct.WY_company_selfrechargeway
	var data []commonstruct.CoSelfRechargeway
	if err := WyMysql.Table(tblname).Where("company_id = ? and state in (7500,7502,7503)", uuid).Order("sort").Find(&data).Error; err != nil {
		beego.Error("GetCoSelfrechargeway err ", err)
		return data, err
	}
	return data, nil
}

func GetThirdRechargeway(masterid int64) ([]commonstruct.CoRechargeWay, error) {
	var info []commonstruct.CoRechargeWay
	err := WyMysql.Table(commonstruct.WY_company_rechargeway).Where("company_id = ? and state in (7500,7502,7503)", masterid).
		Order("sort").Find(&info).Error
	return info, err
}

func GetRechargewayinfo(servicetype int64) (commonstruct.RechargeWayinfo, error) {
	var data commonstruct.RechargeWayinfo
	err := WyMysql.Table(commonstruct.WY_gm_config_recharge_way).Where("service_type = ?", servicetype).Find(&data).Error
	return data, err
}

/*************************
* 更新用户日统计
 ************************/
func UpsertDayStatistic(uuid int64, date int64, column string, optype string, value float64) error {
	var info commonstruct.DayStatistic

	if retinfo := WyMysql.Table(commonstruct.WY_tmp_user_daystatistic).Where("uuid = ? and date = ?", uuid, date).
		Find(&info); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			info.Uuid = uuid
			info.Date = date

			if err := WyMysql.Table(commonstruct.WY_tmp_user_daystatistic).Create(&info).Error; err != nil {
				beego.Error("create err ", info, err)
				return err
			}
		} else {
			beego.Error("UpsertDayStatistic err ", uuid, date, column, value, retinfo.Error.Error())
			return retinfo.Error
		}
	}

	var updateValues map[string]interface{}
	switch optype {
	case "add":
		updateValues = map[string]interface{}{
			column: gorm.Expr(fmt.Sprintf("%v + ?", column), value),
		}
	case "update":
		updateValues = map[string]interface{}{
			column: value,
		}
	default:
		return errors.New(fmt.Sprintf("error optype == %v", optype))
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_user_daystatistic).
		Where("uuid = ? and date = ? ", uuid, date).Update(updateValues).Error; err != nil {
		beego.Error("UpsertDayStatistic err ", uuid, date, column, value, err)
	}

	return nil
}

func GetUserDaystatistc(uuid int64, date int64) (commonstruct.DayStatistic, error) {
	var info commonstruct.DayStatistic

	err := WyMysql.Table(commonstruct.WY_tmp_user_daystatistic).Where("uuid = ? and date = ?", uuid, date).
		Find(&info).Error

	return info, err
}

func GetMoneyInout(orderid int64) (commonstruct.MoneyInOut, error) {
	var oldinfo commonstruct.MoneyInOut
	err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).
		Find(&oldinfo).Error
	return oldinfo, err
}

func CheckRechargeorder(orderid int64, op_result int64, state int, amount float64, oldgold float64, newgold float64, expinfo string) error {
	var log commonstruct.MoneyInOut
	tx := WyMysql.Begin()

	if err := tx.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Find(&log).Error; err != nil {
		beego.Error("CheckRechargeorder err", err)
		tx.Rollback()
		return err
	} else {
		if log.Checked > 0 { // 订单已处理
			tx.Rollback()
			return errors.New("订单已处理")
		} else {
			updateValues1 := map[string]interface{}{
				"checked":   1,
				"op_result": op_result,
				"expinfo":   gorm.Expr("CONCAT_WS('|',expinfo,?)", expinfo),
				"res_time":  commonfunc.GetNowtime(),
				"state":     state,
				"old_gold":  oldgold,
				"new_gold":  newgold,
			}
			if err := tx.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Update(updateValues1).Error; err != nil {
				beego.Error("CheckRechargeorder err", err)
				tx.Rollback()
				return err
			}

			updateValues := map[string]interface{}{
				"checked": 1,
				"amount":  amount,
			}
			if err := tx.Table(commonstruct.WY_tmp_user_recharge).Where("order_id = ?", orderid).Update(updateValues).Error; err != nil {
				beego.Error("CheckRechargeorder err", err)
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("InitProfit commit err ", err)
		tx.Rollback()
		return err
	}
	return nil
}
func BKCheckRechargeorder(orderid int64, op_result int64, state int, amount float64, oldgold float64, newgold float64, expinfo string) error {
	tx := WyMysql.Begin()

	retinfo := tx.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Update("checked", 1)
	if err := retinfo.Error; err != nil {
		beego.Error("CheckRechargeorder err", err)
		tx.Rollback()
		return err
	}

	if retinfo.RowsAffected == 0 {
		beego.Error("order is checked ")
		tx.Rollback()
		return errors.New("订单已处理")
	} else {
		beego.Error("CheckRechargeorder", orderid)
	}

	updateValues1 := map[string]interface{}{
		"op_result": op_result,
		"expinfo":   gorm.Expr("CONCAT_WS('|',expinfo,?)", expinfo),
		"res_time":  commonfunc.GetNowtime(),
		"state":     state,
		"old_gold":  oldgold,
		"new_gold":  newgold,
	}
	if err := tx.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Update(updateValues1).Error; err != nil {
		beego.Error("CheckRechargeorder err", err)
		tx.Rollback()
		return err
	}

	updateValues := map[string]interface{}{
		"checked": 1,
		"amount":  amount,
	}
	if err := tx.Table(commonstruct.WY_tmp_user_recharge).Where("order_id = ?", orderid).Update(updateValues).Error; err != nil {
		beego.Error("CheckRechargeorder err", err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("InitProfit commit err ", err)
		tx.Rollback()
		return err
	}
	return nil
}

func CleartagExec(tblname string, teamids []int64) error {
	if err := WyMysql.Table(tblname).Delete(nil, "dest_id in (?)", teamids).Error; err != nil {
		beego.Error("CleartagExec", tblname, teamids, " err ", err)
		return errors.New(fmt.Sprintf("删除%v失败", tblname))
	}
	return nil
}

func ClearExec(tblname string, teamids []int64) error {
	if err := WyMysql.Table(tblname).Delete(nil, "uuid in (?)", teamids).Error; err != nil {
		beego.Error("ClearExec", tblname, teamids, " err ", err)
		return errors.New(fmt.Sprintf("删除%v失败", tblname))
	}
	return nil
}

func ClearCoExec(tblname string, teamids []int64) error {
	if err := WyMysql.Table(tblname).Delete(nil, "company_id in (?)", teamids).Error; err != nil {
		beego.Error("ClearExec", tblname, teamids, " err ", err)
		return errors.New(fmt.Sprintf("删除%v失败", tblname))
	}
	return nil
}

func DeleteGame(roomid int64) error {
	if err := WyMysql.Table(commonstruct.WY_gm_config_room).Delete(nil, "id = ?", roomid).Error; err != nil {
		beego.Error("DeleteGame", commonstruct.WY_gm_config_room, roomid, " err ", err)
		return err
	}
	return nil
}

func SQLExec(sqlstr string) error {
	if err := WyMysql.Exec(sqlstr).Error; err != nil {
		beego.Error("DeleteUserbkset err ", sqlstr, err)
		return err
	}
	return nil
}

/***********
* 获取代理赔率信息
 ***********/
// func GetAgentOdds(uuid int64, port string) commonstruct.AgentOdds {
// 	var list commonstruct.AgentOdds
// 	if err := WyMysql.Table(commonstruct.WY_user_odds).Where("uuid = ? and port = ? and in_use = 1", uuid, port).Find(&list).Error; err != nil {
// 		beego.Error("GetAgentOddsinfo err", err)
// 	}
// 	return list
// }

func GetPortclass(lotterydalei string, settletype string) []commonstruct.PortClass {
	var classlist []commonstruct.PortClass
	if err := WyMysql.Table(commonstruct.WY_gm_config_portclass).
		Where("lottery_dalei = ? and settle_type = ?", lotterydalei, settletype).Find(&classlist).Error; err != nil {
	}
	return classlist
}

func GetDyncamount(uuid int64, roomid int64) commonstruct.DyncAmount {
	var oddsinfo commonstruct.DyncAmount
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_user_dyncamount).
		Where("uuid = ? and room_id = ?", uuid, roomid).Find(&oddsinfo); retinfo.Error != nil {
		if !retinfo.RecordNotFound() {
			beego.Error("GetDyncamount err", uuid, roomid, retinfo.Error.Error())
		}
	}
	return oddsinfo
}

func GetPortDalei(lotterydalei string, settletype string) []commonstruct.PortClass {
	var classlist []commonstruct.PortClass
	if err := WyMysql.Table(commonstruct.WY_gm_config_portclass).Select("distinct game_dalei").
		Where("lottery_dalei = ? and settle_type = ?", lotterydalei, settletype).Find(&classlist).Error; err != nil {
	}
	return classlist
}

func GetChanglongoddsSet(uuid int64, roomid int64) (commonstruct.ChanglongOdds, error) {
	var oddsinfo commonstruct.ChanglongOdds
	err := WyMysql.Table(commonstruct.WY_company_changlongodds).
		Where("uuid = ? and room_id = ?", uuid, roomid).Find(&oddsinfo).Error
	if err != nil {
		beego.Error("GetChanglongoddsSet err", uuid, roomid, err.Error())
	}
	return oddsinfo, err
}

func GetUserlastlogininfo(uuid int64) (commonstruct.CLoginInfo, error) {
	var oplogs commonstruct.CLoginInfo
	err := WyMysql.Table(commonstruct.WY_tmp_user_logininfo).
		Where("uuid = ?", uuid).Order("id desc").Limit(1).
		Find(&oplogs).Error
	if err != nil {
		if err.Error() != "record not found" {
			beego.Error("GetUserlastlogininfo err", err.Error())
		}
	}
	return oplogs, err
}

func GetOpPagenum(uuids []int64, optype int64, begintime int64, endtime int64, pagecount int) (int64, int64) {
	var oplogs commonstruct.UserOpLog
	if optype == 0 {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_userop).Select("count(*) as id").
			Where("uuid in (?) and time between ? and ?", uuids, begintime, endtime).
			Find(&oplogs).Error; err != nil {
			beego.Error("GetOpLogs", err.Error())
		}
	} else {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_userop).Select("count(*) as id").
			Where("uuid in (?) and op_type = ?  and time between ? and ?", uuids, optype, begintime, endtime).
			Find(&oplogs).Error; err != nil {
			beego.Error("GetOpLogs", err.Error())
		}
	}

	return int64(oplogs.ID), int64(math.Ceil(float64(oplogs.ID) / float64(pagecount)))
}

func GetOpLogs(uuids []int64, optype int64, offset int, pagecount int, begintime int64, endtime int64) []commonstruct.UserOpLog {
	var oplogs []commonstruct.UserOpLog
	if optype == 0 {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_userop).Where("uuid in (?) and time between ? and ?", uuids, begintime, endtime).
			Order("time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&oplogs).Error; err != nil {
			beego.Error("GetOpLogs", err.Error())
		}
	} else {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_userop).Where("uuid in (?) and op_type = ? and time between ? and ?", uuids, optype, begintime, endtime).
			Order("time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&oplogs).Error; err != nil {
			beego.Error("GetOpLogs", err.Error())
		}
	}

	return oplogs
}

func GetSufInoutStatisticS(idlist []int64) []commonstruct.InoutStatistic {
	var list []commonstruct.InoutStatistic

	selectinfo := `uuid,
	sum(in_count) as in_count,
	sum(in_amount) as in_amount,
	sum(out_count) as out_count,
	sum(out_amount) as out_amount`

	if err := WyMysql.Table(commonstruct.WY_tmp_inout_statistic).Select(selectinfo).
		Where("uuid in (?)", idlist).Group("uuid").
		Find(&list).Error; err != nil {
		beego.Error("GetSufInoutStatisticS err", idlist, err)
	}
	return list
}

/***********
* 查询单日统计
 ***********/
func GetUserTeamstats(companyid int64, date int64) commonstruct.UserTeamstats {
	var stats commonstruct.UserTeamstats
	if err := WyMysql.Table(commonstruct.WY_tmp_user_teamstats).
		Where("uuid = ? and date = ?", companyid, date).Find(&stats).Error; err != nil {
		if err.Error() != "record not found" {
			beego.Error("GetCostats err", err)
		} else {
			stats.Uuid = companyid
			stats.Date = date
		}
	}
	return stats
}

/***********
* 更新公司的单日统计数据
 ***********/
func UpsertUserteamstats(companyid int64, date int64, column string, value interface{}) error {
	var info commonstruct.UserTeamstats
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_user_teamstats).Where("uuid = ? and date = ?", companyid, date).
		Find(&info); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			info.Uuid = companyid
			info.Date = date

			switch column {
			case "regnum":
				info.RegNum = int64(value.(int))
			case "highestnum":
				info.HighestNum = int64(value.(int))
			case "ordernum":
				info.OrderNum = int64(value.(int))
			case "wallet_amount":
				info.WalletAmount = value.(float64)
			case "unsettle_amount":
				info.UnsettleAmount = value.(float64)
			case "unwithdraw_amount":
				info.UnwithdrawAmount = value.(float64)
			case "recharge_amount":
				info.RechargeAmount = value.(float64)
			case "mgr_remit":
				info.MgrRemit = value.(float64)
			case "mgr_increase":
				info.MgrIncrease = value.(float64)
			case "mgr_decrease":
				info.MgrDecrease = value.(float64)
			case "withdraw_amount":
				info.WithdrawAmount = value.(float64)
			case "sys_revoke":
				info.SysRevoke = value.(float64)
			case "sys_back":
				info.SysBack = value.(float64)
			case "order_amount":
				info.OrderAmount = value.(float64)
			case "revoke_amount":
				info.RevokeAmount = value.(float64)
			case "settle_amount":
				info.SettleAmount = value.(float64)
			case "wager_amount":
				info.WagerAmount = value.(float64)
			case "gross_profit":
				info.GrossProfit = value.(float64)
			case "shui_amount":
				info.ShuiAmount = value.(float64)
			case "pure_profit":
				info.PureProfit = value.(float64)
			default:
				beego.Error("err column ", column)
				return nil
			}
			if err := WyMysql.Table(commonstruct.WY_tmp_user_teamstats).Create(&info).Error; err != nil {
				beego.Error("create err ", info, err)
				return err
			}
		} else {
			beego.Error("UpsertUserteamstats err ", companyid, date, column, value, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	var updateValues map[string]interface{}
	switch column {
	case "regnum":
		updateValues = map[string]interface{}{
			"reg_num": gorm.Expr("reg_num + ?", value),
		}
	case "highestnum":
		updateValues = map[string]interface{}{
			"highest_num": gorm.Expr("highest_num + ?", value),
		}
	case "ordernum":
		updateValues = map[string]interface{}{
			"order_num": gorm.Expr("order_num + ?", value),
		}
	case "wallet_amount":
		updateValues = map[string]interface{}{
			"wallet_amount": value,
		}
	case "unsettle_amount":
		updateValues = map[string]interface{}{
			"unsettle_amount": value,
		}
	case "unwithdraw_amount":
		updateValues = map[string]interface{}{
			"unwithdraw_amount": value,
		}
	case "recharge_amount":
		updateValues = map[string]interface{}{
			"recharge_amount": value,
		}
	case "mgr_remit":
		updateValues = map[string]interface{}{
			"mgr_remit": value,
		}
	case "mgr_increase":
		updateValues = map[string]interface{}{
			"mgr_increase": value,
		}
	case "mgr_decrease":
		updateValues = map[string]interface{}{
			"mgr_decrease": value,
		}
	case "withdraw_amount":
		updateValues = map[string]interface{}{
			"withdraw_amount": value,
		}
	case "sys_revoke":
		updateValues = map[string]interface{}{
			"sys_revoke": value,
		}
	case "sys_back":
		updateValues = map[string]interface{}{
			"sys_back": value,
		}
	case "order_amount":
		updateValues = map[string]interface{}{
			"order_amount": value,
		}
	case "revoke_amount":
		updateValues = map[string]interface{}{
			"revoke_amount": value,
		}
	case "settle_amount":
		updateValues = map[string]interface{}{
			"settle_amount": value,
		}
	case "wager_amount":
		updateValues = map[string]interface{}{
			"wager_amount": value,
		}
	case "gross_profit":
		updateValues = map[string]interface{}{
			"gross_profit": value,
		}
	case "shui_amount":
		updateValues = map[string]interface{}{
			"shui_amount": value,
		}
	case "pure_profit":
		updateValues = map[string]interface{}{
			"pure_profit": value,
		}
	default:
		beego.Error("err column ", column)
		return nil
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_user_teamstats).
		Where("uuid = ? and date = ? ", companyid, date).Update(updateValues).Error; err != nil {
		beego.Error("UpsertUserteamstats err ", companyid, date, err)
		return err
	}
	return nil
}

/****************************
* 获取公司日统计值
****************************/
func GetCoStatisticByDate(teamlist []int64, begindate int64, enddate int64) []commonstruct.DayStatistic {
	var statistic []commonstruct.DayStatistic
	if len(teamlist) <= 0 {
		return statistic
	}

	selectinfo := `
	date,
	sum(recharge_succamount) as recharge_succamount,
	sum(withdraw_succamount) as withdraw_succamount`

	if err := WyMysql.Table(commonstruct.WY_tmp_user_daystatistic).Select(selectinfo).
		Where("uuid in (?) and date between ? and ?", teamlist, begindate, enddate).Group("date").
		Find(&statistic).Error; err != nil {
		beego.Error("GetCoStatisticByDate err", teamlist, begindate, enddate, err)
	}
	return statistic
}

func GetTeamSumamount(sufids []int64) commonstruct.UserMoney {
	var moneyinfo commonstruct.UserMoney
	selectinfo := `sum(cash) as cash`
	if err := WyMysql.Table(commonstruct.WY_user_money).Select(selectinfo).Where("uuid in (?)", sufids).
		Find(&moneyinfo).Error; err != nil {
		beego.Error("GetTeamAmount err", err)
	}
	return moneyinfo
}

func GetUnWithdrawOrder(orderid int64) commonstruct.MoneyInOut {
	// 查询订单记录
	var logs commonstruct.MoneyInOut
	if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Find(&logs).Error; err != nil {
		beego.Error("GetUnWithdrawOrder err", orderid, err)
	}
	return logs
}

func GetWithdrawInfo(orderid int64) commonstruct.MoneyInOut {
	//	// 查询订单记录
	//	var logs commonstruct.Withdraw
	//	if err := WyMysql.Table(commonstruct.WY_gm_user_withdraw).Where("order_id = ?", orderid).Find(&logs).Error; err != nil {
	//		beego.Error("GetWithdrawInfo err", orderid, err)
	//	}
	// 查询订单记录
	var logs commonstruct.MoneyInOut
	if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Find(&logs).Error; err != nil {
		beego.Error("GetWithdrawInfo err", orderid, err)
	}
	return logs
}

/***********
* 获取日初第一条资金记录
 ***********/
func GetFristlog(uuid int64, begindate int64) commonstruct.MoneyUpdateLog {
	selectarg := fmt.Sprintf("id = ( select min(id) from wy_tmp_log_money where uuid = %v and time > %v)",
		uuid, commonfunc.GetBegintime(begindate))

	// 查询订单记录
	var statistic commonstruct.MoneyUpdateLog
	if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Where(selectarg).Find(&statistic).Error; err != nil {
		if err.Error() != "record not found" {
			beego.Error("GetFristlog err", uuid, begindate, err)
		}
	}

	return statistic
}

/***********
* 获取最后一条资金变动记录
 ***********/
func GetLatestLog(uuid int64, enddate int64) commonstruct.MoneyUpdateLog {
	selectarg := fmt.Sprintf("id = ( select max(id) from wy_tmp_log_money where uuid = %v and time <= %v )",
		uuid, commonfunc.GetEndtime(enddate))

	// 查询订单记录
	var statistic commonstruct.MoneyUpdateLog
	if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Where(selectarg).Find(&statistic).Error; err != nil {
		if err.Error() != "record not found" {
			beego.Error("GetLatestLog err", uuid, enddate, err)
		}
	}

	return statistic
}

/***********
* 获取资金变动统计
 ***********/
func GetLogstats(uuid int64, begindate int64, enddate int64) []commonstruct.MoneyUpdateLog {
	// 查询订单记录
	var statistic []commonstruct.MoneyUpdateLog
	if err := WyMysql.Table(commonstruct.WY_tmp_log_money).Select("op_type ,sum(op_gold) as op_gold,count(*) as id").
		Where("uuid = ? and time between ? and ?  ", uuid, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate)).Group("op_type").Find(&statistic).Error; err != nil {
		beego.Error("GetLogstats err", uuid, begindate, enddate, err)
	}

	return statistic
}

func GetExpgameList() []commonstruct.ExpGame {
	var list []commonstruct.ExpGame
	if err := WyMysql.Table(commonstruct.WY_gm_config_expgame).Where("in_use = 1").Find(&list).Error; err != nil {
		beego.Error("GetExpgameList", err.Error())
	}
	return list
}

/***********
* 查询彩种的刷新过期时间
 ***********/
func GetMaxlttime(tablesuf string) commonstruct.LotteryOpentime {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_lotterytime_, tablesuf)

	var Opentime commonstruct.LotteryOpentime
	if err := WyMysql.Table(tblname).Select("id,expect,unix_timestamp(opentimestamp) as opentimestamp").
		Where(fmt.Sprintf("opentimestamp = (select max(opentimestamp) from %v)", tblname)).Find(&Opentime).Error; err != nil {
		beego.Error("GetMaxlttime err", tablesuf, err)
	}
	return Opentime
}

func GetUserOpStatistic(uuid int64) commonstruct.OpStatistic {
	var user commonstruct.OpStatistic
	if err := WyMysql.Table(commonstruct.WY_tmp_user_op_statistic).Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		beego.Error("GetUserOpStatistic err", uuid, err)
	}
	return user
}

func BKGetMoneyInoutPageinfo(uuids []int64, begindate int64, enddate int64, reqtype int, way int, state int) (int64, int64) {
	// 查询订单记录
	var logs commonstruct.Withdraw

	switch reqtype {
	case -1: //全部查询
		switch way {
		case -1:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ?", uuids, begindate, enddate).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and state = ?", uuids, begindate, enddate, state).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			}
		default:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and way = ?", uuids, begindate, enddate, way).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and way = ? and state = ?", uuids, begindate, enddate, way, state).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			}
		}
	default:
		switch way {
		case -1:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and req_type = ?", uuids, begindate, enddate, reqtype).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and state = ?", uuids, begindate, enddate, reqtype, state).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			}
		default:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and way = ?", uuids, begindate, enddate, reqtype, way).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and way = ? and state = ?", uuids, begindate, enddate, reqtype, way, state).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			}
		}
	}
	return logs.OrderID, int64(math.Ceil(float64(logs.OrderID) / 20))
}

/***********
* 获取用户组资金变动日志
 ***********/
func BKGetMoneyInoutLog(uuids []int64, begindate int64, enddate int64, offset int, reqtype int, way int, state int) []commonstruct.MoneyInOut {
	// 查询订单记录
	var logs []commonstruct.MoneyInOut
	//	var sqlArg string

	switch reqtype {
	case -1: //全部查询
		switch way {
		case -1:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ?", uuids, begindate, enddate).
					Order("req_time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and state = ?", uuids, begindate, enddate, state).
					Order("req_time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			}
		default:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and way = ?", uuids, begindate, enddate, way).
					Order("req_time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and way = ? and state = ?", uuids, begindate, enddate, way, state).
					Order("req_time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			}
		}
	default:
		switch way {
		case -1:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and req_type = ?", uuids, begindate, enddate, reqtype).
					Order("req_time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and state = ?", uuids, begindate, enddate, reqtype, state).
					Order("req_time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			}
		default:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and way = ?", uuids, begindate, enddate, reqtype, way).
					Order("req_time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and way = ? and state = ?", uuids, begindate, enddate, reqtype, way, state).
					Order("req_time desc").Limit(PageCount).Offset((offset - 1) * PageCount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			}
		}
	}
	return logs
}

func GetWininfoHisByDate(tablesuf string, begindate int64, enddate int64, pagecount int64, pagenum int64) []commonstruct.Cli_Wininfo {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_num_, tablesuf)

	// 中奖列表
	var WininfoS []commonstruct.Cli_Wininfo
	if err := WyMysql_ltdata.Table(tblname).Where("opentimestamp between ? and ?", commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate)).
		Order("expect desc").Limit(int(pagecount)).Offset((int(pagenum) - 1) * int(pagecount)).Find(&WininfoS).Error; err != nil {
		beego.Error("GetWininfoHisByDate %v date %v %v err %v", tablesuf, begindate, enddate, err)
		return WininfoS
	}
	return WininfoS
}

func GetWininfoHisPageinfo(tablesuf string, begindate int64, enddate int64, pagecount int64) (int, int) {
	var logs commonstruct.Cli_Wininfo
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_num_, tablesuf)

	if err := WyMysql_ltdata.Table(tblname).Select("Count(*) as opentimestamp").
		Where("opentimestamp between ? and ?", commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate)).
		Find(&logs).Error; err != nil {
		beego.Error("GetWininfoHisPageinfo err", tablesuf, err)
	}

	beego.Error("GetWininfoHisPageinfo === ", logs.Opentimestamp, tablesuf, begindate, enddate, pagecount)

	return int(logs.Opentimestamp), int(math.Ceil(float64(logs.Opentimestamp) / float64(pagecount)))
}

// func GetMoneyInoutPageinfo(uuids int64, begindate int64, enddate int64, reqtype int) (int64, int64) {
// 	// 查询订单记录
// 	var logs commonstruct.Withdraw

// 	switch reqtype {
// 	case -1: //全部查询
// 		if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
// 			Where("uuid = ? and req_time between ? and ?", uuids, begindate, enddate).
// 			Find(&logs).Error; err != nil {
// 			beego.Error("GetMoneyInoutPageinfo err", uuids, err)
// 		}
// 	default:
// 		if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
// 			Where("uuid = ? and req_time between ? and ? and req_type = ?", uuids, begindate, enddate, reqtype).
// 			Find(&logs).Error; err != nil {
// 			beego.Error("GetMoneyInoutPageinfo err", uuids, err)
// 		}
// 	}
// 	return int64(logs.OrderID), int64(math.Ceil(float64(logs.OrderID) / 20))
// }

func GetMoneyInoutPageinfo(uuids []int64, begindate int64, enddate int64, pagecount int, reqtype int, way int, state int) (int64, int64) {
	// 查询订单记录
	var logs commonstruct.Withdraw

	switch reqtype {
	case -1: //全部查询
		switch way {
		case -1:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ?", uuids, begindate, enddate).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and state = ?", uuids, begindate, enddate, state).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			}
		default:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and way = ?", uuids, begindate, enddate, way).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and way = ? and state = ?", uuids, begindate, enddate, way, state).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			}
		}
	default:
		switch way {
		case -1:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and req_type = ?", uuids, begindate, enddate, reqtype).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and state = ?", uuids, begindate, enddate, reqtype, state).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			}
		default:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and way = ?", uuids, begindate, enddate, reqtype, way).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and way = ? and state = ?", uuids, begindate, enddate, reqtype, way, state).
					Find(&logs).Error; err != nil {
					beego.Error("GetMoneyInoutPageinfo err", uuids, err)
				}
			}
		}
	}
	return logs.OrderID, int64(math.Ceil(float64(logs.OrderID) / float64(pagecount)))
}

/***********
* 获取用户组资金变动日志
 ***********/
func GetMoneyInoutLog(uuids []int64, begindate int64, enddate int64, offset int, pagecount int, reqtype int, way int, state int) []commonstruct.MoneyInOut {
	// 查询订单记录
	var logs []commonstruct.MoneyInOut
	//	var sqlArg string

	switch reqtype {
	case -1: //全部查询
		switch way {
		case -1:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ?", uuids, begindate, enddate).
					Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and state = ?", uuids, begindate, enddate, state).
					Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			}
		default:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and way = ?", uuids, begindate, enddate, way).
					Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and way = ? and state = ?", uuids, begindate, enddate, way, state).
					Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			}
		}
	default:
		switch way {
		case -1:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and req_type = ?", uuids, begindate, enddate, reqtype).
					Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and state = ?", uuids, begindate, enddate, reqtype, state).
					Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			}
		default:
			switch state {
			case -1:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and way = ?", uuids, begindate, enddate, reqtype, way).
					Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			default:
				if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
					Where("uuid in (?) and req_time between ? and ? and req_type = ? and way = ? and state = ?", uuids, begindate, enddate, reqtype, way, state).
					Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
					beego.Error("GetUnWithdraw err", uuids, err)
				}
			}
		}
	}
	return logs
}
func BKGetWininfoHisByDate(tablesuf string, offset int, begindate string, enddate string) []commonstruct.Cli_Wininfo {
	tblname := fmt.Sprintf("%s%s", commonstruct.WY_num_, tablesuf)

	timeMin := commonfunc.Date20060102ToBjTime(begindate) // time.Parse("20060102", begindate) //年月日
	timeEnd := commonfunc.Date20060102ToBjTime(enddate)   // time.Parse("20060102", enddate)       //年月日
	timeMax := timeEnd.AddDate(0, 0, 1)
	// 中奖列表
	var WininfoS []commonstruct.Cli_Wininfo
	if err := WyMysql.Table(tblname).Where("opentimestamp between ? and  ?", timeMin.Unix(), timeMax.Unix()).
		Order("opentimestamp desc").Limit(20).Offset((offset - 1) * 20).Find(&WininfoS).Error; err != nil {
		beego.Error("GetWininfoHisByDate %v date %v %v err %v", tablesuf, begindate, enddate, err)
		return WininfoS
	}
	return WininfoS
}

/***********
* 查询公司可用的第三方支付通道
 ***********/
func GetCoRechargeway(masterid int64) ([]commonstruct.CoRechargeWay, error) {
	var info []commonstruct.CoRechargeWay
	err := WyMysql.Table(commonstruct.WY_company_rechargeway).Where("company_id = ?", masterid).Order("sort").Find(&info).Error
	return info, err
}

/***********
* 设置第三方支付通道详情
 ***********/
func GetRechargewayInfo(wayid int64) (commonstruct.RechargeWayinfo, error) {
	var info commonstruct.RechargeWayinfo
	err := WyMysql.Table(commonstruct.WY_gm_config_recharge_way).Where("service_type = ?", wayid).Find(&info).Error
	return info, err
}

func GetWXUrl(code string) (commonstruct.WXUrl, error) {
	var info commonstruct.WXUrl
	err := WyMysql.Table(commonstruct.WY_gm_config_wxurl).Where("code = ?", code).Find(&info).Error
	return info, err
}

/***********
* 获取指定日期零点的钱包余额
 ***********/
func Get24hmoney(uuid int64, date int64) float64 {
	prelog := GetLatestLog(uuid, date)
	if prelog.ID > 0 {
		return prelog.NewGold
	}

	suflog := GetFristlog(uuid, date)
	if suflog.ID > 0 {
		return suflog.OldGold
	}

	usermoney := GetUsermoney(uuid)
	return usermoney.Cash
}

func GetLevelcfgs(uuid int64) ([]commonstruct.LevelConfig, error) {
	var info []commonstruct.LevelConfig
	err := WyMysql.Table(commonstruct.WY_tmp_company_level).Where("uuid = ?", uuid).Order("level").Find(&info).Error
	return info, err
}

/************************************
* 查询公司活动
************************************/
func QueryCompanypromotion(companyid int64) ([]commonstruct.CPromotionInfo, error) {

	var newinfo []commonstruct.CPromotionInfo
	err := WyMysql.Table(commonstruct.WY_company_promotion).Where("company_id = ?", companyid).Find(&newinfo).Error
	return newinfo, err
}

func GetGamelistByType(platform string, gametype string) ([]commonstruct.CRoomInfo, error) {
	var info []commonstruct.CRoomInfo
	err := WyMysql.Table(commonstruct.WY_gm_config_room).
		Where("game_dalei = ? and game_type = ?", platform, gametype).Find(&info).Error
	return info, err
}

func GetClosevalue(preid int64, roomid int64) int64 {
	var value commonstruct.Closevalue
	if retinfo := WyMysql.Table(commonstruct.WY_gm_config_closevalue).Select("value").Where("uuid = ? and roomid = ?", preid, roomid).Find(&value); retinfo.Error != nil {
		if !retinfo.RecordNotFound() {
			beego.Error("GetClosevalue err ", preid, roomid, retinfo.Error.Error())
		}
	}
	return value.Value
}

/**********************************
* 获取所有后台目录权限
**********************************/
func GetConfigbkset() ([]commonstruct.NavisetCfg, error) {
	var infos []commonstruct.NavisetCfg
	if err := WyMysql.Table(commonstruct.WY_gm_config_bkset).Find(&infos).Error; err != nil {
		beego.Error("GetConfigbkset err", err)
		return infos, err
	}
	return infos, nil
}

/***********
* 查询所有的第三方支付通道
 ***********/
func GetAllRechargeway() ([]commonstruct.RechargeWayinfo, error) {
	var info []commonstruct.RechargeWayinfo
	err := WyMysql.Table(commonstruct.WY_gm_config_recharge_way).Find(&info).Error
	return info, err
}

func GetRankinglist(column string, begindate int64, enddate int64) ([]commonstruct.DayStatistic, error) {
	var info []commonstruct.DayStatistic
	err := WyMysql.Table(commonstruct.WY_tmp_user_daystatistic).Select(fmt.Sprintf("uuid,sum(%v) as %v", column, column)).
		Where("date BETWEEN ? and ?", begindate, enddate).Group(column).Order(fmt.Sprintf("%v desc", column)).Limit(10).Find(&info).Error
	return info, err
}

/***********
* 获取公司入款通道
 ***********/
func GetCoSelfRechargeway(masterid int64) ([]commonstruct.CoSelfRechargeway, error) {
	var info []commonstruct.CoSelfRechargeway
	err := WyMysql.Table(commonstruct.WY_company_selfrechargeway).Where("company_id = ?", masterid).Order("sort").Find(&info).Error
	return info, err
}

/***********
* 设置公司入款通道状态
 ***********/
func SetCoSelfRechargewaystate(wayid int64, column string, value string) error {
	if err := WyMysql.Table(commonstruct.WY_company_selfrechargeway).
		Where("id = ? ", wayid).Update(column, value).Error; err != nil {
		beego.Error("SetCoSelfRechargewaystate err ", wayid, column, value, err)
		return err
	}
	return nil
}

/***********
* 设置公司的第三方支付通道状态
 ***********/
func SetCoRechargewaystate(companyid int64, wayid int64, value string) error {
	if err := WyMysql.Table(commonstruct.WY_company_rechargeway).
		Where("company_id = ? and way_id = ? ", companyid, wayid).Update("state", value).Error; err != nil {
		beego.Error("SetCoRechargewaystate err ", companyid, wayid, err)
		return err
	}
	return nil
}

/***********
* 设置公司的第三方支付通道排序
 ***********/
func SetCoRechargewaysort(companyid int64, value string) error {
	wayids := strings.Split(value, ",")
	if len(wayids) <= 0 {
		return nil
	}

	tx := WyMysql.Begin()
	for pos, wayidstr := range wayids {
		wayid, _ := strconv.Atoi(wayidstr)
		if err := tx.Table(commonstruct.WY_company_rechargeway).Where("company_id = ? and way_id = ? ", companyid, wayid).
			Update("sort", pos+1).Error; err != nil {
			beego.Error("SetCoRechargewaysort err ", err)
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("SetCoRechargewaysort commit err ", err)
		tx.Rollback()
		return err
	}
	return nil
}

// 修改代理的水赔比
func UpdateAgentodds(uuid int64, pan string, column string, value interface{}) error {
	var updateValues map[string]interface{}
	switch column {
	case "odds":
		updateValues = map[string]interface{}{
			"pre_odds": value,
		}
	// case "shuilottery":
	// 	updateValues = map[string]interface{}{
	// 		"shui_lottery": value,
	// 	}
	default:
		beego.Error("err column ", column)
		return nil
	}

	if err := WyMysql.Table(commonstruct.WY_user_odds).Where("uuid = ? and port = ?", uuid, pan).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsertAgentOdds err ", err)
		return err
	}

	return nil
}

// 设置房间开关
func SetRoomInUse(uuid int64, roomid int64, inuse int64) error {
	if err := WyMysql.Table(commonstruct.WY_company_game).Where("company_id = ? and room_id = ?", uuid, roomid).
		Update("in_use", inuse).Error; err != nil {
		beego.Error("SetRoomInUse err", err)
		return err
	}
	return nil
}

/************************************
* 修改公司VIP升级设置
************************************/
func UpdateLevelcfgs(uuid int64, level int64, column string, value int64) error {

	updateValues := map[string]interface{}{
		column: value,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_company_level).
		Where("uuid = ? and level = ?", uuid, level).Update(updateValues).Error; err != nil {
		beego.Error("UpdateCoquota err ", uuid, level, err)
		return err
	}
	return nil
}

// 修改公司的游戏设置
func UpdateCompanyGameset(uuid int64, roomid int64, InUse int64, sortid int64, Closevalue int64, TemaClosevalue int64, Openvalue int64, IsHot int64, IsChedan int64) error {
	updateValues := map[string]interface{}{
		"in_use":          InUse,
		"sort_id":         sortid,
		"closevalue":      Closevalue,
		"tema_closevalue": TemaClosevalue,
		"openvalue":       Openvalue,
		"is_hot":          IsHot,
		"is_chedan":       IsChedan,
	}

	if err := WyMysql.Table(commonstruct.WY_company_game).
		Where("company_id = ? and room_id = ? ", uuid, roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateCompanyGameset err ", uuid, roomid, err)
		return err
	}
	return nil
}

// 修改公司的自开设置
func UpdateCompanylimitinfo(masterid int64, uuid int64, account string, value int64) error {

	switch value {
	case 0: // 移除用户限制
		if err := WyMysql.Table(commonstruct.WY_company_limitinfo).Delete(nil, "uuid = ?", uuid).Error; err != nil {
			beego.Error("Delete Userlimit", commonstruct.WY_company_limitinfo, uuid, " err ", err)
			return err
		}
	default:

		var oldinfo commonstruct.LimitUserinfo

		if retinfo := WyMysql.Table(commonstruct.WY_company_limitinfo).Where("uuid = ? ", uuid).
			Find(&oldinfo); retinfo.Error != nil {
			if retinfo.RecordNotFound() {
				var limitinfo commonstruct.LimitUserinfo
				limitinfo.CompanyID = masterid
				limitinfo.Uuid = uuid
				limitinfo.Account = account
				limitinfo.LimitFlag = value
				limitinfo.LimitTime = commonfunc.GetNowtime()
				if err := WyMysql.Table(commonstruct.WY_company_limitinfo).Create(&limitinfo).Error; err != nil {
					beego.Error("new limitinfo err %v", limitinfo, err)
				}
			} else {
				beego.Error("new limitinfo err ", uuid, retinfo.Error.Error())
				return retinfo.Error
			}
			return nil
		}

		updateValues := map[string]interface{}{
			"limit_flag": value,
			"limit_time": commonfunc.GetNowtime(),
		}
		if err := WyMysql.Table(commonstruct.WY_company_limitinfo).Where("uuid = ?", uuid).
			Update(updateValues).Error; err != nil {
			beego.Error("update limitinfo err ", err)
			return err
		}
	}
	return nil
}

func DeleteCompanygame(masterid int64, roomid int64) error {
	cogamesql := fmt.Sprintf("delete from %v where company_id = '%v' and room_id = '%v';", commonstruct.WY_company_game,
		masterid, roomid)
	if err := WyMysql.Exec(cogamesql).Error; err != nil {
		beego.Error("DeleteCompanygame err ", cogamesql, err)
		return err
	}

	cooddssql := fmt.Sprintf("delete from %v where company_id = '%v' and room_id = '%v';", commonstruct.WY_company_portclass,
		masterid, roomid)
	if err := WyMysql.Exec(cooddssql).Error; err != nil {
		beego.Error("DeleteCompanygame err ", cooddssql, err)
		return err
	}

	return nil
}

func DeleteCompanygameset(masterid int64, roomid int64) error {
	cogamesql := fmt.Sprintf("delete from %v where company_id = '%v' and room_id = '%v';", commonstruct.WY_company_game,
		masterid, roomid)
	if err := WyMysql.Exec(cogamesql).Error; err != nil {
		beego.Error("DeleteCompanygame err ", cogamesql, err)
		return err
	}
	return nil
}

//
func UpdateSufCompanygame(masterid int64, roomid int64, inuse int64) error {
	updateValues := map[string]interface{}{
		"pre_kaiguan": inuse,
	}

	if err := WyMysql.Table(commonstruct.WY_company_game).
		Where("company_id = ? and room_id = ? ", masterid, roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateCompanygame err ", masterid, roomid, err)
	}

	return nil
}

// 创建新的游戏房间
func Createroom(roominfo commonstruct.CRoomInfo) (commonstruct.CRoomInfo, error) {

	newroomdb := WyMysql.Table(commonstruct.WY_gm_config_room).Create(&roominfo)
	if err := newroomdb.Error; err != nil {
		beego.Error("Createroom err ", err)
		return roominfo, errors.New("创建role base表失败")
	}

	newroominfo := newroomdb.Value.(*commonstruct.CRoomInfo) // 新注册用户ID
	return *newroominfo, nil
}

func InitDefaultodds(roomid int64, portinfo commonstruct.PortClass) error {

	var oldOdds commonstruct.DefaultOdds

	if retinfo := WyMysql.Table(commonstruct.WY_gm_config_defaultodds).
		Where("room_id = ? and port_id = ?", roomid, portinfo.ID).
		Find(&oldOdds); retinfo.Error != nil {

		if retinfo.RecordNotFound() {
			oldOdds.RoomID = roomid
			oldOdds.PortID = portinfo.ID
			oldOdds.DefaultOdds = portinfo.DefaultOdds

			if err := WyMysql.Table(commonstruct.WY_gm_config_defaultodds).Create(&oldOdds).Error; err != nil {
				beego.Error("InitDefaultodds err", oldOdds, err)
				return err
			}
		} else {
			beego.Error("InitDefaultodds err ", oldOdds, retinfo.Error.Error())
			return retinfo.Error
		}
	}
	return nil
}

func DeleteDefaultodds(roomid int64) error {
	err := WyMysql.Table(commonstruct.WY_gm_config_defaultodds).Delete(nil, "room_id = ?", roomid).Error
	if err != nil {
		beego.Error("DeleteDefaultodds ", commonstruct.WY_company_limitinfo, roomid, " err ", err)
	}
	return err
}

func UpdateUsergameset(info commonstruct.UserGameset) error {

	updateValues := map[string]interface{}{
		// "zhancheng":              info.Zhancheng,
		// "suf_min_zhancheng":      info.SufMinZhancheng,
		// "suf_max_zhancheng":      info.SufMaxZhancheng,
		"zhancheng_next":         info.ZhanchengNext,
		"suf_min_zhancheng_next": info.SufMinZhanchengNext,
		"suf_max_zhancheng_next": info.SufMaxZhanchengNext,
		"pans":                   info.Pans,
		"buhuo_flag":             info.BuhuoFlag,
		"in_use":                 info.InUse,
	}

	if err := WyMysql.Table(commonstruct.WY_user_gameset).Where("uuid = ? and room_id = ?", info.Uuid, info.RoomID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateUsergameset err ", err)
		return err
	}
	return nil
}

// 修改用户的单列游戏设置
func UpdateUsergamesetCloumn(uuid int64, roomid int64, column string, newvalue float64) error {

	var updateValues map[string]interface{}
	switch column {
	case "zhancheng":
		updateValues = map[string]interface{}{
			"zhancheng": newvalue,
		}
	case "zhancheng_next":
		updateValues = map[string]interface{}{
			"zhancheng_next": newvalue,
		}
	case "suf_min_zhancheng":
		updateValues = map[string]interface{}{
			"suf_min_zhancheng": newvalue,
		}
	case "suf_min_zhancheng_next":
		updateValues = map[string]interface{}{
			"suf_min_zhancheng_next": newvalue,
		}
	case "suf_max_zhancheng":
		updateValues = map[string]interface{}{
			"suf_max_zhancheng": newvalue,
		}
	case "suf_max_zhancheng_next":
		updateValues = map[string]interface{}{
			"suf_max_zhancheng_next": newvalue,
		}
	default:
		beego.Error("UpdateUsergamesetCloumn err column", column)
		return errors.New(fmt.Sprintf("UpdateUsergamesetCloumn err column [%v]", column))
	}

	// beego.Error("UpdateUsergamesetCloumn", uuid, roomid, updateValues)
	if err := WyMysql.Table(commonstruct.WY_user_gameset).Where("uuid = ? and room_id = ?", uuid, roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateUsergamesetCloumn err ", err)
		return err
	}
	return nil
}

// 关闭掉整个团队的补货
func UpdateTeambuhuoflag(teamidlist []int64, roomid int64) error {

	updateValues := map[string]interface{}{
		"buhuo_flag": 0,
	}

	if err := WyMysql.Table(commonstruct.WY_user_gameset).Where("uuid in (?) and room_id = ?", teamidlist, roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateTeambuhuoflag err ", err)
		return err
	}
	return nil
}

// 修改用户的权限列表
func UpsertUsernavi(info commonstruct.UserNavi) error {
	numtblname := fmt.Sprintf("%v%v", commonstruct.WY_user_navi_, info.Uuid/500)

	var oldinfo commonstruct.UserNavi
	if retinfo := WyMysql.Table(numtblname).
		Where("uuid = ? and navi_id = ?", info.Uuid, info.NaviID).
		Find(&oldinfo); retinfo.Error != nil {
		if retinfo.RecordNotFound() {

			if err := WyMysql.Table(numtblname).Create(&info).Error; err != nil {
				beego.Error("create err ", oldinfo, err)
				return err
			}
		} else {
			beego.Error("UpsertCompanyodds err ", info, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	updateValues := map[string]interface{}{
		"value": info.Value,
	}

	if err := WyMysql.Table(numtblname).Where("uuid = ? and navi_id = ?", info.Uuid, info.NaviID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateUsernavi err ", err)
		return err
	}

	return nil
}

// 关闭掉团队权限
func CloseUsersnavi(teamidlist []int64, naviid int64) error {
	for _, uuid := range teamidlist {
		numtblname := fmt.Sprintf("%v%v", commonstruct.WY_user_navi_, uuid/500)

		updateValues := map[string]interface{}{
			"value": 0,
		}

		if err := WyMysql.Table(numtblname).Where("uuid = ? and navi_id = ?", uuid, naviid).
			Update(updateValues).Error; err != nil {

		}
	}
	return nil
}

func UpdateSysnaviset(info commonstruct.BKNavi) error {
	updateValues := map[string]interface{}{
		"admin":   info.Admin,
		"saler":   info.Saler,
		"company": info.Company,
		"agent":   info.Agent,
		"player":  info.Player,
	}

	if err := WyMysql.Table(commonstruct.WY_gm_config_bknavi).Where("id = ?", info.ID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateSysnaviset err ", err)
		return err
	}
	return nil
}

func UpdatePortclass(portid int64, desodds float64) error {
	updateValues := map[string]interface{}{
		"default_odds": gorm.Expr("default_odds + ?", desodds),
	}
	if err := WyMysql.Table(commonstruct.WY_gm_config_portclass).
		Where("id = ?", portid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdatePortclass err \n", portid, desodds, err)
	}
	return nil
}

func UpsertSalerbaseodds(uuid int64, portid int64, desodds float64) error {

	var Dyncinfo commonstruct.SalerBaseodds
	if retinfo := WyMysql.Table(commonstruct.WY_user_settletype).
		Find(&Dyncinfo, "uuid = ? and portid = ?", uuid, portid); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			Dyncinfo.Uuid = uuid
			portclassinfo, _ := GetPortclassByID(portid)
			Dyncinfo.SettleType = portclassinfo.SettleType
			Dyncinfo.LotteryDalei = portclassinfo.LotteryDalei
			Dyncinfo.GameDalei = portclassinfo.GameDalei
			Dyncinfo.GameXiaolei = portclassinfo.GameXiaolei

			Dyncinfo.Portid = portid
			Dyncinfo.DesOdds = desodds
			if err := WyMysql.Table(commonstruct.WY_user_settletype).Create(&Dyncinfo).Error; err != nil {
				beego.Error("UpsertSalerbaseodds err ", Dyncinfo, err)
				return nil
			}
		} else {
			beego.Error("UpsertSalerbaseodds err \n", uuid, portid, desodds, retinfo.Error.Error())
			return retinfo.Error
		}
	}

	updateValues := map[string]interface{}{
		"des_odds": desodds,
	}
	if err := WyMysql.Table(commonstruct.WY_user_settletype).
		Where("uuid = ? and portid = ?", uuid, portid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateSalerbaseodds err \n", uuid, portid, desodds, err)
	}
	return nil
}

func UpdateGamedescription(roomid int64, description string) error {
	updateValues := map[string]interface{}{
		"description": description,
	}
	if err := WyMysql.Table(commonstruct.WY_gm_config_room).
		Where("id = ?", roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateGamedescription err \n", roomid, description, err)
	}
	return nil
}

func UpdateGamevalidtime(roomid int64, invalidtime int64) error {
	updateValues := map[string]interface{}{
		"in_valid_time": invalidtime,
	}
	if err := WyMysql.Table(commonstruct.WY_gm_config_room).
		Where("id = ?", roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateGamevalidtime err \n", roomid, invalidtime, err)
	}
	return nil
}

func AddLoginInfo(uuid int64, account string, ip string, ipplace string, platform string, reqhost string, useragent string) error {
	var info commonstruct.CLoginInfo
	info.Uuid = uuid
	info.Account = account
	info.LoginTime = commonfunc.GetNowtime()
	if strings.Contains(ipplace, "邻水县") {
		info.IP = "47.242.144.152"
		info.IPPlace = "中国香港阿里云"
	} else if strings.Contains(ipplace, "新乡市") {
		info.IP = "47.242.144.152"
		info.IPPlace = "中国香港阿里云"
	} else {
		info.IP = ip
		info.IPPlace = ipplace
	}
	info.Platform = platform
	info.ReqHost = reqhost
	info.Useragent = useragent

	err := WyMysql.Table(commonstruct.WY_tmp_user_logininfo).Create(&info).Error
	if err != nil {
		beego.Error("AddLoginInfo err %v", err)
	}
	return err
}

func GetUserLogininfosPageinfo(destuuid int64, pagecount int64, begintime int64, endtime int64) (int64, int64) {
	var infos commonstruct.CLoginInfo
	if err := WyMysql.Table(commonstruct.WY_tmp_user_logininfo).Select("count(*) as id").
		Where("uuid = ? and login_time between ? and ?", destuuid, begintime, endtime).
		Find(&infos).Error; err != nil {
		beego.Error("GetUserLogininfosPageinfo err", destuuid, begintime, endtime, err)
	}
	return int64(infos.ID), int64(math.Ceil(float64(infos.ID) / float64(pagecount)))
}

func GetUserLogininfoS(uuid int64, begintime int64, endtime int64, pagecount int64, pagenum int64) ([]commonstruct.CLoginInfo, error) {
	var infos []commonstruct.CLoginInfo
	err := WyMysql.Table(commonstruct.WY_tmp_user_logininfo).
		Where("uuid = ? and login_time between ? and ?", uuid, begintime, endtime).
		Order("id desc").Limit(int(pagecount)).Offset(int(pagenum-1) * int(pagecount)).Find(&infos).Error
	if err != nil {
		beego.Error("err %v", err)
	}
	return infos, err
}

func GetUserLoginips(uuid int64) ([]commonstruct.CLoginInfo, error) {
	var infos []commonstruct.CLoginInfo
	err := WyMysql.Table(commonstruct.WY_tmp_user_logininfo).Select("distinct ip").
		Where("uuid = ?", uuid).Find(&infos).Error
	if err != nil {
		beego.Error("err %v", err)
	}
	return infos, err
}

func CheckUsergameset(uuid int64) error {

	if roominfos, _ := GetRoominfoS(); len(roominfos) == 0 {
		beego.Error("GetRoominfoS err")
	} else {
		for _, roominfo := range roominfos {
			var oldinfo commonstruct.UserGameset
			if retinfo := WyMysql.Table(commonstruct.WY_user_gameset).
				Where("uuid = ? and room_id = ?", uuid, roominfo.ID).
				Find(&oldinfo); retinfo.Error != nil {
				if retinfo.RecordNotFound() {

					oldinfo.Uuid = uuid
					oldinfo.Platform = roominfo.GameDalei
					oldinfo.RoomID = roominfo.ID
					if roominfo.InValidTime > commonfunc.GetNowtime() {
						oldinfo.InUse = 1
					} else {
						oldinfo.InUse = 0
					}

					oldinfo.Zhancheng = 0
					oldinfo.SufMinZhancheng = 100
					oldinfo.SufMaxZhancheng = 100
					oldinfo.Pans = "A,B,C,D"

					if err := WyMysql.Table(commonstruct.WY_user_gameset).Create(&oldinfo).Error; err != nil {
						beego.Error("create err ", oldinfo, err)
						return err
					}
				} else {
					beego.Error("CheckUsergameset err ", uuid, retinfo.Error.Error())
					return retinfo.Error
				}
			}
		}
	}
	return nil
}

// 获取不想被客户看到的IP列表
func GetElideipS() ([]commonstruct.IPInfo, error) {
	var infos []commonstruct.IPInfo
	err := WyMysql.Table(commonstruct.WY_gm_config_ip_elide).Find(&infos).Error
	if err != nil {
		beego.Error("GetElideipS err ", err)
		return nil, err
	}
	return infos, nil
}

func GetRoominfo(gameres string) commonstruct.CRoomInfo {
	var roominfo commonstruct.CRoomInfo
	if err := WyMysql.Table(commonstruct.WY_gm_config_room).
		Where("name_eng = ?", gameres).Find(&roominfo).Error; err != nil {

		beego.Error("GetRoomgame err ", err)
	}

	return roominfo
}

// 获取下注项的实时变赔
func GetDynimicOdds(masterid int64, roomid int64, expect string, port string, itemid int64) (commonstruct.DynimicOddsInfo, error) {
	var DynimicOddsInfo commonstruct.DynimicOddsInfo
	err := WyMysql.Table(commonstruct.WY_tmp_user_dyncitem).
		Where("uuid = ? and room_id = ? and expect = ? and port = ? and item_id = ?", masterid, roomid, expect, port, itemid).
		Find(&DynimicOddsInfo).Error
	return DynimicOddsInfo, err
}

// 获取用户的玩法设置
func GetUserPortclassset(uuid int64, roomid int64, portid int64) (commonstruct.UserPortclassset, error) {
	userinfo, _ := GetUserinfoByUuid(uuid)

	var info commonstruct.UserPortclassset
	switch userinfo.RoleType {
	case "company":
		var coset commonstruct.CompanyPortinfo
		err := WyMysql.Table(commonstruct.WY_company_portclass).
			Where("company_id = ? and room_id = ? and port_id = ?", uuid, roomid, portid).
			Find(&coset).Error
		info.Uuid = coset.CompanyID
		info.RoomID = coset.RoomID
		info.PortID = coset.PortID
		info.IsTransfer = coset.IsTransfer
		info.TuishuiA = coset.TuishuiA
		info.TuishuiB = coset.TuishuiB
		info.TuishuiC = coset.TuishuiC
		info.TuishuiD = coset.TuishuiD
		info.MinAmount = coset.MinAmount
		info.MaxAmount = coset.MaxAmount
		info.MaxAmountExpect = coset.MaxAmountExpect
		info.TransferAmount = coset.TransferAmount
		info.WarningAmount = int64(coset.WarningAmount)
		return info, err
	default:
		tblname := GetUserPortclasssetTablename(uuid)
		err := WyMysql.Table(tblname).
			Where("uuid = ? and room_id = ? and port_id = ?", uuid, roomid, portid).
			Find(&info).Error
		return info, err
	}
}

// 获取游戏某期的全部变赔
func GetDynimicoddsList(masterid int64, roomid int64, expect string, port string) ([]commonstruct.DynimicOddsInfo, error) {
	var oddsinfo []commonstruct.DynimicOddsInfo
	err := WyMysql.Table(commonstruct.WY_tmp_user_dyncitem).
		Where("uuid = ? and room_id = ? and expect = ? and port = ?", masterid, roomid, expect, port).Find(&oddsinfo).Error
	if err != nil {

	}
	return oddsinfo, err
}

func UpsertDynimicOdds(uuid int64, roomid int64, expect string, itemid int64, changerate float64, validtime int64) error {
	var Dyncinfo commonstruct.DynimicOddsInfo
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_user_dyncitem).
		Find(&Dyncinfo, "uuid = ? and room_id = ? and expect = ? and item_id = ?", uuid, roomid, expect, itemid); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			Dyncinfo.Uuid = uuid
			Dyncinfo.RoomID = roomid
			Dyncinfo.Expect = expect
			Dyncinfo.ItemID = itemid
			Dyncinfo.DynimicOdds = changerate
			Dyncinfo.ValidTime = validtime
			if err := WyMysql.Table(commonstruct.WY_tmp_user_dyncitem).Create(&Dyncinfo).Error; err != nil {
				beego.Error("UpsertDynimicOdds err ", Dyncinfo, err)
				return nil
			}
		} else {
			beego.Error("UpsertDynimicOdds err \n", uuid, roomid, expect, itemid, changerate, retinfo.Error.Error())
			return retinfo.Error
		}
	}

	updateValues := map[string]interface{}{
		"dynimic_odds": changerate,
		"valid_time":   validtime,
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_user_dyncitem).
		Where("uuid = ? and room_id = ? and expect = ? and item_id = ?", uuid, roomid, expect, itemid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsertDynimicOdds err", uuid, roomid, expect, itemid, changerate, err)
	}
	return nil
}

// 更新长龙变赔信息
func UpdateDyncchanglong(uuid int64, roomid int64, dyncinfo string) error {
	var oldDync commonstruct.DyncChanglong
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_user_dyncchanglong).Where("uuid = ? and room_id = ?", uuid, roomid).Find(&oldDync); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			oldDync.Uuid = uuid
			oldDync.RoomID = roomid
			oldDync.DyncInfo = dyncinfo

			if err := WyMysql.Table(commonstruct.WY_tmp_user_dyncchanglong).Create(&oldDync).Error; err != nil {
				beego.Error("new Dyncchanglong err ", uuid, roomid, err)
				return err
			}
			return nil
		} else {
			beego.Error("UpdateDyncchanglong err ", retinfo.Error.Error())
			return retinfo.Error
		}
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_dyncchanglong).Where("uuid = ? and room_id = ?", uuid, roomid).Update("dync_info", dyncinfo).Error; err != nil {
		beego.Error("UpdateDyncchanglong err ", err)
		return err
	}
	return nil
}

// 更新货量变赔信息
func UpdateDyncamount(uuid int64, roomid int64, dyncinfo string) error {

	var oldDync commonstruct.DyncAmount
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_user_dyncamount).Where("uuid = ? and room_id = ?", uuid, roomid).Find(&oldDync); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			oldDync.Uuid = uuid
			oldDync.RoomID = roomid
			oldDync.DyncInfo = dyncinfo

			if err := WyMysql.Table(commonstruct.WY_tmp_user_dyncamount).Create(&oldDync).Error; err != nil {
				beego.Error("new dyncamount err ", uuid, roomid, err)
				return err
			}
			return nil
		} else {
			beego.Error("UpdateDyncamount err ", retinfo.Error.Error())
			return retinfo.Error
		}
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_dyncamount).Where("uuid = ? and room_id = ?", uuid, roomid).Update("dync_info", dyncinfo).Error; err != nil {
		beego.Error("UpdateDyncamount err ", err)
		return err
	}
	return nil
}

// 更新封盘时间
func UpdateClosetime(uuid int64, roomid int64, closetime int64) error {

	var oldTime commonstruct.Closevalue
	if retinfo := WyMysql.Table(commonstruct.WY_gm_config_closevalue).Where("uuid = ? and roomid = ?", uuid, roomid).Find(&oldTime); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			oldTime.Uuid = uuid
			oldTime.Roomid = roomid
			oldTime.Value = closetime

			if err := WyMysql.Table(commonstruct.WY_gm_config_closevalue).Create(&oldTime).Error; err != nil {
				beego.Error("new Closertime err ", uuid, roomid, err)
				return err
			}
			return nil
		} else {
			beego.Error("UpdateClosetime err ", retinfo.Error.Error())
			return retinfo.Error

		}
	}

	if err := WyMysql.Table(commonstruct.WY_gm_config_closevalue).Where("uuid = ? and roomid = ?", uuid, roomid).Update("value", closetime).Error; err != nil {
		beego.Error("UpdateClosetime err ", err)
		return err
	}
	return nil
}

// 更新货量走飞信息
func UpdateOrdertransfer(uuid int64, roomid int64, transferinfo string) error {

	var oldDync commonstruct.Transfer
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_user_ordertransfer).Where("uuid = ? and room_id = ?", uuid, roomid).Find(&oldDync); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			oldDync.Uuid = uuid
			oldDync.RoomID = roomid
			oldDync.TransferInfo = transferinfo

			if err := WyMysql.Table(commonstruct.WY_tmp_user_ordertransfer).Create(&oldDync).Error; err != nil {
				beego.Error("new Ordertransfer err ", uuid, roomid, err)
				return err
			}
			return nil
		} else {
			beego.Error("UpdateOrdertransfer err ", retinfo.Error.Error())
			return retinfo.Error
		}
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_ordertransfer).Where("uuid = ? and room_id = ?", uuid, roomid).Update("transfer_info", transferinfo).Error; err != nil {
		beego.Error("UpdateOrdertransfer err ", err)
		return err
	}
	return nil
}

/***********
* 给客户写信
 ***********/
func AddFeedbackmsg(uuid int64, srcid int64, Title string, Detail string) error {
	var info commonstruct.FeedbackMsg
	info.Time = commonfunc.GetNowtime()
	info.Uuid = uuid
	info.SrcID = srcid
	info.Title = Title
	info.Detail = Detail
	info.Flag = 0

	if err := WyMysql.Table(commonstruct.WY_tmp_user_retmsg).Create(&info).Error; err != nil {
		beego.Error("create err ", info, err)
		return err
	}
	return nil
}

func UpdateUnRecharge(orderid int64, op_result int64, state int, oldgold float64, newgold float64, expinfo string) error {
	updateValues1 := map[string]interface{}{
		"checked":   1,
		"op_result": op_result,
		"expinfo":   gorm.Expr("CONCAT_WS('|',expinfo,?)", expinfo),
		"res_time":  commonfunc.GetNowtime(),
		"state":     state,
		"old_gold":  oldgold,
		"new_gold":  newgold,
	}
	if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Where("order_id = ?", orderid).Update(updateValues1).Error; err != nil {
		beego.Error("UpdateUnRecharge err", err)
		return err
	}
	return nil
}

/***********
* 获取用户组注单统计
 ***********/
func GetProfitResult(uuids []int64, gametype string, begindate int64, enddate int64) []commonstruct.LotdataStatistic {
	var tblname string
	switch gametype {
	case "wylottery":
		tblname = "wy_gm_profit_wylottery"
	case "wyactual":
		tblname = "WY_tmp_profit_wyactual"
	case "bbinsport":
		tblname = "wy_exp_profit_bbinsport"
	case "bbinslots":
		tblname = "wy_exp_profit_bbinslots"
	case "bbinlive":
		tblname = "wy_exp_profit_bbinlive"
	case "sabasport":
		tblname = "wy_exp_profit_sabasport"
	case "sabaslots":
		tblname = "wy_exp_profit_sabaslots"
	case "sabalive":
		tblname = "wy_exp_profit_sabalive"
	}

	var ret []commonstruct.LotdataStatistic
	if err := WyMysql.Table(tblname).Where("uuid in (?) and date between ? and ?", uuids, begindate, enddate).Order("date,uuid").
		Find(&ret).Error; err != nil {
		beego.Error("GetProfitResult err ", err)
	}
	return ret
}

// /***********
// * 查询单个用户的注单统计(按日分割)
//  ***********/
// func SaveRetshui(settleResult []commonstruct.LotdataStatistic) bool {
// 	tx := WyMysql.Begin()

// 	for _, v := range settleResult {
// 		var Money commonstruct.UserMoney
// 		shui := v.ProfitShui + v.ProfitShuiDiff
// 		if _, err := TXUpdateBalance(tx, v.Uuid, "wylottery", shui); err != nil {
// 			beego.Error("Update Usershui err", v.Uuid, shui)
// 			tx.Rollback()
// 			return false
// 		} else {
// 			if shui != 0 {
// 				if err := TXAddUserMoneylog(tx, v.Uuid, commonstruct.UptType_Gongzi, shui, Money.Cash-shui, Money.Cash, fmt.Sprintf("用户%v工资", v.Date)); err != nil {
// 					tx.Rollback()
// 					return false
// 				}
// 			}
// 		}
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		beego.Error("SaveRetshui commit err ", err)
// 		tx.Rollback()
// 		return false
// 	}
// 	return true
// }

func TXAddUserMoneylog(tx *gorm.DB, uuid int64, walletname string, optype commonstruct.MoneyUpdateType, OpGold float64, OldGold float64, NewGold float64, Desc string) error {
	var uptlog commonstruct.MoneyUpdateLog
	uptlog.Uuid = uuid
	uptlog.Time = commonfunc.GetNowtime()
	uptlog.WalletName = walletname
	uptlog.OpType = optype
	uptlog.OpGold = OpGold
	uptlog.NewGold = NewGold
	uptlog.OldGold = OldGold
	uptlog.Opinfo = Desc
	if err := tx.Table(commonstruct.WY_tmp_log_money).Create(&uptlog).Error; err != nil {
		beego.Error("AddUserMoneylog err ", uptlog, err)
		return err
	}
	return nil
}

/***********
* 新建公司入款通道
 ***********/
func CreateSelfRechargeway(newway commonstruct.CoSelfRechargeway) error {
	if err := WyMysql.Table(commonstruct.WY_company_selfrechargeway).Create(&newway).Error; err != nil {
		beego.Error("create err ", newway, err)
		return err
	}
	return nil
}

/***********
* 设置公司入款通道排序
 ***********/
func SetCoSelfRechargewaysort(value string) error {
	wayids := strings.Split(value, ",")
	if len(wayids) <= 0 {
		return nil
	}

	tx := WyMysql.Begin()
	for pos, wayidstr := range wayids {
		wayid, _ := strconv.Atoi(wayidstr)
		if err := tx.Table(commonstruct.WY_company_selfrechargeway).Where("id = ? ", wayid).
			Update("sort", pos+1).Error; err != nil {
			beego.Error("SetCoSelfRechargewaysort err ", err)
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("SetCoSelfRechargewaysort commit err ", err)
		tx.Rollback()
		return err
	}
	return nil
}

func UpsertCompanyrechargeway(info commonstruct.CoRechargeWay) error {
	if err := WyMysql.Table(commonstruct.WY_company_rechargeway).Create(&info).Error; err != nil {
		beego.Error("UpsertCompanyrechargeway err", info, err)
		return err
	}
	return nil
}

// 系统关闭游戏设置
func DeleteCompanyPortclassS(uuid int64, roomid int64) error {

	if err := WyMysql.Table(commonstruct.WY_company_portclass).Delete(nil, "company_id = ? and room_id = ?", uuid, roomid).Error; err != nil {
		beego.Error("DeleteCompanyPortclassS err ", uuid, roomid, err)
		return err
	}
	return nil
}

//
func UpsertCompanyPortclass(portinfo commonstruct.CompanyPortinfo) error {
	var oldinfo commonstruct.CompanyPortinfo
	if retinfo := WyMysql.Table(commonstruct.WY_company_portclass).
		Where("company_id = ? and room_id = ? and port_id = ?", portinfo.CompanyID, portinfo.RoomID, portinfo.PortID).
		Find(&oldinfo); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_company_portclass).Create(&portinfo).Error; err != nil {
				beego.Error("create err ", oldinfo, err)
				return err
			}
		} else {
			beego.Error("UpsertCompanyodds err ", portinfo, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	updateValues := map[string]interface{}{
		"default_odds":      Round4s5r(portinfo.DefaultOdds, 4),
		"odds_a":            Round4s5r(portinfo.OddsA, 4),
		"tuishui_a":         Round4s5r(portinfo.TuishuiA, 4),
		"odds_b":            Round4s5r(portinfo.OddsB, 4),
		"tuishui_b":         Round4s5r(portinfo.TuishuiB, 4),
		"odds_c":            Round4s5r(portinfo.OddsC, 4),
		"tuishui_c":         Round4s5r(portinfo.TuishuiC, 4),
		"odds_d":            Round4s5r(portinfo.OddsD, 4),
		"tuishui_d":         Round4s5r(portinfo.TuishuiD, 4),
		"kaiguan":           1,
		"min_amount":        portinfo.MinAmount,
		"max_amount":        portinfo.MaxAmount,
		"max_amount_expect": portinfo.MaxAmountExpect,
		"warning_amount":    portinfo.WarningAmount,
		"port_switch":       portinfo.PortSwitch,
	}

	if err := WyMysql.Table(commonstruct.WY_company_portclass).
		Where("company_id = ? and room_id = ? and port_id = ?", portinfo.CompanyID, portinfo.RoomID, portinfo.PortID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsertCompanyodds err ", portinfo.CompanyID, portinfo.PortID, err)
	}

	return nil
}

// 修改插入 公司加时配置
func UpsertCompanyJiashiPortclass(portinfo commonstruct.CompanyJiashiPortinfo) error {
	var oldinfo commonstruct.CompanyPortinfo
	if retinfo := WyMysql.Table(commonstruct.WY_company_portclass_jiashi).
		Where("company_id = ? and room_id = ? and port_id = ?", portinfo.CompanyID, portinfo.RoomID, portinfo.PortID).
		Find(&oldinfo); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_company_portclass_jiashi).Create(&portinfo).Error; err != nil {
				beego.Error("create err ", oldinfo, err)
				return err
			}
		} else {
			beego.Error("UpsertCompanyJiashiPortclass err ", portinfo, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	updateValues := map[string]interface{}{
		"odds_des_a":        Round4s5r(portinfo.OddsDesA, 4),
		"odds_des_b":        Round4s5r(portinfo.OddsDesB, 4),
		"odds_des_c":        Round4s5r(portinfo.OddsDesC, 4),
		"odds_des_d":        Round4s5r(portinfo.OddsDesD, 4),
		"min_amount":        portinfo.MinAmount,
		"max_amount":        portinfo.MaxAmount,
		"max_amount_expect": portinfo.MaxAmountExpect,
	}

	if err := WyMysql.Table(commonstruct.WY_company_portclass_jiashi).
		Where("company_id = ? and room_id = ? and port_id = ?", portinfo.CompanyID, portinfo.RoomID, portinfo.PortID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsertCompanyJiashiPortclass err ", portinfo.CompanyID, portinfo.PortID, err)
	}

	return nil
}

func UpdateCompanyPortclassColumn(portinfo commonstruct.CompanyPortinfo, columntype string) error {

	var updateValues map[string]interface{}
	switch columntype {
	case "transferamount":
		updateValues = map[string]interface{}{
			"transfer_amount": portinfo.TransferAmount,
			"is_transfer":     portinfo.IsTransfer,
		}
	case "warningamount":
		updateValues = map[string]interface{}{
			"warning_amount":     portinfo.WarningAmount,
			"warning_loopamount": portinfo.WarningLoopamount,
		}
	}

	if err := WyMysql.Table(commonstruct.WY_company_portclass).
		Where("company_id = ? and room_id = ? and port_id = ?", portinfo.CompanyID, portinfo.RoomID, portinfo.PortID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateCompanyTransferset err ", portinfo.CompanyID, portinfo.PortID, err)
	}

	return nil
}

// func UpdateCompanyWarningset(portinfo commonstruct.CompanyPortinfo) error {

// 	if err := WyMysql.Table(commonstruct.WY_company_portclass).
// 		Where("company_id = ? and room_id = ? and port_id = ?", portinfo.CompanyID, portinfo.RoomID, portinfo.PortID).
// 		Update(updateValues).Error; err != nil {
// 		beego.Error("UpdateCompanyWarningset err ", portinfo.CompanyID, portinfo.PortID, err)
// 	}

// 	return nil
// }

func GetActiveuuids(begindate int64, enddate int64) ([]commonstruct.DayStatistic, error) {
	var info []commonstruct.DayStatistic
	err := WyMysql.Table(commonstruct.WY_tmp_user_daystatistic).Select("distinct uuid").
		Where("date BETWEEN ? and ?", begindate, enddate).Find(&info).Error
	return info, err
}

func GetUserDatastats(uuid int64, column string, begindate int64, enddate int64) (commonstruct.DayStatistic, error) {
	var info commonstruct.DayStatistic
	err := WyMysql.Table(commonstruct.WY_tmp_user_daystatistic).Select(fmt.Sprintf("sum(%v) as %v", column, column)).
		Where("date BETWEEN ? and ?", begindate, enddate).Find(&info).Error
	return info, err
}

func GetLevelcfg(uuid int64, level int64) (commonstruct.LevelConfig, error) {
	var info commonstruct.LevelConfig
	err := WyMysql.Table(commonstruct.WY_gm_config_level).Where("uuid = ? and level = ?", uuid, level).Find(&info).Error
	return info, err
}

/**************************************
* 获取团队未提现金额
 *************************************/
func GetDateUnwithdraw(masterid int64, date int64) (commonstruct.MoneyInOut, error) {

	timeMin := commonfunc.GetBegintime(date)
	selectarg := fmt.Sprintf("master_id = %v and req_type = 6602 and req_time < %v and (res_time > %v or res_time = 0)",
		masterid, timeMin, timeMin)

	// 查询订单记录
	var logs commonstruct.MoneyInOut
	err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("sum(amount) as amount").Where(selectarg).Find(&logs).Error

	return logs, err
}

func GetTeamDaystatis(uuids []int64, date int64, column string) commonstruct.DayStatistic {
	var info commonstruct.DayStatistic
	selectinfo := fmt.Sprintf(`sum(%v) as %v`, column, column)
	if err := WyMysql.Table(commonstruct.WY_tmp_user_daystatistic).Select(selectinfo).
		Where("uuid in (?) and date = ?", uuids, date).Find(&info).Error; err != nil {
		beego.Error("GetTeamDaystatis err", err)
	}
	return info
}

func GetNews() []string {
	var ret []string
	var newS []commonstruct.NewsInfo

	if err := WyMysql.Table(commonstruct.WY_tmp_tool_news).Select("info").Limit(1).Find(&newS).Error; err != nil {
		beego.Error("GetNews err %v\n", err)
		return ret
	}

	for _, v := range newS {
		ret = append(ret, v.Info)
	}

	return ret
}

func GetBilishiDate(inittime int64) int64 {
	var optiem time.Time
	if inittime == 0 {
		optiem = time.Now()
	} else {
		optiem = commonfunc.StrToBjTime_YL(fmt.Sprintf("%d", inittime))
	}

	cstBls := time.FixedZone("CST", 2*3600)
	nowtime, _ := strconv.ParseInt(optiem.In(cstBls).Format("20060102"), 10, 64)

	// h, _ := time.ParseDuration("1h")
	// result := optiem.Add(-6 * h)
	// retdate := result.Format("20060102")
	// nowtime, _ := strconv.ParseInt(retdate, 10, 64)
	return nowtime
}

// 更新订单数据
func UpdateOrdercolumn(orderid int64, column string, value interface{}) error {

	updateValues := map[string]interface{}{
		column: value,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("orderid = ?", orderid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateOrdercolumn err ", orderid, column, value, err)
		return err
	}
	return nil
}

// 更新订单状态
func UpdateOrderstate(tblname string, roomid int64, expect string, state int) error {
	if err := WyMysql.Table(tblname).Where("roomid = ? and expect = ? and state = ?", roomid, expect, commonstruct.Ret_UnSettle).
		Update("state", state).Error; err != nil {
		beego.Error("UpdateOrderstate err ", roomid, expect, err)
		return err
	}
	return nil
}

func GetSettleitemS(dalei string, settletype string) ([]commonstruct.CGameItem, error) {
	var infos []commonstruct.CGameItem
	err := WyMysql.Table(commonstruct.WY_gm_config_item).Where("lottery_dalei = ? and game_type = ?", dalei, settletype).Find(&infos).Error
	return infos, err
}

// 查询未结算期号信息
func GetServermd5key(servername string) (commonstruct.ServerMD5, error) {
	var info commonstruct.ServerMD5
	err := WyMysql.Table(commonstruct.WY_gm_config_serverkey).Where("server_name = ?", servername).Find(&info).Error
	return info, err
}

func GetSettleitemByID(itemid int64) (commonstruct.CGameItem, error) {
	var infos commonstruct.CGameItem
	err := WyMysql.Table(commonstruct.WY_gm_config_item).Where("id = ? ", itemid).Find(&infos).Error
	return infos, err
}

// 获取公司玩法信息
func GetCompanyOddsInfo(companyId int64, roomid int64, portId int64) (commonstruct.CompanyPortinfo, error) {
	var info commonstruct.CompanyPortinfo

	err := WyMysql.Table(commonstruct.WY_company_portclass).Where("company_id = ? and room_id = ? and port_id = ?", companyId, roomid, portId).Find(&info).Error

	if err != nil {
		beego.Error("GetCompanyOddsInfo err", err.Error())
	}

	return info, err
}

func LogMode(b bool) {
	WyMysql.LogMode(b)
}

func GetGormDB() *gorm.DB {
	return WyMysql
}

// 结算日志
func AddSettlelog(ResID int64, begintime int64, roomid int64, expect string, expinfo string) {
	var oplog commonstruct.SettleLog
	oplog.ResID = ResID
	oplog.EndTime = commonfunc.GetNowtime()
	oplog.BeginTime = begintime
	oplog.RoomID = int64(roomid)
	oplog.Expect = expect
	oplog.ExpInfo = expinfo
	if err := WyMysql.Table(commonstruct.WY_tmp_log_settle).Create(&oplog).Error; err != nil {
		beego.Error("err %v", err)
	}
}

// 開獎、中獎信息 寫入訂單表
func SetOrderWininfo(roomid int64, expect string, code string, i_arrivetime int64) error {
	return nil
}

// 保存国家玩法订单的中奖信息
func SetOrderMultiWin(roomname string, roomid int64, orderid int64, wininfo []commonstruct.WinMulti) error {

	return nil
}

// 获取多人单个游戏的设置，主要用于占成计算
func GetUsersGamesetS(uuids []int64, roomid int64) ([]commonstruct.UserGameset, error) {
	var info []commonstruct.UserGameset
	err := WyMysql.Table(commonstruct.WY_user_gameset).Where("uuid in (?) and room_id = ?", uuids, roomid).Find(&info).Error
	if err != nil {
		beego.Error("GetUserGameset err", uuids, roomid, err)
	}
	return info, err
}

func GetUserGameset(uuid int64, roomid int64) (commonstruct.UserGameset, error) {
	var info commonstruct.UserGameset
	err := WyMysql.Table(commonstruct.WY_user_gameset).Where("uuid = ? and room_id = ?", uuid, roomid).Find(&info).Error
	// if err != nil {
	// 	beego.Error("GetUserGameset err", uuid, roomid, err)
	// }
	return info, err
}

func GetWaitupdateSet(column string) ([]commonstruct.UserGameset, error) {

	var info []commonstruct.UserGameset
	err := WyMysql.Table(commonstruct.WY_user_gameset).Where(fmt.Sprintf("%v != %v_next", column, column)).Find(&info).Error
	// if err != nil {
	// 	beego.Error("GetUserGameset err", uuid, roomid, err)
	// }
	return info, err
}

func UpdateUsergameSetByNext(uuid int64, roomid int64, column string) error {

	updateValues := map[string]interface{}{
		column: gorm.Expr(fmt.Sprintf("%v_next", column)),
	}

	if err := WyMysql.Table(commonstruct.WY_user_gameset).
		Where("uuid = ? and room_id = ?", uuid, roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateUsergameSetByNext err ", uuid, roomid, column, err)
		return err
	}
	return nil
}

func UpsertUserlotdataItem(newdata commonstruct.LotdataUseritem) {

	newdata.Date = GetBilishiDate(0)

	var olddata commonstruct.LotdataUseritem
	tblname := commonstruct.WY_tmp_user_lotdata_item
	if retinfo := WyMysql.Table(tblname).
		Where("uuid = ? and date = ? and room_id = ? and pan = ? and expect = ? and port_id = ? and item_id = ?", newdata.Uuid, newdata.Date, newdata.RoomID, newdata.Pan, newdata.Expect, newdata.PortID, newdata.ItemID).
		Find(&olddata); retinfo.Error != nil {
		if retinfo.RecordNotFound() {

			if err := WyMysql.Table(tblname).Create(&newdata).Error; err != nil {
				beego.Error("Create err ", newdata, err)
			}
		} else {
			beego.Error("UpsetUserlotdataItem err \n", newdata, retinfo.Error.Error())
		}
		return
	}

	updateValues := map[string]interface{}{
		"order_num":         gorm.Expr("order_num + ?", newdata.OrderNum),
		"order_amount":      gorm.Expr("order_amount + ?", newdata.OrderAmount),
		"zhanchenghuoliang": gorm.Expr("zhanchenghuoliang + ?", newdata.Zhanchenghuoliang),
		"xiajibuhuo":        gorm.Expr("xiajibuhuo + ?", newdata.Xiajibuhuo),
		"zidongbuchu":       gorm.Expr("zidongbuchu + ?", newdata.Zidongbuchu),
		"shizhanhuoliang":   gorm.Expr("shizhanhuoliang + ?", newdata.Shizhanhuoliang),
		"shoudongbuchu":     gorm.Expr("shoudongbuchu + ?", newdata.Shoudongbuchu),
	}

	if err := WyMysql.Table(tblname).
		Where("uuid = ? and date = ? and room_id = ? and pan = ? and expect = ? and port_id = ? and item_id = ?", newdata.Uuid, newdata.Date, newdata.RoomID, newdata.Pan, newdata.Expect, newdata.PortID, newdata.ItemID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsetUserlotdataItem err ", newdata, err)
	}
}

func UpsertUserlotdataTema(newdata commonstruct.LotdataUsertema) {

	newdata.Date = GetBilishiDate(0)

	var olddata commonstruct.LotdataUsertema
	tblname := commonstruct.WY_tmp_user_lotdata_tema
	if retinfo := WyMysql.Table(tblname).
		Where("uuid = ? and date = ? and room_id = ? and expect = ? and port_id = ? and item_id = ?", newdata.Uuid, newdata.Date, newdata.RoomID, newdata.Expect, newdata.PortID, newdata.ItemID).
		Find(&olddata); retinfo.Error != nil {
		if retinfo.RecordNotFound() {

			if err := WyMysql.Table(tblname).Create(&newdata).Error; err != nil {
				beego.Error("Create err ", newdata, err)
			}
		} else {
			beego.Error("UpsetUserlotdataItem err \n", newdata, retinfo.Error.Error())
		}
		return
	}

	updateValues := map[string]interface{}{
		"order_num":    gorm.Expr("order_num + ?", newdata.OrderNum),
		"order_amount": gorm.Expr("order_amount + ?", newdata.OrderAmount),
	}

	if err := WyMysql.Table(tblname).
		Where("uuid = ? and date = ? and room_id = ? and expect = ? and port_id = ? and item_id = ?", newdata.Uuid, newdata.Date, newdata.RoomID, newdata.Expect, newdata.PortID, newdata.ItemID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsetUserlotdataItem err ", newdata, err)
	}
}

func GetPreDate(date int64) int64 {
	loc := time.FixedZone("CST", 8*3600) //重要：获取时区
	t, _ := time.ParseInLocation("20060102", fmt.Sprintf("%d", date), loc)
	nexttime := t.AddDate(0, 0, -1)
	predate, _ := strconv.ParseInt(nexttime.Format("20060102"), 10, 64)
	return predate
}

func GetNextDate(date int64) int64 {
	loc := time.FixedZone("CST", 8*3600) //重要：获取时区
	t, _ := time.ParseInLocation("20060102", fmt.Sprintf("%d", date), loc)
	nexttime := t.AddDate(0, 0, 1)
	nextdate, _ := strconv.ParseInt(nexttime.Format("20060102"), 10, 64)
	return nextdate
}

func GetDynimicsetS(roomid int64, expect string) []commonstruct.DynimicOddsInfo {
	var oddsinfo []commonstruct.DynimicOddsInfo
	if err := WyMysql.Table(commonstruct.WY_company_dyncitem).
		Where("room_id = ? and expect = ?", roomid, expect).Find(&oddsinfo).Error; err != nil {
	}
	return oddsinfo
}

func GetUserRoomPortclasssetS(uuid int64, roomid int64) ([]commonstruct.UserPortclassset, error) {
	var info []commonstruct.UserPortclassset
	tblname := GetUserPortclasssetTablename(uuid)

	var err error
	if roomid == -1 {
		err = WyMysql.Table(tblname).
			Where("uuid = ?", uuid).
			Find(&info).Error
	} else {
		err = WyMysql.Table(tblname).
			Where("uuid = ? and room_id = ?", uuid, roomid).
			Find(&info).Error
	}
	// if err != nil {
	// 	beego.Error("GetUserRoomPortclasssetS err", uuid, roomid, err)
	// }
	return info, err
}

func GetPreZhanchenginfoS(uuid int64, preidlist []int64, roomid int64) ([]commonstruct.AgentZhanchenginfo, error) {
	var retinfo []commonstruct.AgentZhanchenginfo

	// 加上当前玩家ID ，才能凑出整条线
	preidlist = append(preidlist, uuid)

	if usersgamesets, err := GetUsersGamesetS(preidlist, roomid); err != nil {
		return nil, err
	} else {
		// 所有上级的游戏设置 写入缓存
		preid_gameset := make(map[int64]commonstruct.UserGameset)
		for _, usergameset := range usersgamesets {
			preid_gameset[usergameset.Uuid] = usergameset
		}

		{ // 先写入 会员的数据
			var agentinfo commonstruct.AgentZhanchenginfo
			agentinfo.Uuid = uuid
			userinfo, _ := GetUserinfoByUuid(uuid)
			agentinfo.Account = userinfo.Account
			agentinfo.RoleType = userinfo.RoleType
			agentinfo.BaseZhancheng = 0
			retinfo = append(retinfo, agentinfo)
		}

		var sufzhancheng float64 = 0 // 所有下级的总占成

		for i := len(preidlist) - 2; i > 1; i-- { // 直接从代理开始计算 并且不计算管理员一级

			// 当前要计算的代理 占成
			usergameset := preid_gameset[preidlist[i]]

			// usergameset.SufMinZhancheng // 上级给我的最低占成
			// usergameset.Zhancheng       // 上级想在我这里的占成

			// 我想在下级那里占的成数
			sufiwant := preid_gameset[preidlist[i+1]].Zhancheng

			var agentinfo commonstruct.AgentZhanchenginfo
			agentinfo.Uuid = preidlist[i]
			userinfo, _ := GetUserinfoByUuid(preidlist[i])

			agentinfo.Account = userinfo.Account
			agentinfo.RoleType = userinfo.RoleType

			// 如果上级设置的最低占成 - 下级所有的占成 > 我想要的占成 =>  我必须补齐占成  否则，就按我说的办
			if (usergameset.SufMinZhancheng - sufzhancheng) > sufiwant {
				if (sufzhancheng + sufiwant) > usergameset.SufMaxZhancheng {
					agentinfo.BaseZhancheng = usergameset.SufMaxZhancheng - sufzhancheng
				} else {
					agentinfo.BaseZhancheng = usergameset.SufMinZhancheng - sufzhancheng
				}
			} else {
				if (sufzhancheng + sufiwant) > usergameset.SufMaxZhancheng {
					agentinfo.BaseZhancheng = usergameset.SufMaxZhancheng - sufzhancheng
				} else {
					agentinfo.BaseZhancheng = sufiwant
				}
			}

			sufzhancheng = sufzhancheng + agentinfo.BaseZhancheng

			retinfo = append(retinfo, agentinfo)
		}
	}

	return retinfo, nil
}

// 用户游戏设置
func InitUserGameset(uuid int64, info commonstruct.UserGameset) error {

	if retinfo := WyMysql.Table(commonstruct.WY_user_gameset).Where("uuid = ? and room_id = ?", uuid, info.RoomID).
		Find(&info); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_user_gameset).Create(&info).Error; err != nil {
				beego.Error("create err ", info, err)
				return err
			}
		} else {
			beego.Error("InitUserGameset err ", uuid, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	return errors.New("record is exist!")
}

// 系统关闭游戏设置
func DeleteUserGameset(destuuid int64, roomid int64) error {
	if err := WyMysql.Table(commonstruct.WY_user_gameset).Delete(nil, "uuid = ? and room_id = ?", destuuid, roomid).Error; err != nil {
		beego.Error("DeleteUserGameset", destuuid, roomid, err)
		return err
	}
	return nil
}

// 系统关闭游戏设置
func DeleteUsersGameset(roomid int64) error {
	// updateValues := map[string]interface{}{
	// 	column: value,
	// }
	// if err := WyMysql.Table(commonstruct.WY_user_gameset).
	// 	Where("room_id = ?", roomid).
	// 	Update(updateValues).Error; err != nil {
	// 	beego.Error("UpdateUserGameset err \n", roomid, column, value, err)
	// }

	if err := WyMysql.Table(commonstruct.WY_user_gameset).Delete(nil, "room_id = ?", roomid).Error; err != nil {
		beego.Error("DeleteUsersGameset", roomid, err)
		return err
	}
	return nil
}

// 系统关闭游戏设置
func DeleteCompanyGameset(roomid int64) error {
	if err := WyMysql.Table(commonstruct.WY_company_game).Delete(nil, "room_id = ?", roomid).Error; err != nil {
		beego.Error("DeleteCompanyGameset", roomid, err)
		return err
	}
	return nil
}

func GetRoomInfoByID(roomid int64) commonstruct.CRoomInfo {
	var roominfo commonstruct.CRoomInfo
	if err := WyMysql.Table(commonstruct.WY_gm_config_room).Where("id = ?", roomid).Find(&roominfo).Error; err != nil {
		beego.Error("err %v", err)
	}
	return roominfo
}

// 获取系统权限目录
func GetSysnaviS() ([]commonstruct.BKNavi, error) {
	var infos []commonstruct.BKNavi
	if err := WyMysql.Table(commonstruct.WY_gm_config_bknavi).Order("group_id,id").
		Find(&infos).Error; err != nil {
		beego.Error("GetSysCompanybknavis err", err)
		return infos, err
	}
	return infos, nil
}

func InitUsernavi(uuid int64, roletype string) error {

	if naviinfoS, err := GetSysnaviS(); err != nil {
		beego.Error("GetRoletypebknavis err", err.Error())
	} else {
		usernavitablename := GetUsernaviTablename(uuid)

		if err := CheckUsernaviTablename(usernavitablename); err != nil {
			beego.Error("CheckUsernaviTablename err", usernavitablename, err.Error())
			return err
		} else {
			for _, naviinfo := range naviinfoS {
				var oldinfo commonstruct.UserNavi
				if retinfo := WyMysql.Table(usernavitablename).
					Where("uuid = ? and navi_id = ?", uuid, naviinfo.ID).
					Find(&oldinfo); retinfo.Error != nil {
					if retinfo.RecordNotFound() {
						oldinfo.Uuid = uuid
						oldinfo.NaviID = naviinfo.ID
						oldinfo.GroupID = naviinfo.GroupID

						switch roletype {
						case "company":
							oldinfo.Value = naviinfo.Company
						case "expmanager":
							oldinfo.Value = 0
						case "gm":
							oldinfo.Value = naviinfo.Admin
						case "saler":
							oldinfo.Value = naviinfo.Saler
						case "agent":
							oldinfo.Value = naviinfo.Agent
						default:
							beego.Error("roletype err", roletype)
							return errors.New("roletype err")
						}

						beego.Error("create new navi ", oldinfo)

						if err := WyMysql.Table(usernavitablename).Create(&oldinfo).Error; err != nil {
							beego.Error("create err ", oldinfo, err)
							return err
						}
					} else {
						beego.Error("CheckUsernavi err ", uuid, roletype, retinfo.Error.Error())
						return retinfo.Error
					}
				}
			}
		}
	}
	return nil
}

/************************************
* 新增公司VIP升级设置
************************************/
func CreateLevelcfg(uuid int64, level int64, validamount float64, recharge_succamount float64) error {

	var newinfo commonstruct.LevelConfig

	newinfo.Uuid = uuid
	newinfo.Level = level
	newinfo.Validamount = validamount
	newinfo.RechargeSuccamount = recharge_succamount

	if err := WyMysql.Table(commonstruct.WY_tmp_company_level).Create(&newinfo).Error; err != nil {
		beego.Error("CreateLevelcfg err ", newinfo, err)
		return err
	}
	return nil
}

/************************************
* 新增公司公告
************************************/
func CreateCompanynotice(companyid int64, platform string, gametype string, showtype string, begintime int64, endtime int64, info string) error {

	var newinfo commonstruct.CNoticeInfo

	newinfo.CompanyID = companyid
	newinfo.ShowType = showtype
	newinfo.GameType = gametype
	newinfo.Optime = commonfunc.GetNowtime()
	newinfo.Begintime = begintime
	newinfo.Endtime = endtime
	newinfo.Platform = platform
	newinfo.Info = info

	if err := WyMysql.Table(commonstruct.WY_company_notice).Create(&newinfo).Error; err != nil {
		beego.Error("CreateCompanynotice err ", newinfo, err)
		return err
	}
	return nil
}

/************************************
* 修改公司公告
************************************/
func UpdateCompanynotice(id int64, platform string, gametype string, showtype string, begintime int64, endtime int64, info string) error {

	updateValues := map[string]interface{}{
		"show_type": showtype,
		"game_type": gametype,
		"optime":    commonfunc.GetNowtime(),
		"begintime": begintime,
		"endtime":   endtime,
		"platform":  platform,
		"info":      info,
	}
	if err := WyMysql.Table(commonstruct.WY_company_notice).
		Where("id = ?", id).Update(updateValues).Error; err != nil {
		beego.Error("UpdateCompanynotice err ", id, err)
	}
	return nil
}

// 新增公司公告
func GetCompanynotice(companyid int64, platform string, gametype string, showtype string) ([]commonstruct.CNoticeInfo, error) {
	var selectarg string
	switch platform {
	case "all":
		switch showtype {
		case "all":
			selectarg = fmt.Sprintf("company_id = %v and game_type = '%v'", companyid, gametype)
		default:
			selectarg = fmt.Sprintf("company_id = %v and game_type = '%v' and (show_type = '%v' or show_type = 'all')", companyid, gametype, showtype)
		}
	default:
		switch showtype {
		case "all":
			selectarg = fmt.Sprintf("company_id = %v and (platform = '%v' or platform = 'all') and game_type = '%v'", companyid, platform, gametype)
		default:
			selectarg = fmt.Sprintf("company_id = %v and (platform = '%v' or platform = 'all') and game_type = '%v' and (show_type = '%v' or show_type = 'all')", companyid, platform, gametype, showtype)
		}
	}

	// beego.Error("selectarg == ", selectarg)

	var newinfo []commonstruct.CNoticeInfo
	err := WyMysql.Table(commonstruct.WY_company_notice).Where(selectarg).Find(&newinfo).Error
	return newinfo, err
}

// 删除公司公告
func DeleteCompanynotice(companyid int64, id int64) error {
	if err := WyMysql.Table(commonstruct.WY_company_notice).Delete(nil, "company_id = ? and id = ?", companyid, id).Error; err != nil {
		beego.Error("DeleteCompanynotice", commonstruct.WY_company_notice, companyid, id, " err ", err)
		return err
	}
	return nil
}

// 新增公司活动
func NewPromotion(companyid int64, title string, logourl string, optime int64, info string) error {
	// 插入数据
	promotion := &commonstruct.CPromotionInfo{
		CompanyID: companyid,
		Title:     title,
		LogoUrl:   logourl,
		Optime:    optime,
		Begintime: commonfunc.GetNowtime(),
		Endtime:   20990101000000,
		Info:      info,
	}

	if err := WyMysql.Table(commonstruct.WY_company_promotion).Create(promotion).Error; err != nil {
		beego.Error("NewPromotion err ", promotion, err)
		return err
	}
	return nil
}

func InitUsermoney(uuid int64, cash float64, gaopin float64, tcp3 float64, fc3d float64, taiwanliuhe float64, taiwandaletou float64, xianggangliuhe float64, aomenliuhe float64) error {

	var Usermoney commonstruct.UserMoney
	Usermoney.Uuid = uuid
	Usermoney.Cash = cash
	Usermoney.Gaopin = gaopin
	Usermoney.GaopinLimit = gaopin
	Usermoney.Tcp3 = tcp3
	Usermoney.Tcp3Limit = tcp3
	Usermoney.Fc3d = fc3d
	Usermoney.Fc3dLimit = fc3d
	Usermoney.Taiwanliuhe = taiwanliuhe
	Usermoney.TaiwanliuheLimit = taiwanliuhe
	Usermoney.Taiwandaletou = taiwandaletou
	Usermoney.TaiwandaletouLimit = taiwandaletou
	Usermoney.Xianggangliuhe = xianggangliuhe
	Usermoney.XianggangliuheLimit = xianggangliuhe
	Usermoney.Aomenliuhe = aomenliuhe
	Usermoney.AomenliuheLimit = aomenliuhe

	if cash > 0.000001 {
		AddUserMoneylog(uuid, "wylottery", commonstruct.UptType_Benifit, cash, 0, cash, "新用户注册彩金")
	}
	return WyMysql.Table(commonstruct.WY_user_money).Create(&Usermoney).Error
}

/***********
* 获取系统默认公司信息
 ***********/
func GetDefaultCompanyinfo() (commonstruct.DefaultCompanyBase, error) {
	var data commonstruct.DefaultCompanyBase
	err := WyMysql.Table(commonstruct.WY_gm_config_defaultcompanyinfo).Find(&data).Error
	if err != nil {
		beego.Error("GetDefaultCompanyinfo err", err.Error())
	}

	return data, err
}

func InitCompanyOdds(uuid int64, roomid int64, oddsinfo commonstruct.SalerBaseodds, baseoddsper float64) error {

	var CoOdds commonstruct.CompanyPortinfo
	if retinfo := WyMysql.Table(commonstruct.WY_company_portclass).
		Where("company_id = ? and room_id = ? and port_id = ?", uuid, roomid, oddsinfo.Portid).
		Find(&CoOdds); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			CoOdds.CompanyID = uuid
			CoOdds.RoomID = int64(roomid)
			CoOdds.PortID = oddsinfo.Portid

			defaultodds := oddsinfo.DefaultOdds * baseoddsper

			if oddsinfo.DefaultOdds > 1 {
				CoOdds.OddsA = Round4s5r(defaultodds-0.01, 4)
				CoOdds.TuishuiA = 0.005
				CoOdds.OddsB = Round4s5r(defaultodds-0.02, 4)
				CoOdds.TuishuiB = 0.01
				CoOdds.OddsC = Round4s5r(defaultodds-0.04, 4)
				CoOdds.TuishuiC = 0.02
				CoOdds.OddsD = Round4s5r(defaultodds-0.06, 4)
				CoOdds.TuishuiD = 0.03
			}

			CoOdds.DefaultOdds = defaultodds
			CoOdds.MinAmount = 1
			CoOdds.Kaiguan = 1
			CoOdds.MaxAmount = 5000
			CoOdds.MaxAmountExpect = 50000
			CoOdds.TransferAmount = 0
			CoOdds.WarningAmount = 0
			CoOdds.PortSwitch = 0

			if err := WyMysql.Table(commonstruct.WY_company_portclass).Create(&CoOdds).Error; err != nil {
				beego.Error("InitCompanyOdds err", CoOdds, err)
				return err
			}
		} else {
			beego.Error("UpsertCompanygame err ", CoOdds, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}
	return nil
}

// 用户玩法设置
func InitUserPortclassset(uuid int64, info commonstruct.UserPortclassset) error {
	tblname := GetUserPortclasssetTablename(uuid)

	if retinfo := WyMysql.Table(tblname).Where("uuid = ? and room_id = ? and port_id = ?", uuid, info.RoomID, info.PortID).
		Find(&info); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(tblname).Create(&info).Error; err != nil {
				beego.Error("InitUserPortclassset err ", info, err)
				return err
			}
		} else {
			beego.Error("InitUserPortclassset err ", uuid, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	return errors.New("record is exist!")
}

func GetPortclassS(lotterydalei string, settletype string) ([]commonstruct.PortClass, error) {
	var classlist []commonstruct.PortClass
	err := WyMysql.Table(commonstruct.WY_gm_config_portclass).
		Where("lottery_dalei = ? and settle_type = ? and sys_kaiguan = 1", lotterydalei, settletype).Find(&classlist).Error
	if err != nil {
	}
	return classlist, err
}

func CreateSalerBaseodds(newinfo commonstruct.SalerBaseodds) error {
	var oldinfo commonstruct.SalerBaseodds
	if retinfo := WyMysql.Table(commonstruct.WY_user_settletype).
		Where("uuid = ? and settle_type = ? and portid = ?", newinfo.Uuid, newinfo.SettleType, newinfo.Portid).
		Find(&oldinfo); retinfo.Error != nil {

		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_user_settletype).Create(&newinfo).Error; err != nil {
				beego.Error("create err ", newinfo, err)
				return err
			}
		} else {
			beego.Error("CreateSalerBaseodds err ", newinfo.Uuid, retinfo.Error.Error())
			return retinfo.Error
		}
	}
	return nil
}

func GetSalerBaseoddsS(uuid int64, settletype string) ([]commonstruct.SalerBaseodds, error) {
	var classlist []commonstruct.SalerBaseodds
	err := WyMysql.Table(commonstruct.WY_user_settletype).
		Where("uuid = ? and settle_type = ?", uuid, settletype).Find(&classlist).Error
	if err != nil {
		beego.Error("GetSalerBaseoddsS err ", uuid, settletype, err)
	}
	return classlist, err
}

func GetSalerBaseodds(uuid int64, portid int64) (commonstruct.SalerBaseodds, error) {
	var classlist commonstruct.SalerBaseodds
	err := WyMysql.Table(commonstruct.WY_user_settletype).
		Where("uuid = ? and portid = ?", uuid, portid).Find(&classlist).Error
	if err != nil {
		beego.Error("GetSalerBaseodds err ", uuid, portid, err)
	}
	return classlist, err
}

//
func CreateCompanygame(masterid int64, roomid int64, inuse int64, opentime int64) error {
	var oldinfo commonstruct.CompanyGame
	if retinfo := WyMysql.Table(commonstruct.WY_company_game).
		Where("company_id = ? and room_id = ? ", masterid, roomid).
		Find(&oldinfo); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			oldinfo.CompanyID = masterid
			oldinfo.Port = "A"
			oldinfo.RoomID = roomid
			roominfo := GetRoomInfoByID(roomid)
			oldinfo.Platform = roominfo.GameDalei
			oldinfo.OpTime = commonfunc.GetNowtime()
			oldinfo.ValidTime = 20990101000000
			oldinfo.InUse = inuse
			oldinfo.PreKaiguan = inuse
			oldinfo.SortID = roomid
			oldinfo.Openvalue = opentime

			if err := WyMysql.Table(commonstruct.WY_company_game).Create(&oldinfo).Error; err != nil {
				beego.Error("create err ", oldinfo, err)
				return err
			}
		} else {
			beego.Error("CreateCompanygame err ", oldinfo, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}
	return nil
}

//
func CreateCompanygameset(newinfo commonstruct.CompanyGame) error {
	var oldinfo commonstruct.CompanyGame
	if retinfo := WyMysql.Table(commonstruct.WY_company_game).
		Where("company_id = ? and room_id = ? ", newinfo.CompanyID, newinfo.RoomID).
		Find(&oldinfo); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_company_game).Create(&newinfo).Error; err != nil {
				beego.Error("create err ", newinfo, err)
				return err
			}
		} else {
			beego.Error("CreateCompanygameset err ", oldinfo, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}
	return errors.New("存在该游戏记录")
}

// 返佣日志
func AddRetCashWait(MasterID int64, date int64, platform string, gametype string, expinfo string) error {

	var oldlog commonstruct.RetCashWait
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_wait_retcash).Where("master_id = ? and date = ? and platform = ? and gametype = ?", MasterID, date, platform, gametype).
		Find(&oldlog); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			var oplog commonstruct.RetCashWait
			oplog.MasterID = MasterID
			oplog.Date = date
			oplog.Platform = platform
			oplog.Gametype = gametype
			oplog.Expinfo = expinfo
			if err := WyMysql.Table(commonstruct.WY_tmp_wait_retcash).Create(&oplog).Error; err != nil {
				beego.Error("err %v", err)
				return err
			}
			return nil
		} else {
			beego.Error("AddRetCashWait err ", retinfo.Error.Error())
			return retinfo.Error
		}
	}
	return errors.New("已经执行过返佣")
}

/***********
* 设置彩种的结算号码
 ***********/
func UpsertOpencode(roomid int64, expect string, code string) error {
	roominfo := GetRoomInfoByID(int64(roomid))
	if roominfo.ID == 0 {
		return errors.New("房间ID 错误")
	}

	tblname := fmt.Sprintf("wy_num_%v", roominfo.NameENG)

	var oldinfo commonstruct.DBWininfo
	if retinfo := WyMysql_ltdata.Table(tblname).Find(&oldinfo, "expect = ?", expect); retinfo.Error != nil {
		if retinfo.RecordNotFound() {

			oldinfo.Expect = expect
			oldinfo.Opencode = code
			oldinfo.Opentime = commonfunc.GetBjTime20060102150405(time.Now())
			oldinfo.Opentimestamp = commonfunc.GetNowtime()
			oldinfo.Res = "manager"

			sql := fmt.Sprintf("INSERT INTO `%v` (`expect`,`opencode`,`opentime`,`opentimestamp`) VALUES ('%v','%v','%v','%v')",
				tblname, oldinfo.Expect, oldinfo.Opencode, oldinfo.Opentime, oldinfo.Opentimestamp)
			return WyMysql_ltdata.Exec(sql).Error
		} else {
			return retinfo.Error
		}
	}

	updateValues := map[string]interface{}{
		"opencode": code,
		"res":      "manager",
	}
	if err := WyMysql_ltdata.Table(tblname).
		Where("expect = ?", expect).Update(updateValues).Error; err != nil {
		beego.Error("UpsertOpencode err ", tblname, expect, err)
		return err
	}

	return nil
}

func GetExpectWishtime(tablesuf string, expect string) (commonstruct.LotteryOpentime, error) {
	var Opentime commonstruct.LotteryOpentime
	if tablesuf == "" {
		beego.Error("为空调用===", GetFuncName(3))
		return Opentime, errors.New("表名不能为空")
	}

	tblname := fmt.Sprintf("%s%s", commonstruct.WY_lotterytime_, tablesuf)

	if err := WyMysql_ltdata.Table(tblname).Select("id,expect,unix_timestamp(opentimestamp) as opentimestamp").
		Where("expect = ?", expect).Find(&Opentime).Error; err != nil {
		beego.Error("GetExpectWishtime err ", err, tablesuf)
		return Opentime, err
	}
	return Opentime, nil
}

func GetFuncName(level int) string {
	// 1 为调用的上一级
	var strRet string
	for i := level + 1; i >= 1; i-- {
		pc, _, line, _ := runtime.Caller(i)
		f := runtime.FuncForPC(pc)
		strRet = strRet + f.Name() + "--->" + strconv.Itoa(line) + " $ "
	}
	return strRet
}

// 获取公司玩法设置
func GetCompanyoddsS(uuid int64, gameid int64) ([]commonstruct.CompanyPortinfo, error) {
	var portinfoS []commonstruct.CompanyPortinfo
	err := WyMysql.Table(commonstruct.WY_company_portclass).Where("company_id = ? and room_id = ?", uuid, gameid).Find(&portinfoS).Error
	if err != nil {
		beego.Error("GetCompanyoddsS err ", uuid, gameid, err)
	}
	return portinfoS, err
}

// 获取公司加时玩法设置
func GetCompanyJiashiportclassS(uuid int64, gameid int64) ([]commonstruct.CompanyJiashiPortinfo, error) {
	var portinfoS []commonstruct.CompanyJiashiPortinfo
	err := WyMysql.Table(commonstruct.WY_company_portclass_jiashi).Where("company_id = ? and room_id = ?", uuid, gameid).Find(&portinfoS).Error
	if err != nil {
		beego.Error("GetCompanyJiashiportclassS err ", uuid, gameid, err)
	}
	return portinfoS, err
}

/*
 根据lv2描述获取权限信息
*/
func GetNaviinfoBylv2(navi_lv1 string, navi_lv2 string) (commonstruct.BKNavi, error) {
	var infos commonstruct.BKNavi
	err := WyMysql.Table(commonstruct.WY_gm_config_bknavi).Where("navi_lv1 = ? and navi_lv2 = ?", navi_lv1, navi_lv2).Find(&infos).Error
	return infos, err
}

/*
 获取权限ID的权限信息
*/
func GetNaviinfo(naviid int64) (commonstruct.BKNavi, error) {
	var infos commonstruct.BKNavi
	err := WyMysql.Table(commonstruct.WY_gm_config_bknavi).Where("id = ?", naviid).Find(&infos).Error
	return infos, err
}

func GetIteminfoSByPortclass(platform string, settletype string, gamedalei string, gamexiaolei string) ([]commonstruct.CGameItem, error) {
	var info []commonstruct.CGameItem
	err := WyMysql.Table(commonstruct.WY_gm_config_item).
		Where("lottery_dalei = ? and game_type = ? and game_dalei = ? and game_xiaolei = ?", platform, settletype, gamedalei, gamexiaolei).
		Find(&info).Error
	return info, err
}

// 获取单人的多个游戏设置
func GetUserGamesetS(uuid int64) ([]commonstruct.UserGameset, error) {
	var info []commonstruct.UserGameset
	err := WyMysql.Table(commonstruct.WY_user_gameset).Where("uuid = ? and in_use = 1", uuid).Find(&info).Error
	if err != nil {
		beego.Error("GetUserGameset err", uuid, err)
	}
	return info, err
}

//
func DeleteUserPortclasssetS(uuid int64, roomid int64) error {
	numtblname := GetUserPortclasssetTablename(uuid)
	var err error
	if roomid == -1 {
		if err = WyMysql.Table(numtblname).Delete(nil, "uuid = ?", uuid).Error; err != nil {
			beego.Error("DeleteUserPortclasssetS err ", uuid, err)
			return err
		}
	} else {

		if err = WyMysql.Table(numtblname).Delete(nil, "uuid = ? and room_id = ?", uuid, roomid).Error; err != nil {
			beego.Error("DeleteUserPortclasssetS err ", uuid, roomid, err)
			return err
		}
	}
	return nil
}

func UpsertUserPortclassset(portinfo commonstruct.UserPortclassset, columntype string) error {

	numtblname := GetUserPortclasssetTablename(portinfo.Uuid)
	var oldset commonstruct.UserPortclassset
	if retinfo := WyMysql.Table(numtblname).
		Where("uuid = ? and room_id = ? and port_id = ?", portinfo.Uuid, portinfo.RoomID, portinfo.PortID).
		Find(&oldset); retinfo.Error != nil {
		if retinfo.RecordNotFound() {

			if err := WyMysql.Table(numtblname).Create(&portinfo).Error; err != nil {
				beego.Error("create err ", portinfo, err)
				return err
			}
		} else {
			beego.Error("UpsertUserPortclassset err ", portinfo, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	var updateValues map[string]interface{}
	switch columntype {
	case "tuishui":
		updateValues = map[string]interface{}{
			"tuishui_a": portinfo.TuishuiA,
			"tuishui_b": portinfo.TuishuiB,
			"tuishui_c": portinfo.TuishuiC,
			"tuishui_d": portinfo.TuishuiD,
		}
	case "amount":
		updateValues = map[string]interface{}{
			"min_amount":        portinfo.MinAmount,
			"max_amount":        portinfo.MaxAmount,
			"max_amount_expect": portinfo.MaxAmountExpect,
		}
	case "tuishui,amount":
		updateValues = map[string]interface{}{
			"tuishui_a":         portinfo.TuishuiA,
			"tuishui_b":         portinfo.TuishuiB,
			"tuishui_c":         portinfo.TuishuiC,
			"tuishui_d":         portinfo.TuishuiD,
			"min_amount":        portinfo.MinAmount,
			"max_amount":        portinfo.MaxAmount,
			"max_amount_expect": portinfo.MaxAmountExpect,
		}
	case "transferamount":
		updateValues = map[string]interface{}{
			"transfer_amount": portinfo.TransferAmount,
			"is_transfer":     portinfo.IsTransfer,
		}
	case "warningamount":
		updateValues = map[string]interface{}{
			"warning_amount":     portinfo.WarningAmount,
			"warning_loopamount": portinfo.WarningLoopamount,
		}
	default:
		return errors.New(fmt.Sprintf("unknown columntype [%v]", columntype))
	}

	if err := WyMysql.Table(numtblname).
		Where("uuid = ? and room_id = ? and port_id = ?", portinfo.Uuid, portinfo.RoomID, portinfo.PortID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateUserPortclassset err ", portinfo.Uuid, portinfo.PortID, err)
	}
	return nil
}

func UpdateUserTuishuiset(uuid int64, roomid int64, portid int64, pan string, value float64) error {
	var updateValues map[string]interface{}
	switch pan {
	case "A":
		updateValues = map[string]interface{}{
			"tuishui_a": value,
		}
	case "B":
		updateValues = map[string]interface{}{
			"tuishui_b": value,
		}
	case "C":
		updateValues = map[string]interface{}{
			"tuishui_c": value,
		}
	case "D":
		updateValues = map[string]interface{}{
			"tuishui_d": value,
		}
	case "MinAmount":
		updateValues = map[string]interface{}{
			"min_amount": value,
		}
	case "MaxAmount":
		updateValues = map[string]interface{}{
			"max_amount": value,
		}
	case "MaxAmountExpect":
		updateValues = map[string]interface{}{
			"max_amount_expect": value,
		}
	default:
		beego.Error("UpdateUserTuishuiset err", pan)
	}

	numtblname := GetUserPortclasssetTablename(uuid)

	if err := WyMysql.Table(numtblname).
		Where("uuid = ? and room_id = ? and port_id = ?", uuid, roomid, portid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateUserPortclassset err ", uuid, roomid, portid, err)
	}

	beego.Error("超出上级退水被强平", uuid, roomid, portid, updateValues)

	return nil
}

func GetItemDynimicset(uuid int64, roomid int64, expect string, itemid int64) commonstruct.DynimicOddsInfo {
	var oddsinfo commonstruct.DynimicOddsInfo
	if err := WyMysql.Table(commonstruct.WY_company_dyncitem).
		Where("uuid = ? and room_id = ? and expect = ? and item_id = ?", uuid, roomid, expect, itemid).Find(&oddsinfo).Error; err != nil {
	}
	return oddsinfo
}

func UpdateDyncitem(uuid int64, roomid int64, expect string, itemid int64, changerate float64, validtime int64) (float64, error) {
	var Dyncinfo commonstruct.DynimicOddsInfo
	if retinfo := WyMysql.Table(commonstruct.WY_company_dyncitem).Find(&Dyncinfo, "uuid = ? and room_id = ? and expect= ? and item_id = ?", uuid, roomid, expect, itemid); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if changerate > 0 {
				return 0, errors.New("赔率不能高于初始值")
			}

			Dyncinfo.Uuid = uuid
			Dyncinfo.RoomID = roomid
			Dyncinfo.ItemID = itemid
			Dyncinfo.ValidTime = validtime
			Dyncinfo.DynimicOdds = changerate
			Dyncinfo.Expect = expect
			newinfodb := WyMysql.Table(commonstruct.WY_company_dyncitem).Create(&Dyncinfo)

			if err := newinfodb.Error; err != nil {
				beego.Error("NewOrderUser err ", err)
				return 0, newinfodb.Error
			} else {
				newinfo := newinfodb.Value.(*commonstruct.DynimicOddsInfo)
				return newinfo.DynimicOdds, nil
			}
		} else {
			beego.Error("UpdateDync_item err \n", uuid, roomid, expect, itemid, changerate, retinfo.Error.Error())
			return 0, retinfo.Error
		}
	}

	if (Dyncinfo.DynimicOdds + changerate) > 0 {
		return 0, errors.New("赔率不能高于初始值")
	}

	updateValues := map[string]interface{}{
		"dynimic_odds": gorm.Expr("dynimic_odds + ?", changerate),
		"valid_time":   validtime,
	}

	if err := WyMysql.Table(commonstruct.WY_company_dyncitem).
		Where("uuid = ? and room_id = ? and expect = ? and item_id = ?", uuid, roomid, expect, itemid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateDyncitem err \n", uuid, roomid, expect, itemid, changerate, err)
		return 0, err
	} else {
		if err := WyMysql.Table(commonstruct.WY_company_dyncitem).Find(&Dyncinfo, "uuid = ? and room_id = ? and expect= ? and item_id = ?", uuid, roomid, expect, itemid).Error; err != nil {
			return 0, err
		} else {
			return Dyncinfo.DynimicOdds, nil
		}
	}
}

// 修改公司的开封盘时间
func UpdateCompanyClosetime(uuid int64, roomid int64, Closevalue int64, TemaClosevalue int64, Openvalue int64, JsClosevalue int64, JsTemaClosevalue int64) error {
	updateValues := map[string]interface{}{
		"openvalue":         Openvalue,
		"closevalue":        Closevalue,
		"tema_closevalue":   TemaClosevalue,
		"js_closetime":      JsClosevalue,
		"js_tema_closetime": JsTemaClosevalue,
	}

	if err := WyMysql.Table(commonstruct.WY_company_game).
		Where("company_id = ? and room_id = ? ", uuid, roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateCompanyClosetime err ", uuid, roomid, err)
		return err
	}
	return nil
}

/*
 获取用户的单个后台目录
*/
func GetUsernaviinfo(uuid int64, naviid int64) (commonstruct.UserNavi, error) {
	var infos commonstruct.UserNavi
	usernavitablename := fmt.Sprintf("%v%v", commonstruct.WY_user_navi_, uuid/500)
	if err := WyMysql.Table(usernavitablename).Where("uuid = ? and navi_id = ?", uuid, naviid).Find(&infos).Error; err != nil {
		beego.Error("GetUsernaviS err", uuid, usernavitablename, naviid, err)
		return infos, err
	}
	return infos, nil
}

func GetUnfinishedLog(uuid int64, offset int, pagecount int, optypes []commonstruct.MoneyUpdateType, btimelimit bool /*是否需要设置时间区段*/) []commonstruct.MoneyInOut {
	// 查询订单记录
	var logs []commonstruct.MoneyInOut
	if btimelimit {
		begintime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().AddDate(0, 0, -2)), 10, 64)
		endtime := commonfunc.GetNowdate()*1000000 + 235959

		if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
			Where("master_id = ? and checked = 0 and req_type in (?) and req_time BETWEEN ? and ?", uuid, optypes, begintime, endtime).
			Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
			beego.Error("GetUnfinishedLog err", uuid, optypes, err)
		}
	} else {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).
			Where("master_id = ? and checked = 0 and req_type in (?)", uuid, optypes).
			Order("req_time desc").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&logs).Error; err != nil {
			beego.Error("GetUnfinishedLog err", uuid, optypes, err)
		}
	}

	return logs
}

// 查询未审核的提现订单记录
func GetUnfinishedlogPageinfo(uuid int64, pagecount int, optypes []commonstruct.MoneyUpdateType, btimelimit bool /*是否需要设置时间区段*/) (int64, int64) {
	var logs commonstruct.MoneyInOut
	if btimelimit {
		begintime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().AddDate(0, 0, -2)), 10, 64)
		endtime := commonfunc.GetNowdate()*1000000 + 235959

		if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
			Where("master_id = ? and checked = 0 and req_type in (?) and req_time BETWEEN ? and ?", uuid, optypes, begintime, endtime).
			Find(&logs).Error; err != nil {
			beego.Error("GetUnfinishedlogPageinfo err", uuid, optypes, err)
		}
	} else {
		if err := WyMysql.Table(commonstruct.WY_tmp_log_moneyinout).Select("Count(*) as order_id").
			Where("master_id = ? and checked = 0 and req_type in (?)", uuid, optypes).
			Find(&logs).Error; err != nil {
			beego.Error("GetUnfinishedlogPageinfo err", uuid, optypes, err)
		}
	}
	return logs.OrderID, int64(math.Ceil(float64(logs.OrderID) / float64(pagecount)))
}

func GetSufstatisticByDate(tblname string, sufids []int64, begindate int64, enddate int64) []commonstruct.LotdataStatistic {
	var list []commonstruct.LotdataStatistic

	if len(sufids) == 1 {
		if err := WyMysql.Table(tblname).
			Where("uuid = ? and date >= ? and date <= ?", sufids[0], begindate, enddate).
			Find(&list).Error; err != nil {
			beego.Error("GetSufLotterystatisticByDate err", begindate, enddate, err)
		}
	} else {
		selectinfo := `uuid,
			sum(order_num) as order_num,
			sum(order_amount) as order_amount,
			sum(settled_num) as settled_num,
			sum(settled_amount) as settled_amount,
			sum(valid_amount) as valid_amount,
			sum(wager) as wager,
			sum(profit_wager) as profit_wager,
			sum(tuishui) as tuishui,
			sum(profit_tuishui) as profit_tuishui,
			sum(shizhanhuoliang) as shizhanhuoliang,
			sum(shizhanshuying) as shizhanshuying,
			sum(shizhantuishui) as shizhantuishui,
			sum(shizhanpeicha) as shizhanpeicha`

		if err := WyMysql.Table(tblname).Select(selectinfo).
			Where("uuid in (?) and date between ? and ? ", sufids, begindate, enddate).
			Group("uuid").Find(&list).Error; err != nil {
			beego.Error("GetSufLotterystatisticByDate err", begindate, enddate, err)
		}
	}
	return list
}

// 获取公司的游戏设置
func GetCompanyGameinfo(uuid int64, roomid int64) commonstruct.CompanyGame {
	var list commonstruct.CompanyGame
	if uuid == 0 {
		return list
	}

	if err := WyMysql.Table(commonstruct.WY_company_game).Where("company_id = ? and room_id = ?", uuid, roomid).Find(&list).Error; err != nil {
		beego.Error("GetCompanyGameinfo err", uuid, roomid, err)
		return list
	}
	return list
}

func GetAgentoddsByUuid(uuid int64, port string) commonstruct.AgentOdds {
	var oddsinfo commonstruct.AgentOdds
	if err := WyMysql.Table(commonstruct.WY_user_odds).Where("uuid = ? and port = ?", uuid, port).Find(&oddsinfo).Error; err != nil {
		beego.Error("GetAgentoddsByUuid err", uuid, err)
	}
	return oddsinfo
}

// 查询团队最大值
func GetTeamMaxvalue(uuids []int64, pan string) commonstruct.AgentOdds {

	// selectarg := `max(odds) as odds,
	// max(shui_lottery) as shui_lottery,
	// max(shui_actual) as shui_actual,
	// max(shui_electric) as shui_electric,
	// max(shui_card) as shui_card,
	// max(shui_sport) as shui_sport`
	selectarg := `max(odds) as odds`

	var retinfo commonstruct.AgentOdds
	if err := WyMysql.Table(commonstruct.WY_user_odds).Select(selectarg).Where("uuid in (?) and port = ?", uuids, pan).Find(&retinfo).
		Error; err != nil {
		beego.Error("GetTeamMaxvalue err ", err)
	}
	return retinfo
}

// 查询团队最大值
func GetPreMinvalue(uuids []int64, pan string) commonstruct.AgentOdds {

	// selectarg := `min(odds) as odds,
	// min(shui_lottery) as shui_lottery,
	// min(shui_actual) as shui_actual,
	// min(shui_electric) as shui_electric,
	// min(shui_card) as shui_card,
	// min(shui_sport) as shui_sport`
	selectarg := `min(odds) as odds`

	var retinfo commonstruct.AgentOdds
	if err := WyMysql.Table(commonstruct.WY_user_odds).Select(selectarg).Where("uuid in (?) and port = ?", uuids, pan).Find(&retinfo).
		Error; err != nil {
		beego.Error("GetPreMinvalue err ", err)
	}
	return retinfo
}

/*
WY_user_base:

查询荷官列表 GetDealers(masterid int64) ([]commonstruct.Users, error)
查询多个用户的资料 GetUserbaseS(idlist []int64) []commonstruct.Users
查询子账号列表 GetExpmanagers(preid int64, pagecount int, pagenum int) ([]commonstruct.Users, error)
查询子账号页码信息 GetExpmanagersPageinfo(preid int64, pagecount int) (int64, int64)
查询直属用户列表 GetDirectRoleS(uuid int64, roletype string, offset int, pagecount int) ([]commonstruct.Users, error)
查询直属用户页码信息 GetDirectRoleSPageinfo(uuid int64, roletype string, pagecount int) (int64, int64)

查询直属用户人数 GetDirectNum(uuid int64, roletype string) (int64, error)
创建新用户基本信息 CreateRolebase(tx *gorm.DB, Newuser commonstruct.Users) (commonstruct.Users, error)
更新用户VIP等级 UpdateUserlevel(uuid int64, newlevel int64) error
更新用户基本信息 UpdateUserbase(uuid int64, column string, value interface{}) error

通过ID获取用户信息 GetUserinfoByUuid(uuid int64) (commonstruct.Users, error)
通过账号获取用户信息 GetUserinfoByAccount(account string, masterid int64) (commonstruct.Users, error)
通过WXID 获取用户信息 GetUserinfoByWxid(wechat string) (commonstruct.Users, error)
修改用户推广码 SetExtcode(uuid int64, extcode string) error
绑定微信 BindWechat(username string, wechat string) error
判断账号是否存在 IsExistaccount(account string)  error

获取单日注册数 GetDateRegnumS(teamids []int64, date int64) ([]commonstruct.Users, error)
*/

// 查询荷官列表
func GetDealers(masterid int64) ([]commonstruct.Users, error) {
	var infos []commonstruct.Users
	if err := WyMysql.Table(commonstruct.WY_user_base).Where("master_id = ? and is_general = ?", masterid, commonstruct.Role_Dealer).Order("uuid").
		Find(&infos).Error; err != nil {
		beego.Error("GetDealers err", masterid, err)
		return infos, err
	}
	return infos, nil
}

// 查询多个用户的资料
func GetUserbaseS(idlist []int64) []commonstruct.Users {
	var list []commonstruct.Users
	if err := WyMysql.Table(commonstruct.WY_user_base).Where("uuid in (?)", idlist).Find(&list).Error; err != nil {
		beego.Error("GetUserbaseS err", idlist, err)
	}
	return list
}

// 查询子账号列表
func GetExpmanagers(preid int64, pagecount int, pagenum int) ([]commonstruct.Users, error) {
	var infos []commonstruct.Users
	if err := WyMysql.Table(commonstruct.WY_user_base).Where("pre_id = ? and role_type = 'expmanager'", preid).
		Order("uuid").Limit(pagecount).Offset((pagenum - 1) * pagecount).
		Find(&infos).Error; err != nil {
		beego.Error("GetExpmanagers err", preid, err)
		return infos, err
	}
	return infos, nil
}

// 查询子账号页码信息
func GetExpmanagersPageinfo(preid int64, pagecount int) (int64, int64) {
	var infos commonstruct.Users
	if err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").
		Where("pre_id = ? and is_general = ?", preid, commonstruct.Role_Assistant).
		Find(&infos).Error; err != nil {
		beego.Error("GetExpmanagers err", preid, err)
	}
	return int64(infos.Uuid), int64(math.Ceil(float64(infos.Uuid) / float64(pagecount)))
}

// 查询直属用户列表
func GetDirectRoleS(uuid int64, roletype string, limited string, sorttype string, offset int, pagecount int) ([]commonstruct.Users, error) {
	var list []commonstruct.Users

	var selectarg string
	switch roletype {
	case "all":
		switch limited {
		case "-1":
			selectarg = fmt.Sprintf("pre_id = %v", uuid)
		case "1,2":
			selectarg = fmt.Sprintf("pre_id = %v and limited in (1,2)", uuid)
		default:
			selectarg = fmt.Sprintf("pre_id = %v and limited = %v", uuid, limited)
		}
	default:
		switch limited {
		case "-1":
			selectarg = fmt.Sprintf("pre_id = %v and role_type = '%v'", uuid, roletype)
		case "1,2":
			selectarg = fmt.Sprintf("pre_id = %v and role_type = '%v' and limited in (1,2)", uuid, roletype)
		default:
			selectarg = fmt.Sprintf("pre_id = %v and role_type = '%v' and limited = %v", uuid, roletype, limited)
		}
	}

	var err error
	switch sorttype {
	case "createtime":
		err = WyMysql.Table(commonstruct.WY_user_base).Where(selectarg).
			Order("op_time desc,uuid").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&list).Error
	default:
		err = WyMysql.Table(commonstruct.WY_user_base).Where(selectarg).
			Order("last_logintime desc,uuid").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&list).Error
	}

	return list, err
}

// 查询直属用户页码信息
func GetDirectRoleSPageinfo(uuid int64, roletype string, limited string, pagecount int) (int64, int64) {
	var list commonstruct.Users
	var selectarg string
	switch roletype {
	case "all":
		switch limited {
		case "-1":
			selectarg = fmt.Sprintf("pre_id = %v", uuid)
		case "1,2":
			selectarg = fmt.Sprintf("pre_id = %v and limited in (1,2)", uuid)
		default:
			selectarg = fmt.Sprintf("pre_id = %v and limited = %v", uuid, limited)
		}
	default:
		switch limited {
		case "-1":
			selectarg = fmt.Sprintf("pre_id = %v and role_type = '%v'", uuid, roletype)
		case "1,2":
			selectarg = fmt.Sprintf("pre_id = %v and role_type = '%v' and limited in (1,2)", uuid, roletype)
		default:
			selectarg = fmt.Sprintf("pre_id = %v and role_type = '%v' and limited = %v", uuid, roletype, limited)
		}
	}

	if err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").Where(selectarg).Find(&list).Error; err != nil {
		beego.Error("GetDirectRoleSPageinfo err", uuid, err)
	}
	return int64(list.Uuid), int64(math.Ceil(float64(list.Uuid) / float64(pagecount)))
}

// 查询直属用户人数
func GetDirectNum(uuid int64, roletype string) (int64, error) {

	var selectarg string
	switch roletype {
	case "all":
		selectarg = fmt.Sprintf("pre_id = '%v'", uuid)
	default:
		selectarg = fmt.Sprintf("pre_id = '%v' and role_type = '%v'", uuid, roletype)
	}

	var info commonstruct.BranchUsers
	err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").
		Where(selectarg).Find(&info).Error

	if err != nil {
		beego.Error("GetDirectNum err = ", err.Error())
	}

	return info.Uuid, err
}

// 创建新用户基本信息
func CreateRolebase(tx *gorm.DB, Newuser commonstruct.Users) (commonstruct.Users, error) {

	newuserdb := tx.Table(commonstruct.WY_user_base).Create(&Newuser)
	if err := newuserdb.Error; err != nil {
		beego.Error("Create userbase err ", err)
		tx.Rollback()
		return Newuser, errors.New("创建role base表失败")
	}

	userinfo := newuserdb.Value.(*commonstruct.Users) // 新注册用户ID
	Newuser.Uuid = userinfo.Uuid
	return Newuser, nil
}

// 更新用户VIP等级
func UpdateUserlevel(uuid int64, newlevel int64) error {
	var updateValues map[string]interface{}

	updateValues = map[string]interface{}{
		"vip_level": newlevel,
		"up_time":   commonfunc.GetNowtime(),
	}

	if err := WyMysql.Table(commonstruct.WY_user_base).Where("uuid = ?", uuid).Update(updateValues).Error; err != nil {
		beego.Error("UpdateUserlevel err", err)
		return err
	}
	return nil
}

// 更新用户基本信息
func UpdateUserbase(uuid int64, column string, value interface{}) error {
	updateValues := map[string]interface{}{
		column: value,
	}

	if err := WyMysql.Table(commonstruct.WY_user_base).Where("uuid = ?", uuid).Update(updateValues).Error; err != nil {
		beego.Error("UpdateUserbase err ", uuid, updateValues, err)
		return err
	}
	return nil
}

// 批量更新用户基本信息
func UpdateUserbaseS(uuids []int64, column string, value interface{}) error {
	updateValues := map[string]interface{}{
		column: value,
	}

	if err := WyMysql.Table(commonstruct.WY_user_base).Where("uuid in (?)", uuids).Update(updateValues).Error; err != nil {
		beego.Error("UpdateUserbaseS err ", uuids, updateValues, err)
		return err
	}
	return nil
}

// 合计用户基本信息
func HejiUserbase(uuid int64, column string, value interface{}) error {
	updateValues := map[string]interface{}{
		column: gorm.Expr(fmt.Sprintf("%v + ?", column), value),
	}

	if err := WyMysql.Table(commonstruct.WY_user_base).Where("uuid = ?", uuid).Update(updateValues).Error; err != nil {
		return err
	}
	return nil
}

// 合计用户基本信息
func ResetUserErrpwdcount() error {
	if err := WyMysql.Table(commonstruct.WY_user_base).Update("err_pwd_count", 0).Error; err != nil {
		return err
	}
	return nil
}

// 通过ID获取用户信息
func GetUserinfoByUuid(uuid int64) (commonstruct.Users, error) {
	var user commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Where("uuid = ?", uuid).Find(&user).Error
	return user, err
}

// 获取所有试玩账号
func GetShiwanUserinfoS() ([]commonstruct.Users, error) {
	var user []commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Where("is_shiwan = 1").Find(&user).Error
	return user, err
}

// 通过账号获取用户信息
func GetUserinfoByAccount(straccount string, masterid int64) (commonstruct.Users, error) {
	account := strings.ToLower(straccount)
	var user commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Where("account = ? and (master_id = ? or master_id = 0)", account, masterid).Find(&user).Error
	if err != nil {
		beego.Error("GetUserinfoByAccount", account, masterid, err)
	}
	return user, err
}

// 通过WXID 获取用户信息
func GetUserinfoByWxid(wechat string) (commonstruct.Users, error) {
	var user commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Where("wechat = ?", wechat).Find(&user).Error
	if err != nil {
		beego.Error("GetUserinfoByWxid %v", err)
	}
	return user, err
}

// 修改用户推广码
func SetExtcode(uuid int64, extcode string) error {
	err := WyMysql.Table(commonstruct.WY_user_base).Where("uuid = ?", uuid).Update("ext_code", extcode).Error
	return err
}

// 绑定微信
func BindWechat(username string, wechat string) error {
	if err := WyMysql.Table(commonstruct.WY_user_base).Where("account = ?", username).Update("wechat", wechat).Error; err != nil {
		beego.Error("BindWechat err ", username, wechat, err)
		return err
	}
	return nil
}

// 获取单日注册数
func GetDateRegnum(teamids []int64, date int64) int64 {
	beego.Error(teamids, date, commonfunc.GetBegintime(date), commonfunc.GetEndtime(date))
	var users commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").Where("uuid in (?) and op_time between ? and ?", teamids, commonfunc.GetBegintime(date), commonfunc.GetEndtime(date)).Find(&users).Error
	if err != nil {
		beego.Error("GetDateRegnumS err", date, err)
	}
	return users.Uuid
}

// 获取单日登录数
func GetDateLoginnum(teamids []int64, date int64) int64 {

	var users commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").Where("uuid in (?) and last_logintime between ? and ?", teamids, commonfunc.GetBegintime(date), commonfunc.GetEndtime(date)).Find(&users).Error
	if err != nil {
		beego.Error("GetDateRegnumS err", date, err)
	}
	return users.Uuid
}

// 获取单日下单数
func GetDateOrdernum(teamids []int64, date int64) int64 {

	var users commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").Where("uuid in (?) and last_ordertime between ? and ?", teamids, commonfunc.GetBegintime(date), commonfunc.GetEndtime(date)).Find(&users).Error
	if err != nil {
		beego.Error("GetDateRegnumS err", date, err)
	}
	return users.Uuid
}

// 判断账号是否存在
func IsExistaccount(straccount string) (commonstruct.Users, error) {
	account := strings.ToLower(straccount)
	var user commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Where("account = ?", account).Find(&user).Error
	if err != nil {
		beego.Error("IsExistaccount err %v", err)
	}
	return user, err
}

/*
WY_company_base:

修改公司外接游戏列表 SetExpgamelist(uuid int64, gamelist string) error
通过ID获取公司基本信息 GetCompanybase(uuid int64) (commonstruct.CompanyBase, error)
新建公司基本信息 InitCompanyBase(uuid int64, account string, oddstype string, maxonlinenum int64, defaultinfo commonstruct.DefaultCompanyBase) error
更新公司基本信息 UpdateCompanybase(companyid int64, column string, value interface{}) error

通过账号获取公司基本信息 GetCompanybaseByAccount(account string) (commonstruct.CompanyBase, error)
获取所有公司基本信息 GetCompanybaseS() ([]commonstruct.CompanyBase, error)
通过ID获取公司基本信息 GetCompanybaseByid(uuid int64) (commonstruct.CompanyBase, error)

*/

// 修改公司外接游戏列表
func SetExpgamelist(uuid int64, gamelist string) error {
	if err := WyMysql.Table(commonstruct.WY_company_base).Where("uuid = ?", uuid).Update("expgame_list", gamelist).Error; err != nil {
		beego.Error("SetExpgamelist err", err)
		return err
	}
	return nil
}

// 新建公司基本信息
func InitCompanyBase(uuid int64, account string, nickname string, baseoddsper float64, maxonlinenum int64, defaultinfo commonstruct.DefaultCompanyBase, wallettype string) error {

	var baseinfo commonstruct.CompanyBase
	baseinfo.Uuid = uuid
	baseinfo.Account = account
	// baseinfo.RegOdds = defaultinfo.RegOdds
	// baseinfo.RegShui = defaultinfo.RegShuilt
	// baseinfo.MinOdds = defaultinfo.MinOdds
	// baseinfo.MaxOdds = defaultinfo.MaxOdds
	baseinfo.NameCN = nickname
	baseinfo.MinAmount = defaultinfo.MinAmount
	baseinfo.Baseoddsper = baseoddsper
	baseinfo.MaxOnlinenum = maxonlinenum
	baseinfo.InitPwd = defaultinfo.InitPwd
	baseinfo.InitCash = defaultinfo.InitCash
	baseinfo.FixedamountUrl = defaultinfo.FixedamountUrl
	baseinfo.RechargeUrl = defaultinfo.RechargeUrl
	baseinfo.CallbackUrl = defaultinfo.CallbackUrl
	baseinfo.SignType = defaultinfo.SignType
	baseinfo.FrontUrls = defaultinfo.FrontUrls
	baseinfo.BackUrls = defaultinfo.BackUrls
	baseinfo.PcShiwan = 2
	baseinfo.WalletSort = wallettype

	if err := WyMysql.Table(commonstruct.WY_company_base).Create(&baseinfo).Error; err != nil {
		beego.Error("InitCompanyBase err", baseinfo, err)
		return err
	}
	return nil
}

// 更新公司基本信息
func UpdateCompanybase(companyid int64, column string, value interface{}) error {
	updateValues := map[string]interface{}{
		column: value,
	}

	if err := WyMysql.Table(commonstruct.WY_company_base).Where("uuid = ?", companyid).Update(updateValues).Error; err != nil {
		beego.Error("UpdateCompanybase err ", companyid, updateValues, err)
		return err
	}
	return nil
}

// 更新公司基本信息
func UpdateCompanyDaohangkey(strcompanyaccount string, column string, value interface{}) error {
	companyaccount := strings.ToLower(strcompanyaccount)
	updateValues := map[string]interface{}{
		column: value,
	}

	if err := WyMysql_cokey.Table(commonstruct.Daohang_key_url).Where("company_account = ?", companyaccount).Update(updateValues).Error; err != nil {
		beego.Error("UpdateCompanyDaohangkey err ", companyaccount, updateValues, err)
		return err
	}
	return nil
}

// 整表修改公司基本信息
func UpdateCompanybaseS(newinfo commonstruct.CompanyBase) error {

	updateValues := map[string]interface{}{
		"name_cn":          newinfo.NameCN,
		"operate_state":    newinfo.OperateState,
		"update_zhancheng": newinfo.UpdateZhancheng,
		"front_urls":       newinfo.FrontUrls,
		"back_urls":        newinfo.BackUrls,
		"update_buhuo":     newinfo.UpdateBuhuo,
		"qiantai_logo":     newinfo.QiantaiLogo,
		"houtai_logo":      newinfo.HoutaiLogo,
		"qiantai_ico":      newinfo.QiantaiIco,
		"houtai_ico":       newinfo.HoutaiIco,
		"qiantai_title":    newinfo.QiantaiTitle,
		"houtai_title":     newinfo.HoutaiTitle,
		"pc_login_type":    newinfo.PcLoginType,
		"pc_guodu_type":    newinfo.PcGuoduType,
		"pc_game_type":     newinfo.PcGameType,
		"mbd_login_type":   newinfo.MbdLoginType,
		"mbd_game_type":    newinfo.MbdGameType,
		"adm_login_type":   newinfo.AdmLoginType,
		"adm_game_type":    newinfo.AdmGameType,
		"skin_type":        newinfo.SkinType,
		"pc_shiwan":        newinfo.PcShiwan,
		"zhudanyuming":     newinfo.Zhudanyuming,
		"expcodeurl1":      newinfo.Expcodeurl1,
		"expcodeurl2":      newinfo.Expcodeurl2,
		"expcodeurl3":      newinfo.Expcodeurl3,
		"expcodeurl4":      newinfo.Expcodeurl4,
		"expcodeurl5":      newinfo.Expcodeurl5,
		"expcodeurl6":      newinfo.Expcodeurl6,
	}

	if err := WyMysql.Table(commonstruct.WY_company_base).Where("uuid = ?", newinfo.Uuid).Update(updateValues).Error; err != nil {
		beego.Error("UpdateCompanybaseS err ", newinfo.Uuid, err)
		return err
	}
	return nil
}

// 通过账号获取公司基本信息
func GetCompanybaseByAccount(straccount string) (commonstruct.CompanyBase, error) {
	account := strings.ToLower(straccount)
	var datas commonstruct.CompanyBase
	err := WyMysql.Table(commonstruct.WY_company_base).Where("account = ?", account).Find(&datas).Error
	return datas, err
}

// 更新公司基本信息
func GetCompanyKeyinfoByAccount(straccount string) (commonstruct.KeyJumpurl, error) {
	account := strings.ToLower(straccount)
	var retinfo commonstruct.KeyJumpurl
	if err := WyMysql_cokey.Table(commonstruct.Daohang_key_url).Where("company_account = ?", account).Find(&retinfo).Error; err != nil {
		beego.Error("GetCompanyKeyinfo err ", account, err)
		return retinfo, err
	}
	return retinfo, nil
}

// 更新公司基本信息
func GetCompanyKeyinfoByKey(key string) (commonstruct.KeyJumpurl, error) {
	var retinfo commonstruct.KeyJumpurl
	if err := WyMysql_cokey.Table(commonstruct.Daohang_key_url).Where("daohang_key = ?", key).Find(&retinfo).Error; err != nil {
		beego.Error("GetCompanyKeyinfo err ", key, err)
		return retinfo, err
	}
	return retinfo, nil
}

// 通过urlkey 获取公司资料
func GetCompanybaseByUrlkey(urlkey string) (commonstruct.CompanyBase, error) {
	var datas commonstruct.CompanyBase
	err := WyMysql.Table(commonstruct.WY_company_base).Where("daohang_key = ?", urlkey).Find(&datas).Error
	return datas, err
}

// 获取所有公司基本信息
func GetCompanybaseS() ([]commonstruct.CompanyBase, error) {
	var infos []commonstruct.CompanyBase
	if err := WyMysql.Table(commonstruct.WY_company_base).Find(&infos).Error; err != nil {
		beego.Error("GetCompanybaseS err", err)
		return infos, err
	}
	return infos, nil
}

// 通过ID获取公司基本信息
func GetCompanybaseByid(uuid int64) (commonstruct.CompanyBase, error) {
	var info commonstruct.CompanyBase
	if uuid == 0 {
		return info, errors.New("uuid is zero!")
	}

	if err := WyMysql.Table(commonstruct.WY_company_base).Where("uuid = ?", uuid).Find(&info).Error; err != nil {
		beego.Error("GetCompanyodds err", uuid, err)
		return info, err
	}
	return info, nil
}

func UpsertChanglongoddsSet(uuid int64, roomid int64, updateinfos commonstruct.ChanglongOdds) error {

	var oldlog commonstruct.ChanglongOdds
	if retinfo := WyMysql.Table(commonstruct.WY_company_changlongodds).
		Where("uuid = ? and room_id = ?", uuid, roomid).Find(&oldlog); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_company_changlongodds).Create(&updateinfos).Error; err != nil {
				beego.Error("create err ", updateinfos, err)
				return err
			}
		} else {
			beego.Error("UpsertChanglongoddsSet err ", updateinfos, retinfo.Error.Error())
			return retinfo.Error
		}
		return nil
	}

	updateValues := map[string]interface{}{
		"liankai1":  updateinfos.Liankai1,
		"liankai2":  updateinfos.Liankai2,
		"liankai3":  updateinfos.Liankai3,
		"liankai4":  updateinfos.Liankai4,
		"liankai5":  updateinfos.Liankai5,
		"liankai6":  updateinfos.Liankai6,
		"liankai7":  updateinfos.Liankai7,
		"liankai8":  updateinfos.Liankai8,
		"liankai9":  updateinfos.Liankai9,
		"liankai10": updateinfos.Liankai10,
		"liankai11": updateinfos.Liankai11,
		"liankai12": updateinfos.Liankai12,
		"liankai13": updateinfos.Liankai13,
		"liankai14": updateinfos.Liankai14,
		"liankai15": updateinfos.Liankai15,
		"liankai16": updateinfos.Liankai16,
		"liankai17": updateinfos.Liankai17,
		"liankai18": updateinfos.Liankai18,
		"liankai19": updateinfos.Liankai19,
		"liankai20": updateinfos.Liankai20,

		"weikai1":  updateinfos.Weikai1,
		"weikai2":  updateinfos.Weikai2,
		"weikai3":  updateinfos.Weikai3,
		"weikai4":  updateinfos.Weikai4,
		"weikai5":  updateinfos.Weikai5,
		"weikai6":  updateinfos.Weikai6,
		"weikai7":  updateinfos.Weikai7,
		"weikai8":  updateinfos.Weikai8,
		"weikai9":  updateinfos.Weikai9,
		"weikai10": updateinfos.Weikai10,
		"weikai11": updateinfos.Weikai11,
		"weikai12": updateinfos.Weikai12,
		"weikai13": updateinfos.Weikai13,
		"weikai14": updateinfos.Weikai14,
		"weikai15": updateinfos.Weikai15,
		"weikai16": updateinfos.Weikai16,
		"weikai17": updateinfos.Weikai17,
		"weikai18": updateinfos.Weikai18,
		"weikai19": updateinfos.Weikai19,
		"weikai20": updateinfos.Weikai20,
	}

	if err := WyMysql.Table(commonstruct.WY_company_changlongodds).
		Where("uuid = ? and room_id = ?", uuid, roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateChanglongoddsSet err \n", uuid, roomid, updateinfos, err)
	}
	return nil
}

// 获取系统用户列表
func GetSysusers(pagenum int) ([]commonstruct.Users, error) {
	var infos []commonstruct.Users
	if err := WyMysql.Table(commonstruct.WY_user_base).
		Order("uuid").Limit(500).Offset((pagenum - 1) * 500).
		Find(&infos).Error; err != nil {
		beego.Error("GetSysusers err", err)
		return infos, err
	}
	return infos, nil
}

// 获取系统用户数
func GetSysusernum() (int64, int64) {
	var infos commonstruct.Users
	if err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").
		Find(&infos).Error; err != nil {
		beego.Error("GetSysusernum err", err)
	}
	return int64(infos.Uuid), int64(math.Ceil(float64(infos.Uuid) / float64(500)))
}

// 四舍五入
func Round4s5r(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	if f >= 0.5/pow10_n {
		return math.Trunc((f)*pow10_n+0.5) / pow10_n
	} else if f < 0.5/pow10_n && f > -0.5/pow10_n {
		return 0
	} else {
		return math.Trunc((f)*pow10_n-0.5) / pow10_n
	}
}

// 团队登录人数
func GetTeamLoginnum(teamidlist []int64, date int64) (commonstruct.Users, error) {
	begintime := commonfunc.GetBegintime(date)
	endtime := commonfunc.GetEndtime(date)

	var stats commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").
		Where("uuid in (?) and last_logintime between ? and ?", teamidlist, begintime, endtime).Find(&stats).Error
	if err != nil {
		beego.Error("GetTeamLoginnum err", err.Error())
	}
	return stats, nil
}

// 团队注册人数
func GetTeamRegnum(teamidlist []int64, date int64) (commonstruct.Users, error) {
	begintime := commonfunc.GetBegintime(date)
	endtime := commonfunc.GetEndtime(date)

	var stats commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").
		Where("uuid in (?) and op_time between ? and ?", teamidlist, begintime, endtime).Find(&stats).Error
	if err != nil {
		beego.Error("GetTeamRegnum err", err.Error())
	}
	return stats, nil
}

// 团队下注人数
func GetTeamOrderusernum(teamidlist []int64, date int64) (commonstruct.Users, error) {
	begintime := commonfunc.GetBegintime(date)
	endtime := commonfunc.GetEndtime(date)

	var stats commonstruct.Users
	err := WyMysql.Table(commonstruct.WY_user_base).Select("count(*) as uuid").
		Where("uuid in (?) and last_ordertime between ? and ?", teamidlist, begintime, endtime).Find(&stats).Error
	if err != nil {
		beego.Error("GetTeamOrderusernum err", err.Error())
	}
	return stats, nil
}

// 获取DB sessioninfo
func GetUserSession(uuid int64) (commonstruct.SessionValue, error) {
	var info commonstruct.SessionValue
	err := WyMysql.Table(commonstruct.WY_tmp_session).Where("uuid = ?", uuid).Find(&info).Error
	if err != nil {
		beego.Error("GetUserSession err", uuid, err.Error())
	}
	return info, nil
}

// 获取DB sessioninfo
func GetUserSessionBySid(sid string) (commonstruct.SessionValue, error) {
	var info commonstruct.SessionValue
	err := WyMysql.Table(commonstruct.WY_tmp_session).Where("s_id = ?", sid).Find(&info).Error
	if err != nil {
		beego.Error("GetUserSession err", err.Error())
	}
	return info, nil
}

func GetUserSessionS() ([]commonstruct.SessionValue, error) {
	var infos []commonstruct.SessionValue
	err := WyMysql.Table(commonstruct.WY_tmp_session).Find(&infos).Error
	if err != nil {
		beego.Error("GetUserSessionS err", err.Error())
	}
	return infos, nil
}

func UpsertSession(sessioninfo commonstruct.SessionValue) error {
	var Dyncinfo commonstruct.SessionValue
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_session).
		Find(&Dyncinfo, "uuid = ?", sessioninfo.Uuid); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_tmp_session).Create(&sessioninfo).Error; err != nil {
				beego.Error("UpsertSession err ", sessioninfo, err)
				return nil
			}
		} else {
			beego.Error("UpsertSession err \n", sessioninfo, retinfo.Error.Error())
			return retinfo.Error
		}
	}

	updateValues := map[string]interface{}{
		"s_id":      sessioninfo.SID,
		"account":   sessioninfo.Account,
		"expiry":    sessioninfo.Expiry,
		"ip":        sessioninfo.IP,
		"ip_place":  sessioninfo.IPPlace,
		"is_mobile": sessioninfo.IsMobile,
		"logintime": sessioninfo.Logintime,
		"ordertime": sessioninfo.Ordertime,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_session).
		Where("uuid = ?", sessioninfo.Uuid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsertSession err", sessioninfo, err)
	}
	return nil
}

func UpdateSession(uuid int64, sessioninfo commonstruct.SessionValue, column string) error {
	var updateValues map[string]interface{}
	switch column {
	case "ordertime":
		updateValues = map[string]interface{}{
			"ordertime": sessioninfo.Ordertime,
		}
	case "expiry":
		updateValues = map[string]interface{}{
			"expiry": sessioninfo.Expiry,
		}
	default:
		beego.Error("unknown column", column)
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_session).
		Where("uuid = ?", sessioninfo.Uuid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsertSession err", sessioninfo, err)
	}
	return nil
}

// 删除在线用户信息
func DeleteSession(uuid int64) error {
	sql := fmt.Sprintf("delete from %v where uuid = %v;", commonstruct.WY_tmp_session, uuid)
	if err := WyMysql.Exec(sql).Error; err != nil {
		beego.Error("DeleteSession err ", err)
	}
	return nil
}

// 删除在线用户信息
func DeleteSessionBySID(sid string) error {
	sql := fmt.Sprintf("delete from %v where s_id = %v;", commonstruct.WY_tmp_session, sid)
	if err := WyMysql.Exec(sql).Error; err != nil {
		beego.Error("DeleteSessionBySID err ", err)
	}
	return nil
}

// 获取需要执行的风控检测
func GetInuseFengkongSetS() ([]commonstruct.FengkongSet, error) {
	var infos []commonstruct.FengkongSet
	err := WyMysql.Table(commonstruct.WY_tmp_company_fengkongset).Where("qiuhao_inuse = 1").Find(&infos).Error
	if err != nil {
		beego.Error("GetInuseFengkongSetS err", err.Error())
	}
	return infos, nil
}

func GetFengkongSet(companyid int64, roomid int64) (commonstruct.FengkongSet, error) {
	var infos commonstruct.FengkongSet
	err := WyMysql.Table(commonstruct.WY_tmp_company_fengkongset).Where("company_id = ? and room_id = ?", companyid, roomid).Find(&infos).Error
	if err != nil {
		beego.Error("GetFengkongSet err", err.Error())
	}
	return infos, nil
}

func UpsertFengkongSet(companyid int64, roomid int64, updateinfo commonstruct.FengkongSet) error {

	var oldinfo commonstruct.FengkongSet
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_company_fengkongset).
		Find(&oldinfo, "company_id = ? and room_id = ?", companyid, roomid); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_tmp_company_fengkongset).Create(&updateinfo).Error; err != nil {
				beego.Error("UpsertFengkongSet err ", updateinfo, err)
				return nil
			}
		} else {
			beego.Error("UpsertFengkongSet err \n", updateinfo, retinfo.Error.Error())
			return retinfo.Error
		}
	}

	updateValues := map[string]interface{}{
		"qiuhao_inuse":  updateinfo.QiuhaoInuse,
		"qiuhao_num":    updateinfo.QiuhaoNum,
		"qiuhao_amount": updateinfo.QiuhaoAmount,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_company_fengkongset).
		Where("company_id = ? and room_id = ?", companyid, roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateFengkongSet err", updateinfo, err)
	}
	return nil
}

func UpsertFengkonglog(newlog commonstruct.FengkongLog) error {

	var oldinfo commonstruct.FengkongLog
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_company_fengkonglog).
		Find(&oldinfo, "uuid = ? and room_id = ? and expect = ?", newlog.Uuid, newlog.RoomID, newlog.Expect); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_tmp_company_fengkonglog).Create(&newlog).Error; err != nil {
				beego.Error("UpsertFengkonglog err ", newlog, err)
				return nil
			}
		} else {
			beego.Error("UpsertFengkonglog err \n", newlog, retinfo.Error.Error())
			return retinfo.Error
		}
	}
	return errors.New("record is exist!")
}

func GetFengkonglogS(companyid int64) ([]commonstruct.FengkongLog, error) {

	var RetS []commonstruct.FengkongLog

	err := WyMysql.Table(commonstruct.WY_tmp_company_fengkonglog).Where("company_id = ?", companyid).Order("id desc").Find(&RetS).Error
	if err != nil {
		beego.Error("GetFengkonglogS err", err.Error())
	}

	return RetS, nil
}

func GetFengkongcount(uuid int64) int64 {
	var RetS commonstruct.FengkongLog
	err := WyMysql.Table(commonstruct.WY_tmp_company_fengkonglog).Select("count(*) as id").Where("uuid = ?", uuid).Find(&RetS).Error
	if err != nil {
		// beego.Error("GetFengkongcount err", err.Error())
		return 0
	}
	return RetS.ID
}

func GetFengkongmarkLog(uuid int64) (commonstruct.FengkongmarkLog, error) {
	var RetS commonstruct.FengkongmarkLog
	err := WyMysql.Table(commonstruct.WY_tmp_user_fengkongmark_log).Where("uuid = ? and mark_date = ?", uuid, commonfunc.GetNowdate()).Find(&RetS).Error
	return RetS, err
}

func UpsertFengkongmarklog(newlog commonstruct.FengkongmarkLog) error {
	var oldinfo commonstruct.FengkongmarkLog
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_user_fengkongmark_log).
		Find(&oldinfo, "uuid = ? and mark_date = ?", newlog.Uuid, newlog.MarkDate); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_tmp_user_fengkongmark_log).Create(&newlog).Error; err != nil {
				beego.Error("UpsertFengkongmarklog err ", newlog, err)
				return nil
			}
		} else {
			beego.Error("UpsertFengkongmarklog err \n", newlog, retinfo.Error.Error())
			return retinfo.Error
		}
	}
	return errors.New("record is exist!")
}

func CreateGentouLog(newlog commonstruct.GentouLog) error {
	if err := WyMysql.Table(commonstruct.WY_tmp_log_gentou).Create(&newlog).Error; err != nil {
		beego.Error("CreateGentouLog err ", newlog, err)
		return err
	}
	return nil
}

// func CreateGentouPlan(newlog commonstruct.GentouPlan) error {

// 	if err := WyMysql.Table(commonstruct.WY_tmp_user_gentouplan).Create(&newlog).Error; err != nil {
// 		beego.Error("CreateGentouPlan err ", newlog, err)
// 		return err
// 	}

// 	return nil
// }

func UpsertGentouPlan(newlog commonstruct.GentouPlan) error {

	if newlog.PlanID == 0 {
		if err := WyMysql.Table(commonstruct.WY_tmp_user_gentouplan).Create(&newlog).Error; err != nil {
			beego.Error("CreateGentouPlan err ", newlog, err)
			return err
		} else {
			return nil
		}
	}

	var oldinfo commonstruct.GentouPlan
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_user_gentouplan).
		Find(&oldinfo, "plan_id = ?", newlog.PlanID); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(commonstruct.WY_tmp_user_gentouplan).Create(&newlog).Error; err != nil {
				beego.Error("CreateGentouPlan err ", newlog, err)
				return err
			}
		} else {
			beego.Error("UpsertGentouPlan err \n", newlog, retinfo.Error.Error())
			return retinfo.Error
		}
	}

	updateValues := map[string]interface{}{
		"gentou_uuid":         newlog.GentouUuid,
		"gentou_account":      newlog.GentouAccount,
		"gentou_game_list":    newlog.GentouGameList,
		"gentou_type":         newlog.GentouType,
		"gentou_percent":      newlog.GentouPercent,
		"touzhu_accountinfo1": newlog.TouzhuAccountinfo1,
		"touzhu_accountinfo2": newlog.TouzhuAccountinfo2,
		"touzhu_accountinfo3": newlog.TouzhuAccountinfo3,
		"touzhu_accountinfo4": newlog.TouzhuAccountinfo4,
		"touzhu_accountinfo5": newlog.TouzhuAccountinfo5,
		"touzhu_accountinfo6": newlog.TouzhuAccountinfo6,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_gentouplan).
		Where("plan_id = ?", newlog.PlanID).
		Update(updateValues).Error; err != nil {
		beego.Error("UpsertGentouPlan err", newlog, err)
	}
	return nil
}

// 获取代理的跟投计划
func GetGentouplanS(uuid int64) ([]commonstruct.GentouPlan, error) {

	var RetS []commonstruct.GentouPlan

	err := WyMysql.Table(commonstruct.WY_tmp_user_gentouplan).Where("uuid = ?", uuid).Order("plan_id desc").Find(&RetS).Error
	if err != nil {
		beego.Error("GetGentouplanS err", err.Error())
	}

	return RetS, nil
}

func DeleteGentouplan(planid int64) error {
	if err := WyMysql.Table(commonstruct.WY_tmp_user_gentouplan).Delete(nil, "plan_id = ?", planid).Error; err != nil {
		beego.Error("DeleteGentouplan err", planid, err)
		return errors.New(fmt.Sprintf("删除%v失败", planid))
	}
	return nil
}

// 获取用户的被跟投计划
func GetBeigentouplanS(uuid int64) ([]commonstruct.GentouPlan, error) {
	var RetS []commonstruct.GentouPlan
	if err := WyMysql.Table(commonstruct.WY_tmp_user_gentouplan).Where("gentou_uuid = ?", uuid).Find(&RetS).Error; err != nil {
		beego.Error("GetBeigentouplanS err ", err)
		return nil, err
	}
	return RetS, nil
}

// 获取上级跟用户的计划
func GetGentouplanByBothuuid(higherid int64, beigentou_uuid int64) (commonstruct.GentouPlan, error) {
	var RetS commonstruct.GentouPlan
	err := WyMysql.Table(commonstruct.WY_tmp_user_gentouplan).Where("uuid = ? and gentou_uuid", higherid, beigentou_uuid).Find(&RetS).Error
	if err != nil {
		beego.Error("GetBeigentouplanS err ", err)
		return RetS, err
	}
	return RetS, nil
}

// 增加撤单日志
func AddRevokelog(uuid int64, orderid int64) error {
	var oplog commonstruct.RevokeLog
	oplog.OrderID = orderid
	oplog.Uuid = uuid
	oplog.OpTime = commonfunc.GetNowtime()
	if err := WyMysql.Table(commonstruct.WY_tmp_log_revoke).Create(&oplog).Error; err != nil {
		beego.Error("err %v", err)
		return err
	}
	return nil
}
