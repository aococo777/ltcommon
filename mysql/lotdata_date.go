package wymysql

import (
	// "encoding/json"
	"github.com/aococo777/ltcommon/commonfunc"
	"github.com/aococo777/ltcommon/commonstruct"
	"math"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
WY_tmp_user_lotdata_data:

新增用户统计数据 InitProfit(uuid int64, preid int64, masterid int64, roominfo commonstruct.CRoomInfo, expect string)
修改用户日+游戏统计 UpsertUserlotdataDate(newdata commonstruct.LotdataStatistic)

获取多用户所有游戏的统计 GetAgentTeamstatistic(teamlist []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic
获取多用户多个游戏的统计  GetTeamstatistic(tblname string, sufids []int64, roomids []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic
获取多用户所有游戏统计 GetSufStatisticByDate(sufids []int64, begindate int64, enddate int64) []commonstruct.LotdataStatistic
获取多用户所有游戏按日统计 GetTeamStatisticByDate(teamlist []int64, begindate int64, enddate int64) []commonstruct.LotdataStatistic
获取多用户多日期统计 GetTeamLtstatistic(sufids []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic

获取代理线单个游戏的下注用户列表 GetAgentGameuserS(uuid int64, roomid int64, begindate int64, enddate int64) []int64
获取代理线多个游戏的下注用户列表 GetAgentMultigameuserS(uuid int64, roomids []int64, begindate int64, enddate int64, offset int, pagecount int) []int64
获取代理线多个游戏的下注用户页码信息 GetAgentMultigameusersPageinfo(uuid int64, roomids []int64, begindate int64, enddate int64, pagecount int) (int64, int64)

获取单用户的统计 GetLtprofitStats(uuid int64, begindate int64, enddate int64) commonstruct.LotdataStatistic
获取单用户按游戏合计 GetUserLotStatisticGroupbyRoomid(uuid int64, begindate int64, enddate int64) []commonstruct.LotdataStatistic
获取单用户输赢统计 GetUserLotStatistic(uuid int64, begindate int64, enddate int64) commonstruct.LotdataStatistic
获取单用户多游戏按游戏统计 GetUserGamestatisticS(uuid int64, roomids []int64, begindate int64, enddate int64) ([]commonstruct.LotdataStatistic, error)
获取单用户所有游戏合计 GetSelfStatistic(sufids int64, tblname string, begindate int64, enddate int64) commonstruct.LotdataStatistic
获取单用户所有游戏按日统计 GetUserDaystatisticS(uuid int64, begindate int64, enddate int64) ([]commonstruct.LotdataStatistic, error)
获取单用户单下注玩法统计 GetUserlotdataByPortid(uuid int64, roomid int64, expect string, portid int64) commonstruct.LotdataUseritem
获取单用户单游戏单期统计 GetUserlotdataByExpect(uuid int64, roomid int64, expect string) commonstruct.LotdataUseritem

获取公司的周活跃用户输赢统计 GetWeekActiveusers(masterid int64) ([]commonstruct.LotdataStatistic, error)

保存用户的结算数据 SaveLotdataStatistic(tx *gorm.DB, roomid int64, stats commonstruct.ExpectStatistic) error
撤单修改统计 UpsertRevokeamount(uuid int64, preid int64, masterid int64, roomid int64, expect string, revokeamount float64)
清空单期结算数据 DeleteProfit(roomid int64, expect string)
*/

// 获取代理线的统计
func GetAgentTeamstatistic(teamlist []int64, roomids []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic {
	var statistic commonstruct.LotdataStatistic
	if len(teamlist) <= 0 {
		return statistic
	}

	selectinfo := `
		sum(order_num) as order_num,
		sum(order_amount) as order_amount,
		sum(settled_num) as settled_num,
		sum(settled_amount) as settled_amount,
		sum(valid_amount) as valid_amount,
		sum(wager) as wager,
		sum(profit_wager) as profit_wager,
		sum(tuishui) as tuishui,
		sum(profit_tuishui) as profit_tuishui`

	if roomids == nil {
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
			Where("uuid in (?) and date between ? and ? ", teamlist, begindate, enddate).
			Find(&statistic).Error; err != nil {
			beego.Error("GetTeamStatistic err", teamlist, begindate, enddate, err)
		}
	} else {
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
			Where("uuid in (?) and room_id in (?) and date between ? and ? ", teamlist, roomids, begindate, enddate).
			Find(&statistic).Error; err != nil {
			beego.Error("GetTeamStatistic err", teamlist, begindate, enddate, err)
		}
	}
	return statistic
}

// 获取代理下所有下注的用户数量
func GetAgentTeamOrderusernum(teamlist []int64, roomids []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic {
	var statistic commonstruct.LotdataStatistic
	if len(teamlist) <= 0 {
		return statistic
	}

	selectinfo := `count(distinct uuid) as uuid`

	if roomids == nil {
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
			Where("uuid in (?) and date between ? and ? and order_num > 0", teamlist, begindate, enddate).
			Find(&statistic).Error; err != nil {
			beego.Error("GetTeamStatistic err", teamlist, begindate, enddate, err)
		}
	} else {
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
			Where("uuid in (?) and room_id in (?) and date between ? and ? and order_num > 0", teamlist, roomids, begindate, enddate).
			Find(&statistic).Error; err != nil {
			beego.Error("GetTeamStatistic err", teamlist, begindate, enddate, err)
		}
	}
	return statistic
}

// 获取代理下所有下注订单数
func GetAgentTeamOrdernum(teamlist []int64, roomids []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic {
	var statistic commonstruct.LotdataStatistic
	if len(teamlist) <= 0 {
		return statistic
	}

	selectinfo := `sum(order_num) as order_num`

	if roomids == nil {
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
			Where("uuid in (?) and date between ? and ?", teamlist, begindate, enddate).
			Find(&statistic).Error; err != nil {
			beego.Error("GetTeamStatistic err", teamlist, begindate, enddate, err)
		}
	} else {
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
			Where("uuid in (?) and room_id in (?) and date between ? and ?", teamlist, roomids, begindate, enddate).
			Find(&statistic).Error; err != nil {
			beego.Error("GetTeamStatistic err", teamlist, begindate, enddate, err)
		}
	}
	return statistic
}

// 获取多个用户多个游戏的统计
func GetTeamstatistic(tblname string, sufids []int64, roomids []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic {
	var retinfo commonstruct.LotdataStatistic
	if len(sufids) <= 0 {
		return retinfo
	}

	selectinfo := `uuid,
	sum(order_num) as order_num,
	sum(order_amount) as order_amount,
	sum(settled_num) as settled_num,
	sum(settled_amount) as settled_amount,
	sum(valid_amount) as valid_amount,
	sum(revoke_amount) as revoke_amount,
	sum(wager) as wager,
	sum(profit_wager) as profit_wager,
	sum(tuishui) as tuishui,
	sum(profit_tuishui) as profit_tuishui,
	sum(buchu) as buchu,
	sum(buchuyingkui) as buchuyingkui,
	sum(shizhanbuhuo) as shizhanbuhuo,
	sum(shizhanbuhuoyingkui) as shizhanbuhuoyingkui`

	if err := WyMysql_read.Table(tblname).Select(selectinfo).
		Where("uuid in (?) and room_id in (?) and date between ? and ? ", sufids, roomids, begindate, enddate).
		Group("uuid").Find(&retinfo).Error; err != nil {
		beego.Error("GetTeamstatisticByDate err", begindate, enddate, err)
	}

	return retinfo
}

// 获取代理线单个游戏的下注用户列表
func GetAgentGameuserS(uuid int64, roomid int64, begindate int64, enddate int64) []int64 {
	var list []commonstruct.DayStatistic
	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select("distinct uuid").
		Where("pre_id = ? and room_id = ? and date between ? and ?", uuid, roomid, begindate, enddate).
		Order("uuid desc").Find(&list).Error; err != nil {
		beego.Error("GetAgentGameuserS err", uuid, err)
	}

	var ret []int64
	for _, user := range list {
		if user.Uuid != commonstruct.SYSTEMID {
			ret = append(ret, user.Uuid)
		}
	}
	return ret
}

// 获取单个用户的统计 roomids 可以为nil
func GetLtprofitStats(uuid int64, roomids []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic {

	selectinfo := `
	sum(order_num) as order_num,
	sum(order_amount) as order_amount,
	sum(settled_num) as settled_num,
	sum(settled_amount) as settled_amount,
	sum(valid_amount) as valid_amount,
	sum(revoke_amount) as revoke_amount,
	sum(wager) as wager,
	sum(profit_wager) as profit_wager,
	sum(tuishui) as tuishui,
	sum(profit_tuishui) as profit_tuishui,
	sum(yingshouxiaxian) as yingshouxiaxian,
	sum(shizhanhuoliang) as shizhanhuoliang,
	sum(shizhanshuying) as shizhanshuying,
	sum(shizhanjieguo) as shizhanjieguo,
	sum(shizhantuishui) as shizhantuishui,
	sum(yingkuijieguo) as yingkuijieguo,
	sum(shangjiaohuoliang) as shangjiaohuoliang,
	sum(shangjijiaoshou) as shangjijiaoshou,
	sum(shouhuo) as shouhuo,
	sum(shouhuoyingkui) as shouhuoyingkui,
	sum(buchu) as buchu,
	sum(buchuyingkui) as buchuyingkui,
	sum(settle_shoudongbuchu) as settle_shoudongbuchu,
	sum(settle_shoudongbuchuyingkui) as settle_shoudongbuchuyingkui`

	var ret commonstruct.LotdataStatistic

	if roomids == nil {
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
			Where("uuid = ? and date between ? and ?", uuid, begindate, enddate).
			Find(&ret).Error; err != nil {
			beego.Error("GetLtprofitStats err ", uuid, begindate, enddate, err)
		}
	} else {
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
			Where("uuid = ? and room_id in (?) and date between ? and ?", uuid, roomids, begindate, enddate).
			Find(&ret).Error; err != nil {
			beego.Error("GetLtprofitStats err ", uuid, begindate, enddate, err)
		}
	}

	return ret
}

// 获取代理线多个游戏的下注用户列表
func GetAgentMultigameuserS(uuid int64, roomids []int64, begindate int64, enddate int64, offset int, pagecount int) []int64 {
	var list []commonstruct.DayStatistic

	if roomids == nil { // 用于算出 今日输赢
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select("distinct uuid").
			Where("pre_id = ? and date between ? and ?", uuid, begindate, enddate).
			Order("uuid desc").Limit(pagecount).Offset((offset - 1) * pagecount).
			Find(&list).Error; err != nil {
			beego.Error("GetAgentMultigameuserS err", uuid, err)
		}
	} else {
		if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select("distinct uuid").
			Where("pre_id = ? and room_id in(?) and date between ? and ?", uuid, roomids, begindate, enddate).
			Order("uuid desc").Limit(pagecount).Offset((offset - 1) * pagecount).
			Find(&list).Error; err != nil {
			beego.Error("GetAgentMultigameuserS err", uuid, err)
		}
	}

	var ret []int64

	for _, user := range list {
		if user.Uuid != commonstruct.SYSTEMID {
			ret = append(ret, user.Uuid)
		}
	}
	return ret
}

// 获取代理线多个游戏的下注用户页码信息
func GetAgentMultigameusersPageinfo(uuid int64, roomids []int64, begindate int64, enddate int64, pagecount int) (int64, int64) {
	var list commonstruct.DayStatistic
	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select("count(distinct uuid) as uuid").
		Where("pre_id = ? and room_id in(?) and date between ? and ? and uuid != ? and uuid != ?", uuid, roomids, begindate, enddate, uuid, commonstruct.SYSTEMID).
		Find(&list).Error; err != nil {
		beego.Error("GetSufIDList err", uuid, err)
	}

	return int64(list.Uuid), int64(math.Ceil(float64(list.Uuid) / float64(pagecount)))
}

// 获取多个用户所有游戏统计
func GetSufStatisticByDate(sufids []int64, begindate int64, enddate int64) []commonstruct.LotdataStatistic {
	var list []commonstruct.LotdataStatistic
	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Where("uuid in (?) and date between ? and ?", sufids, begindate, enddate).
		Find(&list).Error; err != nil {
		beego.Error("GetSufStatisticByDate err", begindate, enddate, err)
	}
	return list
}

// 修改用户日+游戏 统计
func UpsertUserlotdataDate(newdata commonstruct.LotdataStatistic) {

	newdata.Date = GetBilishiDate(0)

	var result commonstruct.LotdataStatistic
	tblname := commonstruct.WY_tmp_user_lotdata_data
	if retinfo := WyMysql.Table(tblname).
		Where("uuid = ? and date = ? and room_id = ?", newdata.Uuid, newdata.Date, newdata.RoomID).Find(&result); retinfo.Error != nil {
		if retinfo.RecordNotFound() {
			if err := WyMysql.Table(tblname).Create(&newdata).Error; err != nil {
				beego.Error("Create err ", result, err)
			}
		} else {
			beego.Error("UpsetUserlotdataDate err \n", newdata.Uuid, newdata.Date, newdata.RoomID, retinfo.Error.Error())
		}
		return
	}

	updateValues := map[string]interface{}{
		"order_num":       gorm.Expr("order_num + ?", newdata.OrderNum),
		"order_amount":    gorm.Expr("order_amount + ?", newdata.OrderAmount),
		"shizhanhuoliang": gorm.Expr("shizhanhuoliang + ?", newdata.Shizhanhuoliang),
		"shoudongbuchu":   gorm.Expr("shoudongbuchu + ?", newdata.Shoudongbuchu),
	}

	if err := WyMysql.Table(tblname).Where("uuid = ? and date = ? and room_id = ?", newdata.Uuid, newdata.Date, newdata.RoomID).Update(updateValues).Error; err != nil {
		beego.Error("UpsetUserlotdataDate err ", newdata.Uuid, newdata.Date, newdata.RoomID, newdata.OrderNum, newdata.OrderAmount, err)
	}
}

// 修改撤单金额
func UpsertRevokeamount(uuid int64, preid int64, masterid int64, roomid int64, expect string, revokeamount float64) {
	date := commonfunc.GetNowdate()
	updateValues := map[string]interface{}{
		"revoke_amount": gorm.Expr("revoke_amount + ?", revokeamount),
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_lotdata_data).Where("uuid = ? and room_id = ? and expect = ?", uuid, roomid, expect).Update(updateValues).Error; err != nil {
		beego.Error("UpsertRevokeamount err ", uuid, date, revokeamount, err)
	}
}

// 多用户所有游戏按日统计
func GetTeamStatisticByDate(teamlist []int64, begindate int64, enddate int64) []commonstruct.LotdataStatistic {
	var statistic []commonstruct.LotdataStatistic
	if len(teamlist) <= 0 {
		return statistic
	}

	selectinfo := `
	date,
	sum(order_amount) as order_amount,
	sum(order_num) as order_num,
	sum(settled_amount) as settled_amount,
	sum(settled_num) as settled_num,
	sum(valid_amount) as valid_amount,
	sum(wager) as wager,
	sum(profit_wager) as profit_wager,
	sum(tuishui) as tuishui,
	sum(profit_tuishui) as profit_tuishui`

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
		Where("uuid in (?) and date between ? and ?", teamlist, begindate, enddate).Group("date").
		Find(&statistic).Error; err != nil {
		beego.Error("GetTeamStatisticByDate err", teamlist, begindate, enddate, err)
	}
	return statistic
}

// 单用户游戏单期合计
func GetUserGameStatisticByExpect(uuid int64, roomid int64, expect string) commonstruct.LotdataUseritem {
	var list commonstruct.LotdataUseritem

	selectarg := `sum(order_amount) as order_amount,
	sum(order_num) as order_num,
	sum(valid_amount) as valid_amount,
	sum(wager) as wager,
	sum(tuishui) as tuishui`

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).Select(selectarg).
		Where("uuid = ? and room_id = ? and expect = ?", uuid, roomid, expect).
		Find(&list).Error; err != nil {
		beego.Error("GetUserGameStatisticByExpect err", err.Error())
	}

	list.Uuid = uuid
	list.RoomID = roomid
	list.Expect = expect

	return list
}

// 单用户按游戏合计
func GetUserLotStatisticGroupbyRoomid(uuid int64, begindate int64, enddate int64) []commonstruct.LotdataStatistic {
	var list []commonstruct.LotdataStatistic

	selectarg := `room_id,
	sum(order_amount) as order_amount,
	sum(order_num) as order_num,
	sum(valid_amount) as valid_amount,
	sum(wager) as wager,
	sum(tuishui) as tuishui`

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectarg).
		Where("uuid = ? and date between ? and ?", uuid, begindate, enddate).
		Group("room_id").Find(&list).Error; err != nil {
		beego.Error("GetRechargeNavidetail err", err.Error())
	}

	for i := 0; i < len(list); i++ {
		list[i].Uuid = uuid
	}

	return list
}

// 获取单个用户输赢统计
func GetUserLotStatistic(uuid int64, begindate int64, enddate int64) commonstruct.LotdataStatistic {
	var ret commonstruct.LotdataStatistic
	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).
		Select("sum(wager) as wager,sum(tuishui) as tuishui,sum(settled_amount) as settled_amount").
		Where("uuid = ? and date between ? and ?", uuid, begindate, enddate).
		Find(&ret).Error; err != nil {
		beego.Error("GetUserLotStatistic err ", uuid, begindate, enddate, err)
	}
	return ret
}

// 查询用户多游戏按游戏统计
func GetUserGamestatisticS(uuid int64, roomids []int64, begindate int64, enddate int64) ([]commonstruct.LotdataStatistic, error) {
	var StatisticsS []commonstruct.LotdataStatistic

	selectarg := `uuid,room_id,room_eng,room_cn,
	sum(order_amount) as order_amount,
	sum(order_num) as order_num,
	sum(settled_amount) as settled_amount,
	sum(settled_num) as settled_num,
	sum(valid_amount) as valid_amount,
	sum(wager) as wager,
	sum(profit_wager) as profit_wager,
	sum(tuishui) as tuishui,
	sum(profit_tuishui) as profit_tuishui`

	var err error
	if len(roomids) > 0 {
		err = WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectarg).
			Where("uuid = ? and date between ? and ? and room_id in (?)", uuid, begindate, enddate, roomids).
			Group("uuid,room_id,room_eng,room_cn").
			Order("room_id").Find(&StatisticsS).Error
	} else {
		err = WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectarg).
			Where("uuid = ? and date between ? and ?", uuid, begindate, enddate).
			Group("uuid,room_id,room_eng,room_cn").
			Order("room_id").Find(&StatisticsS).Error

	}

	if err != nil {
		beego.Error("GetUserStatisticsByDate err ", err)
	}
	return StatisticsS, err
}

// 单用户所有游戏合计
func GetSelfStatistic(sufids int64, tblname string, begindate int64, enddate int64) commonstruct.LotdataStatistic {
	var list commonstruct.LotdataStatistic

	selectinfo := `
		sum(order_num) as order_num,
		sum(order_amount) as order_amount,
		sum(settled_num) as settled_num,
		sum(settled_amount) as settled_amount,
		sum(valid_amount) as valid_amount,
		sum(wager) as wager,
		sum(profit_wager) as profit_wager,
		sum(tuishui) as tuishui,
		sum(profit_tuishui) as profit_tuishui,
		sum(yingshouxiaxian) as yingshouxiaxian,
		sum(shizhanhuoliang) as shizhanhuoliang,
		sum(shizhanshuying) as shizhanshuying,
		sum(shizhanjieguo) as shizhanjieguo,
		sum(shizhantuishui) as shizhantuishui,
		sum(shizhanpeicha) as shizhanpeicha,
		sum(yingkuijieguo) as yingkuijieguo,
		sum(shangjiaohuoliang) as shangjiaohuoliang,
		sum(shangjijiaoshou) as shangjijiaoshou`

	if err := WyMysql_read.Table(tblname).
		Select(selectinfo).
		Where("uuid = ? and date between ? and ?", sufids, begindate, enddate).
		Find(&list).Error; err != nil {
		beego.Error("GetSufStatisticByDate err", begindate, enddate, err)
	}
	return list
}

// 单用户所有游戏按日统计
func GetUserDaystatisticS(uuid int64, begindate int64, enddate int64) ([]commonstruct.LotdataStatistic, error) {
	var StatisticsS []commonstruct.LotdataStatistic

	selectarg := `uuid,date,
	sum(order_amount) as order_amount,
	sum(order_num) as order_num,
	sum(settled_amount) as settled_amount,
	sum(settled_num) as settled_num,
	sum(valid_amount) as valid_amount,
	sum(wager) as wager,
	sum(profit_wager) as profit_wager,
	sum(tuishui) as tuishui,
	sum(profit_tuishui) as profit_tuishui`

	err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectarg).
		Where("uuid = ? and date between ? and ?", uuid, begindate, enddate).
		Group("uuid,date").
		Order("date desc").Find(&StatisticsS).Error
	if err != nil {
		beego.Error("GetUserStatisticsByDate err ", err)
	}
	return StatisticsS, err
}

// 多用户多日期统计
func GetTeamLtstatistic(sufids []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic {
	var retinfo commonstruct.LotdataStatistic
	if len(sufids) <= 0 {
		return retinfo
	}

	selectinfo := `
	sum(order_num) as order_num,
	sum(order_amount) as order_amount,
	sum(settled_amount) as settled_amount,
	sum(settled_num) as settled_num,
	sum(yingkuijieguo) as yingkuijieguo,
	sum(wager) as wager`

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
		Where("uuid in (?) and date between ? and ? ", sufids, begindate, enddate).Find(&retinfo).Error; err != nil {
		beego.Error("GetTeamLtstatistic err", begindate, enddate, err)
	}
	return retinfo
}

// 保存用户的结算数据
func SaveLotdataStatistic(tx *gorm.DB, roomid int64, arrivetime int64, stats commonstruct.UserItemstatistic) error {
	date := GetBilishiDate(arrivetime)

	var info commonstruct.LotdataStatistic
	if retinfo := tx.Table(commonstruct.WY_tmp_user_lotdata_data).
		Where("uuid = ? and date = ? and room_id = ?", stats.Uuid, date, roomid).
		Find(&info); retinfo.Error != nil {

		if retinfo.RecordNotFound() {
			info.Uuid = stats.Uuid
			info.Date = date
			info.RoomID = roomid

			info.SettledNum = stats.Settlenum
			info.SettledAmount = stats.Settledamount
			info.ValidAmount = stats.Validamount
			info.Wager = stats.Wager
			info.Tuishui = stats.Shuiamount

			info.ProfitWager = stats.ProfitWager
			info.ProfitTuishui = stats.Porfittuishui

			info.Yingshouxiaxian = stats.Yingshouxiaxian
			info.Shizhanhuoliang = stats.Shizhanhuoliang
			info.Shizhanshuying = stats.Shizhanshuying
			info.Shizhanjieguo = stats.Shizhanjieguo
			info.Shizhantuishui = stats.Shizhantuishui
			info.Shizhanpeicha = stats.Shizhanpeicha
			info.Shangjiaohuoliang = stats.Shangjiaohuoliang
			info.Shangjijiaoshou = stats.Shangjijiaoshou

			info.Shouhuo = stats.Shouhuo
			info.Shouhuoyingkui = stats.Shouhuoyingkui
			info.Buchu = stats.Buchu
			info.Buchuyingkui = stats.Buchuyingkui

			info.SettleShoudongbuchu = stats.Shoudongbuchu
			info.SettleShoudongbuchuyingkui = stats.Shoudongbuchuyingkui
			info.Shizhanbuhuo = stats.Shizhanbuhuo
			info.Shizhanbuhuoyingkui = stats.Shizhanbuhuoyingkui

			if err := tx.Table(commonstruct.WY_tmp_user_lotdata_data).Create(&info).Error; err != nil {
				beego.Error("create err ", info, err)
				return err
			}
		} else {
			beego.Error("SaveLotdataStatistic err ", stats.Uuid, date, roomid, retinfo.Error.Error())
			return retinfo.Error
		}
	} else {
		updateValues := map[string]interface{}{
			"settled_num":    gorm.Expr("settled_num + ?", stats.Settlenum),
			"settled_amount": gorm.Expr("settled_amount + ?", stats.Settledamount),
			"valid_amount":   gorm.Expr("valid_amount + ?", stats.Validamount),
			"wager":          gorm.Expr("wager + ?", stats.Wager),
			"tuishui":        gorm.Expr("tuishui + ?", stats.Shuiamount),

			"profit_wager":   gorm.Expr("profit_wager + ?", stats.ProfitWager),
			"profit_tuishui": gorm.Expr("profit_tuishui + ?", stats.Porfittuishui),

			"yingshouxiaxian": gorm.Expr("yingshouxiaxian + ?", stats.Yingshouxiaxian),
			"shizhanhuoliang": gorm.Expr("shizhanhuoliang + ?", stats.Shizhanhuoliang),
			"shizhanshuying":  gorm.Expr("shizhanshuying + ?", stats.Shizhanshuying),
			"shizhanjieguo":   gorm.Expr("shizhanjieguo + ?", stats.Shizhanjieguo),
			"shizhantuishui":  gorm.Expr("shizhantuishui + ?", stats.Shizhantuishui),
			"shizhanpeicha":   gorm.Expr("shizhanpeicha + ?", stats.Shizhanpeicha),
			"yingkuijieguo":   gorm.Expr("yingkuijieguo + ?", stats.Yingkuijieguo),

			"shangjiaohuoliang": gorm.Expr("shangjiaohuoliang + ?", stats.Shangjiaohuoliang),
			"shangjijiaoshou":   gorm.Expr("shangjijiaoshou + ?", stats.Shangjijiaoshou),

			"shouhuo":        gorm.Expr("shouhuo + ?", stats.Shouhuo),
			"shouhuoyingkui": gorm.Expr("shouhuoyingkui + ?", stats.Shouhuoyingkui),
			"buchu":          gorm.Expr("buchu + ?", stats.Buchu),
			"buchuyingkui":   gorm.Expr("buchuyingkui + ?", stats.Buchuyingkui),

			"settle_shoudongbuchu":        gorm.Expr("settle_shoudongbuchu + ?", stats.Shoudongbuchu),
			"settle_shoudongbuchuyingkui": gorm.Expr("settle_shoudongbuchuyingkui + ?", stats.Shoudongbuchuyingkui),
			"shizhanbuhuo":                gorm.Expr("shizhanbuhuo + ?", stats.Shizhanbuhuo),
			"shizhanbuhuoyingkui":         gorm.Expr("shizhanbuhuoyingkui + ?", stats.Shizhanbuhuoyingkui),
		}

		if err := tx.Table(commonstruct.WY_tmp_user_lotdata_data).Where("uuid = ? and date = ? and room_id = ? ", stats.Uuid, date, roomid).
			Update(updateValues).Error; err != nil {
			beego.Error("SaveLotdataStatistic err ", stats.Uuid, stats, err)
			return err
		}
	}
	return nil
}

// 撤销未结算订单统计
func RevokeUnsettleLotdata(orderinfo commonstruct.OrderS) error {
	date := GetBilishiDate(orderinfo.Optime)

	updateValues := map[string]interface{}{
		"order_num":    gorm.Expr("order_num + ?", -1),
		"order_amount": gorm.Expr("order_amount + ?", -orderinfo.Amount),
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_lotdata_data).Where("uuid = ? and date = ? and room_id = ? ", orderinfo.Uuid, date, orderinfo.Roomid).
		Update(updateValues).Error; err != nil {
		beego.Error("RevokeUnsettleLotdata err ", orderinfo.Uuid, orderinfo.Orderid, err)
		return err
	}

	return nil
}

// 保存用户的结算数据
func RevokeUnsettleLotItem(tx *gorm.DB, roomid int64, arrivetime int64, stats commonstruct.UnsettleRevokedata) error {
	date := GetBilishiDate(arrivetime)

	updateValues := map[string]interface{}{
		"order_num":         gorm.Expr("order_num + ?", stats.OrderNum),
		"order_amount":      gorm.Expr("order_amount + ?", stats.OrderAmount),
		"zhanchenghuoliang": gorm.Expr("zhanchenghuoliang + ?", stats.Zhanchenghuoliang),
		"shizhanhuoliang":   gorm.Expr("shizhanhuoliang + ?", stats.Shizhanhuoliang),
		"xiajibuhuo":        gorm.Expr("xiajibuhuo + ?", stats.Xiajibuhuo),
		"zidongbuchu":       gorm.Expr("zidongbuchu + ?", stats.Zidongbuchu),
		"shoudongbuchu":     gorm.Expr("shoudongbuchu + ?", stats.Shoudongbuchu),
	}

	if err := tx.Table(commonstruct.WY_tmp_user_lotdata_item).
		Where("uuid = ? and date = ? and room_id = ?  and pan = ? and expect = ? and item_id = ?", stats.Uuid, date, roomid, stats.Pan, stats.Expect, stats.ItemID).
		Update(updateValues).Error; err != nil {
		beego.Error("RevokeUnsettleLotItem err ", stats.Uuid, stats, err)
		return err
	}
	return nil
}

// 保存用户的结算数据
func SaveLotItemStatistic(tx *gorm.DB, roomid int64, arrivetime int64, stats commonstruct.UserItemstatistic) error {
	date := GetBilishiDate(arrivetime)

	// statsbyte, _ := json.Marshal(&stats)
	// beego.Error("SaveLotItemStatistic == ", string(statsbyte))

	var info commonstruct.LotdataUseritem
	if retinfo := tx.Table(commonstruct.WY_tmp_user_lotdata_item).
		Where("uuid = ? and date = ? and room_id = ? and pan = ? and expect = ? and item_id = ?", stats.Uuid, date, roomid, stats.Pan, stats.Expect, stats.ItemID).
		Find(&info); retinfo.Error != nil {

		if retinfo.RecordNotFound() {
			info.Uuid = stats.Uuid
			info.Date = date
			info.RoomID = roomid
			info.Pan = stats.Pan
			info.Expect = stats.Expect
			info.PortID = stats.PortID
			info.ItemID = stats.ItemID

			info.SettledNum = stats.Settlenum
			info.SettledAmount = stats.Settledamount
			info.ValidAmount = stats.Validamount
			info.Wager = stats.Wager
			info.Tuishui = stats.Shuiamount

			info.ProfitWager = stats.ProfitWager
			info.ProfitTuishui = stats.Porfittuishui

			info.Yingshouxiaxian = stats.Yingshouxiaxian
			info.SettleShizhanhuoliang = stats.Shizhanhuoliang
			info.Shizhanshuying = stats.Shizhanshuying
			info.Shizhanjieguo = stats.Shizhanjieguo
			info.Shizhantuishui = stats.Shizhantuishui
			info.Shizhanpeicha = stats.Shizhanpeicha
			info.Yingkuijieguo = stats.Yingkuijieguo
			info.Shangjiaohuoliang = stats.Shangjiaohuoliang
			info.Shangjijiaoshou = stats.Shangjijiaoshou

			info.Shouhuo = stats.Shouhuo
			info.Shouhuoyingkui = stats.Shouhuoyingkui
			info.Buchu = stats.Buchu
			info.Buchuyingkui = stats.Buchuyingkui

			info.SettleShoudongbuchu = stats.Shoudongbuchu
			info.SettleShoudongbuchuyingkui = stats.Shoudongbuchuyingkui
			info.Shizhanbuhuo = stats.Shizhanbuhuo
			info.Shizhanbuhuoyingkui = stats.Shizhanbuhuoyingkui

			if err := tx.Table(commonstruct.WY_tmp_user_lotdata_item).Create(&info).Error; err != nil {
				beego.Error("create err ", info, err)
				return err
			}
		} else {
			beego.Error("SaveLotdataStatistic err ", stats.Uuid, date, roomid, retinfo.Error.Error())
			return retinfo.Error
		}
	} else {
		updateValues := map[string]interface{}{
			"settled_num":    gorm.Expr("settled_num + ?", stats.Settlenum),
			"settled_amount": gorm.Expr("settled_amount + ?", stats.Settledamount),
			"valid_amount":   gorm.Expr("valid_amount + ?", stats.Validamount),
			"wager":          gorm.Expr("wager + ?", stats.Wager),
			"tuishui":        gorm.Expr("tuishui + ?", stats.Shuiamount),

			"profit_wager":   gorm.Expr("profit_wager + ?", stats.ProfitWager),
			"profit_tuishui": gorm.Expr("profit_tuishui + ?", stats.Porfittuishui),

			"yingshouxiaxian":        gorm.Expr("yingshouxiaxian + ?", stats.Yingshouxiaxian),
			"settle_shizhanhuoliang": gorm.Expr("settle_shizhanhuoliang + ?", stats.Shizhanhuoliang),
			"shizhanshuying":         gorm.Expr("shizhanshuying + ?", stats.Shizhanshuying),
			"shizhanjieguo":          gorm.Expr("shizhanjieguo + ?", stats.Shizhanjieguo),
			"shizhantuishui":         gorm.Expr("shizhantuishui + ?", stats.Shizhantuishui),
			"shizhanpeicha":          gorm.Expr("shizhanpeicha + ?", stats.Shizhanpeicha),
			"yingkuijieguo":          gorm.Expr("yingkuijieguo + ?", stats.Yingkuijieguo),

			"shangjiaohuoliang": gorm.Expr("shangjiaohuoliang + ?", stats.Shangjiaohuoliang),
			"shangjijiaoshou":   gorm.Expr("shangjijiaoshou + ?", stats.Shangjijiaoshou),

			"shouhuo":        gorm.Expr("shouhuo + ?", stats.Shouhuo),
			"shouhuoyingkui": gorm.Expr("shouhuoyingkui + ?", stats.Shouhuoyingkui),
			"buchu":          gorm.Expr("buchu + ?", stats.Buchu),
			"buchuyingkui":   gorm.Expr("buchuyingkui + ?", stats.Buchuyingkui),

			"settle_shoudongbuchu":        gorm.Expr("settle_shoudongbuchu + ?", stats.Shoudongbuchu),
			"settle_shoudongbuchuyingkui": gorm.Expr("settle_shoudongbuchuyingkui + ?", stats.Shoudongbuchuyingkui),
			"shizhanbuhuo":                gorm.Expr("shizhanbuhuo + ?", stats.Shizhanbuhuo),
			"shizhanbuhuoyingkui":         gorm.Expr("shizhanbuhuoyingkui + ?", stats.Shizhanbuhuoyingkui),
		}

		if err := tx.Table(commonstruct.WY_tmp_user_lotdata_item).
			Where("uuid = ? and date = ? and room_id = ?  and pan = ? and expect = ? and item_id = ?", stats.Uuid, date, roomid, stats.Pan, stats.Expect, stats.ItemID).
			Update(updateValues).Error; err != nil {
			beego.Error("SaveLotdataStatistic err ", stats.Uuid, stats, err)
			return err
		}
	}
	return nil
}

// 清空结算数据 date表
func ClearLotdataDate(roomid int64, expect string) error {
	updateValues := map[string]interface{}{
		"settled_num":                 0,
		"settled_amount":              0,
		"valid_amount":                0,
		"wager":                       0,
		"tuishui":                     0,
		"yingshouxiaxian":             0,
		"shizhanhuoliang":             0,
		"shizhanshuying":              0,
		"shizhanjieguo":               0,
		"shizhantuishui":              0,
		"profit_tuishui":              0,
		"profit_wager":                0,
		"shizhanpeicha":               0,
		"yingkuijieguo":               0,
		"shangjiaohuoliang":           0,
		"shangjijiaoshou":             0,
		"shouhuo":                     0,
		"shouhuoyingkui":              0,
		"buchu":                       0,
		"buchuyingkui":                0,
		"settle_shoudongbuchu":        0,
		"settle_shoudongbuchuyingkui": 0,
		"shizhanbuhuo":                0,
		"shizhanbuhuoyingkui":         0,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_lotdata_data).Where("room_id = ? and expect = ?", roomid, expect).
		Update(updateValues).Error; err != nil {
		beego.Error("DeleteProfit err ", err)
		return err
	}
	return nil
}

// 清空结算数据 date表
func ClearLotdataItem(roomid int64, expect string) error {
	updateValues := map[string]interface{}{
		"settled_num":                 0,
		"settled_amount":              0,
		"valid_amount":                0,
		"wager":                       0,
		"tuishui":                     0,
		"yingshouxiaxian":             0,
		"shizhanhuoliang":             0,
		"shizhanshuying":              0,
		"shizhanjieguo":               0,
		"shizhantuishui":              0,
		"profit_tuishui":              0,
		"profit_wager":                0,
		"shizhanpeicha":               0,
		"yingkuijieguo":               0,
		"shangjiaohuoliang":           0,
		"shangjijiaoshou":             0,
		"shouhuo":                     0,
		"shouhuoyingkui":              0,
		"buchu":                       0,
		"buchuyingkui":                0,
		"settle_shoudongbuchu":        0,
		"settle_shoudongbuchuyingkui": 0,
		"shizhanbuhuo":                0,
		"shizhanbuhuoyingkui":         0,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_user_lotdata_item).Where("room_id = ? and expect = ?", roomid, expect).
		Update(updateValues).Error; err != nil {
		beego.Error("DeleteProfit err ", err)
		return err
	}
	return nil
}

// 单用户单下注项统计
func GetUserlotdataByItemid(uuid int64, roomid int64, expect string, itemid int64) commonstruct.LotdataUseritem {
	var data commonstruct.LotdataUseritem

	selectinfo := `
	sum(shizhanhuoliang) as shizhanhuoliang,
	sum(zidongbuchu) as zidongbuchu,
	sum(shoudongbuchu) as shoudongbuchu,
	sum(order_num) as order_num,
	sum(order_amount) as order_amount`

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).
		Select(selectinfo).
		Where("uuid = ? and room_id = ? and expect = ? and item_id = ?", uuid, roomid, expect, itemid).
		Find(&data).Error; err != nil {
		beego.Error("GetUserlotdataByItemid err", uuid, roomid, expect, itemid, err)
	}
	return data
}

// 单用户单下注项按盘口统计
func GetUserlotdataByPan_itemid(uuid int64, roomid int64, pan string, expect string, itemid int64) commonstruct.LotdataUseritem {

	var datas commonstruct.LotdataUseritem
	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).
		Where("uuid = ? and room_id = ? and pan = ? and expect = ? and item_id = ?", uuid, roomid, pan, expect, itemid).
		Find(&datas).Error; err != nil {

		if err.Error() != "record not found" {
			beego.Error("GetUserlotdataByPan_itemid err", uuid, roomid, expect, pan, itemid, err)
		}
	}
	return datas
}

// 多用户单下注项按盘口统计
func GetUserslotdataByItemid(uuids []int64, roomid int64, pan string, expect string, itemid int64) commonstruct.LotdataUseritem {
	selectinfo := `
	sum(order_num) as order_num,
	sum(order_amount) as order_amount`

	var datas commonstruct.LotdataUseritem
	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).
		Select(selectinfo).
		Where("uuid in (?) and room_id = ? and pan = ? and expect = ? and item_id = ?", uuids, roomid, pan, expect, itemid).
		Find(&datas).Error; err != nil {
		beego.Error("GetUserslotdataByItemid err", uuids, roomid, expect, itemid, err)
	}
	return datas
}

// 直属用户的下注统计
func GetDirectUserslotdataByItemid(uuid int64, roletype string, roomid int64, pan string, expect string, itemid int64) commonstruct.LotdataUseritem {
	selectinfo := `
	sum(order_num) as order_num,
	sum(order_amount) as order_amount`

	var datas commonstruct.LotdataUseritem
	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).
		Select(selectinfo).
		Where("pre_id = ? and role_type = ? and room_id = ? and pan = ? and expect = ? and item_id = ?", uuid, roletype, roomid, pan, expect, itemid).
		Find(&datas).Error; err != nil {
		beego.Error("GetUserslotdataByItemid err", uuid, roomid, expect, itemid, err)
	}
	return datas
}

