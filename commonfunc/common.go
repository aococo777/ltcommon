package commonfunc

/*
// 对象是否在数组内
func InIDArr(arr []int64, uuid int64) bool
// 2006-01-02 15:04:05 转时间戳
func DateTime_Str2int64(datetime string) int64
// 2006-01-02 转时间戳
func Date_Str2int64(date string) int64
// 返回函数调用层级
func GetFuncName(level int) string
// 去重和去零
func RemoveDuplicatesAndZero(a []int64) (ret []int64)
// KeyS 是否在 FormS 内
func IsCorrectForms(KeyS *[]string, FormS *[]string) bool
// http请求返回
func SendMsgtoClient(Ctx *context.Context, Status bool, Error_code commonstruct.ErrorType, msg string)
// 获取当前时间
func GetNowtime() int64
// 获取当前日期
func GetNowdate() int64

func GetMoneytypeStr(moneytype commonstruct.MoneyType) string
// 保留小数位数
func Round(f float64, n int) float64

func GetTablesuf(uuid int64) string
// 字符串规则
func RegexpStr(str string) error
*/

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"github.com/aococo777ltcommon/commonstruct"
	"math"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func InIDArr(arr []int64, uuid int64) bool {
	if len(arr) <= 0 || uuid <= 0 {
		return false
	}
	for _, v := range arr {
		if v == uuid {
			return true
		}
	}
	return false
}

func DateTime_Str2int64(datetime string) int64 {
	timeUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
	return timeUnix.Unix()
}

func Date_Str2int64(date string) int64 {
	timeUnix, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return timeUnix.Unix()
}

