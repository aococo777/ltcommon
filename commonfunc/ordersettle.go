package commonfunc

import (
	"encoding/json"
	"github.com/aococo777/ltcommon/commonstruct"

	"github.com/astaxie/beego"
)

// // 获取订单的退款数据
// func GetRevokeAmountdata(orderinfo commonstruct.OrderS) []commonstruct.UserItemstatistic {
// 	var retdataS []commonstruct.UserItemstatistic

// 	// 会员部分
// 	var playerstats commonstruct.UserItemstatistic // 会员部分合计
// 	playerstats.Uuid = orderinfo.Uuid
// 	playerstats.Pan = orderinfo.Pan
// 	playerstats.Expect = orderinfo.Expect
// 	playerstats.PortID = orderinfo.PortID
// 	playerstats.ItemID = orderinfo.ItemID
// 	playerstats.Settlenum = 1
// 	playerstats.Settledamount = orderinfo.Amount
// 	playerstats.Validamount = orderinfo.Amount
// 	playerstats.Wager = -(orderinfo.Wager)
// 	playerstats.Shuiamount = -(orderinfo.ShuiAmount)

// 	retdataS = append(retdataS, playerstats)

// 	// 代理部分
// 	var agentperS []commonstruct.AgentZhanchenginfo
// 	if err := json.Unmarshal([]byte(orderinfo.ZhanchengInfo), &agentperS); err != nil {
// 		beego.Error("Unmarshal err", orderinfo.Orderid, orderinfo.ZhanchengInfo)
// 		return retdataS
// 	}

// 	var sufZhancheng float64 = 0
// 	// var sufid int64 = 0
// 	var sufshangjiao float64 = 0
// 	for i := len(agentperS) - 1; i > -1; i-- {
// 		if agentperS[i].Uuid == orderinfo.Uuid {
// 			// 会员初次产生上交货量
// 			sufshangjiao = playerstats.Validamount - playerstats.Wager - playerstats.Shuiamount
// 			// sufid = agentperS[i].Uuid
// 		} else {
// 			// if sufid != orderinfo.Uuid { // 代理部分合计
// 			var agentstats commonstruct.UserItemstatistic
// 			agentstats.Uuid = agentperS[i].Uuid
// 			agentstats.Pan = orderinfo.Pan
// 			agentstats.Expect = orderinfo.Expect
// 			agentstats.PortID = orderinfo.PortID
// 			agentstats.ItemID = orderinfo.ItemID

// 			agentstats.Yingshouxiaxian = sufshangjiao
// 			agentstats.Shizhanhuoliang = playerstats.Validamount * agentperS[i].ZhanchengPer / 100
// 			agentstats.Shizhanshuying = (playerstats.Validamount - playerstats.Wager) * agentperS[i].ZhanchengPer / 100
// 			agentstats.Shizhantuishui = -agentstats.Shizhanhuoliang * (agentperS[i].Tuishui - agentperS[i].ProfitTuishui)
// 			agentstats.Shizhanpeicha = 0
// 			agentstats.ProfitWager = 0
// 			agentstats.Shizhanjieguo = agentstats.Shizhanshuying + agentstats.Shizhantuishui
// 			sufZhancheng = sufZhancheng + agentperS[i].ZhanchengPer

// 			agentstats.Shangjiaohuoliang = playerstats.Validamount * (1 - sufZhancheng/100)
// 			agentstats.Porfittuishui = agentstats.Shangjiaohuoliang * agentperS[i].ProfitTuishui
// 			agentstats.Yingkuijieguo = -(agentstats.Shizhanjieguo + agentstats.Porfittuishui)

// 			agentstats.Shouhuo = playerstats.Validamount * agentperS[i].ShouhuoPer / 100
// 			agentstats.Shouhuoyingkui = (playerstats.Validamount-playerstats.Wager)*agentperS[i].ShouhuoPer/100 +
// 				(-agentstats.Shouhuo * (agentperS[i].Tuishui - agentperS[i].ProfitTuishui)) // 收货后，要给出去的退水

// 			agentstats.Buchu = -playerstats.Validamount * agentperS[i].BuchuPer / 100 // 补出金额为负数
// 			agentstats.Buchuyingkui = (playerstats.Wager-playerstats.Validamount)*agentperS[i].BuchuPer/100 +
// 				(-agentstats.Buchu * agentperS[i].Tuishui) // 补出后，我要拿到的退水

// 			if orderinfo.RoleType == "agent" && agentstats.Uuid == orderinfo.Uuid {
// 				agentstats.Shoudongbuchu = -playerstats.Validamount
// 				agentstats.Shoudongbuchuyingkui = (playerstats.Wager - playerstats.Validamount) + playerstats.Validamount*(agentperS[i].Tuishui-agentperS[i].ProfitTuishui)
// 			}
// 			agentstats.Shizhanbuhuo = agentstats.Shouhuo + agentstats.Buchu
// 			agentstats.Shizhanbuhuoyingkui = agentstats.Shouhuoyingkui + agentstats.Buchuyingkui