// 多用户单下注项按盘口统计
func GetDangeuserlotdataByPortid(uuids int64, roomid int64, portid int64, begindate int64, enddate int64) commonstruct.LotdataUseritem {
	selectinfo := `
	sum(wager) as wager,
	sum(valid_amount) as valid_amount,
	sum(tuishui) as tuishui,
	sum(profit_tuishui) as profit_tuishui,
	sum(settle_shizhanhuoliang) as settle_shizhanhuoliang,
	sum(shouhuo) as shouhuo,
	sum(shouhuoyingkui) as shouhuoyingkui,
	sum(buchu) as buchu,
	sum(buchuyingkui) as buchuyingkui,
	sum(settle_shoudongbuchu) as settle_shoudongbuchu,
	sum(settle_shoudongbuchuyingkui) as settle_shoudongbuchuyingkui,
	sum(shangjiaohuoliang) as shangjiaohuoliang,
	sum(shangjijiaoshou) as shangjijiaoshou`

	var datas commonstruct.LotdataUseritem
	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).
		Select(selectinfo).
		Where("uuid = ? and room_id = ?  and port_id = ? and date between ? and ?", uuids, roomid, portid, begindate, enddate).
		Find(&datas).Error; err != nil {
		beego.Error("GetDangeuserlotdataByPortid err", uuids, roomid, portid, err)
	}
	return datas
}