func GetFuncName(level int) string {
	var strRet string
	for i := level + 1; i >= 1; i-- {
		pc, _, line, _ := runtime.Caller(i)
		f := runtime.FuncForPC(pc)
		strRet = strRet + f.Name() + fmt.Sprintf("--->%d\n", line)
	}
	return strRet
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

// 表单规范检测
func IsCorrectForms(KeyS *[]string, FormS *[]string) bool {
	if len(*KeyS) > len(*FormS) {
		return false
	}
	for _, kUser := range *KeyS {
		bCorrectForm := false
		for _, kWeb := range *FormS {
			if strings.EqualFold(kUser, kWeb) {
				bCorrectForm = true
				break
			}
		}
		if !bCorrectForm {
			return false
		}
	}
	return true
}

func SendMsgtoClient(Ctx *context.Context, Status bool, Error_code commonstruct.ErrorType, msg string) {
	var ret commonstruct.RetMsg
	ret.Status = Status
	ret.Error_code = Error_code
	ret.Msg = msg
	retbyte, _ := json.Marshal(ret)
	Ctx.ResponseWriter.Write(retbyte)
}

// 获取北京时间
func BeijingTime() time.Time {
	// var cstSh, _ = time.LoadLocation("Local") //上海
	// cstSh, err := time.LoadLocation("Asia/Shanghai")
	return time.Now()
}

func GetBjHour() int64 {
	cstSh := time.FixedZone("CST", 8*3600)
	ret := time.Now().In(cstSh).Hour()
	return int64(ret)
}

func GetBjTime060102(timearg time.Time) string {
	cstSh := time.FixedZone("CST", 8*3600)
	ret := timearg.In(cstSh).Format("060102")
	return ret
}

func GetBjYearDay(timearg time.Time) int {
	cstSh := time.FixedZone("CST", 8*3600)
	ret := timearg.In(cstSh).YearDay()
	return ret
}

func GetBjTime20060102(timearg time.Time) string {
	cstSh := time.FixedZone("CST", 8*3600)
	ret := timearg.In(cstSh).Format("20060102")
	return ret
}

func GetBjTime2006_01_02(timearg time.Time) string {
	cstSh := time.FixedZone("CST", 8*3600)
	ret := timearg.In(cstSh).Format("2006-01-02")
	return ret
}

func GetBjTime20060102150405(timearg time.Time) string {
	cstSh := time.FixedZone("CST", 8*3600)
	ret := timearg.In(cstSh).Format("20060102150405")
	return ret
}

func GetBjTime2006_01_02_15_04_05(timearg time.Time) string {
	cstSh := time.FixedZone("CST", 8*3600)
	ret := timearg.In(cstSh).Format("2006-01-02 15:04:05")
	return ret
}

// 2006-01-02 15:04:05
func StrToBjTime(drawtime string) time.Time {
	timeLayout := "2006-01-02 15:04:05"                           //转化所需模板
	loc := time.FixedZone("CST", 8*3600)                          //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, drawtime, loc) //
	return theTime
}

// 2006-01-02
func DateStrToBjTime(drawtime string) time.Time {
	timeLayout := "2006-01-02"                                    //转化所需模板
	loc := time.FixedZone("CST", 8*3600)                          //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, drawtime, loc) //
	return theTime
}

// 20060102
func Date20060102ToBjTime(drawtime string) time.Time {
	timeLayout := "20060102"                                      //转化所需模板
	loc := time.FixedZone("CST", 8*3600)                          //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, drawtime, loc) //
	return theTime
}

// 20060102150405
func StrToBjTime_YL(drawtime string) time.Time {
	timeLayout := "20060102150405"                                //转化所需模板
	loc := time.FixedZone("CST", 8*3600)                          //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, drawtime, loc) //
	return theTime
}

func GetNowtime() int64 {
	// 这个方式最优 不依赖于系统时区和设置。 只要 时间戳对就行
	cstSh := time.FixedZone("CST", 8*3600)
	// shtime, _ := strconv.ParseInt(time.Now().In(cstSh).Format("20060102150405"), 10, 64)

	// LoadLocation 依赖于tzdata数据库。 windows不带
	// cstSh, err := time.LoadLocation("Local")
	// if err != nil {
	// 	beego.Error("time.LoadLocation err", err.Error())
	// }
	nowtime, _ := strconv.ParseInt(time.Now().In(cstSh).Format("20060102150405"), 10, 64)
	return nowtime
}

func GetNowdate() int64 {
	cstSh := time.FixedZone("CST", 8*3600)
	// if err != nil {
	// 	beego.Error("time.LoadLocation err", err.Error())
	// }
	nowtime, _ := strconv.ParseInt(time.Now().In(cstSh).Format("20060102"), 10, 64)
	return nowtime
}

func GetBilishiDate(inittime int64) int64 {
	var optiem time.Time
	if inittime == 0 {
		optiem = time.Now()
	} else {
		optiem = StrToBjTime_YL(fmt.Sprintf("%d", inittime))
	}

	cstBls := time.FixedZone("CST", 2*3600)
	nowtime, _ := strconv.ParseInt(optiem.In(cstBls).Format("20060102"), 10, 64)
	return nowtime
}

func GetBegintime(date int64) int64 {
	var rettime int64
	if GetNowdate() == date && GetBjHour() < 6 { // 今日六点前获取的时间为昨日六点时刻
		loc := time.FixedZone("CST", 8*3600) //重要：获取时区
		t, _ := time.ParseInLocation("20060102", fmt.Sprintf("%d", date), loc)
		nexttime := t.AddDate(0, 0, -1)
		nextdate, _ := strconv.ParseInt(GetBjTime20060102(nexttime), 10, 64)
		rettime = nextdate*1000000 + 60000

	} else {
		rettime = date*1000000 + 60000
	}

	return rettime
}

func GetEndtime(date int64) int64 {
	var rettime int64

	if GetNowdate() == date && GetBjHour() < 6 { // 今日六点前获取的时间为今日六点时刻
		rettime = date*1000000 + 60000
	} else {
		loc := time.FixedZone("CST", 8*3600) //重要：获取时区
		t, _ := time.ParseInLocation("20060102", fmt.Sprintf("%d", date), loc)
		nexttime := t.AddDate(0, 0, 1)
		nextdate, _ := strconv.ParseInt(GetBjTime20060102(nexttime), 10, 64)
		rettime = nextdate*1000000 + 60000
	}

	return rettime
}

func GetMoneytypeStr(moneytype commonstruct.MoneyType) string {
	var ret string
	switch moneytype {
	case commonstruct.MoneyType_Cash:
		ret = "cash"
		break
	case commonstruct.MoneyType_Credit:
		ret = "credit"
		break
	case commonstruct.MoneyType_Virtual:
		ret = "virtual"
		break
	}
	return ret
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func GetTablesuf(uuid int64) string {

	suf := fmt.Sprintf("%v", uuid%100)

	return suf
}

func RegexpStr(str string) error {

	if len(str) > 12 {
		return errors.New("账号长度必须小于等于12")
	}
	if len(str) < 4 {
		return errors.New("账号长度必须大于等于4")
	}

	chr := regexp.MustCompile("[a-zA-Z]")
	posinfo := chr.FindStringIndex(str)
	if len(posinfo) != 2 {
		return errors.New("第一个必须为字母")
	} else {
		if posinfo[0] != 0 {
			return errors.New("第一个必须为字母")
		}
	}

	allr := regexp.MustCompile("[a-zA-Z0-9]+")
	strs := allr.FindAllString(str, -1)
	if len(strs) == 1 {
		if str == strs[0] {
			return nil
		}
	}
	return errors.New("只允许字母和数字组合")
}

func GetUrlImgBase64(path string) (baseImg string, err error) {

	//获取本地文件
	file, err := os.Open(path)
	if err != nil {
		beego.Error("获取本地图片失败", path, err.Error())
		err = errors.New("获取本地图片失败")
		return
	}
	defer file.Close()
	imgByte, _ := ioutil.ReadAll(file)

	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
	mimeType := http.DetectContentType(imgByte) //取图片类型
	switch mimeType {
	case "image/jpeg":
		baseImg = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/png":
		baseImg = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/vnd.microsoft.icon":
		baseImg = "data:image/x-icon;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/gif":
		baseImg = "data:image/gif;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/bmp":
		baseImg = "data:image/bmp;base64," + base64.StdEncoding.EncodeToString(imgByte)
	default:
		beego.Error("文件格式", path, mimeType)
		if strings.Contains(path, ".ico") {
			baseImg = "data:image/x-icon;base64," + base64.StdEncoding.EncodeToString(imgByte)
		} else {
			beego.Error("未知的文件类型")
		}

		// 1.png data:image/png;base64,
		// 2.jpg data:image/jpeg;base64,
		// 3.gif data:image/gif;base64,
		// 4.svg data:image/svg+xml;base64,
		// 5.ico data:image/x-icon;base64,
		// 6.bmp data:image/bmp;base64,
	}

	return
}

func StrToIntarr(strarg string, sep string) []int64 {

	var retarr []int64
	strArr := strings.Split(strarg, sep)
	for _, strtmp := range strArr {
		inttmp, _ := strconv.ParseInt(strtmp, 10, 64)
		retarr = append(retarr, inttmp)
	}
	return retarr
}
