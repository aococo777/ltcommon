package commonclass

import (
	"fmt"
	"github.com/aococo777ltcommon/commonstruct"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

// 游戏结果分类
type GentouplanManager struct {
	DestUuid_SrcuuidS map[int64][]int64                  // 下单用户对应的跟投对象
	uuid_Gentouplan   map[string]commonstruct.GentouPlan // 跟投计划信息()

	mysql  *gorm.DB
	gtLock sync.Mutex // 生肖锁
}

func (this *GentouplanManager) Init(DB *gorm.DB) error {
	beego.Error("GentouplanManager Init enter!")

	this.mysql = DB

	go this.TimerUpdate()

	return nil
}

func (this *GentouplanManager) TimerUpdate() {
	this.UpdateGentouplanS()
	CheckTimer := time.NewTicker(20 * time.Second) // 刷新跟投计划
	for {
		select {
		case <-CheckTimer.C:
			this.UpdateGentouplanS()
		}
	}
}

// 获取item列表 (每日凌晨需要重新刷新,新年需要修改年肖项)
func (this *GentouplanManager) UpdateGentouplanS() {
	beego.Error("UpdateGentouplanS enter ~~~~~~")

	this.gtLock.Lock()
	defer this.gtLock.Unlock()

	this.DestUuid_SrcuuidS = make(map[int64][]int64)
	this.uuid_Gentouplan = make(map[string]commonstruct.GentouPlan)

	// 初始化结果分类表
	var PlanS []commonstruct.GentouPlan
	if err := this.mysql.Table(commonstruct.WY_tmp_user_gentouplan).Find(&PlanS).Error; err != nil {
		beego.Error("UpdateGameItem err ", err)
		return
	}

	for _, v := range PlanS {
		this.AddGentouuuidS(v.GentouUuid, v.Uuid)

		roomids := strings.Split(v.GentouGameList, ",")
		for _, roomidstr := range roomids {
			roomid, _ := strconv.ParseInt(roomidstr, 10, 64)
			this.AddGentouplan(v.Uuid, v.GentouUuid, roomid, v)
		}
	}
}

func (this *GentouplanManager) AddGentouuuidS(destuuid int64, uuid int64) {
	if value, ok := this.DestUuid_SrcuuidS[destuuid]; ok {
		value = append(value, uuid)
		this.DestUuid_SrcuuidS[destuuid] = value
	} else {
		this.DestUuid_SrcuuidS[destuuid] = []int64{uuid}
	}
}

// 初始化即时注单分组
func (this *GentouplanManager) GetGentouuuidS(uuid int64) []int64 {
	this.gtLock.Lock()
	defer this.gtLock.Unlock()
	if info, ok := this.DestUuid_SrcuuidS[uuid]; ok {
		return info
	}
	return nil
}

// 代理 跟 会员 在XX游戏 投
func (this *GentouplanManager) AddGentouplan(srcuuid int64, destuuid int64, roomid int64, planinfo commonstruct.GentouPlan) {

	gentoukey := fmt.Sprintf("%v_%v_%v", srcuuid, destuuid, roomid)

	if _, ok := this.uuid_Gentouplan[gentoukey]; ok {
		// value = append(value, uuid)
		// this.DestUuid_SrcuuidS[gentoukey] = value
	} else {
		this.uuid_Gentouplan[gentoukey] = planinfo
	}
}

// 初始化即时注单分组
func (this *GentouplanManager) GetGentouplan(srcuuid int64, destuuid int64, roomid int64) (commonstruct.GentouPlan, bool) {
	gentoukey := fmt.Sprintf("%v_%v_%v", srcuuid, destuuid)

	this.gtLock.Lock()
	defer this.gtLock.Unlock()
	info, ok := this.uuid_Gentouplan[gentoukey]

	return info, ok
}
