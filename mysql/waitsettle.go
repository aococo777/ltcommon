package wymysql

import (
	"errors"
	"fmt"
	"github.com/aococo777/ltcommon/commonfunc"
	"github.com/aococo777/ltcommon/commonstruct"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
WY_tmp_wait_settle:

获取游戏未结算期数 GetRoomWaitsettleinfo(roomid int64) commonstruct.SettleWait
保存开奖号码 UpsertWaitSettle(roomid int64, expect string, code string) error
保存开奖号码 UpdateSettleCode(roomid int64, expect string, code string, i_arrivetime int64) error
保存结算记录日志 UpdateSettleExpinfo(roomid int64, expect string, expinfo string) error
保存结算信息 UpdateSettlewaitInfo(orderid int64, action string, result int64, expinfo string) error

获取游戏单期结算记录 GetWaitsettleinfo(roomid int64, expect string) (commonstruct.SettleWait, error)
获取未统计的记录 GetUnStatisticWait() (bool, commonstruct.SettleWait)
获取未派彩的记录 GetUnPaicaiWait() (bool, commonstruct.SettleWait)
获取未结算的记录 GetUnSettleWait() (bool, commonstruct.SettleWait)
查询未结算期号信息 GetErrsettles(offset int, roomlist []int64, errtype string, begindate int64, enddate int64) ([]commonstruct.SettleWait, error)
查询错误结算分页信息 GetErrsettlePageinfo(roomlist []int64, errtype string, begindate int64, enddate int64) (int64, int64)
提前开奖详情 GetAdvanceopeninfo(offset int) []commonstruct.SettleWait
提前开奖分页详情 GetAdvanceopenPageinfo() (int64, int64)
查询未结算的期号数 GetUnsettlenum() int64

分页查询未结算的详情  GetUnsettleinfo(roomid int64, offset int, pagecount int) []commonstruct.SettleWait
未结算的期号分页详情 GetUnsettlePageinfo(roomid int64, pagecount int) (int64, int64)
设置彩种的结算号码 SetSettlecode(roomid int64, expect string, code string) error
重置彩种的结算号码、标志位  ResetOpencode(roomid int64, expect string, code string) error
新增结算 AddSettleWait(roomid int64, expect string, wishtime int64) error
删除结算信息 ClearSettledinfo()
是否新增结算记录 IsAddSettleWait(roomid int64, expect string) bool
*/

// 获取游戏未结算期数
func GetRoomWaitsettleinfo(roomid int64) commonstruct.SettleWait {
	// 中奖列表
	var settleinfo commonstruct.SettleWait
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Select("count(*) as id").Where("room_id = ? and issettled = 0", roomid).Find(&settleinfo); retinfo.Error != nil {
		if !retinfo.RecordNotFound() {
			beego.Error("GetWaitsettleinfo err", roomid, retinfo.Error.Error())
		}
	}
	return settleinfo
}

// 保存开奖号码
func UpdateSettleCode(roomid int64, expect string, code string, i_arrivetime int64) error {
	tx := WyMysql.Begin()
	var log commonstruct.SettleWait

	if err := tx.Table(commonstruct.WY_tmp_wait_settle).Where("room_id = ? and expect = ?", roomid, expect).Find(&log).Error; err != nil {
		beego.Error("UpdateSettleCode err", err)
		tx.Rollback()
		return err
	} else {
		if log.Issettled > 0 { // 订单已处理
			tx.Rollback()
			return errors.New("该期已处理")
		} else {
			expinfo := fmt.Sprintf("获取到开奖号码%v=>%v", i_arrivetime, code)
			updateValues := map[string]interface{}{
				"code":       code,
				"arrivetime": i_arrivetime,
				"expinfo":    gorm.Expr("CONCAT_WS('|',expinfo,?)", expinfo),
			}

			if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("room_id = ? and expect = ?", roomid, expect).Update(updateValues).Error; err != nil {
				beego.Error("UpdateSettleCode err", updateValues, err)
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("UpdateSettleCode commit err ", err)
		tx.Rollback()
		return err
	}
	return nil
}

// 保存结算记录日志
func UpdateSettleExpinfo(roomid int64, expect string, expinfo string) error {
	updateValues := map[string]interface{}{
		"expinfo": gorm.Expr("CONCAT_WS('|',expinfo,?)", expinfo),
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("room_id = ? and expect = ?", roomid, expect).Update(updateValues).Error; err != nil {
		beego.Error("SetSettlecode err", updateValues, err)
		return err
	}
	return nil
}

// 获取游戏单期结算记录
func GetWaitsettleinfo(roomid int64, expect string) (commonstruct.SettleWait, error) {
	var Ret commonstruct.SettleWait
	err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("room_id = ? and expect = ?", roomid, expect).Find(&Ret).Error
	if err != nil {
		beego.Error("GetWaitSettleinfo err", roomid, expect, err)
	}
	return Ret, err
}

// 获取未拿到号码的记录
func GetUncodeWaits() ([]commonstruct.SettleWait, error) {

	nowtime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().Add(-25*time.Minute)), 10, 64)

	var Ret []commonstruct.SettleWait
	err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("arrivetime  = 0 and wishtime < ?", nowtime).Order("id").Find(&Ret).Error
	if err != nil {
		beego.Error("GetUncodeWaits err", err.Error())
		return Ret, err
	}
	return Ret, err
}

// 获取未统计的记录
func GetUnStatisticWait() (bool, commonstruct.SettleWait) {
	var Ret commonstruct.SettleWait
	var count int64
	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("issettled = ? and isstatistic = 0", commonstruct.SaveResult_Success).Order("id").Limit(1).Find(&Ret).Count(&count).Error; err != nil {
		return false, Ret
	}
	if count == 0 {
		return false, Ret
	}
	return true, Ret
}

// 获取未派彩的记录
func GetUnPaicaiWait() (bool, commonstruct.SettleWait) {
	var Ret commonstruct.SettleWait
	var count int64
	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("issettled = ? and isretcash = 0", commonstruct.SaveResult_Success).Order("id").Limit(1).Find(&Ret).Count(&count).Error; err != nil {
		return false, Ret
	}
	if count == 0 {
		return false, Ret
	}
	return true, Ret
}

// 获取未结算的记录
func GetUnSettleWait() (bool, commonstruct.SettleWait) {
	var Ret commonstruct.SettleWait
	var count int64
	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("issettled = 0 and arrivetime > 0").Order("id").Limit(1).Find(&Ret).Count(&count).Error; err != nil {
		return false, Ret
	}
	if count == 0 {
		return false, Ret
	}
	return true, Ret
}

// 保存结算信息
func UpdateSettlewaitInfo(orderid int64, action string, result int64, expinfo string) error {
	var updateValues map[string]interface{}
	nowtime := commonfunc.GetNowtime()
	addinfo := fmt.Sprintf("%v=>%v", nowtime, expinfo)
	switch action {
	case "settle":
		updateValues = map[string]interface{}{
			"issettled":    result,
			"settled_time": nowtime,
			"expinfo":      gorm.Expr("CONCAT_WS('|',expinfo,?)", addinfo),
		}
	case "retcash":
		updateValues = map[string]interface{}{
			"isretcash":    result,
			"retcash_time": nowtime,
			"expinfo":      gorm.Expr("CONCAT_WS('|',expinfo,?)", addinfo),
		}
	case "statistic":
		updateValues = map[string]interface{}{
			"isstatistic":    result,
			"statistic_time": nowtime,
			"expinfo":        gorm.Expr("CONCAT_WS('|',expinfo,?)", addinfo),
		}
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("id = ?", orderid).Update(updateValues).Error; err != nil {
		beego.Error("UpdateSettleinfo err", updateValues, err)
		return err
	}
	return nil
}

// 查询错误结算期号信息
func GetErrsettles(offset int, roomlist []int64, errtype string, begindate int64, enddate int64) ([]commonstruct.SettleWait, error) {

	today := commonfunc.GetNowdate()
	var begintime, endtime int64
	begintime = commonfunc.GetBegintime(begindate)
	if enddate == today {
		endtime, _ = strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().Add(-10*time.Minute)), 10, 64)
	} else {
		endtime = commonfunc.GetEndtime(enddate)
	}

	var unsettles []commonstruct.SettleWait

	switch errtype {
	case "lostcode":
		if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).
			Where("wishtime between ? and ? and issettled = 0 and room_id in (?)", begintime, endtime, roomlist).Order("wishtime desc").
			Limit(20).Offset((offset - 1) * 20).Find(&unsettles).Error; err != nil {
			beego.Error("GetErrsettles err", begintime, endtime, roomlist, offset, err)
			return nil, errors.New(fmt.Sprintf("wrong errtype %v", errtype))
		}
	case "errcode":
		if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).
			Where("wishtime between ? and ? and issettled in (7307,7308) and room_id in (?)", begintime, endtime, roomlist).Order("wishtime desc").
			Limit(20).Offset((offset - 1) * 20).Find(&unsettles).Error; err != nil {
			beego.Error("GetErrsettles err", begintime, endtime, roomlist, offset, err)
			return nil, errors.New(fmt.Sprintf("wrong errtype %v", errtype))
		}
	default:
		return nil, errors.New(fmt.Sprintf("wrong errtype %v", errtype))
	}
	return unsettles, nil
}