// 多用户单下注项按玩法统计
func GetUserslotdataGroupbyPortid(uuids []int64, roomid int64, begindate int64, enddate int64) ([]commonstruct.LotdataUseritem, error) {
	selectinfo := `
	port_id,
	sum(order_num) as order_num,
	sum(order_amount) as order_amount,
	sum(settled_num) as settled_num,
	sum(settled_amount) as settled_amount,
	sum(valid_amount) as valid_amount,
	sum(revoke_amount) as revoke_amount,
	sum(wager) as wager,
	sum(profit_wager) as profit_wager,
	sum(tuishui) as tuishui,
	sum(profit_tuishui) as profit_tuishui,
	sum(yingshouxiaxian) as yingshouxiaxian,
	sum(shizhanhuoliang) as shizhanhuoliang,
	sum(shizhanshuying) as shizhanshuying,
	sum(shizhanjieguo) as shizhanjieguo,
	sum(shizhantuishui) as shizhantuishui,
	sum(yingkuijieguo) as yingkuijieguo,
	sum(shangjiaohuoliang) as shangjiaohuoliang,
	sum(shangjijiaoshou) as shangjijiaoshou`

	var datas []commonstruct.LotdataUseritem
	err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).
		Select(selectinfo).
		Where("uuid in (?) and room_id = ? and date between ? and ?", uuids, roomid, begindate, enddate).
		Group("port_id").
		Find(&datas).Error
	if err != nil {
		beego.Error("GetUserslotdataGroupbyPortid err", uuids, roomid, begindate, enddate, err)
	}
	return datas, err
}

