package wymysql

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aococo777ltcommon/commonfunc"
	"github.com/aococo777ltcommon/commonstruct"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
WY_tmp_lotteryorder:

获取用户下注项合计 GetOrderinfoByItemid(uuid int64, roomid int64, expect string, itemid int64) commonstruct.OrderS
根据订单号获取订单 GetOrderinfoByID(orderid int64) (commonstruct.OrderS, error)
未结算的订单页码详情 GetUnusualorderSPageinfo(roomlist []int64, pagecount int) (int64, int64)
分页查询未结算的详情 GetUnusualorderS(offset int, roomlist []int64, pagecount int) ([]commonstruct.OrderS, error)
获取整期所有未结算订单 GetUnusualorderSByExpect(roomid int64, expect string) ([]commonstruct.OrderS, error)
获取游戏所有未结算订单 GetUnusualorderSByroomid(roomid int64) ([]commonstruct.OrderS, error)
通过订单号获取未结算订单 GetUnusualorderSByorderid(roomid int64, orderid int64) ([]commonstruct.OrderS, error)
获取即时注单下注项的订单列表 GetItemorderS(uuids []int64, roomid int64, expect string, itemid int64, offset int, pagecount int) ([]commonstruct.OrderS, error)
获取即时注单下注项的订单分页信息 GetItemorderSPageinfo(uuids []int64, roomid int64, expect string, itemid int64, offset int, pagecount int) (int, int)
获取玩法分类货量统计 GetPortstatsS(uuids []int64, roomid int64, begintime int64, endtime int64, orderstate int, settlestate int) ([]commonstruct.OrderS, error)
获取一期所有订单 GetUnsettleOrdersByExpect(roomid int64, expect string) ([]commonstruct.OrderS, error)
保存是否有中奖 SetOrderWin(tablesuf string, roomid int64, orderid int64, winid int64)
清理详单记录 ClearOrders(roomids []int64)
连码订单保存 SettleLianmaorder(orderid int64, iswin int64, winlist string, winodds float64, wager float64, opencode string) error
两面中奖订单保存 SettleLiangmianWinorder(roomid int64, expect string, itemid int64, winlist string, opencode string) error
连码输单保存 SettleLiangmianLostorder(roomid int64, expect string, itemid int64, opencode string) error
打平订单处理 SettleLiangmianPingorder(roomid int64, expect string, itemid int64, opencode string) error
保存中奖号码 SaveOrderOpencode(roomid int64, expect string, opencode string, iteminfo string) error
连码订单获取 GetLianmaOrderS(roomid int64, expect string) ([]commonstruct.OrderS, error)
结算时获取当期的所有未结算的两面玩法下注项 GetOrderSByXCItemID(roomid int64, expect string) ([]commonstruct.OrderS, error)
获取用户下注项统计 GetUserOrderstats(uuid int64, roomid int64, expect string) ([]commonstruct.OrderS, error)
保存订单派彩金额 TXUpdateOrderwager(tx *gorm.DB, orderid int64, wager float64) error
获取用户日统计 GetUserOrderstatistic(uuid int64, date int64) ([]commonstruct.OrderS, error)
游戏单期按下注项统计 GetOrderStatistic(roomid int64, expect string) []commonstruct.OrderS
订单作废 RevokeCoorder(roomid int64, orderid int64, amount float64) error
根据游戏和订单号获取订单 GetOrdersByOrderid(roomid int64, orderid int64) (commonstruct.OrderS, error)
查询单期订单页码信息 GetExpectorderPageinfo(roomid int64, expect string, masterid int64, pagecount int) (int, int)
查询单期订单列表 GetExpectorderS(roomid int64, expect string, masterid int64, pagenum int, pagecount int) []commonstruct.OrderS
根据游戏和期号获取订单 GetExpectOrders(masterid int64, roomid int64, expect string) ([]commonstruct.OrderS, error)
获取用户日初未结算订单 GetDateUnsettle(uuid int64, begindate int64) (commonstruct.OrderS, error)
获取当前下注金额 GetItemamount(roomid int64, masterid int64, expect string) []commonstruct.OrderS
获取下注分页信息 BKGetOrdersdetailSPageinfo(uuid int64, roomids []int64, begintime int64, endtime int64) (int, int)
根据状态获取订单分页信息 GetOrdersdetailSPageinfo(uuids []int64, roomids []int64, pagecount int, begintime int64, endtime int64, orderstate int, settlestate int) (int, int)
获取多用户未结算订单金额 GetTeamUnsettleamount(sufids []int64) commonstruct.OrderS
重置用户组的结算订单状态 ClearSettleinfo(roomid int64, expect string) error
新建订单 NewOrderUser(tx *gorm.DB, neworder commonstruct.OrderS) (commonstruct.OrderS, error)
获取已结算订单 GetTodayorders(uuid int64, issettled int, begintime int64, endtime int64) ([]commonstruct.OrderS, error)
获取已结算订单 GetOrdersdetailS(uuid int64, roomid int64, daleis []string, issettled int, offset int, begintime int64, endtime int64) ([]commonstruct.OrderS, error)
分页获取单用户不同状态订单列表 GetLotOrderinfoS(uuid int64, roomid int64, issettled int64, begindate int64, enddate int64, pagenum int64, pagecount int64) ([]commonstruct.OrderS, error)
分页获取单用户不同状态订单分页信息 GetLotOrderinfoSPageinfo(uuid int64, roomid int64, issettled int64, begindate int64, enddate int64) (int, int)
删除三个月前订单 ClearExpireorder()
按期号查询注单 GetOrdersByExpect(roomid int64, expect string) ([]commonstruct.OrderS, error)
单用户多游戏订单分页信息 GetOrdershisSPageinfo(uuid int64, roomids []int64, begintime int64, endtime int64) (int, int)
分页查询多用户多游戏订单列表 GetOrderSHisByDate(uuids []int64, roomids []int64, offset int, pagecount int, begintime int64, endtime int64, orderstate int, settlestate int) ([]commonstruct.OrderS, error)
单用户单期订单统计 GetWagerstatsByExpect(roomid int64, expect string) ([]commonstruct.OrderS, error)
代理直属的订单列表 GetOrdersByPreid(preid int64, roomid int64, expect string) ([]commonstruct.OrderS, error)
撤单 RevokeOrder(roomid int64, orderid int64) error
获取单个订单信息 GetOrdersinfo(roomid int64, orderid int64) (commonstruct.OrderS, error)
下注玩法订单分页信息 GetPortOrdersdetailSPageinfo(uuids []int64, roomids []int64, portid int, pagecount int, begintime int64, endtime int64, orderstate int, settlestate int) (int, int)
下注玩法订单列表 GetPortOrderSHisByDate(uuids []int64, roomids []int64, portid int, offset int, pagecount int, begintime int64, endtime int64, orderstate int, settlestate int) ([]commonstruct.OrderS, error)
查询用户未计算金额合计  QueryUnsettlemoney(uuid int64) float64
获取未统计报表的订单 GetUnstatsorderSByExpect(roomid int64, expect string) ([]commonstruct.OrderS, error)
获取时间段内的下注用户 func GetOrderuuids(begindate int64, enddate int64) []int64
获取时间段内的下注用户 func GetOrderuuidsByExpect(roomid int64, expect string) []int64