// 查询错误结算分页信息
func GetErrsettlePageinfo(roomlist []int64, errtype string, begindate int64, enddate int64) (int64, int64) {

	today := commonfunc.GetNowdate()
	var begintime, endtime int64
	begintime = commonfunc.GetBegintime(begindate)
	if enddate == today {
		endtime, _ = strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().Add(-10*time.Minute)), 10, 64)
	} else {
		endtime = commonfunc.GetEndtime(enddate)
	}

	var unsettles commonstruct.SettleWait
	switch errtype {
	case "lostcode":
		if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Select("count(*) as id").
			Where("wishtime between ? and ? and issettled = 0 and room_id in (?)", begintime, endtime, roomlist).Find(&unsettles).Error; err != nil {
			beego.Error("GetErrsettlePageinfo err", err)
		}
	case "errcode":
		if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Select("count(*) as id").
			Where("wishtime between ? and ? and issettled in (7307,7308) and room_id in (?)", begintime, endtime, roomlist).Find(&unsettles).Error; err != nil {
			beego.Error("GetErrsettlePageinfo err", err)
		}
	default:
		return 0, 0
	}
	return int64(unsettles.ID), int64(math.Ceil(float64(unsettles.ID) / 20))
}

// 提前开奖详情
func GetAdvanceopeninfo(offset int) []commonstruct.SettleWait {

	nowtime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().AddDate(0, 0, -7)), 10, 64)

	var unsettles []commonstruct.SettleWait
	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).
		Where("arrivetime < wishtime and arrivetime > ?", nowtime).Order("wishtime desc").
		Limit(20).Offset((offset - 1) * 20).Find(&unsettles).Error; err != nil {
		beego.Error("GetUnsettleinfo err", err)
	}
	return unsettles
}