// 获取直属下级的 交收统计部分
func GetDirectuserStats(uuids []int64, roomid int64, portid int64, begindate int64, enddate int64) (commonstruct.LotdataUseritem, error) {
	selectinfo := `
	sum(profit_tuishui) as profit_tuishui,
	sum(yingshouxiaxian) as yingshouxiaxian,
	sum(settle_shizhanhuoliang) as settle_shizhanhuoliang,
	sum(shizhanshuying) as shizhanshuying,
	sum(shizhanjieguo) as shizhanjieguo,
	sum(shizhantuishui) as shizhantuishui,
	sum(yingkuijieguo) as yingkuijieguo,
	sum(shangjiaohuoliang) as shangjiaohuoliang,
	sum(shangjijiaoshou) as shangjijiaoshou,
	sum(shouhuo) as shouhuo,
	sum(shouhuoyingkui) as shouhuoyingkui,
	sum(buchu) as buchu,
	sum(buchuyingkui) as buchuyingkui`

	var datas commonstruct.LotdataUseritem
	err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).
		Select(selectinfo).
		Where("uuid in (?) and room_id = ? and port_id = ? and date between ? and ?", uuids, roomid, portid, begindate, enddate).
		Find(&datas).Error
	if err != nil {
		beego.Error("GetDirectuserStats err", uuids, roomid, portid, begindate, enddate, err)
	}
	return datas, err
}