*/

var dancibishu float64 = 500

// 获取用户下注项合计
func GetOrderinfoByItemid(uuid int64, roomid int64, expect string, itemid int64) commonstruct.LotdataUseritem {
	var result commonstruct.LotdataUseritem
	if retinfo := WyMysql_read.Table(commonstruct.WY_tmp_user_lotdata_item).
		Where("uuid = ? and room_id = ? and expect = ? and item_id = ?", uuid, roomid, expect, itemid).
		Find(&result); retinfo.Error != nil {
		beego.Error("GetOrderinfoByItemid err", retinfo.Error)
	}
	return result
}

// 根据订单号获取订单
func GetOrderinfoByID(orderid int64) (commonstruct.OrderS, error) {
	var orders commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("orderid = ?", orderid).Find(&orders).Error
	return orders, err
}

// 未结算的订单页码详情
func GetUnusualorderSPageinfo(roomlist []int64, pagecount int) (int64, int64) {

	nowtime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().Add(-30*time.Minute)), 10, 64)

	var unsettles commonstruct.OrderS

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("optime < ? and state = 0 and roomid in (?)", nowtime, roomlist).Find(&unsettles).Error; err != nil {
		beego.Error("GetUnusualorderSPageinfo err", err)
	}

	return int64(unsettles.Orderid), int64(math.Ceil(float64(unsettles.Orderid) / float64(pagecount)))
}

// 分页查询未结算的详情
func GetUnusualorderS(offset int, roomlist []int64, pagecount int) ([]commonstruct.OrderS, error) {

	nowtime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().Add(-30*time.Minute)), 10, 64)

	var unsettles []commonstruct.OrderS

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("optime < ? and state = 0 and roomid in (?) ", nowtime, roomlist).Order("optime desc").
		Limit(pagecount).Offset((offset - 1) * pagecount).Find(&unsettles).Error

	return unsettles, err
}

// 获取整期所有未结算订单
func GetUnusualorderSByExpect(roomid int64, expect string) ([]commonstruct.OrderS, error) {

	nowtime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().Add(-30*time.Minute)), 10, 64)

	var unsettles []commonstruct.OrderS

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("optime < ? and state = 0 and roomid = ? and expect = ?", nowtime, roomid, expect).Order("optime desc").Find(&unsettles).Error

	return unsettles, err
}

// 获取游戏所有未结算订单
func GetUnusualorderSByroomid(roomid int64) ([]commonstruct.OrderS, error) {

	nowtime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().Add(-30*time.Minute)), 10, 64)

	var unsettles []commonstruct.OrderS

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("optime < ? and state = 0 and roomid = ?", nowtime, roomid).Order("optime desc").Find(&unsettles).Error
	return unsettles, err
}

// 通过订单号获取未结算订单
func GetUnusualorderSByorderid(roomid int64, orderid int64) ([]commonstruct.OrderS, error) {
	var unsettles []commonstruct.OrderS

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("state = 0 and roomid = ? and orderid = ?", roomid, orderid).Order("optime desc").Find(&unsettles).Error

	return unsettles, err
}

// 获取即时注单下注项的订单
func GetItemorderS(uuids []int64, roomid int64, expect string, itemid int64, offset int, pagecount int) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("uuid in (?) and roomid = ? and expect = ? and item_id = ? and state = 0", uuids, roomid, expect, itemid).
		Order("orderid").Limit(pagecount).Offset((offset - 1) * pagecount).
		Find(&orders).Error
	return orders, err
}

func GetItemorderSPageinfo(uuids []int64, roomid int64, expect string, itemid int64, offset int, pagecount int) (int, int) {
	var logs commonstruct.OrderS

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("uuid in (?) and roomid = ? and expect = ? and item_id = ? and state = 0", uuids, roomid, expect, itemid).
		Find(&logs).Error; err != nil {
		beego.Error("GetItemorderSPageinfo err", uuids, roomid, expect, itemid, err)
	}
	return int(logs.Orderid), int(math.Ceil(float64(logs.Orderid) / float64(pagecount)))
}

// 获取玩法分类货量统计
func GetPortstatsS(uuids []int64, roomid int64, begintime int64, endtime int64, orderstate int, settlestate int) ([]commonstruct.OrderS, error) {
	var sqlArg string
	switch orderstate { // 是否中奖
	case -1:
		switch settlestate { // 是否结算
		case -1:
			sqlArg = ""
		default:
			sqlArg = fmt.Sprintf("state = %v and ", settlestate)
		}
	case 0:
		switch settlestate {
		case -1:
			sqlArg = "wager = 0 and "
		default:
			sqlArg = fmt.Sprintf("wager = 0 and state = %v and ", settlestate)
		}
	case 1:
		switch settlestate {
		case -1:
			sqlArg = "wager > 0 and "
		default:
			sqlArg = fmt.Sprintf("wager > 0 and state = %v and ", settlestate)
		}
	default:
		return nil, errors.New("错误的订单状态")
	}

	var orders []commonstruct.OrderS

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Select("roomid, port_id,count(*) as orderid, sum(amount) as amount,sum(wager) as wager").
		Where(fmt.Sprintf("%v%v", "uuid in (?) and roomid = ? and optime between ? and ? ", sqlArg), uuids, roomid, begintime, endtime).
		Group("roomid,port_id").
		Find(&orders).Error
	return orders, err
}

// 获取一期所有订单
func GetZengliangOrderSByExpect(roomid int64, expect string, orderid int64) ([]commonstruct.OrderS, error) {
	var OrderS []commonstruct.OrderS
	err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).
		Where("roomid = ? and expect = ? and orderid > ?", roomid, expect, orderid).
		Find(&OrderS).Error

	if err != nil {
		beego.Error("GetOrderS err", roomid, expect, orderid, err)
	}
	return OrderS, err
}

// 保存是否有中奖
func SetOrderWin(tablesuf string, roomid int64, orderid int64, winid int64) {
	//	tblname := fmt.Sprintf("%s%s", commonstruct.WY_gm_order_, tablesuf)
	var winlist []int64
	winlist = append(winlist, winid)
	winbyte, _ := json.Marshal(winlist)
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("orderid = ? and roomid = ?", orderid, roomid).
		Update("is_win", string(winbyte)).Error; err != nil {
		beego.Error("SetOrderWin err", tablesuf, orderid, err)
	}
}

