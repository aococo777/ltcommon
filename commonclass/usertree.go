package commonclass

import (
	"encoding/json"
	"errors"
	"fmt"
	"ltback/src/ltcommon/commonfunc"
	"ltback/src/ltcommon/commonstruct"
	"sync"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

/******************************************************
好记性不如烂笔头！

该类实现功能: (提高系统运行效率)
	1、服务启动的时候，读取wy_user_branch表的用户树 至内存中;
		对外导出接口 常用:GetMasterID,GetPreID,GetPreIDList
					GetMasterList,GetZhanchengPer
	2、在内存中没有读到数据时，UpdateUsertree(uuid)更新单个用户的数据

******************************************************/

type UserTree struct {
	mysql       *gorm.DB
	PreIDList   map[int64][]int64 // 上级列表
	MasterID    map[int64]int64   // 公司ID
	PreID       map[int64]int64   // 上级ID
	SalerIDList map[int64][]int64 // 销售列表
	snlock      sync.Mutex
}

func (this *UserTree) Init(db *gorm.DB) error {
	this.snlock.Lock()
	defer this.snlock.Unlock()
	this.mysql = db
	this.PreIDList = make(map[int64][]int64)
	this.MasterID = make(map[int64]int64)
	this.PreID = make(map[int64]int64)
	this.SalerIDList = make(map[int64][]int64)

	// 先获取companylist
	var companybaseS []commonstruct.CompanyBase
	if err := this.mysql.Table(commonstruct.WY_company_base).Find(&companybaseS).Error; err != nil {
		beego.Error("GetCompanybaseS err", err)
		return err
	}
	// for _, companybase := range companybaseS {
	// 	var userbranch commonstruct.BranchUsers
	// 	if err := this.mysql.Table(commonstruct.WY_user_branch).Select("uuid,saler_id,saler_list").
	// 		Where("uuid = ?", companybase.Uuid).Find(&userbranch).Error; err != nil {
	// 		beego.Error("getuserbranch err", companybase.Uuid, err)
	// 		return err
	// 	} else {
	// 		var IDList []int64
	// 		if err := json.Unmarshal(userbranch.SalerList, &IDList); err != nil {
	// 			beego.Error(fmt.Sprintf("InitSalerIDList [%v] Unmarshal err %v", userbranch.Uuid, err))
	// 			continue
	// 		}
	// 		this.SalerIDList[userbranch.Uuid] = IDList
	// 	}
	// }

	beego.Error("initsalerlist finished")

	var PreBranch []commonstruct.BranchUsers
	if err := this.mysql.Table(commonstruct.WY_user_branch).Select("uuid,master_id,pre_id,pre_list").
		Where("master_id > 0").Find(&PreBranch).Error; err != nil {
		beego.Error("getuserbranch err", err)
		return err
	}
	for _, v := range PreBranch {
		if v.PreID > 0 {
			// if salerlist, ok := this.SalerIDList[v.MasterID]; ok {
			// for _, salerid := range salerlist {
			// 	this.PreIDList[v.Uuid] = append(this.PreIDList[v.Uuid], salerid)
			// }

			var preidlist []int64
			if err := json.Unmarshal(v.PreList, &preidlist); err != nil {
				beego.Error("InitPreIDList [%v] Unmarshal err %v\n", v, err)
				continue
			}

			for _, preid := range preidlist {
				this.PreIDList[v.Uuid] = append(this.PreIDList[v.Uuid], preid)
			}
			this.MasterID[v.Uuid] = v.MasterID
			this.PreID[v.Uuid] = v.PreID

			beego.Error(fmt.Sprintf("uuid:%v =>[%v] prelist => [%v]", v.Uuid, v.MasterID, this.PreIDList[v.Uuid]))
			// }
		}
	}
	return nil
}

// 获取站长ID
func (this *UserTree) GetMasterID(uuid int64) int64 {
	this.snlock.Lock()
	defer this.snlock.Unlock()
	if MasterID, ok := this.MasterID[uuid]; ok {
		return MasterID
	} else {
		this.updateUsertree(uuid)
		if MasterID, ok := this.MasterID[uuid]; ok {
			return MasterID
		}
	}
	return 0
}

// 获取站长ID
func (this *UserTree) GetPreID(uuid int64) (int64, error) {
	this.snlock.Lock()
	defer this.snlock.Unlock()
	if MasterID, ok := this.PreID[uuid]; ok {
		return MasterID, nil
	}
	err := this.updateUsertree(uuid)
	if MasterID, ok := this.PreID[uuid]; ok {
		return MasterID, nil
	}
	return 0, err
}

// 此IDlist 为 从上而下
func (this *UserTree) GetPreIDList(uuid int64) ([]int64, error) {
	if uuid == 0 {
		return nil, errors.New("uuid is zero")
	}

	this.snlock.Lock()
	defer this.snlock.Unlock()
	if List, ok := this.PreIDList[uuid]; ok {
		return List, nil
	}
	err := this.updateUsertree(uuid)
	if List, ok := this.PreIDList[uuid]; ok {
		return List, nil
	}
	return nil, err
}

// 获取代理列表
func (this *UserTree) updateUsertree(uuid int64) error {
	var NewUuidS commonstruct.BranchUsers
	if err := this.mysql.Table(commonstruct.WY_user_branch).Where("uuid = ?", uuid).Find(&NewUuidS).Error; err != nil {
		beego.Error(commonfunc.GetFuncName(3))
		beego.Error("updateUsertree err", uuid, err.Error())
		return err
	}

	if NewUuidS.PreID > 0 {
		var IDList []int64
		if err := json.Unmarshal(NewUuidS.PreList, &IDList); err != nil {
			beego.Error("Unmarshal err", err.Error())
			return err
		}
		this.PreIDList[uuid] = IDList
		this.MasterID[uuid] = NewUuidS.MasterID
		this.PreID[uuid] = NewUuidS.PreID
	}

	switch NewUuidS.RoleType {
	case "gm", "saler":
	default:
		beego.Error("updateUsertree === > ", NewUuidS)
	}
	return nil
}