// 获取直属下级的 交收统计部分
func GetDirectuserGamestats(uuids []int64, roomid int64, begindate int64, enddate int64) (commonstruct.LotdataStatistic, error) {
	selectinfo := `sum(yingkuijieguo) as yingkuijieguo`

	var datas commonstruct.LotdataStatistic
	err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).
		Select(selectinfo).
		Where("uuid in (?) and room_id = ? and date between ? and ?", uuids, roomid, begindate, enddate).
		Find(&datas).Error
	if err != nil {
		beego.Error("GetDirectuserStats err", uuids, roomid, begindate, enddate, err)
	}
	return datas, err
}

// 获取公司的周活跃用户输赢统计
func GetWeekActiveusers(masterid int64) ([]commonstruct.LotdataStatistic, error) {
	var list []commonstruct.LotdataStatistic

	enddate := commonfunc.GetNowdate()
	begindate := commonfunc.GetBjTime20060102(time.Now().AddDate(0, 0, -7))

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select("uuid,sum(profit_total) as profit_total").
		Where("master_id = ? and date between ? and ?", masterid, begindate, enddate).Group("uuid").
		Find(&list).Error; err != nil {
		beego.Error("GetWeekActiveusers err ", masterid, begindate, enddate, err)
		return list, err
	}
	return list, nil
}