// 清理详单记录
func ClearOrders(roomids []int64) {
	cleartime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().AddDate(0, 0, -1)), 10, 64)
	sql := fmt.Sprintf("delete from %v where optime < %v and roomid in (?) and state > 0;", commonstruct.WY_tmp_lotteryorder, cleartime, roomids)
	if err := WyMysql.Exec(sql).Error; err != nil {
		beego.Error("ClearOrders err ", err)
	}
}

// 连码订单保存
func SettleLianmaorder(orderid int64, amount float64, iswin int64, winlist string, winodds float64, wager float64, opencode string) error {
	var updateValues map[string]interface{}
	switch iswin {
	case -1:
		updateValues = map[string]interface{}{
			"is_win":      iswin,                   // 中奖
			"wager":       amount,                  // 派彩金额
			"state":       commonstruct.Ret_System, // 结算标识
			"settle_time": commonfunc.GetNowtime(), // 结算时间
			"opencode":    opencode,                // 开奖号码
			"shui_per":    0,                       // 退水金额
			"shui_amount": 0,                       // 退水金额
		}
	case 0:
		updateValues = map[string]interface{}{
			"state":       commonstruct.Ret_System,        // 结算标识
			"settle_time": commonfunc.GetNowtime(),        // 结算时间
			"opencode":    opencode,                       // 开奖号码
			"shui_amount": gorm.Expr("amount * shui_per"), // 退水金额
		}
	case 1:
		updateValues = map[string]interface{}{
			"is_win":      iswin,                          // 中奖
			"win_list":    winlist,                        // 中奖项
			"winodds":     winodds,                        // 中奖赔率
			"wager":       wager,                          // 派彩金额
			"state":       commonstruct.Ret_System,        // 结算标识
			"settle_time": commonfunc.GetNowtime(),        // 结算时间
			"opencode":    opencode,                       // 开奖号码
			"shui_amount": gorm.Expr("amount * shui_per"), // 退水金额
		}
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("orderid = ? ", orderid).
		Update(updateValues).Error; err != nil {
		beego.Error("SettleLiangmianWinorder err ", orderid, err)
		return err
	}
	return nil
}

// 两面中奖订单保存
func SettleLiangmianWinorder(roomid int64, expect string, itemid int64, winlist string, opencode string) error {
	updateValues := map[string]interface{}{
		"is_win":      1,                              // 中奖
		"win_list":    winlist,                        // 中奖项
		"state":       commonstruct.Ret_System,        // 结算标识
		"settle_time": commonfunc.GetNowtime(),        // 结算时间
		"opencode":    opencode,                       // 开奖号码
		"wager":       gorm.Expr("amount * winodds"),  // 派彩金额
		"shui_amount": gorm.Expr("amount * shui_per"), // 退水金额
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("roomid = ? and expect = ? and state = 0 and item_id = ? ", roomid, expect, itemid).
		Update(updateValues).Error; err != nil {
		beego.Error("SettleLiangmianWinorder err ", roomid, expect, itemid, err)
		return err
	}
	return nil
}

// 两面输单保存
func SettleLiangmianLostorder(roomid int64, expect string, itemid int64, opencode string) error {
	updateValues := map[string]interface{}{
		"state":       commonstruct.Ret_System,        // 结算标识
		"settle_time": commonfunc.GetNowtime(),        // 结算时间
		"opencode":    opencode,                       // 开奖号码
		"shui_amount": gorm.Expr("amount * shui_per"), // 退水金额
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("roomid = ? and expect = ? and state = 0 and item_id = ? ", roomid, expect, itemid).
		Update(updateValues).Error; err != nil {
		beego.Error("SettleLiangmianLostorder err ", roomid, expect, itemid, err)
		return err
	}
	return nil
}

// 两面打平订单处理
func SettleLiangmianPingorder(roomid int64, expect string, itemid int64, opencode string) error {
	updateValues := map[string]interface{}{
		"is_win":      -1,                      // 中奖
		"wager":       gorm.Expr("amount"),     // 派彩金额
		"state":       commonstruct.Ret_System, // 结算标识
		"settle_time": commonfunc.GetNowtime(), // 结算时间
		"opencode":    opencode,                // 开奖号码
		"shui_per":    0,                       // 退水金额
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("roomid = ? and expect = ? and state = 0 and item_id = ? ", roomid, expect, itemid).
		Update(updateValues).Error; err != nil {
		beego.Error("SettleLiangmianPingorder err ", roomid, expect, itemid, err)
		return err
	}
	return nil
}

// 保存中奖号码
func SaveOrderOpencode(roomid int64, expect string, opencode string, iteminfo string) error {
	var updateValues map[string]interface{}
	updateValues = map[string]interface{}{
		"opencode": opencode,
		"iteminfo": iteminfo,
	}

	return WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and expect = ?", roomid, expect).Update(updateValues).Error
}

// 获取当期连码订单总数-用于拆分结算
func GetLianmaOrderSNum(roomid int64, expect string) int64 {

	var OrderS commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Select("count(*) as orderid").
		Where("expect = ? and roomid = ? and state = 0 and official > 0", expect, roomid).Find(&OrderS).Error
	if err != nil {
		beego.Error("GetOrderS err", roomid, expect, err)
	}
	return OrderS.Orderid
}

// 获取增量连码单
func GetZenglianglianmaOrders(roomid int64, expect string, orderid int64) ([]commonstruct.OrderS, error) {
	var OrderS []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("roomid = ? and expect = ? and orderid > ? and official > 0 and state = 0", roomid, expect, orderid).
		Find(&OrderS).Error

	if err != nil {
		beego.Error("GetZenglianglianmaOrders err", roomid, expect, orderid, err)
	}
	return OrderS, err
}

// 连码订单获取
func GetLianmaOrderS(roomid int64, expect string) ([]commonstruct.OrderS, error) {

	var logcount commonstruct.OrderS
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("expect = ? and roomid = ? and official > 0", expect, roomid).
		Find(&logcount).Error; err != nil {
		beego.Error("GetOrdersByExpect count err", roomid, expect, err)
	}

	var orders []commonstruct.OrderS
	count := int(math.Ceil(float64(logcount.Orderid) / dancibishu)) // 需要执行次数
	for i := 0; i < count; i++ {

		var danciorders []commonstruct.OrderS
		if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
			Where("expect = ? and roomid = ? and official > 0", expect, roomid).
			Limit(int(dancibishu)).Offset(i * int(dancibishu)).
			Find(&danciorders).Error; err != nil {
			return orders, err
		} else {
			beego.Error("GetOrdersByExpect ====> ", roomid, expect, logcount.Orderid, count, i, len(danciorders))
			for _, order := range danciorders {
				orders = append(orders, order)
			}
		}
	}

	return orders, nil

	// var OrderS []commonstruct.OrderS
	// err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
	// 	Limit(pagecount).Offset(pagenum * pagecount).Find(&OrderS).Error
	// if err != nil {
	// 	beego.Error("GetLianmaOrderS err", roomid, expect, pagenum, pagecount, err)
	// }
	// return OrderS, err
}

// 结算时，获取当期的所有未结算的两面玩法下注项
func GetOrderSByXCItemID(roomid int64, expect string) ([]commonstruct.OrderS, error) {
	selectarg := `distinct item_id`

	var OrderS []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Select(selectarg).
		Group("item_id").
		Where("expect = ? and roomid = ? and state = 0 and official = 0", expect, roomid).Find(&OrderS).Error
	if err != nil {
		beego.Error("GetOrderS err", roomid, expect, err)
	}
	return OrderS, err
}

// 获取用户下注项统计
func GetUserOrderstats(uuid int64, roomid int64, expect string) ([]commonstruct.OrderS, error) {
	var ret []commonstruct.OrderS
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("item_id,count(*) as orderid,sum(amount) as amount").
		Where("uuid = ? and roomid = ? and expect = ?", uuid, roomid, expect).
		Group("item_id").Find(&ret).Error; err != nil {
		beego.Error("GetUserOrderstats err", uuid, roomid, expect, err)
		return ret, err
	}
	return ret, nil
}

// 保存订单派彩金额
func TXUpdateOrderwager(tx *gorm.DB, orderid int64, wager float64) error {
	return tx.Table(commonstruct.WY_tmp_lotteryorder).Where("orderid = ?", orderid).Update("wager", wager).Error
}

// 获取用户日统计
func GetUserOrderstatistic(uuid int64, date int64) ([]commonstruct.OrderS, error) {
	var infos []commonstruct.OrderS

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("uuid,state,count(*) as orderid,sum(amount) as amount,sum(wager) as amount").
		Where("uuid = ? and settle_time between ? and ?", uuid, commonfunc.GetBegintime(date), commonfunc.GetEndtime(date)).Group("uuid,state").Order("uuid,state").Find(&infos).Error
	return infos, err
}

// 游戏单期按下注项统计
func GetOrderStatistic(roomid int64, expect string) []commonstruct.OrderS {
	var orders []commonstruct.OrderS
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("item_id,sum(amount) as amount").
		Where("roomid = ? and expect = ?", roomid, expect).Group("item_id").Find(&orders).Error; err != nil {
		beego.Error("GetOrderStatistic err ", roomid, expect, err)
	}

	return orders
}

// 订单作废
func RevokeCoorder(roomid int64, orderid int64, amount float64) error {
	var log commonstruct.OrderS
	tx := WyMysql.Begin()

	if err := tx.Table(commonstruct.WY_tmp_lotteryorder).Where("orderid = ? and roomid = ?", orderid, roomid).Find(&log).Error; err != nil {
		beego.Error("RevokeCoorder err", err)
		tx.Rollback()
		return err
	} else {
		updateValues := map[string]interface{}{
			"settle_time": commonfunc.GetNowtime(),
			"is_win":      0,
			"win_list":    "",
			"wager":       amount,
			"shui_per":    0,
			"shui_amount": 0,
			"state":       commonstruct.Ret_UnValid,
		}

		if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and orderid = ?", roomid, orderid).
			Update(updateValues).Error; err != nil {
			beego.Error("RevokeCoorder err ", orderid, roomid)
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("InitProfit commit err ", err)
		tx.Rollback()
		return err
	}
	return nil
}

// 根据游戏和订单号获取订单
func GetOrdersByOrderid(roomid int64, orderid int64) (commonstruct.OrderS, error) {
	var orders commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and orderid = ?", roomid, orderid).Find(&orders).Error
	return orders, err
}

// 查询单期订单页码信息
func GetExpectorderPageinfo(roomid int64, expect string, masterid int64, pagecount int) (int, int) {

	var logs commonstruct.OrderS
	//	var sqlArg string
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("roomid = ? and expect = ? and master_id = ?", roomid, expect, masterid).
		Find(&logs).Error; err != nil {
		beego.Error("GetExpectorderPageinfo err", roomid, expect, err)
	}
	return int(logs.Orderid), int(math.Ceil(float64(logs.Orderid) / float64(pagecount)))
}

// 查询单期订单列表
func GetExpectorderS(roomid int64, expect string, masterid int64, pagenum int, pagecount int) []commonstruct.OrderS {
	var orders []commonstruct.OrderS
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and expect = ? and master_id = ?", roomid, expect, masterid).
		Order("optime desc,orderid").Limit(pagecount).Offset((pagenum - 1) * pagecount).
		Find(&orders).Error; err != nil {
		beego.Error("GetExpectorderS err", roomid, expect, pagenum, err)
		return orders
	}
	return orders
}

// 根据游戏和期号获取订单
func GetExpectAllorderS(masterid int64, roomid int64, expect string) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and master_id = ? and expect = ?", roomid, masterid, expect).Find(&orders).Error
	return orders, err
}

// 获取用户日初未结算订单
func GetDateUnsettle(uuid int64, begindate int64) (commonstruct.OrderS, error) {

	timeMin := commonfunc.GetBegintime(begindate)
	selectarg := fmt.Sprintf("master_id = %v and optime < %v and ((settle_time = 0 and state = 0) or settle_time > %v)",
		uuid, timeMin, timeMin)

	// 查询订单记录
	var logs commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("sum(amount) as amount").Where(selectarg).Find(&logs).Error

	return logs, err
}

// 获取当前下注金额
func GetItemamount(roomid int64, masterid int64, expect string) []commonstruct.OrderS {
	var datas []commonstruct.OrderS
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("item_id,sum(amount) as amount").
		Where("roomid = ? and master_id = ? and expect = ?", roomid, masterid, expect).Group("item_id").
		Find(&datas).Error; err != nil {
		beego.Error("GetItemamount err", err)
	}
	return datas
}

// 获取下注分页信息
func BKGetOrdersdetailSPageinfo(uuid int64, roomids []int64, begintime int64, endtime int64) (int, int) {
	// 查询订单记录
	var logs commonstruct.OrderS
	//	var sqlArg string
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("uuid = ? and roomid in (?) and optime between ? and ?", uuid, roomids, begintime, endtime).
		Find(&logs).Error; err != nil {
		beego.Error("GetOrdersdetailSPageinfo err", uuid, roomids, begintime, endtime, err)
	}
	return int(logs.Orderid), int(math.Ceil(float64(logs.Orderid) / 20))
}

// 根据状态获取订单分页信息
func GetOrdersdetailSPageinfo(uuids []int64, roomids []int64, pagecount int, begintime int64, endtime int64, orderstate int, settlestate int) (int64, int64) {
	// 查询订单记录
	var logs commonstruct.OrderS

	var sqlArg string
	switch orderstate {
	case -1:
		switch settlestate {
		case -1:
			sqlArg = ""
		default:
			sqlArg = fmt.Sprintf("state = %v and ", settlestate)
		}
	case 0:
		switch settlestate {
		case -1:
			sqlArg = "wager = 0 and "
		default:
			sqlArg = fmt.Sprintf("wager = 0 and state = %v and ", settlestate)
		}
	case 1:
		switch settlestate {
		case -1:
			sqlArg = "wager > 0 and "
		default:
			sqlArg = fmt.Sprintf("wager > 0 and state = %v and ", settlestate)
		}
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where(fmt.Sprintf("%v%v", sqlArg, "uuid in (?) and roomid in (?) and optime between ? and ?"), uuids, roomids, begintime, endtime).
		Find(&logs).Error; err != nil {
		beego.Error("GetOrdersdetailSPageinfo err", uuids, roomids, begintime, endtime, err)
	}
	return int64(logs.Orderid), int64(math.Ceil(float64(logs.Orderid) / float64(pagecount)))
}

// 获取多用户未结算订单金额
func GetTeamUnsettleamount(sufids []int64) commonstruct.OrderS {
	var info commonstruct.OrderS
	selectinfo := `sum(amount) as amount`
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select(selectinfo).Where("uuid in (?) and state = 0", sufids).
		Find(&info).Error; err != nil {
		beego.Error("GetTeamUnsettleamount err", err)
	}
	return info
}

// 重置已结算订单的状态
func ResetLotteryorderState(roomid int64, expect string) error {
	updateValues := map[string]interface{}{
		"state":       0,
		"settle_time": 0,
		"opencode":    "",
		"is_win":      0,
		"win_list":    "",
		"wager":       0,
		"shui_amount": 0,
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and expect = ? and state = 1", roomid, expect).
		Update(updateValues).Error; err != nil {
		beego.Error("ClearSettleinfo err ", err)
		return err
	}

	return nil
}

// 新建订单
func NewOrderUser(tx *gorm.DB, neworder commonstruct.OrderS) (commonstruct.OrderS, error) {

	newinfodb := tx.Table(commonstruct.WY_tmp_lotteryorder).Create(&neworder)
	if err := newinfodb.Error; err != nil {
		beego.Error("NewOrderUser err ", err)
		tx.Rollback()
		return neworder, errors.New(fmt.Sprintf("create userorder err %v", err.Error()))
	}

	newinfo := newinfodb.Value.(*commonstruct.OrderS)
	return *newinfo, nil
}

// 获取已结算订单
func GetTodayorders(uuid int64, issettled int, begintime int64, endtime int64) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("uuid = ?  and state >= ? and optime between ? and ?", uuid, issettled, begintime, endtime).
		Order("optime desc,orderid").Find(&orders).Error
	return orders, err
}

// 获取已结算订单
func GetOrdersdetailS(uuid int64, roomid int64, daleis []string, issettled int, offset int, begintime int64, endtime int64) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("uuid = ?  and state >= ? and lottery_dalei in (?) and optime between ? and ?", uuid, issettled, daleis, begintime, endtime).
		Order("optime desc,orderid").Limit(20).Offset((offset - 1) * 20).
		Find(&orders).Error
	return orders, err
}

// 分页获取单用户不同状态订单列表
func GetLotOrderinfoS(uuid int64, roomid int64, issettled int64, begindate int64, enddate int64, pagenum int64, pagecount int64) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	var sqlArg string
	if roomid == 0 {
		switch issettled {
		case -1:
			sqlArg = fmt.Sprintf("uuid = %v and optime between %v and %v", uuid, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		case 1:
			sqlArg = fmt.Sprintf("uuid = %v and state in (1,3) and optime between %v and %v", uuid, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		default:
			sqlArg = fmt.Sprintf("uuid = %v and state = %v and optime between %v and %v", uuid, issettled, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		}
	} else {
		switch issettled {
		case -1:
			sqlArg = fmt.Sprintf("uuid = %v and roomid = %v and optime between %v and %v", uuid, roomid, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		case 1:
			sqlArg = fmt.Sprintf("uuid = %v and roomid = %v and state in (1,3) and optime between %v and %v", uuid, roomid, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		default:
			sqlArg = fmt.Sprintf("uuid = %v and roomid = %v and state = %v and optime between %v and %v", uuid, roomid, issettled, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		}
	}

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where(sqlArg).
		Order("optime desc,orderid").Limit(int(pagecount)).Offset(int((pagenum - 1) * pagecount)).
		Find(&orders).Error

	if err != nil {
		beego.Error("GetLotOrderinfoS err", err.Error())
	}
	return orders, err
}

// 分页获取单用户不同状态订单分页信息
func GetLotOrderinfoSPageinfo(uuid int64, roomid int64, issettled int64, begindate int64, enddate int64, pagecount int64) (int, int) {
	var logs commonstruct.OrderS
	var sqlArg string
	if roomid == 0 {
		switch issettled {
		case -1:
			sqlArg = fmt.Sprintf("uuid = %v and optime between %v and %v", uuid, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		default:
			sqlArg = fmt.Sprintf("uuid = %v and state = %v and optime between %v and %v", uuid, issettled, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		}
	} else {
		switch issettled {
		case -1:
			sqlArg = fmt.Sprintf("uuid = %v and roomid = %v and optime between %v and %v", uuid, roomid, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		default:
			sqlArg = fmt.Sprintf("uuid = %v and roomid = %v and state = %v and optime between %v and %v", uuid, roomid, issettled, commonfunc.GetBegintime(begindate), commonfunc.GetEndtime(enddate))
		}
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").Where(sqlArg).
		Find(&logs).Error; err != nil {
		beego.Error("GetLotOrderinfoSPageinfo err", uuid, err)
	}
	return int(logs.Orderid), int(math.Ceil(float64(logs.Orderid) / float64(pagecount)))
}

// 分页获取单用户未结算订单列表
func GetUnsettleOrderinfoS(uuid int64, pagenum int64, pagecount int64) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("uuid = %v and state = 0").
		Order("optime desc,orderid").Limit(int(pagecount)).Offset(int((pagenum - 1) * pagecount)).
		Find(&orders).Error

	if err != nil {
		beego.Error("GetLotOrderinfoS err", err.Error())
	}
	return orders, err
}

// 分页获取单用户未结算订单分页信息
func GetUnsettleOrderinfoSPageinfo(uuid int64, pagecount int64) (int, int) {
	var logs commonstruct.OrderS
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").Where("uuid = %v and state = 0").
		Find(&logs).Error; err != nil {
		beego.Error("GetUnsettleOrderinfoSPageinfo err", uuid, err)
	}
	return int(logs.Orderid), int(math.Ceil(float64(logs.Orderid) / float64(pagecount)))
}

// 删除三个月前订单
func ClearExpireorder() {
	threemouth, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().AddDate(0, -3, 0)), 10, 64)
	sql := fmt.Sprintf("delete from %v where state > 0 and optime < %v;", commonstruct.WY_tmp_lotteryorder, threemouth)
	if err := WyMysql.Exec(sql).Error; err != nil {
		beego.Error("ClearExpireorder err ", err)
	}
}

// 按期号查询注单
func GetOrdersByExpect(roomid int64, expect string) ([]commonstruct.OrderS, error) {

	var logcount commonstruct.OrderS
	if err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("roomid = ? and expect = ?", roomid, expect).
		Find(&logcount).Error; err != nil {
		beego.Error("GetOrdersByExpect count err", roomid, expect, err)
	}

	var orders []commonstruct.OrderS
	count := int(math.Ceil(float64(logcount.Orderid) / dancibishu)) // 需要执行次数
	for i := 0; i < count; i++ {

		var danciorders []commonstruct.OrderS
		if err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).
			Where("roomid = ? and expect = ?", roomid, expect).
			Limit(int(dancibishu)).Offset(i * int(dancibishu)).
			Find(&danciorders).Error; err != nil {
			return orders, err
		} else {
			beego.Error("GetOrdersByExpect ====> ", roomid, expect, logcount.Orderid, count, i, len(danciorders))

			for _, order := range danciorders {
				orders = append(orders, order)
			}
		}
	}

	// if err != nil {
	// 	beego.Error("GetOrdersByExpect err", roomid, expect, err)
	// }

	return orders, nil
}

// 获取当期打和单
func GetDaheordersByExpect(roomid int64, expect string) ([]commonstruct.OrderS, error) {

	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and expect = ? and is_win = -1", roomid, expect).
		Find(&orders).Error

	if err != nil {
		beego.Error("GetOrdersByExpect err", roomid, expect, err)
	}

	return orders, err
}

// 获取未统计订单总数-用于拆分统计
func GetUnstatsorderSNum(roomid int64, expect string) int64 {

	var OrderS commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Select("count(*) as orderid").
		Where("roomid = ? and expect = ? and state = 1", roomid, expect).
		Find(&OrderS).Error
	if err != nil {
		beego.Error("GetOrderS err", roomid, expect, err)
	}
	return OrderS.Orderid
}

// 获取未统计报表的订单
func GetUnstatsorderSByExpect(roomid int64, expect string, offset int, pagecount int) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("roomid = ? and expect = ? and state = 1", roomid, expect).
		Limit(pagecount).Offset(offset * pagecount).
		Find(&orders).Error

	if err != nil {
		beego.Error("GetOrdersByExpect err", roomid, expect, err)
	}

	return orders, err
}

// 单用户多游戏订单分页信息
func GetOrdershisSPageinfo(uuid int64, roomids []int64, begintime int64, endtime int64) (int, int) {
	// 查询订单记录
	var logs commonstruct.OrderS
	//	var sqlArg string
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("uuid = ? and roomid in (?) and optime between ? and ?", uuid, roomids, begintime, endtime).
		Find(&logs).Error; err != nil {
		beego.Error("GetOrdershisSPageinfo err", uuid, roomids, begintime, endtime, err)
	}
	return int(logs.Orderid), int(logs.Orderid/20 + 1)
}

// 分页查询多用户多游戏订单列表
func GetBuhuoorderSHisByDate(uuids []int64, roomids []int64, begintime int64, endtime int64) ([]commonstruct.OrderS, error) {

	// 查询订单记录
	var logcount commonstruct.OrderS
	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("uuid in (?) and roomid in (?) and is_zidongbuhuo = 1 and optime between ? and ?", uuids, roomids, begintime, endtime).
		Find(&logcount).Error; err != nil {
		beego.Error("GetBuhuoorderSHisByDate err", uuids, roomids, begintime, endtime, err)
		return nil, err
	}

	if logcount.Orderid > 5000 {
		return nil, errors.New(fmt.Sprintf("数据量过长[%d],请缩短查询期限!", logcount.Orderid))
	} else {
		var orders []commonstruct.OrderS

		count := int(math.Ceil(float64(logcount.Orderid) / dancibishu)) // 需要执行次数
		for i := 0; i < count; i++ {

			var danciorders []commonstruct.OrderS
			if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
				Where("uuid in (?) and roomid in (?) and is_zidongbuhuo = 1 and optime between ? and ?", uuids, roomids, begintime, endtime).
				Limit(int(dancibishu)).Offset(i * int(dancibishu)).
				Order("optime desc,orderid").
				Find(&danciorders).Error; err != nil {
				return orders, err
			} else {
				beego.Error("GetBuhuoorderSHisByDate ====> ", begintime, endtime, logcount.Orderid, count, i, len(danciorders))

				for _, order := range danciorders {
					orders = append(orders, order)
				}
			}
		}
		return orders, nil
	}
}

// 分页查询多用户多游戏订单列表
func GetOrderSHisByDate(uuids []int64, roomids []int64, offset int, pagecount int, begintime int64, endtime int64, orderstate int, settlestate int) ([]commonstruct.OrderS, error) {

	var sqlArg string
	switch orderstate {
	case -1:
		switch settlestate {
		case -1:
			sqlArg = ""
		default:
			sqlArg = fmt.Sprintf("state = %v and ", settlestate)
		}
	case 0:
		switch settlestate {
		case -1:
			sqlArg = "is_win = 0 and "
		default:
			sqlArg = fmt.Sprintf("is_win = 0 and state = %v and ", settlestate)
		}
	case 1:
		switch settlestate {
		case -1:
			sqlArg = "wager > 0 and "
		default:
			sqlArg = fmt.Sprintf("is_win = 1 and state = %v and ", settlestate)
		}
	default:
		return nil, errors.New("错误的订单状态")
	}

	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where(fmt.Sprintf("%v%v", sqlArg, "uuid in (?) and roomid in (?) and optime between ? and ?"), uuids, roomids, begintime, endtime).
		Order("optime desc,orderid").Limit(pagecount).Offset((offset - 1) * pagecount).
		Find(&orders).Error

	if err != nil {
		beego.Error(sqlArg, uuids, roomids, begintime, endtime, err.Error())
	}
	return orders, err
}

//
func GetOrdersSumAmountByDate(uuids []int64, roomids []int64, begintime int64, endtime int64, orderstate int, settlestate int) (commonstruct.OrderS, error) {
	var order commonstruct.OrderS

	var sqlArg string
	switch orderstate {
	case -1:
		switch settlestate {
		case -1:
			sqlArg = "state != 3 and "
		default:
			sqlArg = fmt.Sprintf("state = %v and ", settlestate)
		}
	case 0:
		switch settlestate {
		case -1:
			sqlArg = "is_win = 0 and state != 3 and "
		default:
			sqlArg = fmt.Sprintf("is_win = 0 and state = %v and ", settlestate)
		}
	case 1:
		switch settlestate {
		case -1:
			sqlArg = "wager > 0 and state != 3 and "
		default:
			sqlArg = fmt.Sprintf("is_win = 1 and state = %v and ", settlestate)
		}
	default:
		return order, errors.New("错误的订单状态")
	}

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Select("sum(amount) as amount").
		Where(fmt.Sprintf("%v%v", sqlArg, "uuid in (?) and roomid in (?) and optime between ? and ?"), uuids, roomids, begintime, endtime).
		Find(&order).Error

	if err != nil {
		beego.Error(sqlArg, uuids, roomids, begintime, endtime, err.Error())
	}
	return order, err
}

// 分页查询代理的即时注单
func GetAgentJishizhudan(uuids []int64, roomid int64, expect string, offset int, pagecount int) ([]commonstruct.OrderS, error) {

	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("uuid in (?) and roomid = ? and expect = ? and state = 0", uuids, roomid, expect).
		Order("optime desc,orderid").Limit(pagecount).Offset((offset - 1) * pagecount).
		Find(&orders).Error

	if err != nil {
		beego.Error(uuids, roomid, expect, err.Error())
	}
	return orders, err
}

// 根据状态获取订单分页信息
func GetAgentJishizhudanPageinfo(uuids []int64, roomid int64, expect string, pagecount int) (int, int) {
	// 查询订单记录
	var logs commonstruct.OrderS

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("uuid in (?) and roomid = ? and expect = ? and state = 0", uuids, roomid, expect).
		Find(&logs).Error; err != nil {
		beego.Error("GetAgentJishizhudanPageinfo err", uuids, roomid, expect, err)
	}
	return int(logs.Orderid), int(math.Ceil(float64(logs.Orderid) / float64(pagecount)))
}

// 单用户单期订单统计
func GetWagerstatsByExpect(roomid int64, expect string) ([]commonstruct.OrderS, error) {
	// 查询订单记录
	var logs []commonstruct.OrderS
	selectinfo := `
	uuid,
	count(*) as orderid,
	sum(amount) as amount,
	sum(wager) as wager,
	sum(shui_amount) as shui_amount`

	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select(selectinfo).
		Where("roomid = ? and expect = ? and state = 1 and is_win in (0,1)", roomid, expect).
		Group("uuid").
		Find(&logs).Error
	if err != nil {
		beego.Error("GetWagerstatsByExpect err", roomid, expect, err)
	}

	return logs, err
}

// 代理直属的订单列表
func GetOrdersByPreid(preid int64, roomid int64, expect string) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("pre_id = ? and roomid = ? and expect = ?", preid, roomid, expect).Order("uuid,optime desc").Find(&orders).Error
	//		if err.Error() == "record not found" {
	//			return 0, "账号不存在"
	//		}
	return orders, err
}

// 撤单
func RevokeOrder(roomid int64, orderid int64) error {
	tx := WyMysql.Begin()

	var order commonstruct.OrderS
	if err := tx.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and orderid = ?", roomid, orderid).
		Find(&order).Error; err != nil {
		beego.Error("getorder err ", orderid, roomid, err)
		tx.Rollback()
		return err
	}
	if order.SettleTime > 0 {
		tx.Rollback()
		beego.Error("RevokeOrder err ", orderid, order.SettleTime)
		return errors.New("订单状态异常")
	}

	updateValues1 := map[string]interface{}{
		"state":       commonstruct.Ret_Revoke,
		"settle_time": commonfunc.GetNowtime(),
		"wager":       order.Amount,
	}

	if err := tx.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and orderid = ?", roomid, orderid).
		Update(updateValues1).Error; err != nil {
		beego.Error("RevokeOrder err ", orderid, roomid)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		beego.Error("RevokeOrder commit err ", err)
		tx.Rollback()
		return err
	}
	return nil
}

// 获取单个订单信息
func GetOrdersinfo(roomid int64, orderid int64) (commonstruct.OrderS, error) {
	var orders commonstruct.OrderS
	err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).Where("roomid = ? and orderid = ?", roomid, orderid).Find(&orders).Error
	//		if err.Error() == "record not found" {
	//			return 0, "账号不存在"
	//		}
	return orders, err
}

// 下注玩法订单分页信息
func GetPortOrdersdetailSPageinfo(uuids []int64, roomid int64, portid int64, settlestate int64, pagecount int, begintime int64, endtime int64) (int, int) {
	// 查询订单记录
	var logs commonstruct.OrderS

	switch settlestate {
	case 0:
		if err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
			Where("uuid in (?) and roomid = ? and state = 0 and optime between ? and ? and port_id = ?", uuids, roomid, begintime, endtime, portid).
			Find(&logs).Error; err != nil {
			beego.Error("GetOrdersdetailSPageinfo err", uuids, roomid, begintime, endtime, err)
		}
	default:
		if err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
			Where("uuid in (?) and roomid = ? and optime between ? and ? and port_id = ?", uuids, roomid, begintime, endtime, portid).
			Find(&logs).Error; err != nil {
			beego.Error("GetOrdersdetailSPageinfo err", uuids, roomid, begintime, endtime, err)
		}
	}

	return int(logs.Orderid), int(math.Ceil(float64(logs.Orderid) / float64(pagecount)))
}

// 下注玩法订单列表
func GetPortOrderSHisByDate(uuids []int64, roomid int64, portid int64, settlestate int64, offset int, pagecount int, begintime int64, endtime int64) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	var err error

	switch settlestate {
	case 0:
		err = WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).
			Where("uuid in (?) and roomid =? and state = 0 and optime between ? and ? and port_id = ?", uuids, roomid, begintime, endtime, portid).
			Order("optime desc,orderid").Limit(pagecount).Offset((offset - 1) * pagecount).
			Find(&orders).Error
	default:
		err = WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).
			Where("uuid in (?) and roomid =? and optime between ? and ? and port_id = ?", uuids, roomid, begintime, endtime, portid).
			Order("optime desc,orderid").Limit(pagecount).Offset((offset - 1) * pagecount).
			Find(&orders).Error
	}
	return orders, err
}

// 下注玩法订单分页信息
func GetPortJishizhudansPageinfo(uuids []int64, roomid int64, expect string, portids []int64, pagecount int) (int, int) {
	// 查询订单记录
	var logs commonstruct.OrderS

	if err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("uuid in (?) and roomid = ? and expect = ? and port_id in (?)", uuids, roomid, expect, portids).
		Find(&logs).Error; err != nil {
		beego.Error("GetOrdersdetailSPageinfo err", uuids, roomid, err)
	}

	return int(logs.Orderid), int(math.Ceil(float64(logs.Orderid) / float64(pagecount)))
}

// 下注玩法订单列表
func GetPortJishizhudans(uuids []int64, roomid int64, expect string, portids []int64, offset int, pagecount int) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).
		Where("uuid in (?) and roomid = ? and expect = ? and port_id in (?)", uuids, roomid, expect, portids).
		Order("optime desc,orderid").Limit(pagecount).Offset((offset - 1) * pagecount).
		Find(&orders).Error

	return orders, err
}

// 下注玩法订单分页信息
func GetItemJishizhudansPageinfo(uuids []int64, roomid int64, expect string, itemids []int64, pagecount int) (int, int) {
	// 查询订单记录
	var logs commonstruct.OrderS

	if err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).Select("count(*) as orderid").
		Where("uuid in (?) and roomid = ? and expect = ? and item_id in (?)", uuids, roomid, expect, itemids).
		Find(&logs).Error; err != nil {
		beego.Error("GetItemJishizhudansPageinfo err", uuids, roomid, err)
	}

	return int(logs.Orderid), int(math.Ceil(float64(logs.Orderid) / float64(pagecount)))
}