// 提前开奖分页详情
func GetAdvanceopenPageinfo() (int64, int64) {

	nowtime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().AddDate(0, 0, -7)), 10, 64)

	var unsettles commonstruct.SettleWait
	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Select("count(*) as id").
		Where("arrivetime < wishtime and arrivetime > ?", nowtime).Find(&unsettles).Error; err != nil {
		beego.Error("GetUnsettleinfo err", err)
	}
	return int64(unsettles.ID), int64(math.Ceil(float64(unsettles.ID) / 20))
}

// 查询未结算的期号数
func GetUnsettlenum() int64 {

	nowtime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().Add(-10*time.Minute)), 10, 64)

	var unsettles commonstruct.SettleWait
	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Select("count(*) as id").
		Where("wishtime < ? and issettled = 0", nowtime).Find(&unsettles).Error; err != nil {
		beego.Error("GetUnsettleinfo err", err)
	}
	return unsettles.ID
}

// 分页查询未结算的详情
func GetUnsettleinfo(roomid int64, offset int, pagecount int) []commonstruct.SettleWait {
	var unsettles []commonstruct.SettleWait
	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).
		Where("issettled = 0 and room_id = ?", roomid).Order("wishtime desc").
		Limit(pagecount).Offset((offset - 1) * pagecount).Find(&unsettles).Error; err != nil {
		beego.Error("GetUnsettleinfo err", err)
	}
	return unsettles
}