// 按代理获取收补明细
func GetAgentShoubumingxi(uuid int64, roomids []int64, begindate int64, enddate int64) commonstruct.LotdataStatistic {
	var retinfo commonstruct.LotdataStatistic

	selectinfo := `
	sum(shouhuo) as shouhuo,
	sum(shouhuoyingkui) as shouhuoyingkui,
	sum(buchu) as buchu,
	sum(buchuyingkui) as buchuyingkui,
	sum(settle_shoudongbuchu) as settle_shoudongbuchu,
	sum(settle_shoudongbuchuyingkui) as settle_shoudongbuchuyingkui,
	sum(shizhanbuhuo) as shizhanbuhuo,
	sum(shizhanbuhuoyingkui) as shizhanbuhuoyingkui`

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_data).Select(selectinfo).
		Where("uuid = ? and room_id in (?) and date between ? and ? ", uuid, roomids, begindate, enddate).Find(&retinfo).Error; err != nil {
		beego.Error("GetAgentShoubumingxi err", uuid, roomids, begindate, enddate, err)
	}
	return retinfo
}

// 获取代理单个玩法的收补明细
func GetPortShoubumingxi(uuid int64, roomid int64, portid int64, begindate int64, enddate int64) commonstruct.LotdataUseritem {
	var retinfo commonstruct.LotdataUseritem

	selectinfo := `
	sum(shouhuo) as shouhuo,
	sum(shouhuoyingkui) as shouhuoyingkui,
	sum(buchu) as buchu,
	sum(buchuyingkui) as buchuyingkui,
	sum(shizhanbuhuo) as shizhanbuhuo,
	sum(shizhanbuhuoyingkui) as shizhanbuhuoyingkui`

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).Select(selectinfo).
		Where("pre_id = ? and room_id = ? and port_id = ? and date between ? and ? ", uuid, roomid, portid, begindate, enddate).
		Find(&retinfo).Error; err != nil {
		beego.Error("GetAgentShoubumingxi err", uuid, roomid, begindate, enddate, err)
	}
	return retinfo
}