// 下注玩法订单列表
func GetItemJishizhudans(uuids []int64, roomid int64, expect string, itemids []int64, offset int, pagecount int) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql_read.Table(commonstruct.WY_tmp_lotteryorder).
		Where("uuid in (?) and roomid = ? and expect = ? and item_id in (?)", uuids, roomid, expect, itemids).
		Order("optime desc,orderid desc").Limit(pagecount).Offset((offset - 1) * pagecount).
		Find(&orders).Error

	return orders, err
}

// 下注玩法订单列表
func GetHuoliangLog(orderid int64, uuid int64) (commonstruct.HuoliangLog, error) {
	var logs commonstruct.HuoliangLog
	err := WyMysql_read.Table(commonstruct.WY_tmp_log_huoliang).
		Where("order_id = ? and uuid = ? ", orderid, uuid).Find(&logs).Error

	return logs, err
}

// 查询用户未计算金额合计
func QueryUnsettlemoney(uuid int64) float64 {
	var unsettle commonstruct.OrderS

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Select("sum(amount) as amount").Where("uuid = ? and state = 0", uuid).Find(&unsettle).Error; err != nil {
		beego.Error("QueryUnsettlemoney [%v] err %v", uuid, err)
		return 0
	}

	return unsettle.Amount
}

// 获取时间段内的下注用户
func GetOrderuuids(begindate int64, enddate int64) []int64 {
	var orders []commonstruct.OrderS

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Select("distinct uuid").Group("uuid").
		Where("optime between ? and ?", begindate*1000000, enddate*1000000+235959).
		Find(&orders).Error; err != nil {
		beego.Error("GetOrderuuids [%v]-[%v] err %v", begindate, enddate, err)
		return nil
	}

	var retuuids []int64
	for _, orderinfo := range orders {
		retuuids = append(retuuids, orderinfo.Uuid)
	}

	return retuuids
}

