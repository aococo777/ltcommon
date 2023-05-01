package commonclass

import (
	"ltback/src/ltcommon/commonstruct"
	"sync"

	"github.com/jinzhu/gorm"
)

type RechargeWay struct {
	mysql      *gorm.DB
	WayID_Info map[int64]commonstruct.RechargeWayinfo
	snlock     sync.Mutex
}

func (this *RechargeWay) Init(db *gorm.DB) error {
	this.snlock.Lock()
	defer this.snlock.Unlock()
	this.mysql = db
	this.WayID_Info = make(map[int64]commonstruct.RechargeWayinfo)

	// 支付方式的详细内容
	var wayinfos []commonstruct.RechargeWayinfo
	if err := this.mysql.Table(commonstruct.WY_gm_config_recharge_way).Select("id,address,name,navi_type,service_type,img,gold_list,enable_time,scan_type").
		Order("sort").Find(&wayinfos).Error; err != nil {
		return err
	}
	for _, v := range wayinfos {
		//		beego.Informational("way = ", v.ID, v)
		this.WayID_Info[v.ID] = v
	}

	return nil
}

// 获取支付方式详细内容
func (this *RechargeWay) GetRechargewayInfo(wayid int64) (commonstruct.RechargeWayinfo, bool) {
	this.snlock.Lock()
	defer this.snlock.Unlock()
	wayinfo, ok := this.WayID_Info[wayid]
	if !ok {
		return wayinfo, false
	}
	return wayinfo, true
}

// 获取支付方式详细内容
func (this *RechargeWay) GetRechargewayinfoByServicetype(servicetype string) (commonstruct.RechargeWayinfo, bool) {
	this.snlock.Lock()
	defer this.snlock.Unlock()
	for _, v := range this.WayID_Info {

		if v.ServiceType == servicetype {

			return v, true
		}
	}
	var ret commonstruct.RechargeWayinfo
	return ret, false
}