// 			if agentstats.Shangjiaohuoliang < 0.000001 {
// 				agentstats.Shangjijiaoshou = 0
// 			} else {
// 				agentstats.Shangjijiaoshou = agentstats.Yingshouxiaxian - agentstats.Yingkuijieguo
// 			}

// 			retdataS = append(retdataS, agentstats)

// 			sufshangjiao = agentstats.Shangjijiaoshou
// 			// sufid = agentperS[i].Uuid
// 		}
// 	}

// 	return retdataS
// }

// 获取订单的逆报表数据
func GetSettledRevokeLotdata(orderinfo commonstruct.OrderS) []commonstruct.UserItemstatistic {
	var retdataS []commonstruct.UserItemstatistic

	// 会员部分
	var playerstats commonstruct.UserItemstatistic // 会员部分合计
	playerstats.Uuid = orderinfo.Uuid
	playerstats.Pan = orderinfo.Pan
	playerstats.Expect = orderinfo.Expect
	playerstats.PortID = orderinfo.PortID
	playerstats.ItemID = orderinfo.ItemID
	playerstats.Settlenum = 1
	playerstats.Settledamount = orderinfo.Amount
	playerstats.Validamount = orderinfo.Amount
	playerstats.Wager = orderinfo.Wager
	playerstats.Shuiamount = orderinfo.ShuiAmount

	// 代理部分
	var agentperS []commonstruct.AgentZhanchenginfo
	if err := json.Unmarshal([]byte(orderinfo.ZhanchengInfo), &agentperS); err != nil {
		beego.Error("Unmarshal err", orderinfo.Orderid, orderinfo.ZhanchengInfo)
		return retdataS
	}

	var sufZhancheng float64 = 0
	var sufid int64 = 0
	var sufshangjiao float64 = 0
	for i := len(agentperS) - 1; i > -1; i-- {
		if agentperS[i].Uuid == orderinfo.Uuid {
			// 会员初次产生上交货量
			sufshangjiao = playerstats.Validamount - playerstats.Wager - playerstats.Shuiamount
			sufid = agentperS[i].Uuid
		} else {
			// if sufid != orderinfo.Uuid { // 代理部分合计
			var agentstats commonstruct.UserItemstatistic
			agentstats.Uuid = sufid
			agentstats.Pan = orderinfo.Pan
			agentstats.Expect = orderinfo.Expect
			agentstats.PortID = orderinfo.PortID
			agentstats.ItemID = orderinfo.ItemID

			agentstats.Yingshouxiaxian = sufshangjiao
			agentstats.Shizhanhuoliang = playerstats.Validamount * agentperS[i].ZhanchengPer / 100
			agentstats.Shizhanshuying = (playerstats.Validamount - playerstats.Wager) * agentperS[i].ZhanchengPer / 100
			agentstats.Shizhantuishui = -agentstats.Shizhanhuoliang * (agentperS[i].Tuishui - agentperS[i].ProfitTuishui)
			agentstats.Shizhanpeicha = 0
			agentstats.ProfitWager = 0
			agentstats.Shizhanjieguo = agentstats.Shizhanshuying + agentstats.Shizhantuishui
			sufZhancheng = sufZhancheng + agentperS[i].ZhanchengPer

			agentstats.Shangjiaohuoliang = playerstats.Validamount * (1 - sufZhancheng/100)
			agentstats.Porfittuishui = agentstats.Shangjiaohuoliang * agentperS[i].ProfitTuishui
			agentstats.Yingkuijieguo = agentstats.Shizhanjieguo + agentstats.Porfittuishui

			// agentstats.Shouhuo = playerstats.Validamount * agentperS[i].ShouhuoPer / 100
			// agentstats.Shouhuoyingkui = (playerstats.Validamount-playerstats.Wager)*agentperS[i].ShouhuoPer/100 +
			// 	(-agentstats.Shouhuo * (agentperS[i].Tuishui - agentperS[i].ProfitTuishui)) // 收货后，要给出去的退水

			// agentstats.Buchu = -playerstats.Validamount * agentperS[i].BuchuPer / 100 // 补出金额为负数
			// agentstats.Buchuyingkui = (playerstats.Wager-playerstats.Validamount)*agentperS[i].BuchuPer/100 +
			// 	(-agentstats.Buchu * agentperS[i].Tuishui) // 补出后，我要拿到的退水

			if orderinfo.RoleType == "agent" && agentstats.Uuid == orderinfo.Uuid {
				agentstats.Shoudongbuchu = -playerstats.Validamount
				agentstats.Shoudongbuchuyingkui = (playerstats.Wager - playerstats.Validamount) + playerstats.Validamount*(agentperS[i].Tuishui-agentperS[i].ProfitTuishui)
			}
			agentstats.Shizhanbuhuo = agentstats.Shouhuo + agentstats.Buchu
			agentstats.Shizhanbuhuoyingkui = agentstats.Shouhuoyingkui + agentstats.Buchuyingkui

			if agentstats.Shangjiaohuoliang < 0.000001 {
				agentstats.Shangjijiaoshou = 0
			} else {
				agentstats.Shangjijiaoshou = agentstats.Yingshouxiaxian - agentstats.Yingkuijieguo
			}

			sufshangjiao = agentstats.Shangjijiaoshou
			sufid = agentperS[i].Uuid

			agentstats.ProfitWager = -agentstats.ProfitWager
			agentstats.Porfittuishui = -agentstats.Porfittuishui

			agentstats.Yingshouxiaxian = -agentstats.Yingshouxiaxian
			agentstats.Shizhanhuoliang = -agentstats.Shizhanhuoliang
			agentstats.Shizhanshuying = -agentstats.Shizhanshuying
			agentstats.Shizhanjieguo = -agentstats.Shizhanjieguo
			agentstats.Shizhantuishui = -agentstats.Shizhantuishui
			agentstats.Shizhanpeicha = -agentstats.Shizhanpeicha
			agentstats.Yingkuijieguo = -agentstats.Yingkuijieguo

			agentstats.Shangjiaohuoliang = -agentstats.Shangjiaohuoliang
			agentstats.Shangjijiaoshou = -agentstats.Shangjijiaoshou

			agentstats.Shouhuo = -agentstats.Shouhuo
			agentstats.Shouhuoyingkui = -agentstats.Shouhuoyingkui
			agentstats.Buchu = -agentstats.Buchu
			agentstats.Buchuyingkui = -agentstats.Buchuyingkui

			agentstats.Shoudongbuchu = -agentstats.Shoudongbuchu
			agentstats.Shoudongbuchuyingkui = -agentstats.Shoudongbuchuyingkui
			agentstats.Shizhanbuhuo = -agentstats.Shizhanbuhuo
			agentstats.Shizhanbuhuoyingkui = -agentstats.Shizhanbuhuoyingkui
			retdataS = append(retdataS, agentstats)
		}
	}

	playerstats.Settlenum = -playerstats.Settlenum
	playerstats.Settledamount = -playerstats.Settledamount
	playerstats.Validamount = -playerstats.Validamount
	playerstats.Wager = -playerstats.Wager
	playerstats.Shuiamount = -playerstats.Shuiamount

	retdataS = append(retdataS, playerstats)

	return retdataS
}

