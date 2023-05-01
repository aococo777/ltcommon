package commonclass

import (
	"errors"
	"github.com/aococo777/ltcommon/commonstruct"
	"strings"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

type CoBase struct {
	cobase map[int64]commonstruct.CompanyBase

	cbLock sync.Mutex
	mysql  *gorm.DB
}

func (this *CoBase) Init(db *gorm.DB) {
	this.mysql = db
	this.cobase = make(map[int64]commonstruct.CompanyBase)
	//	this.Urlmap = make(map[int64][]string)
	//	this.adminurl = make(map[int64][]string)
	//	this.whiteip = make(map[int64][]string)

	this.InitUrl()
	go this.UpdateUrl()
}

func (this *CoBase) UpdateUrl() {
	SessionTimer := time.NewTicker(300 * time.Second) // 刷新公司域名
	for {
		select {
		case <-SessionTimer.C:
			this.InitUrl()
		}
	}
}

func (this *CoBase) InitUrl() {
	baselist, err := this.GetBaseinfoInDB()
	if err == nil {
		for _, baseinfo := range baselist {
			this.setCompanyBase(baseinfo.Uuid, baseinfo)
		}
	}
}

func (this *CoBase) GetBaseinfoInDB() ([]commonstruct.CompanyBase, error) {
	var ways []commonstruct.CompanyBase
	if err := this.mysql.Table(commonstruct.WY_company_base).Find(&ways).Error; err != nil {
		return ways, err
	}
	return ways, nil
}

func (this *CoBase) GetBaseinfoS() map[int64]commonstruct.CompanyBase {
	this.cbLock.Lock()
	defer this.cbLock.Unlock()
	return this.cobase
}

func (this *CoBase) setCompanyBase(companyid int64, info commonstruct.CompanyBase) {
	this.cbLock.Lock()
	defer this.cbLock.Unlock()
	this.cobase[companyid] = info
}

func (this *CoBase) GetCompanyBase(companyid int64) (commonstruct.CompanyBase, error) {
	this.cbLock.Lock()
	defer this.cbLock.Unlock()
	value, ok := this.cobase[companyid]
	if !ok {
		return value, errors.New("wrong companyid")
	}
	return value, nil
}

func (this *CoBase) GetFronturlS(companyid int64) []string {
	baseinfo, err := this.GetCompanyBase(companyid)
	if err != nil {
		return nil
	} else {
		if baseinfo.FrontUrls != "" {
			urls := strings.Split(baseinfo.FrontUrls, ",")
			return urls
		} else {
			return nil
		}
	}
}

func (this *CoBase) GetBackurlS(companyid int64) []string {
	baseinfo, err := this.GetCompanyBase(companyid)
	if err != nil {
		return nil
	} else {
		if baseinfo.BackUrls != "" {
			urls := strings.Split(baseinfo.BackUrls, ",")
			return urls
		} else {
			return nil
		}
	}
}

func (this *CoBase) GetWhitelist(companyid int64) []string {
	baseinfo, err := this.GetCompanyBase(companyid)
	if err != nil {
		return nil
	} else {
		if baseinfo.WhiteList != "" {
			urls := strings.Split(baseinfo.WhiteList, ",")
			return urls
		} else {
			return nil
		}
	}
}