// 获取代理单个游戏的收补明细合计
func GetPortShoubumingxiS(uuid int64, roomid int64, begindate int64, enddate int64) ([]commonstruct.LotdataUseritem, error) {
	selectinfo := `
	port_id,
	sum(shouhuo) as shouhuo,
	sum(shouhuoyingkui) as shouhuoyingkui,
	sum(buchu) as buchu,
	sum(buchuyingkui) as buchuyingkui,
	sum(shizhanbuhuo) as shizhanbuhuo,
	sum(shizhanbuhuoyingkui) as shizhanbuhuoyingkui`

	var datas []commonstruct.LotdataUseritem
	err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).Select(selectinfo).
		Where("pre_id = ? and room_id = ? and date between ? and ?", uuid, roomid, begindate, enddate).
		Group("port_id").Find(&datas).Error
	if err != nil {
		beego.Error("GetPortShoubumingxiS err", uuid, roomid, begindate, enddate, err)
	}
	return datas, err
}

// 获取单用户单游戏单期统计 GetUserlotdataByExpect(uuid int64, roomid int64, expect string)
func GetUserlotdataByExpect(uuid int64, roomid int64, expect string) commonstruct.LotdataUseritem {
	var statistic commonstruct.LotdataUseritem

	selectinfo := `
		sum(order_num) as order_num,
		sum(order_amount) as order_amount,
		sum(settled_num) as settled_num,
		sum(settled_amount) as settled_amount,
		sum(valid_amount) as valid_amount,
		sum(wager) as wager,
		sum(profit_wager) as profit_wager,
		sum(tuishui) as tuishui,
		sum(profit_tuishui) as profit_tuishui`

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).Select(selectinfo).
		Where("uuid = ? and room_id = ? and expect = ? ", uuid, roomid, expect).
		Find(&statistic).Error; err != nil {
		beego.Error("GetUserlotdataByExpect err", uuid, roomid, expect, err)
	}

	return statistic
}