// 未结算的期号分页详情
func GetUnsettlePageinfo(roomid int64, pagecount int) (int64, int64) {
	var unsettles commonstruct.SettleWait

	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Select("count(*) as id").
		Where("issettled = 0 and room_id = ?", roomid).Find(&unsettles).Error; err != nil {
		beego.Error("GetUnsettlePageinfo err", err)
	}

	return int64(unsettles.ID), int64(math.Ceil(float64(unsettles.ID) / float64(pagecount)))
}

// 设置彩种的结算号码
func SetSettlecode(roomid int64, expect string, code string) error {
	updateValues := map[string]interface{}{
		"code": code,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("room_id = ? and expect = ?", roomid, expect).
		Update(updateValues).Error; err != nil {
		beego.Error("ResetOpencode err ", err)
		return err
	}
	return nil
}

// 重置彩种的结算号码、标志位
func ResetSettleflag(roomid int64, expect string) error {
	updateValues := map[string]interface{}{
		"issettled":      0,
		"settled_time":   0,
		"isstatistic":    0,
		"statistic_time": 0,
		"isretcash":      0,
		"retcash_time":   0,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("room_id = ? and expect = ?", roomid, expect).
		Update(updateValues).Error; err != nil {
		beego.Error("ResetSettleflag err ", err)
		return err
	}

	return nil
}

// 新增结算
func AddSettleWait(roomid int64, expect string, wishtime int64) error {
	if roomid < 0 || expect == "" {
		return errors.New("roomid is zero or expect is null")
	}

	var oplog commonstruct.SettleWait
	oplog.RoomID = roomid
	oplog.Expect = expect
	oplog.Wishtime = wishtime
	if err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Create(&oplog).Error; err != nil {
		beego.Error("err %v", err)
		return err
	}
	return nil
}

// 修改结算预期时间
func DeleteWaitsettle(roomid int64, expect string) error {
	sql := fmt.Sprintf("delete from %v where room_id = %v and expect = '%v';", commonstruct.WY_tmp_wait_settle, roomid, expect)
	if err := WyMysql.Exec(sql).Error; err != nil {
		beego.Error("DeleteWaitsettle err ", roomid, expect, err)
	}
	return nil
}

// 获取游戏某期结算状态
func GetSettlestate(roomid int64, expect string) int64 {
	// 0 等到号码 1 结算中 2 结算完成
	if settleinfo, err := GetSlotsSettleinfo(roomid); err != nil {
		return 0
	} else {
		if settleinfo.Arrivetime == 0 {
			return 0
		} else if settleinfo.Arrivetime > 0 && settleinfo.Isretcash == 0 {
			return 1
		} else if settleinfo.Isretcash > 0 {
			return 2
		}
	}
	return 0
}

// 获取彩种的结算信息
func GetSlotsSettleinfo(roomid int64) (commonstruct.SettleWait, error) {
	var data commonstruct.SettleWait
	err := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("room_id = ?", roomid).Order("id desc").Limit(1).
		Find(&data).Error
	return data, err
}

// 删除结算信息
func ClearSettledinfo() {

	var RoomInfoS []commonstruct.CRoomInfo
	err := WyMysql.Table(commonstruct.WY_gm_config_room).Where("frequency = 'minite' and  in_valid_time > 0").Order("id").Find(&RoomInfoS).Error
	if err != nil {
		beego.Error("GetRoominfoS err ", err)
	}

	for _, roominfo := range RoomInfoS {
		threemouth, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().AddDate(0, 0, -3)), 10, 64)
		sql := fmt.Sprintf("delete from %v where issettled = 7300 and isretcash = 7302 and wishtime < %v and room_id = %v", commonstruct.WY_tmp_wait_settle, threemouth, roominfo.ID)
		if err := WyMysql.Exec(sql).Error; err != nil {
			beego.Error("ClearSettledinfo err ", err)
		}
	}
}

// 是否新增结算记录
func IsAddSettleWait(roomid int64, expect string) (commonstruct.SettleWait, bool) {
	var oplog commonstruct.SettleWait
	if retinfo := WyMysql.Table(commonstruct.WY_tmp_wait_settle).Where("room_id = ? and expect = ?", roomid, expect).
		Find(&oplog); retinfo.Error != nil {
		if !retinfo.RecordNotFound() {
			beego.Error("IsSettled err ", roomid, retinfo.Error.Error())
			return oplog, true
		} else {
			return oplog, false
		}
	}
	return oplog, true
}
