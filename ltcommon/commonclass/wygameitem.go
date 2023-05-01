package commonclass

import (
	"fmt"
	"ltback/src/ltcommon/commonfunc"
	"ltback/src/ltcommon/commonstruct"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

// 游戏结果分类
type GameItemManager struct {
	ItemInfo_ID   map[string]int64                 // 游戏结果分类
	ItemID_Info   map[int64]commonstruct.CGameItem // Item的详细信息
	PortInfo_ID   map[string]int64                 // 盘口设置分类
	PortID_Info   map[int64]commonstruct.PortClass // 盘口ID 的详细信息
	GameENG_ItemS map[string][]int64               // 单个游戏下的所有玩法
	// GameENG_Info  map[string]commonstruct.CRoomGame // 游戏名 匹配 游戏信息
	// GameID_Info   map[int64]commonstruct.CRoomGame  // 游戏ID 获取 游戏信息
	// RoomID_Info   map[int64]commonstruct.CRoomInfo  // 房间ID获取房间信息
	// RoomName_Info map[string]commonstruct.CRoomInfo // 房间名字获取房间信息
	Sebo          map[int64]string  // 色波信息
	LianmaID      map[int64]int64   // 一中一连码
	MultiLianmaID []int64           // 一中多连码
	RiskItemID    map[int64][]int64 // 风险对打ID列表

	JszdGroupIDList map[string][]int64  // 即时注单分组 gametype_groupname => IDList
	JszdGroupList   map[string][]string // 即时注单分组 gametype => groupList

	mysql    *gorm.DB
	itemLock sync.Mutex // 生肖锁
}

func (this *GameItemManager) Init(DB *gorm.DB) error {
	beego.Error("GameItemManager Init enter!")

	this.mysql = DB

	// 初始化盘口设置分类
	this.UpdatePortclass()

	// 初始化色波
	this.UpdateSebo()

	// 初始化 连码下注项 ( 加了一个风险对冲项 )
	this.UpdateLianmaItem()

	// 初始化即时注单分组
	this.UpdateJszdGroup()

	go this.TimerUpdate()

	return nil
}

func (this *GameItemManager) GetAllGameitem() ([]commonstruct.CGameItem, error) {
	var dancibishu float64 = 1000

	var logcount commonstruct.CGameItem
	if err := this.mysql.Table(commonstruct.WY_gm_config_item).Select("count(*) as id").
		Find(&logcount).Error; err != nil {
		beego.Error("GetAllGameitem count err", err)
	}

	var orders []commonstruct.CGameItem
	count := int(math.Ceil(float64(logcount.ID) / dancibishu)) // 需要执行次数
	for i := 0; i < count; i++ {
		var danciorders []commonstruct.CGameItem
		if err := this.mysql.Table(commonstruct.WY_gm_config_item).
			Limit(int(dancibishu)).Offset(i * int(dancibishu)).
			Find(&danciorders).Error; err != nil {
			return orders, err
		} else {
			for _, order := range danciorders {
				orders = append(orders, order)
			}
		}
	}
	return orders, nil
}

// 初始化即时注单分组
func (this *GameItemManager) UpdateJszdGroup() {
	this.JszdGroupIDList = make(map[string][]int64)
	this.JszdGroupList = make(map[string][]string)

	FenleiS, _ := this.GetAllGameitem()

	for _, v := range FenleiS {
		if v.JszdGroupName != "" {
			groupkey := fmt.Sprintf("%v_%v", v.GameType, v.JszdGroupName)
			if value, ok := this.JszdGroupIDList[groupkey]; ok {
				value = append(value, v.ID)
				this.JszdGroupIDList[groupkey] = value
			} else {
				this.JszdGroupIDList[groupkey] = []int64{v.ID}
			}

			if value, ok := this.JszdGroupList[v.GameType]; ok {
				bhave := false
				for _, groupname := range value {
					if groupname == v.JszdGroupName {
						bhave = true
						break
					}
				}
				if !bhave {
					value = append(value, v.JszdGroupName)
					this.JszdGroupList[v.GameType] = value
				}
			} else {
				this.JszdGroupList[v.GameType] = []string{v.JszdGroupName}
			}
		}
	}
}

func (this *GameItemManager) TimerUpdate() {
	this.UpdateGameItem()
	CheckTimer := time.NewTicker(5 * time.Second) // 刷新iteminfo
	for {
		select {
		case <-CheckTimer.C:
			now := commonfunc.BeijingTime()
			flg := (commonfunc.GetBjHour() == 0 && now.Minute() == 0 && (now.Second() >= 5 && now.Second() < 10))
			if flg {
				// 初始化 所有下注项 ( 需要每日定时获取,有修改年肖普肖的需求 )
				this.UpdateGameItem()
			}
		}
	}
}

// 获取item列表 (每日凌晨需要重新刷新,新年需要修改年肖项)
func (this *GameItemManager) UpdateGameItem() {
	beego.Error("UpdateGameItem enter ~~~~~~")

	this.itemLock.Lock()
	defer this.itemLock.Unlock()

	this.ItemInfo_ID = make(map[string]int64)
	this.ItemID_Info = make(map[int64]commonstruct.CGameItem)
	this.GameENG_ItemS = make(map[string][]int64)

	// 初始化结果分类表

	FenleiS, _ := this.GetAllGameitem()
	beego.Error("CGameItem length ==  ", len(FenleiS))

	for _, v := range FenleiS {
		// 检查item 是否在 portclass 分类里，防止赔率为零
		if err := this.checkiteminportclass(v.LotteryDalei, v.GameType, v.GameDalei, v.GameXiaolei, v.GameMode, v.GameItem); err != nil {
			beego.Error(v.LotteryDalei, v.GameType, v.GameDalei, v.GameXiaolei, v.GameMode, v.GameItem, " 系统未启用玩法 ")
		} else {
			key := fmt.Sprintf("%s_%s_%s_%s_%s_%s", v.LotteryDalei, v.GameType, v.GameDalei, v.GameXiaolei, v.GameMode, v.GameItem)
			this.ItemInfo_ID[key] = v.ID
			this.ItemID_Info[v.ID] = v

			this.AddLotteryGameItem(fmt.Sprintf("%s_%s", v.LotteryDalei, v.GameType), v.ID)
		}
	}
}

func (this *GameItemManager) UpdatePortclass() {
	this.PortInfo_ID = make(map[string]int64)
	this.PortID_Info = make(map[int64]commonstruct.PortClass)

	var FenleiS []commonstruct.PortClass
	if err := this.mysql.Table(commonstruct.WY_gm_config_portclass).Find(&FenleiS).Error; err != nil {
		beego.Error("Init Query WY_gm_config_portclass", err.Error())
		return
	}
	for _, v := range FenleiS {
		key := fmt.Sprintf("%s_%s_%s_%s", v.LotteryDalei, v.SettleType, v.GameDalei, v.GameXiaolei)
		this.PortInfo_ID[key] = v.ID
		this.PortID_Info[v.ID] = v
	}
}

func (this *GameItemManager) UpdateSebo() {
	this.Sebo = make(map[int64]string)
	var FenleiS []commonstruct.BallInfo
	if err := this.mysql.Table(commonstruct.WY_gm_config_ballinfo).Find(&FenleiS).Error; err != nil {
		beego.Error("Init Query WY_gm_config_ballinfo", err.Error())
		return
	}

	for _, v := range FenleiS {
		if v.Balltype == "红波" || v.Balltype == "绿波" || v.Balltype == "蓝波" {
			BallS := strings.Split(v.Balllist, ",")
			for _, ball := range BallS {
				num, _ := strconv.Atoi(ball)
				this.Sebo[int64(num)] = v.Balltype
			}
		}
	}
}

// 初始化 连码下注项 ( 加了一个风险对冲项 )
func (this *GameItemManager) UpdateLianmaItem() {
	this.LianmaID = make(map[int64]int64)
	// 对打下注项列表
	this.RiskItemID = make(map[int64][]int64)

	var FenleiS []commonstruct.CGameItem
	if err := this.mysql.Table(commonstruct.WY_gm_config_item).Find(&FenleiS).Error; err != nil {
		beego.Error("UpdateGameItem err ", err)
		return
	}

	for _, v := range FenleiS {

		// 检查item 是否在 portclass 分类里，防止赔率为零
		if err := this.checkiteminportclass(v.LotteryDalei, v.GameType, v.GameDalei, v.GameXiaolei, v.GameMode, v.GameItem); err != nil {
			beego.Error(v.LotteryDalei, v.GameType, v.GameDalei, v.GameXiaolei, v.GameMode, v.GameItem, err.Error(), " 系统未启用玩法 ")
		} else {
			switch v.Islianma {
			case 0:
				this.AddRiskGroupitemid(v.RiskGroupid, v.ID)
			default:
				this.LianmaID[v.ID] = v.Islianma
			}
		}
	}

	if len(this.LianmaID) > 0 {
		beego.Error("连码长度", len(this.LianmaID))
	} else {
		beego.Error("连码异常 无连码ID")
	}
}

// 检查下注项的分类是否在 portclass 里面
func (this *GameItemManager) checkiteminportclass(platform string, settletype string, game_dalei string, game_xiaolei string, GameMode string, GameItem string) error {

	var infos commonstruct.PortClass
	if err := this.mysql.Table(commonstruct.WY_gm_config_portclass).
		Where("lottery_dalei = ? and settle_type = ? and game_dalei = ? and game_xiaolei = ?", platform, settletype, game_dalei, game_xiaolei).
		Find(&infos).Error; err != nil {
		beego.Error("checkiteminportclass err", platform, settletype, game_dalei, game_xiaolei, GameMode, GameItem, err.Error())
		return err
	}
	return nil
}

// 普通连码
func (this *GameItemManager) IsLianma(itemid int64) int64 {
	value, _ := this.LianmaID[itemid]
	return value
}

// 一中多 连码
func (this *GameItemManager) IsMultiLianma(resultid int64) bool {
	for i := 0; i < len(this.MultiLianmaID); i++ {
		if this.MultiLianmaID[i] == resultid {
			return true
		}
	}
	return false
}

func (this *GameItemManager) AddLotteryGameItem(dalei_gametype string, id int64) {
	if value, ok := this.GameENG_ItemS[dalei_gametype]; ok {
		value = append(value, id)
		this.GameENG_ItemS[dalei_gametype] = value
	} else {
		this.GameENG_ItemS[dalei_gametype] = []int64{id}
	}
}

func (this *GameItemManager) AddRiskGroupitemid(riskgroupid int64, id int64) {
	if value, ok := this.RiskItemID[riskgroupid]; ok {
		value = append(value, id)
		this.RiskItemID[riskgroupid] = value
	} else {
		this.RiskItemID[riskgroupid] = []int64{id}
	}
}

func (this *GameItemManager) GetRiskGroupitemids(itemid int64) []int64 {

	iteminfo := this.GetGameitemInfo(itemid)
	if value, ok := this.RiskItemID[iteminfo.RiskGroupid]; ok {
		return value
	}
	return nil
}

// 获取游戏的所有玩法ID的Map
func (this *GameItemManager) GetGameItemS(dalei_gametype string) map[int64]int64 {
	ret := make(map[int64]int64)

	this.itemLock.Lock()
	defer this.itemLock.Unlock()

	if value, ok := this.GameENG_ItemS[dalei_gametype]; ok {
		for _, v := range value {
			ret[v] = 0
		}
	}
	return ret
}

// 获取游戏的所有玩法ID数组
func (this *GameItemManager) GetGameItemSArr(dalei_gametype string) []int64 {
	this.itemLock.Lock()
	defer this.itemLock.Unlock()

	if value, ok := this.GameENG_ItemS[dalei_gametype]; ok {
		return value
	}
	return nil
}

func (this *GameItemManager) GetGameitemInfo(id int64) commonstruct.CGameItem {
	var ret commonstruct.CGameItem

	this.itemLock.Lock()
	defer this.itemLock.Unlock()
	if info, ok := this.ItemID_Info[id]; ok {
		return info
	} else {
		beego.Error("GetGameitemInfo err", id)
	}
	return ret
}

func (this *GameItemManager) GetPortInfo(id int64) commonstruct.PortClass {
	var ret commonstruct.PortClass
	if info, ok := this.PortID_Info[id]; ok {
		return info
	}
	return ret
}

func (this *GameItemManager) GetGameItemID(platform string, settletype string, GameDalei string, gamexiaolei string, gamemode string, result string) int64 {
	this.itemLock.Lock()
	defer this.itemLock.Unlock()

	key := fmt.Sprintf("%s_%s_%s_%s_%s_%s", platform, settletype, GameDalei, gamexiaolei, gamemode, result)
	itemid := this.ItemInfo_ID[key]

	if itemid == 0 {
		beego.Error("GetGameItemID err", platform, settletype, GameDalei, gamexiaolei, gamemode, result)
	}
	return itemid
}

/**********************************
*通过下注ID 获取到赔率分类
**********************************/
func (this *GameItemManager) GetPortidByItemid(id int64) int64 {
	iteminfo := this.GetGameitemInfo(id)
	ret := this.GetPortID(iteminfo.LotteryDalei, iteminfo.GameType, iteminfo.GameDalei, iteminfo.GameXiaolei)
	return ret
}

/**********************************
*通过分类描述 获取的分类ID
**********************************/
func (this *GameItemManager) GetPortID(dalei string, gametype string, GameDalei string, GameXiaolei string) int64 {
	key := fmt.Sprintf("%s_%s_%s_%s", dalei, gametype, GameDalei, GameXiaolei)

	portid := this.PortInfo_ID[key]
	if portid == 0 {
		beego.Error("GetPortID err", dalei, gametype, GameDalei, GameXiaolei)
	}
	return portid
}

func (this *GameItemManager) GetColor(Num int64) string {
	return this.Sebo[Num]
}

func RemoveDuplicatesAndZero(a []int64) (ret []int64) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || a[i] == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

// 初始化即时注单分组
func (this *GameItemManager) GetJszdGrouplist(gametype string) []string {

	this.itemLock.Lock()
	defer this.itemLock.Unlock()

	if info, ok := this.JszdGroupList[gametype]; ok {
		return info
	} else {
		beego.Error("GetJszdGroupItemlist err", gametype)
	}
	return nil
}

// 初始化即时注单分组
func (this *GameItemManager) GetJszdGroupItemlist(gametype string, groupname string) []int64 {

	groupkey := fmt.Sprintf("%v_%v", gametype, groupname)
	this.itemLock.Lock()
	defer this.itemLock.Unlock()

	if info, ok := this.JszdGroupIDList[groupkey]; ok {
		return info
	} else {
		beego.Error("GetJszdGroupItemlist err", gametype, groupname)
	}
	return nil

	// if value, ok := this.JszdGroupIDList[groupkey]; ok {
	// 	value = append(value, v.ID)
	// 	this.JszdGroupIDList[groupkey] = value
	// } else {
	// 	this.JszdGroupIDList[groupkey] = []int64{v.ID}
	// }
}

// 是否为特码下注项
func (this *GameItemManager) IsTemaitem(itemid int64) bool {
	for temaA, temaB := range TemaItemMap {
		if temaA == itemid || temaB == itemid {
			return true
		}
	}
	return false
}

// 特码B=>特码A
func (this *GameItemManager) GetTemaAitemID(itemid int64) int64 {
	if itemidB, ok := TemaItemMap[itemid]; ok {
		return itemidB
	} else {
		for temaA, temaB := range TemaItemMap {
			if temaB == itemid {
				return temaA
			}
		}
	}
	return itemid
}

var (
	TemaItemMap = map[int64]int64{
		416: 475,
		417: 476,
		418: 477,
		419: 478,
		420: 479,
		421: 480,
		422: 481,
		423: 482,
		424: 483,
		425: 484,
		426: 485,
		427: 486,
		428: 487,
		429: 488,
		430: 489,
		431: 490,
		432: 491,
		433: 492,
		434: 493,
		435: 494,
		436: 495,
		437: 496,
		438: 497,
		439: 498,
		440: 499,
		441: 500,
		442: 501,
		443: 502,
		444: 503,
		445: 504,
		446: 505,
		447: 506,
		448: 507,
		449: 508,
		450: 509,
		451: 510,
		452: 511,
		453: 512,
		454: 513,
		455: 514,
		456: 515,
		457: 516,
		458: 517,
		459: 518,
		460: 519,
		461: 520,
		462: 521,
		463: 522,
		464: 523,
	}

	SscQiu1ItemS = []int64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	SscQiu2ItemS = []int64{
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	}
	SscQiu3ItemS = []int64{
		23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
	}
	SscQiu4ItemS = []int64{
		33, 34, 35, 36, 37, 38, 39, 40, 41, 42,
	}
	SscQiu5ItemS = []int64{
		43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
	}

	Pk10Qiu1ItemS = []int64{
		1070, 1071, 1072, 1073, 1074, 1075, 1076, 1077, 1078, 1079,
	}
	Pk10Qiu2ItemS = []int64{
		1080, 1081, 1082, 1083, 1084, 1085, 1086, 1087, 1088, 1089,
	}
	Pk10Qiu3ItemS = []int64{
		1090, 1091, 1092, 1093, 1094, 1095, 1096, 1097, 1098, 1099,
	}
	Pk10Qiu4ItemS = []int64{
		1100, 1101, 1102, 1103, 1104, 1105, 1106, 1107, 1108, 1109,
	}
	Pk10Qiu5ItemS = []int64{
		1110, 1111, 1112, 1113, 1114, 1115, 1116, 1117, 1118, 1119,
	}
	Pk10Qiu6ItemS = []int64{
		1120, 1121, 1122, 1123, 1124, 1125, 1126, 1127, 1128, 1129,
	}
	Pk10Qiu7ItemS = []int64{
		1130, 1131, 1132, 1133, 1134, 1135, 1136, 1137, 1138, 1139,
	}
	Pk10Qiu8ItemS = []int64{
		1140, 1141, 1142, 1143, 1144, 1145, 1146, 1147, 1148, 1149,
	}
	Pk10Qiu9ItemS = []int64{
		1150, 1151, 1152, 1153, 1154, 1155, 1156, 1157, 1158, 1159,
	}
	Pk10Qiu10ItemS = []int64{
		1160, 1161, 1162, 1163, 1164, 1165, 1166, 1167, 1168, 1169,
	}
)