// 获取特码风控列表 GetUserTemaFengkongList(roomid int64, expect string, amount float64)
func GetUserTemaorderinfoS(roomid int64, expect string, amount float64) []commonstruct.LotdataUsertema {
	var statistic []commonstruct.LotdataUsertema

	selectinfo := `uuid,
		count(*) as order_num` // 满足条件的球号个数

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_tema).Select(selectinfo).
		Where("room_id = ? and expect = ? and order_amount >= ?", roomid, expect, amount).
		Group("uuid").
		Find(&statistic).Error; err != nil {
		beego.Error("GetUserTemaorderinfoS err", roomid, expect, err)
	}

	return statistic
}

// 获取时时彩球号风控列表
func GetUserSscorderinfoS(masterid int64, roomid int64, expect string, itemids []int64, amount float64) []commonstruct.LotdataUseritem {
	var statistic []commonstruct.LotdataUseritem

	selectinfo := `uuid,count(*) as order_num` // 满足条件的球号个数

	if err := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).Select(selectinfo).
		Where("master_id = ? and room_id = ? and expect = ? and item_id in (?) and order_amount >= ?", masterid, roomid, expect, itemids, amount).
		Group("uuid").
		Find(&statistic).Error; err != nil {
		beego.Error("GetUserSscorderinfoS err", roomid, expect, err)
	}

	return statistic
}