// 获取时间段内的下注用户
func GetOrderuuidsByExpect(roomid int64, expect string) []int64 {
	var orders []commonstruct.LotdataUseritem

	if err := WyMysql.Table(commonstruct.WY_tmp_user_lotdata_item).
		Select("distinct uuid").Group("uuid").
		Where("room_id = ? and expect = ? and order_amount > 0", roomid, expect).
		Find(&orders).Error; err != nil {
		beego.Error("GetOrderuuidsByExpect [%v]-[%v] err %v", roomid, expect, err)
		return nil
	}

	var retuuids []int64
	for _, orderinfo := range orders {
		retuuids = append(retuuids, orderinfo.Uuid)
	}

	return retuuids
}

// 获取用户未结算统计
func GetUserUnsettlestatistic(uuids []int64, roomids []int64, begindate int64, enddate int64) commonstruct.OrderS {
	var orders commonstruct.OrderS

	nextdate := GetNextDate(enddate)
	rettime := nextdate*1000000 + 60000

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Select("count(*) as orderid, sum(amount) as amount").
		Where("uuid in (?) and roomid in(?) and state = 0 and optime between ? and ?", uuids, roomids, begindate*1000000+60000, rettime).
		Find(&orders).Error; err != nil {
		beego.Error("GetUserUnsettlestatistic [%v]-[%v] err %v", begindate, enddate, err)
		return orders
	}

	return orders
}