// 获取未结算订单的逆报表数据
func GetUnSettleRevokeLotdata(orderinfo commonstruct.OrderS) []commonstruct.UnsettleRevokedata {
	var retdataS []commonstruct.UnsettleRevokedata

	// 会员部分
	var playerstats commonstruct.UnsettleRevokedata // 会员部分合计
	playerstats.Uuid = orderinfo.Uuid
	playerstats.Pan = orderinfo.Pan
	playerstats.Expect = orderinfo.Expect
	playerstats.PortID = orderinfo.PortID
	playerstats.ItemID = orderinfo.ItemID
	playerstats.OrderNum = -1
	playerstats.OrderAmount = -orderinfo.Amount

	// 代理部分
	var agentperS []commonstruct.AgentZhanchenginfo
	if err := json.Unmarshal([]byte(orderinfo.ZhanchengInfo), &agentperS); err != nil {
		beego.Error("Unmarshal err", orderinfo.Orderid, orderinfo.ZhanchengInfo)
		return retdataS
	}

	for i := len(agentperS) - 1; i > -1; i-- {
		if agentperS[i].Uuid == orderinfo.Uuid {
			// 会员初次产生上交货量
		} else {
			// if sufid != orderinfo.Uuid { // 代理部分合计
			var agentstats commonstruct.UnsettleRevokedata
			agentstats.Uuid = agentperS[i].Uuid
			agentstats.Pan = orderinfo.Pan
			agentstats.Expect = orderinfo.Expect
			agentstats.PortID = orderinfo.PortID
			agentstats.ItemID = orderinfo.ItemID

			agentstats.Zhanchenghuoliang = -orderinfo.Amount * agentperS[i].BaseZhancheng / 100
			agentstats.Shizhanhuoliang = -orderinfo.Amount * agentperS[i].ZhanchengPer / 100
			// agentstats.Xiajibuhuo = -orderinfo.Amount * agentperS[i].ShouhuoPer / 100 // 下级补货
			// agentstats.Zidongbuchu = -orderinfo.Amount * agentperS[i].BuchuPer / 100  // 自动补出

			retdataS = append(retdataS, agentstats)
		}
	}

	retdataS = append(retdataS, playerstats)

	return retdataS
}
