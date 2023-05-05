package commonclass

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aococo777/ltcommon/commonfunc"
	"github.com/aococo777/ltcommon/commonstruct"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

// 游戏结果分类
type ShengxiaoManager struct {
	Shengxiao_Ball map[int64]string // 色波信息
	Nianxiao       string           // 一中一连码
	mysql          *gorm.DB
	shengxiaoLock  sync.Mutex // 生肖锁
}

func (this *ShengxiaoManager) Init(DB *gorm.DB) error {
	beego.Error("ShengxiaoManager Init enter!")

	this.mysql = DB
	this.Shengxiao_Ball = make(map[int64]string)
	go this.TimerUpdate()
	return nil
}

func (this *ShengxiaoManager) TimerUpdate() {
	beego.Error("ShengxiaoManager enter!")
	this.UpdateShengxiaoset()
	CheckTimer := time.NewTicker(10 * time.Second) // 刷新年肖信息和对应球号
	for {
		select {
		case <-CheckTimer.C:
			now := commonfunc.BeijingTime()
			flg := (commonfunc.GetBjHour() == 0 && now.Minute() == 0 && (now.Second() >= 0 && now.Second() < 10))
			if flg {
				this.UpdateShengxiaoset()
			}
		}
	}
}

// 将数据库中的生肖更新到内存中
func (this *ShengxiaoManager) UpdateShengxiaoset() {
	date := commonfunc.GetNowdate()

	var infos []commonstruct.CShengxiaoball
	if err := this.mysql.Table(commonstruct.WY_gm_config_shengxiaoball).
		Where("date <= ?", date).Order("date desc").Limit(12).
		Find(&infos).Error; err != nil {
		beego.Error("GetShengxiaoballinfoByDate err", date, err.Error())
		return
	}

	if len(infos) != 12 {
		beego.Error("WY_gm_config_shengxiaoball 生肖记录数有误", date, len(infos))

	} else {
		this.SetNianxiao(infos[0].Nianxiao)
	}

	for _, shengxiaoballinfo := range infos {
		ballnumlist := strings.Split(shengxiaoballinfo.Balls, ",")
		for _, ballnum := range ballnumlist {
			num, _ := strconv.Atoi(ballnum)
			this.SetBallxiao(int64(num), shengxiaoballinfo.Shengxiao)
		}
	}
}

func (this *ShengxiaoManager) SetNianxiao(xiao string) {
	this.shengxiaoLock.Lock()
	defer this.shengxiaoLock.Unlock()
	this.Nianxiao = xiao
}

func (this *ShengxiaoManager) GetNianxiao() string {
	this.shengxiaoLock.Lock()
	defer this.shengxiaoLock.Unlock()
	ret := this.Nianxiao
	return ret
}

func (this *ShengxiaoManager) SetBallxiao(Num int64, xiao string) {
	this.shengxiaoLock.Lock()
	defer this.shengxiaoLock.Unlock()
	this.Shengxiao_Ball[Num] = xiao
}

func (this *ShengxiaoManager) GetShengxiao(Num int64) string {
	this.shengxiaoLock.Lock()
	defer this.shengxiaoLock.Unlock()

	if info, ok := this.Shengxiao_Ball[Num]; ok {
		return info
	}
	return ""
}

func (this *ShengxiaoManager) GetShengxiaoS() map[int64]string {
	this.shengxiaoLock.Lock()
	defer this.shengxiaoLock.Unlock()

	ret := this.Shengxiao_Ball
	return ret
}

func (this *ShengxiaoManager) UpdateShengxiao() {
	today := commonfunc.GetBjTime20060102(time.Now())
	var infos commonstruct.CShengxiaoball
	if err := this.mysql.Table(commonstruct.WY_gm_config_shengxiaoball).
		Where("date <= ?", today).Order("date desc").Limit(1).
		Find(&infos).Error; err != nil {
		beego.Error("get WY_gm_config_shengxiao err", today)
		return
	}

	i_today, _ := strconv.ParseInt(today, 10, 64)
	if infos.Date == i_today { // 切换生肖日
		beego.Error("切换生肖日", infos.Date, infos.Nianxiao)
		if err := this.NianxiaoToNormal(); err != nil {
			beego.Error("NianxiaoToNormal err", err)
		}
		if err := this.NewNianxiao(infos.Nianxiao); err != nil {
			beego.Error("NewNianxiao err", infos.Nianxiao, err)
		}
	} else {
		beego.Error("非生肖日", infos.Date, i_today)
	}
}

func (this *ShengxiaoManager) NianxiaoToNormal() error {
	updateValues1 := map[string]interface{}{
		"game_xiaolei": "普肖",
	}

	if err := this.mysql.Table(commonstruct.WY_gm_config_item).Where("game_xiaolei in ('年肖','普肖')").
		Update(updateValues1).Error; err != nil {
		beego.Error("NianxiaoToNormal err ", err)
		return err
	}
	return nil
}

func (this *ShengxiaoManager) NewNianxiao(nianxiao string) error {
	updateValues1 := map[string]interface{}{
		"game_xiaolei": "年肖",
	}

	if err := this.mysql.Table(commonstruct.WY_gm_config_item).Where("game_item = ? and game_xiaolei in ('年肖','普肖')", nianxiao).
		Update(updateValues1).Error; err != nil {
		beego.Error("NewNianxiao err ", err)
		return err
	}
	return nil
}