// 按玩法ID获取用户未结算统计
func GetPortUnsettlestatistic(uuids []int64, roomid int64, begindate int64, enddate int64) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS

	nextdate := GetNextDate(enddate)
	rettime := nextdate*1000000 + 60000

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Select("port_id, count(*) as orderid, sum(amount) as amount").
		Where("uuid in (?) and roomid = ? and state = 0 and optime between ? and ?", uuids, roomid, begindate*1000000+60000, rettime).
		Group("port_id").
		Find(&orders).Error; err != nil {
		beego.Error("GetPortUnsettlestatistic [%v]-[%v] err %v", begindate, enddate, err)
		return orders, nil
	}

	return orders, nil
}

// 连码订单保存
func UpdateErrororder(orderid int64, zhanchenginfo string, shui_per float64, shui_amount float64) error {

	updateValues := map[string]interface{}{
		"zhancheng_info": zhanchenginfo, // 中奖
		"shui_per":       shui_per,      // 派彩金额
		"shui_amount":    shui_amount,   // 结算标识
	}

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("orderid = ? ", orderid).
		Update(updateValues).Error; err != nil {
		beego.Error("UpdateErrororder err ", orderid, err)
		return err
	}
	return nil
}

func GetOrdersByTime(begintime int64, endtime int64) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("optime between ? and ?", begintime, endtime).Find(&orders).Error

	if err != nil {
		beego.Error("GetOrdersByTime err", err)
	}

	return orders, err
}

func GetRelatedOrderS(orderid int64) ([]commonstruct.OrderS, error) {
	var orders []commonstruct.OrderS
	err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).Where("parent_id = ?", orderid).Find(&orders).Error

	if err != nil {
		beego.Error("GetOrdersByTime err", err)
	}

	return orders, err
}

// 用户按期号查询注单
func GetUserorderSByExpect(uuid int64, roomid int64, expect string) ([]commonstruct.OrderS, error) {

	var orders []commonstruct.OrderS

	if err := WyMysql.Table(commonstruct.WY_tmp_lotteryorder).
		Where("uuid = ? and roomid = ? and expect = ?", uuid, roomid, expect).
		Find(&orders).Error; err != nil {
		return orders, err
	}
	return orders, nil
}
