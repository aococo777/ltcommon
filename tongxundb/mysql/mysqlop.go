package txmysql

import (
	"fmt"
	"log"
	"github.com/aococo777/ltcommon/commonfunc"
	"github.com/aococo777/ltcommon/commonstruct"
	"github.com/aococo777/ltcommon/tongxundb/tablestruct"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	WyMysql_Tongxun *gorm.DB //
)

const (
	PageCount = 20
)

func init() {
	{
		var err error
		mysqlinfo := new(commonstruct.MysqlDBInfo)
		mysqlinfo.Hostsip = beego.AppConfig.String("mysqlurls_tongxun")
		mysqlinfo.Username = beego.AppConfig.String("mysqluser_tongxun")
		mysqlinfo.Password = beego.AppConfig.String("mysqlpass_tongxun")
		mysqlinfo.DBName = beego.AppConfig.String("mysqldb_tongxun")
		WyMysql_Tongxun, err = InitDB(mysqlinfo)
		if err != nil {
			beego.Error("Init MysqlDB failed! ", mysqlinfo, err)
		}
	}
}

func InitDB(dbinfo *commonstruct.MysqlDBInfo) (*gorm.DB, error) {
	Dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", dbinfo.Username, dbinfo.Password, dbinfo.Hostsip, dbinfo.DBName)
	DB, err := gorm.Open("mysql", Dsn)
	if err != nil {
		return nil, err
	}
	//	DB.LogMode(true)

	file, err := os.Create("DB.log")
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	DB.SetLogger(log.New(file, "", log.LstdFlags|log.Llongfile))

	DB.DB().SetMaxIdleConns(10)
	return DB, nil
}

func GetFuncName(level int) string {
	// 1 为调用的上一级
	var strRet string
	for i := level + 1; i >= 1; i-- {
		pc, _, line, _ := runtime.Caller(i)
		f := runtime.FuncForPC(pc)
		strRet = strRet + f.Name() + "--->" + strconv.Itoa(line) + " $ "
	}
	return strRet
}

func GetUserbaseS() ([]tablestruct.UserInfo, error) {

	var userbase []tablestruct.UserInfo
	if err := WyMysql_Tongxun.Table("user_base").Find(&userbase).Error; err != nil {
		return userbase, err
	}
	return userbase, nil
}

func GetUserbase(account string) (tablestruct.UserInfo, error) {

	var userbase tablestruct.UserInfo
	if err := WyMysql_Tongxun.Table("user_base").Where("account = ?", account).Find(&userbase).Error; err != nil {
		return userbase, err
	}
	return userbase, nil
}

func GetCallmeunresplogS(account string) ([]tablestruct.CallLog, error) {
	var logs []tablestruct.CallLog
	if err := WyMysql_Tongxun.Table("call_log").Where("dest_account = ? and resp_time = 0", account).Order("call_time desc").Find(&logs).Error; err != nil {
		return logs, err
	}
	return logs, nil
}

func GetCallmeresplogS(account string) ([]tablestruct.CallLog, error) {
	var logs []tablestruct.CallLog
	if err := WyMysql_Tongxun.Table("call_log").Where("dest_account = ? and resp_time > 0", account).Order("call_time desc").Find(&logs).Error; err != nil {
		return logs, err
	}
	return logs, nil
}

func GetMycalllogs(account string) ([]tablestruct.CallLog, error) {

	var logs []tablestruct.CallLog
	if err := WyMysql_Tongxun.Table("call_log").Where("src_account = ?", account).Order("call_time desc").Find(&logs).Error; err != nil {
		return logs, err
	}
	return logs, nil
}

func CreateCalllog(newlog tablestruct.CallLog) error {
	if err := WyMysql_Tongxun.Table("call_log").Create(&newlog).Error; err != nil {
		beego.Error("CreateCalllog err %v", err)
	}
	return nil
}

func RespCalllog(logid int64, respinfo string) error {
	var updateValues map[string]interface{}
	updateValues = map[string]interface{}{
		"resp_time": commonfunc.GetNowtime(),
		"respinfo":  respinfo,
	}

	if err := WyMysql_Tongxun.Table("call_log").Where("id = ?", logid).Update(updateValues).Error; err != nil {
		beego.Error("RespCalllog err", err)
		return err
	}
	return nil
}

// 删除三个月前订单
func ClearExpiredata() {
	expiretime, _ := strconv.ParseInt(commonfunc.GetBjTime20060102150405(time.Now().AddDate(0, 0, -5)), 10, 64)
	sql := fmt.Sprintf("delete from call_log where call_time < %v;", expiretime)
	if err := WyMysql_Tongxun.Exec(sql).Error; err != nil {
		beego.Error("ClearExpiredata err ", err)
	}
}
