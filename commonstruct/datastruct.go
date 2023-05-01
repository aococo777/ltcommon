package commonstruct

type Wininfo struct {
	Expect        string `gorm:"primary_key"`
	Opencode      string
	Opentime      string
	Opentimestamp int64
	Expinfo       string `gorm:"type:text"`
	Changlongflag int64  `gorm:"type:tinyint(1);default:0"`
}

//
type Cli_Wininfo struct {
	Expect        string `gorm:"primary_key"` // 期号
	Opencode      string // 号码
	Opentime      string // 时间
	Opentimestamp int64  //
	// Expinfo       string `gorm:"type:text"`
}

// 返回给客户端的结构体
type RetClientWininfo struct {
	Expect        string `gorm:"primary_key"` // 期号
	Opencode      string // 号码
	Opentime      string // 时间
	Opentimestamp int64  //
	SettleState   int64  `json:"settlestate"` //结算状态
}

// ltcode num表的开奖号码结构体
type DBWininfo struct {
	Expect        string `gorm:"primary_key"` // 期号
	Opencode      string // 号码
	Opentime      string // 时间
	Opentimestamp int64  //
	Res           string // 号码来源
}

type MysqlDBInfo struct {
	Hostsip  string
	Username string
	Password string
	DBName   string
}

type CompanyConfig struct { // 公司的水赔比
	Uuid          int64   `gorm:"not null" json:"uuid"` // 代理ID
	DaywagerLevel int64   `json:"daywager_level"`       // 日工资结算层
	WagerRate     float64 `json:"wager_rate"`           // 工资比例(百分比来算)
	LosebackRate  float64 `json:"loseback_rate"`        // 输返比例(百分比来算)
	AllocateLevel int64   `json:"allocate_level"`       // 分配详情
	AllocateRule  string  `json:"allocate_rule"`        // 分配规则
}

type Daywager struct { // 公司的水赔比
	Uuid         int64   `gorm:"not null" json:"uuid"` // 用户ID
	Date         int64   `json:"date"`                 // 日工资结算层
	OrderNum     int64   `json:"order_num"`            // 投注量
	Wager        float64 `json:"wager"`                // 投注金额
	SettledNum   int64   `json:"settle_num"`           // 结算量
	SettledWager float64 `json:"settle_wager"`         // 结算金额
	ProfitWager  float64 `json:"profit_wager"`         // 结算金额
	Selfshui     float64 `json:"selfshui"`             // 自身退水
	Dayshui      float64 `json:"dayshui"`              // 总退水
	Daywager     float64 `json:"daywager"`             // 日工资
	DaywagerFlag int64   `json:"daywager_flag"`        // 日工资发放标志
	Loseback     float64 `json:"loseback"`             // 输返金额
	LosebackFlag int64   `json:"loseback_flag"`        // 输返标志
}

// 权限设置
type Authority struct {
	Uuid         int64 `gorm:"primary_key;not null" json:"uuid"` // 用户ID
	PreID        int64 `json:"preid"`                            // 上级ID
	MasterID     int64 `json:"masterid"`                         // 站长ID
	SetAuthority int64 `json:"set_authority"`                    // 设置权限权
	SetOdds      int64 `json:"setodds"`                          // 设赔权
	SetShui      int64 `json:"setshui"`                          // 设水权
	EarnOdds     int64 `json:"earnodds"`                         // 赚赔权
	EarnShui     int64 `json:"earnshui"`                         // 赚水权
	NewPort      int64 `json:"newport"`                          // 开盘权
	NewUser      int64 `json:"newuser"`                          // 开户权
	Recharge     int64 `json:"recharge"`                         // 充值权
	Allocate     int64 `json:"allocate"`                         // 下分权
	AllocateSelf int64 `json:"allocate_self"`                    // 自身充值权
	Zoufei       int64 `json:"zoufei"`                           // 走飞权
	NewCompany   int64 `json:"newcompany"`                       // 开公司权
	Loginback    int64 `json:"loginback"`                        // 登录后台权
	Closetime    int64 `json:"closetime"`                        // 修改封盘时间
	Withdraw     int64 `json:"withdraw"`                         // 审核提现权
	RetShui      int64 `json:"retshui"`                          // 手动返佣权
	CoManage     int64 `json:"comanage"`                         // 公司管理权
}

// 百家权限
type Authority_Baijia struct {
	Baseop  bool `json:"baseop"`  // 分配ID
	Banroom bool `json:"banroom"` // 封禁房间
}

// 彩票权限
type Authority_Lottery struct {
	Zhuanpei bool `json:"zhuanpei"` // 赚赔
	Shepei   bool `json:"shepei"`   // 设赔
}

//注单  (分表) 分表规则？
type OrderResult struct {
	Orderid    int64   `gorm:"not null"` // 自动生成ID
	SettleTime int64   // 结算时间
	IsWin      string  // 中奖列表
	Wager      float64 // 派彩金额
	WagerTime  int64   // 返现时间
	Expinfo    string  // 扩展信息
}

type OrderUnsettle struct {
	Orderid      int64   `gorm:"not null"` // 来自于订单表生成
	Uuid         int64   // 下单用户ID
	Roomid       int64   // 房间ID
	Level        int64   // 用户层级
	PreID        int64   // 上级代理ID
	MasterID     int64   // 站长ID
	Expect       string  // 期号
	LotteryDalei string  // 游戏大类(guanfang、lottery、真人、竞彩)
	GameDalei    string  // 游戏查询分类(龙虎和、两面等)
	Port         string  // 盘口
	ItemID       int64   // 下单结果ID
	Amount       float64 // 投注金额
	Type         int64   // 注单类型(1走飞，0普通注单)
	ZouFeiinfo   string  `gorm:"type:text"`          // 扩展信息(前端版本等等)
	Orderinfo    string  `gorm:"type:text;not null"` // 投注内容(盘口, 投注金额, 该单注数, 上级给赔, 上级给水, 变赔率)
	Optime       int64   // 投注时间
	SettleTime   int64   // 结算时间
	IsWin        string  // 中奖列表
	RetType      int     // 是否结算
	Expinfo      string  `gorm:"type:text"` // 扩展信息(前端版本等等)
}

type ReCharge struct { // 充值表
	OrderID    int64   `json:"NoOrder"`     // 订单ID
	Time       string  `json:"time_order"`  // 下单时间
	Uuid       int64   `json:"user_id"`     // 用户ID
	PayType    string  `json:"pay_type"`    // 支付方式
	OrderTitle string  `json:"name_goods"`  // 订单内容
	OrderInfo  string  `json:"info_order"`  // 定案内容
	Amount     float64 `json:"money_order"` // 订单金额
	PartnerID  string  `json:"oid_partner"` // 第三方给我们的ID
	Callback   string  `json:"notify_url"`  // 回调接口
	SignType   string  `json:"sign_type"`   // 签名方式
	SignData   string  `json:"sign"`        // 签名数据
	QRCodeUrl  string  `json:"qrcodeurl"`   // 二维码地址
	EnableTime int64   `json:"enable_time"` // 二维码超时时间
	Checked    int64   `json:"checked"`     // 订单结果
}

type RechargeWayinfo struct { // 充值方式的信息
	ID          int64  `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"` // 订单ID
	Address     string `json:"address"`                                       // 下单时间
	RequestUrl  string `json:"requesturl"`                                    // 用户ID
	Name        string `json:"name"`                                          // 方式名称
	Platform    string `json:"platform"`                                      // 平台
	NaviType    string `gorm:"type:varchar(20);not null" json:"navi_type"`    // 一级导航分类
	ServiceType string `json:"service_type"`                                  // 给第三方的paytype
	Sort        string `json:"sort"`                                          // 优先使用排序
	Img         string `json:"img"`                                           // 图片名
	GoldList    string `json:"gold_list"`                                     // 可用金额列表
	EnableTime  int64  `json:"enable_time"`                                   // 二维码过期时间
	ScanType    string `json:"scan_type"`                                     // 扫码类型
}

type OpStatistic struct { // 行为统计
	Uuid       int64   `gorm:"primary_key;not null" json:"uuid"` // 用户ID
	IPList     string  `gorm:"type:text" json:"iplist"`          // 登录ip信息
	Amount     float64 `json:"amount"`                           // 总下注金额
	GameAmount string  `gorm:"type:text" json:"gameamount"`      // 游戏下注额
	Count      int64   `json:"count"`                            // 总登录次数
	DateCount  string  `gorm:"type:text" json:"datecount"`       // 每日登录次数
}

type AuthorityList struct { // 行为统计
	ID        int64  `gorm:"primary_key;not null" json:"id"`             // ID
	Name      string `gorm:"type:text" json:"name"`                      // 名字
	Group     string `gorm:"NOT NULL" json:"group"`                      // 权限分组，显示用
	Sort      int64  `gorm:"not null;DEFAULT 0" json:"sort"`             // 排序，显示用
	ClassName string `gorm:"type:varchar(40);NOT NULL" json:"classname"` // 后端 控制器名
	FunName   string `gorm:"NOT NULL" json:"funname"`                    // 函数名
}

type Datares struct { // 数据来源
	ID      int64  `gorm:"primary_key;not null" json:"id"` // ID
	Uuid    int64  `json:"uuid"`                           // 账号ID
	Expinfo string `json:"expinfo"`
}

type MaxOdds struct {
	PortID       int64   `json:"portid"`
	MaxNum       float64 `json:"maxnum"`
	MinAmount    float64 `json:"minamount"`
	MaxAmount    float64 `json:"maxamount"`
	ExpectAmount float64 `json:"ExpectAmount"`
}

type CompanyOdds struct { // 公司基准赔率
	CompanyID   int64   `gorm:"not null" json:"companyid"` // 公司ID
	Port        string  `gorm:"type:varchar(3);not null" json:"port"`
	RoomID      int64   `json:"roomid"`
	MinOdds     float64 `json:"minodds"` // 允许的最低赔率
	OpTime      int64   `json:"optime"`
	ValidTime   int64   `json:"validtime"`
	InUse       int64   `json:"inuse"`
	MaxoddsInfo string  `gorm:"type:text" json:"maxoddsinfo"` // 最大赔率信息
}

type CompanyPortodds struct { // 公司基准赔率
	CompanyID       int64   `gorm:"not null" json:"companyid"` // 公司ID
	Port            string  `gorm:"type:varchar(3);not null" json:"port"`
	RoomID          int64   `json:"roomid"`
	PortID          int64   `json:"portid"`
	DefaultOdds     float64 `json:"defaultodds"`     // 允许的最低赔率
	MinAmount       float64 `json:"minamount"`       // 单注最低额
	MaxAmount       float64 `json:"maxamount"`       // 单注最大额
	MaxAmountExpect float64 `json:"maxamountexpect"` // 玩法单期最大额
	IsTransfer      int64   `json:"istranser"`       //
	TransferAmount  int64   `json:"transferamount"`  //
	WarningAmount   float64 `json:"warningamount"`   // 警示金额
	PortSwitch      int64   `json:"portswitch"`      // 未知
}

type ProfitFormExt struct { // 外接游戏盈利报表
	ProfitTotal float64 `json:"profit_total"` // 总输赢
	Tip         float64 `json:"tip"`          // 小费
	RedPacket   float64 `json:"redpacket"`    // 红包
	Amount      float64 `json:"amount"`       // 总投注额
	AmountValid float64 `json:"amount_valid"` // 有效投注额
	Ordernum    int64   `json:"ordernum"`     // 投注数量
}

type MoneyInOut struct { // 资金变动日志
	OrderID         int64   `gorm:"primary_key;not null" json:"orderid"` // 订单ID
	Uuid            int64   `json:"uuid"`                                // 用户ID
	Account         string  `json:"account"`                             // 用户账户
	MasterID        int64   `json:"masterid"`                            // 站长ID
	ReqTime         int64   `json:"reqtime"`                             // 下单时间
	ResTime         int64   `json:"restime"`                             // 操作时间
	ReqType         int64   `json:"reqtype"`                             // 类型 (充值、提现等)
	PayType         string  `json:"pay_type"`                            // 用户支付方式(第三方支付通道ID)
	Way             int64   `json:"way"`                                 // 操作方式 (在线取款、在线充值等)
	RechargeInfo    string  `json:"rechargeinfo"`                        // 充值信息
	QRCodeUrl       string  `json:"qrcodeurl"`                           // 二维码地址
	Vercode         string  `json:"vercode"`                             // 验证码
	EnableTime      int64   `json:"enable_time"`                         // 二维码超时时间
	CashType        string  `json:"cashtype"`                            // 币别
	Bankinfo        string  `json:"bankinfo"`                            // 银行卡信息
	Amount          float64 `json:"amount"`                              // 金额
	ProfitCO        float64 `json:"profit_co"`                           // 返佣
	ProfitSaleIn    float64 `json:"profit_salein"`                       // 充值折扣
	ProfitSaleThird float64 `json:"profit_salethird"`                    // 其它折扣
	OpTip           float64 `json:"optip"`                               // 手续费
	OldGold         float64 `json:"oldgold"`                             // 操作前资金
	NewGold         float64 `json:"newgold"`                             // 操作后资金
	State           int64   `json:"state"`                               // 状态
	Checked         int     `json:"checked"`                             // 校验订单
	Operator        string  `json:"operator"`                            // 操作员
	OpResult        int64   `json:"opresult"`                            // 操作结果
	IP              string  `json:"ip"`                                  // IP
	IPPlace         string  `json:"ipplace"`                             // IP归属地
	UserPlatform    string  `json:"userpaltform"`                        // 用户请求平台
	Expinfo         string  `json:"expinfo"`                             // 其他备注信息
}

type Withdraw struct { // 提款
	OrderID         int64   `gorm:"primary_key;AUTO_INCREMENT;not null" json:"orderid"` // 订单ID
	Uuid            int64   `json:"uuid"`                                               // 用户ID
	Account         string  `json:"account"`                                            // 用户账户
	MasterID        int64   `json:"masterid"`                                           // 站长ID
	ReqTime         int64   `json:"reqtime"`                                            // 下单时间
	ResTime         int64   `json:"restime"`                                            // 操作时间
	Way             string  `json:"way"`                                                // 取款方式
	Bankinfo        string  `json:"bankinfo"`                                           // 银行卡信息
	IP              string  `json:"ip"`                                                 // 操作IP
	IPPlace         string  `json:"ipplace"`                                            // IP归属地
	CashType        string  `json:"cashtype"`                                           // 币别
	Amount          float64 `json:"amount"`                                             // 金额
	Checked         int     `json:"checked"`                                            // 校验订单
	Operator        string  `json:"operator"`                                           // 操作员
	OpResult        int64   `json:"opresult"`                                           // 操作结果
	ProfitCO        float64 `json:"profit_co"`                                          // 返佣
	ProfitSaleIn    float64 `json:"profit_salein"`                                      // 充值折扣
	ProfitSaleThird float64 `json:"profit_salethird"`                                   // 其它折扣
	OpTip           float64 `json:"optip"`                                              // 手续费
	OldGold         float64 `json:"oldgold"`                                            // 原金额
	NewGold         float64 `json:"newgold"`                                            // 新金额
	UserPlatform    string  `json:"userpaltform"`                                       // 用户请求平台
	Expinfo         string  `json:"expinfo"`                                            // 扩展信息
}

type WalletTransfer struct { // 钱包转换
	ID         int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"`
	Uuid       int64   `json:"uuid"`       // 用户ID
	Time       int64   `json:"time"`       // 时间
	ResWallet  string  `json:"reswallet"`  // 转出钱包
	ResState   int64   `json:"resstate"`   // 转出状态
	OutTime    int64   `json:"outtime"`    // 转出时间
	DestWallet string  `json:"destwallet"` // 转入钱包
	DestState  int64   `json:"deststate"`  // 转入状态
	InTime     int64   `json:"intime"`     // 转入时间
	Amount     float64 `json:"amount"`     // 金额
	Expinfo    string  `json:"expinfo"`    // 备注信息
}

type InoutStatistic struct { // 钱包转换
	Uuid      int64   `json:"uuid"`      // 用户ID
	Date      int64   `json:"date"`      // 日期
	InCount   int64   `json:"incount"`   // 转入笔数
	InAmount  float64 `json:"inamount"`  // 转入金额
	OutCount  int64   `json:"outcount"`  // 转出笔数
	OutAmount float64 `json:"outamount"` // 转出金额
}

type FeedbackMsg struct { // 反馈信息
	ID     int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"ID"`
	Time   int64  `json:"time"`   // 时间
	Uuid   int64  `json:"uuid"`   // 客户ID
	SrcID  int64  `json:"srcid"`  // 发信者ID
	Title  string `json:"title"`  // 标题
	Detail string `json:"detail"` // 详情
	Flag   int64  `json:"flag"`   // 阅读标志
}

//注单  (分表) 分表规则？
type CoRechargeWay struct {
	CompanyID      int64  `json:"companyid"`             // 公司ID
	WayID          int64  `json:"wayid"`                 // 通道ID
	State          int64  `json:"state"`                 // 通道状态
	Sort           int64  `json:"sort"`                  // 通道排序
	FixedamountUrl string `json:"fixedamount_url"`       // 定额转账接口
	RechargeUrl    string `json:"recharge_url"`          // 普通充值接口
	CallbackUrl    string `json:"callback_url"`          // 支付回调函数
	SignType       string `json:"sign_type"`             // 加密类型
	ThirdAccount   string `json:"third_account"`         // 商户号
	ThirdKey       string `json:"third_key"`             // 加密Key
	EnableTime     int64  `json:"enable_time"`           // 订单有效时间
	Text           string `gorm:"type:text" json:"text"` // 显示文本
}

//注单  (分表) 分表规则？
type CoSelfRechargeway struct {
	ID          int64  `gorm:"primary_key;AUTO_INCREMENT;not null"` // 通道ID
	CompanyID   int64  `json:"companyid"`                           // 公司ID
	WayType     string `json:"waytype"`                             // 付款类型
	Address     string `json:"address"`                             // 网点
	Bankuser    string `json:"bankuser"`                            // 收款人
	Bankname    string `json:"bankname"`                            // 银行名
	Bankaccount string `json:"bankaccount"`                         // 账号
	State       int64  `json:"state"`                               // 状态
	Sort        int64  `json:"sort"`                                // 排序
	Expinfo     string `json:"expinfo"`                             // 扩展信息
}

// 房间表
type CDataUrl struct {
	ID              int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	NameCN          string // 房间名字
	NameENG         string // 房间名字
	Urla            string `gorm:"type:varchar(100)"` // 下载地址
	UrlaRate        int64  // 下载频率
	UrlaLimit       int64  // 是否启用
	UrlaRepair      string `gorm:"type:varchar(100)"` // 修复地址
	UrlaRepairRate  int64  // 下载频率
	UrlaRepairLimit int64  // 是否启用
	Urlb            string `gorm:"type:varchar(100)"` // 下载地址
	UrlbRate        int64  // 下载频率
	UrlbLimit       int64  // 是否启用
	UrlbRepair      string `gorm:"type:varchar(100)"` // 修复地址
	UrlbRepairRate  int64  // 下载频率
	UrlbRepairLimit int64  // 是否启用
}

type CGameDataUrl struct {
	ID         int64  `gorm:"primary_key;AUTO_INCREMENT;not null"` // 用户ID
	NameENG    string // 房间名字 作为账号
	TableSuf   string // 表名后缀
	DataUrl    string `gorm:"type:varchar(100)"` // 数据下载地址
	DataFlag   int64
	DataRate   int64  // 下载周期
	RepairUrl  string `gorm:"type:varchar(100)"` // 数据修复地址
	RepairFlag int64
	RepairRate int64
}

type CShengxiaoball struct {
	ID        int64  `json:"id"`        //
	Date      int64  `json:"date"`      // 春节日期
	Shengxiao string `json:"shengxiao"` // 球号生肖
	Balls     string `json:"balls"`
	Nianxiao  string `json:"nianxiao"` // 年肖
}

// 微信配置表
type WXUrl struct {
	ID   int64  `gorm:"primary_key;AUTO_INCREMENT;not null"` // 通道ID
	Code string `json:"code"`                                // 编码
	Addr string `json:"addr"`                                // 跳转地址
}

type DayStatistic struct { // 日统计
	Uuid               int64   `gorm:"not null"`            // 分配ID
	Date               int64   `gorm:"not null"`            // 日期
	BeginAmount        float64 `json:"begin_amount"`        // 日初钱包金额
	BeginUnsettle      float64 `json:"begin_unsettle"`      // 日初未结算金额
	BeginUnwithdraw    float64 `json:"begin_unwithdraw"`    // 日初未体现金额
	RechargeReqcount   int64   `json:"recharge_reqcount"`   // 申请充值次数
	RechargeReqamount  float64 `json:"recharge_reqamount"`  // 申请充值金额
	RechargeSucccount  int64   `json:"recharge_succcount"`  // 成功充值次数
	RechargeSuccamount float64 `json:"recharge_succamount"` // 成功充值金额
	WithdrawReqcount   int64   `json:"withdraw_reqcount"`   // 申请充值次数
	WithdrawReqamount  float64 `json:"withdraw_reqamount"`  // 申请充值金额
	WithdrawSucccount  int64   `json:"withdraw_succcount"`  // 成功充值次数
	WithdrawSuccamount float64 `json:"withdraw_succamount"` // 成功充值金额
	LimitRechargecount int64   `json:"limitrechargecount"`  // 小于1000的提款次数
	ThirdRepair        float64 `json:"thirdrepair"`         // 第三方补单
	CoRecharge         float64 `json:"corecharge"`          // 公司线下收款
	RemitIn            float64 `json:"remitin"`             // 后台上分
	RemitOut           float64 `json:"remitout"`            // 后台下分
	Benifit            float64 `json:"benifit"`             // 优惠收入
	PreIn              float64 `json:"prein"`               // 上级转入
	SysRet             float64 `json:"sysret"`              // 系统退款
	OutSuf             float64 `json:"outsuf"`              // 转入下级
	EndAmount          float64 `json:"end_amount"`          // 日末钱包金额
	EndUnsettle        float64 `json:"end_unsettle"`        // 日末未结算金额
	EndUnwithdraw      float64 `json:"end_unwithdraw"`      // 日末未体现金额
}

//注单  (分表) 分表规则？
type Conaviset struct {
	ID       int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 自动生成ID
	Uuid     int64  `json:"uuid"`                                          // 用户ID
	MasterID int64  `json:"masterid"`                                      // 公司ID
	Navi1    string `json:"navi1"`                                         // 一级导航
	Navi2    string `json:"navi2"`                                         // 二级导航
	Value    int64  `json:"value"`                                         // 开关值
}

//注单  (分表) 分表规则？
type NavisetCfg struct {
	ID        int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 自动生成ID
	Navi1     string `json:"navi1"`                                         // 一级导航
	Navi2     string `json:"navi2"`                                         // 二级导航
	Value     int64  `json:"value"`                                         // 权限可传性(1,可用不可传 2 可用可传)
	Super     int64  `json:"super"`                                         // 超管
	Company   int64  `json:"company"`                                       // 公司
	Assistant int64  `json:"assistant"`                                     // 子账号
	Agent     int64  `json:"agent"`                                         // 代理
}

//注单  (分表) 分表规则？
type MoneychangType struct {
	ID        int64   `gorm:"primary_key;not null" json:"id"` // 类型ID
	Uuid      int64   `json:"uuid"`                           // 公司ID
	ShortName string  `json:"shortname"`                      // 活动简称
	Money     float64 `json:"money"`                          // 活动金额
	Begindate int64   `json:"begindate"`                      // 起始日期
	Enddate   int64   `json:"enddate"`                        // 截止日期
}

//公司升级规则
type LevelConfig struct {
	ID                 int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 类型ID
	Uuid               int64   `json:"uuid"`                                          // 公司ID
	Level              int64   `json:"level"`                                         // 等级
	Validamount        float64 `json:"validamount"`                                   // 流水金额
	RechargeSuccamount float64 `json:"recharge_succamount"`                           // 成功充值金额
}

// 查询用户下注统计
type Orderstats struct {
	ID          int64 `gorm:"primary_key;not null;AUTO_INCREMENT"` // 自动生成ID
	Date        int64
	RoomID      int64
	Expect      string
	ItemID      int64
	GameDalei   string
	GameMode    string
	GameItem    string
	Amount      float64
	Expinfo     string
	Official    int64
	Orderinfo   string
	OrderInfoOf OfOrderinfo //官方玩法下注详情
}

//注单  (分表) 分表规则？
type CBKItemType struct {
	RoomID     int64  `json:"roomid"`   // 房间ID
	BKTypeID   int64  `json:"bktypeid"` // 后台分类
	BKTypename string `json:"bktypename"`
	ItemID     int64  `json:"itemid"` // 下注项
}

// 服务MD5key 值
type ServerMD5 struct {
	ID         int64  `gorm:"primary_key;not null" json:"id"` //
	ServerName string `json:"servername"`                     //
	Port       string `json:"port"`                           // 服务的端口
	Md5Key     string `json:"md5key"`                         // 服务的MD5签
	InUse      int64  `json:"inuser"`                         // 启用状态
}

// 限红详情
type CCoQuota struct {
	ID       int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` //
	Uuid     int64   `json:"uuid"`                                          //
	MinQuota float64 `json:"minquota"`                                      //
	MaxQuota float64 `json:"maxquota"`                                      //
}

// 用户限红信息
type CUserQuota struct {
	Uuid     int64 `json:"uuid"`     //
	Preid    int64 `json:"preid"`    //
	Masterid int64 `json:"masterid"` //
	QuotaID  int64 `json:"quotaid"`  //
}

type WininfoDZBJ struct { // 电子百家的开奖号码
	RoomID     int64  `json:"roomid"`     // 房间号
	Expect     string `json:"expect"`     // 期号
	ResID      int64  `json:"resid"`      // 荷官ID
	Type       int64  `json:"type"`       // 类型(不记得干嘛了)
	Boot       int64  `json:"boot"`       // 靴数
	Round      int64  `json:"round"`      // 局数
	Opencode   string `json:"opencode"`   // 中奖号码
	Bankeritem string `json:"bankeritem"` // 庄牌
	Playeritem string `json:"palyeritem"` // 闲牌
	Opentime   int64  `json:"opentime"`   // 开奖时间
}

type WininfoBRNN struct { // 百人牛牛开奖号码
	RoomID     int64  `json:"roomid"`     // 房间号
	Expect     string `json:"expect"`     // 期号
	ResID      int64  `json:"resid"`      // 庄家ID
	Opencode   string `json:"opencode"`   // 中奖Item
	Bankeritem string `json:"bankeritem"` // 庄牌
	Tianitem   string `json:"tianitem"`   // 天牌
	Diitem     string `json:"diitem"`     // 地牌
	Renitem    string `json:"renitem"`    // 人牌
	Heitem     string `json:"heitem"`     // 和牌
	Opentime   int64  `json:"opentime"`   // 开奖时间
}

type WininfoHHDZ struct { // 百人牛牛开奖号码
	RoomID    int64  `json:"roomid"`    // 房间号
	Expect    string `json:"expect"`    // 期号
	ResID     int64  `json:"resid"`     // 庄家ID
	Opencode  string `json:"opencode"`  // 中奖Item
	Blackitem string `json:"blackitem"` // 天牌
	Reditem   string `json:"reditem"`   // 地牌
	Luckyodds int64  `json:"luckyodds"` // 幸运一击赔率
	Opentime  int64  `json:"opentime"`  // 开奖时间
}

type WininfoLHDZ struct { // 百人牛牛开奖号码
	RoomID   int64  `json:"roomid"`   // 房间号
	Expect   string `json:"expect"`   // 期号
	ResID    int64  `json:"resid"`    // 庄家ID
	Opencode string `json:"opencode"` // 中奖Item
	Huitem   string `json:"huitem"`   // 天牌
	Longitem string `json:"longitem"` // 地牌
	Opentime int64  `json:"opentime"` // 开奖时间
}

type WininfoFQZS struct { // 飞禽走兽开奖号码
	RoomID   int64  `json:"roomid"`   // 房间号
	Expect   string `json:"expect"`   // 期号
	ResID    int64  `json:"resid"`    // 庄家ID
	Opencode string `json:"opencode"` // 中奖Item
	Opentime int64  `json:"opentime"` // 开奖时间
}

type WininfoWFLM struct { // 五福临门开奖号码
	RoomID    int64  `json:"roomid"`    // 房间号
	Expect    string `json:"expect"`    // 期号
	ResID     int64  `json:"resid"`     // 庄家ID
	Opencode  string `json:"opencode"`  // 中奖Item
	Blackitem string `json:"blackitem"` // 天牌
	Reditem   string `json:"reditem"`   // 地牌
	Luckyodds int64  `json:"luckyodds"` // 幸运一击赔率
	Opentime  int64  `json:"opentime"`  // 开奖时间
}

// 真人百家房间信息表
type CACTRoominfo struct {
	ID           int64   `gorm:"primary_key"` // 房间ID
	NameCN       string  // 房间名字
	NameENG      string  // 房间名字 作为账号
	GameDalei    string  // 游戏大类
	TableSuf     string  // 表名后缀
	SettleType   string  // 结算类型
	Masterid     int64   // 公司
	DataResid    int64   // 数据来源
	CreateTime   int64   // 房间创建时间
	InValidTime  int64   // 房间过期时间
	LimitNum     int64   // 限制人数
	Boot         int64   // 局号
	Round        int64   // 靴号
	MinAmount    float64 // 房间最低限红
	MaxAmount    float64 // 房间最高限红
	Opentime     int64   // 可下注时间   (百家专用)
	Bantime      int64   // 封盘时间
	VideoUrl     string  // 视频源
	ExpInfo      string  // 扩展信息
	VideoTableId int64   // 视频桌号
}

// 电子百家房间信息
type CSLOTRoomInfo struct {
	ID          int64   `gorm:"primary_key;not null"` //
	NameCN      string  // 房间名字
	NameENG     string  // 房间名字 作为账号
	GameDalei   string  // 游戏大类
	TableSuf    string  // 表名后缀
	SettleType  string  // 结算类型
	GameName    string  // 挂靠游戏名
	CreateTime  int64   // 房间创建时间
	InValidTime int64   // 房间过期时间
	MinAmount   float64 // 房间最低限红 (百家专用)
	MaxAmount   float64 // 房间最高限红 (百家专用)
	LimitNum    int64   // 限制人数
	Boot        int64   // 局号
	Round       int64   // 靴号
	Opentime    int64   // 可下注时间   (百家专用)
	Bantime     int64   // 封盘时间   (百家专用)
	Masterid    int64   // 公司id	(所有房间都有)
}

// 奔驰宝马房间信息
type CBZBMRoomInfo struct {
	ID          int64   `gorm:"primary_key;not null"` //
	NameCN      string  // 房间名字
	NameENG     string  // 房间名字 作为账号
	GameDalei   string  // 游戏大类
	TableSuf    string  // 表名后缀
	SettleType  string  // 结算类型
	GameName    string  // 挂靠游戏名
	Masterid    int64   // 公司
	CreateTime  int64   // 房间创建时间
	InValidTime int64   // 房间过期时间
	EnterAmount float64 // 进入要求
	MinAmount   float64 // 房间最低限红
	MaxAmount   float64 // 房间最高限红
	LimitNum    int64   // 限制人数
	Jackpot     float64 // 奖池
	Expect      int64   // 期号
	Opentime    int64   // 可下注时间
	Bantime     int64   // 封盘时间
	Flashtime   int64   //  动画时间
	Waittime    int64   // 新开局等待时间
	Commission  float64 // 佣金提成
}

// 百人牛牛房间信息
type CBRNNRoomInfo struct {
	ID          int64   `gorm:"primary_key;not null"` //
	NameCN      string  // 房间名字
	NameENG     string  // 房间名字 作为账号
	GameDalei   string  // 游戏大类
	TableSuf    string  // 表名后缀
	SettleType  string  // 结算类型
	GameName    string  // 挂靠游戏名
	Masterid    int64   // 公司
	CreateTime  int64   // 房间创建时间
	InValidTime int64   // 房间过期时间
	RoomType    string  // 房间倍率类型
	EnterAmount float64 // 进入要求
	MinAmount   float64 // 房间最低限红
	MaxAmount   float64 // 房间最高限红
	LimitNum    int64   // 限制人数
	SeatCount   int64   // 房间座位数
	Jackpot     float64 // 奖池
	Expect      int64   // 期号
	Opentime    int64   // 可下注时间
	Bantime     int64   // 封盘时间
	Flashtime   int64   //  动画时间
	Waittime    int64   // 新开局等待时间
	Commission  float64 // 新开局等待时间
}

// 奔驰宝马房间信息
type CFQZSRoomInfo struct {
	ID          int64   `gorm:"primary_key;not null"` //
	NameCN      string  // 房间名字
	NameENG     string  // 房间名字 作为账号
	GameDalei   string  // 游戏大类
	TableSuf    string  // 表名后缀
	SettleType  string  // 结算类型
	GameName    string  // 挂靠游戏名
	Masterid    int64   // 公司
	CreateTime  int64   // 房间创建时间
	InValidTime int64   // 房间过期时间
	EnterAmount float64 // 进入要求
	MinAmount   float64 // 房间最低限红
	MaxAmount   float64 // 房间最高限红
	LimitNum    int64   // 限制人数
	Jackpot     float64 // 奖池
	Expect      int64   // 期号
	Opentime    int64   // 可下注时间
	Bantime     int64   // 封盘时间
	Flashtime   int64   //  动画时间
	Waittime    int64   // 新开局等待时间
	Commission  float64 // 佣金提成
}

// 真人百家设置
type CACTBJCfg struct {
	Uuid        int64  `gorm:"primary_key;not null"`
	Account     string // 彩种大类
	PreID       int64  // 上级ID
	MasterID    int64  // 站长ID
	Restriction string // 限红列表
}

// 真人百家设置
type CCOACTBJCfg struct {
	ID             int64   `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Uuid           int64   //
	Account        string  //
	TotalMin       float64 //
	TotalMax       float64 //
	BankerMin      float64 //
	BankerMax      float64 //
	PlayerMin      float64 //
	PlayerMax      float64 //
	TieMin         float64 //
	TieMax         float64 //
	BankerPairsMin float64 //
	BankerPairsMax float64 //
	PlayerPairsMin float64 //
	PlayerPairsMax float64 //
}

// 红黑大战房间信息
type CHHDZRoomInfo struct {
	ID          int64   `gorm:"primary_key;not null"` //
	NameCN      string  // 房间名字
	NameENG     string  // 房间名字 作为账号
	GameDalei   string  // 游戏大类
	TableSuf    string  // 表名后缀
	SettleType  string  // 结算类型
	GameName    string  // 挂靠游戏名
	Masterid    int64   // 公司
	CreateTime  int64   // 房间创建时间
	InValidTime int64   // 房间过期时间
	EnterAmount float64 // 进入要求
	MinAmount   float64 // 房间最低限红
	MaxAmount   float64 // 房间最高限红
	LimitNum    int64   // 限制人数
	Jackpot     float64 // 奖池
	Expect      int64   // 期号
	Opentime    int64   // 可下注时间
	Bantime     int64   // 封盘时间
	Flashtime   int64   //  动画时间
	Waittime    int64   // 新开局等待时间
	Commission  float64 // 佣金提成
}

// 龙虎大战房间信息
type CLHDZRoomInfo struct {
	ID          int64   `gorm:"primary_key;not null"` //
	NameCN      string  // 房间名字
	NameENG     string  // 房间名字 作为账号
	GameDalei   string  // 游戏大类
	TableSuf    string  // 表名后缀
	SettleType  string  // 结算类型
	GameName    string  // 挂靠游戏名
	Masterid    int64   // 公司
	CreateTime  int64   // 房间创建时间
	InValidTime int64   // 房间过期时间
	EnterAmount float64 // 进入要求
	MinAmount   float64 // 房间最低限红
	MaxAmount   float64 // 房间最高限红
	LimitNum    int64   // 限制人数
	Jackpot     float64 // 奖池
	Expect      int64   // 期号
	Opentime    int64   // 可下注时间
	Bantime     int64   // 封盘时间
	Flashtime   int64   //  动画时间
	Waittime    int64   // 新开局等待时间
	Commission  float64 // 佣金提成
}

// 五福临门房间信息
type CWFLMRoomInfo struct {
	ID          int64   `gorm:"primary_key;not null"` //
	NameCN      string  // 房间名字
	NameENG     string  // 房间名字 作为账号
	GameDalei   string  // 游戏大类
	TableSuf    string  // 表名后缀
	SettleType  string  // 结算类型
	GameName    string  // 挂靠游戏名
	Masterid    int64   // 公司
	CreateTime  int64   // 房间创建时间
	InValidTime int64   // 房间过期时间
	EnterAmount float64 // 进入要求
	MinAmount   float64 // 房间最低限红
	MaxAmount   float64 // 房间最高限红
	LimitNum    int64   // 限制人数
	Jackpot     float64 // 奖池
	Commission  float64 // 佣金提成
	Scrollinfo  string  `gorm:"type:text;not null"` // 滚筒设置
}

//黑名单信息
type LimitUserinfo struct {
	ID        int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 自动生成ID
	CompanyID int64  `json:"companyid"`                                     // 公司ID
	Uuid      int64  `json:"uuid"`                                          // 用户ID
	Account   string `json:"account"`                                       // 账户
	LimitFlag int64  `json:"limitflag"`                                     // 限制标志
	LimitTime int64  `json:"limittime"`                                     // 加入时间
}

type DefaultCompanyBase struct {
	CoInitcash     float64 `json:"coinitcash"`                           // 公司的初始金额
	RegOdds        float64 `json:"regodds"`                              // 用户注册的默认赔率
	RegShuilt      float64 `json:"regshuilt"`                            // 用户注册的默认退水
	MinOdds        float64 `json:"minodds"`                              // 公司的最低赔
	MaxOdds        float64 `json:"maxodds"`                              // 公司的最大赔
	MaxLottery     float64 `json:"maxlottery"`                           // 彩票最大佣金
	MaxActual      float64 `json:"maxactual"`                            // 视讯最大佣金
	MaxSlots       float64 `json:"maxslots"`                             // 电子最大佣金
	MaxSports      float64 `json:"maxsports"`                            // 体育最大佣金
	MaxCard        float64 `json:"maxcard"`                              // 棋牌最大佣金
	MinAmount      float64 `json:"minamount"`                            // 最低提现额度
	InitPwd        string  `json:"initpwd"`                              // 公司的初始密码
	InitCash       float64 `json:"initcash"`                             // 注册的初始资金
	FixedamountUrl string  `json:"fixedamount_url"`                      // 查询固定额度
	RechargeUrl    string  `json:"recharge_url"`                         // 支付接口
	CallbackUrl    string  `json:"callback_url"`                         // 充值完成后的回调接口
	SignType       string  `json:"sign_type"`                            // 校验类型
	ThirdAccount   string  `json:"thirdaccount"`                         // 第三方支付平台账号
	ThirdKey       string  `json:"thirdkey"`                             // 支付验证码
	FrontUrls      string  `gorm:"type:varchar(1000)" json:"fronturls"`  // 前端可用域名
	BackUrls       string  `gorm:"type:varchar(1000)" json:"backurls"`   // 后端可用域名
	NaviidList     string  `gorm:"type:varchar(1000)" json:"naviidlist"` // 权限列表
	Issetwhite     int64   `json:"issetwhite"`                           // 是否启用白名单
	WhiteList      string  `gorm:"type:varchar(400)" json:"whitelist"`   // 白名单列表
	ExpgameList    string  `gorm:"type:varchar(400)" json:"expgamelist"` // 外接游戏列表
}

// 用户拥有的后台目录
type UserNavi struct {
	Uuid    int64 `json:"uuid"`     // 用户ID
	NaviID  int64 `json:"naviid"`   // naviid
	GroupID int64 `json:"groupid"`  // groupid
	Value   int64 `json:"masterid"` // 状态值
}

// 超管拥有的权限
type SuperNavi struct {
	ID     int64 `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 自动生成ID
	Uuid   int64 `json:"uuid"`                                          // 用户ID
	NaviID int64 `json:"naviid"`                                        // 一级导航
	Value  int64 `json:"value"`                                         // 开关值
	Level  int64 `json:"level"`                                         // 开关值
	Level1 int64 `json:"level1"`                                        // 开关值
	Level2 int64 `json:"level2"`                                        // 开关值
}

// 用户拥有的后台目录
type BKNavi struct {
	ID      int64  `json:"id"`      // 目录ID
	NaviLv1 string `json:"navilv1"` // 一级目录
	NaviLv2 string `json:"navilv2"` // 二级目录
	NaviLv3 string `json:"navilv3"` // 三级目录
	GroupID int64  `json:"groupid"` // groupid
	Admin   int64  `json:"admin"`   // admin
	Saler   int64  `json:"saler"`   // 销售
	Company int64  `json:"company"` // 公司
	Agent   int64  `json:"agent"`   // 代理
	Player  int64  `json:"player"`  // 代理
}

// 用户拥有的后台目录
type SuperBKNavi struct {
	ID     int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 自动生成ID
	Name   string `json:"uuid"`                                          // 名字
	Level  int64  `json:"preid"`                                         // 层级
	Level1 int64  `json:"masterid"`                                      // 一级目录ID
	Level2 int64  `json:"naviid"`                                        // 二级目录ID
	Value  int64  `json:"value"`                                         // 开关值
}

// 公告信息
type CLoginInfo struct {
	ID        int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // ID
	Uuid      int64  `json:"uuid"`                                          // 用户ID
	LoginTime int64  `json:"logintime"`                                     // 登录时间
	Account   string `json:"account"`                                       // 账户
	IP        string `json:"ip"`                                            // IP地址
	IPPlace   string `json:"ipplace"`                                       // IP归属地
	Platform  string `json:"platform"`                                      // 登录平台(player,agent)
	ReqHost   string `json:"reqhost"`                                       // 登录线路
	Useragent string `json:"useragent"`                                     // 浏览器信息
}

// 公告信息
type CNoticeInfo struct {
	ID        int64  `json:"id"`        // ID
	CompanyID int64  `json:"companyid"` // 公司ID
	Platform  string `json:"platform"`  // 显示平台 (前端、后端、手机端等)
	GameType  string `json:"gametype"`  // 游戏类型 (棋牌、彩票、真人等)
	ShowType  string `json:"showtype"`  // 公告类型 (弹窗、系统讯息、通告等)
	Optime    int64  `json:"optime"`    // 创建/修改时间
	Begintime int64  `json:"begintime"` // 起始时间
	Endtime   int64  `json:"endtime"`   // 结束时间
	Info      string `json:"info"`      // 内容
}

// 公告信息
type CPromotionInfo struct {
	ID        int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // ID
	CompanyID int64  `json:"companyid"`                                     // 名字
	Title     string `json:"title"`                                         // 公告类型 (弹窗、系统讯息、通告等)
	LogoUrl   string `json:"logourl"`                                       // 游戏类型 (棋牌、彩票、真人等)
	Optime    int64  `json:"optime"`                                        // 公告创建日
	Begintime int64  `json:"begintime"`                                     // 公告起始日
	Endtime   int64  `json:"endtime"`                                       // 公告结束日
	Info      string `gorm:"type:text;not null" json:"info"`                // 内容
}

//控奖设置
type CoSetnumcfg struct {
	ID             int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 自动生成ID
	CompanyID      int64   `json:"uuid"`                                          // 公司ID
	RoomID         int64   `json:"roomid"`                                        // 房间ID
	NameENG        string  `json:"nameeng"`                                       // 游戏英文名
	NameCN         string  `json:"namecn"`                                        // 游戏中文名
	ExpectPer      float64 `json:"expectper"`                                     // 控奖期数比例
	RetwagerPerMin float64 `json:"retwagerpermin"`                                // 控奖最低返奖率
	RetwagerPerMax float64 `json:"retwagerpermax"`                                // 控奖最高返奖率
	RandPerMax     float64 `json:"randpermax"`                                    // 随机期最高返奖率
	Commission     float64 `json:"commission"`                                    // 提成比例
}

type Odds_Set struct {
	Superiorbaseodds float64 // 上级两面赔
	Juniorbaseodds   float64 // 下级两面赔
	Superiorodds     float64 // 上级赔率
	Juniorodds       float64 // 下级赔率
	Superiorshui     float64 // 上级给水
	Juniorshui       float64 // 给下级水
	Selfpercent      float64 // 自身分成
	Underpercent     float64 // 自身和下级所占分成
	DynimicOdds      float64 // 动态赔率
}

// 游戏名
type GameName struct {
	GameName string // 游戏名
}

type RetMsg struct {
	Status     bool
	Error_code ErrorType
	Msg        string
}

type ExpUrl struct {
	WebUrl    string `json:"weburl"`
	IPhoneUrl string `json:"iphoneurl"`
}

type RetLogin struct {
	Uuid        int64
	Sid         string
	ShuiLottery float64 `json:"shui_lottery"` // 彩票返点
	Account     string  `json:"account"`      // 账号
	IsResetpwd  bool    `json:"isresetpwd"`   // 首次登录修改密码标识
	IsGeneral   int64   `json:"isgeneral"`    // 弃用
	VIPLevel    int64   `json:"viplevel"`     // VIP等级
	QYType      string  `json:"qytype"`       // 契约模式用户类型 弃用
	MaxActual   float64 `json:"maxactual"`    // 真人百家最大退水
	MaxSlots    float64 `json:"maxslots"`     // 电子 最大退水
	MaxSports   float64 `json:"maxsports"`    // 体育 最大退水
	MaxCard     float64 `json:"maxcard"`      // 棋牌 最大退水
	Oddsinfo    string  `json:"oddsinfo"`     // 赔率
	WalletType  string  `json:"wallettype"`   // 用户钱包模式
	IsShiwan    int64   `json:"isshiwan"`     // 试用账号标识
	Authority   string  `json:"authority"`    // 用户权限列表 弃用
}

type SettleResult struct {
	Settlednum        int64   // 结算笔数
	Youxiaojine       float64 // 结算金额
	Shizhanjine       float64 // 代理所占金额
	Shizhanjieguo     float64 //
	Shizhantuishui    float64 //
	ProfitShui        float64 // 水差
	ProfitShuiDiff    float64 // 下注吃水
	ProfitWager       float64 // 赚赔
	ProfitWagerDiff   float64 // 吃赔
	Yingkuijieguo     float64 // 输赢结果
	UpShizhanjine     float64
	UpShizhanjieguo   float64
	UpShizhantuishui  float64
	UpProfitShui      float64
	UpProfitShuiDiff  float64
	UpProfitWager     float64
	UpProfitWagerDiff float64
	UpYingkuijieguo   float64
	Yingshouhuoliang  float64
	Shangjiaohuoliang float64
	Shangjiaojieguo   float64
}

type Message struct {
	Rows   int64
	Code   string
	Remain string
	Data   []Wininfo
}

type WSMessage_Out struct { // 发送给客户端
	ID        MsgID
	RetCode   ErrorType
	Body      string
	Timestamp int64
}

type WSMessage_In struct { // 客户端发送来的
	ID        MsgID
	Uuid      int64
	Body      string
	Timestamp int64
}

type WSMessage_Tcp struct { // 服务器内部适用.
	ID      MsgID     `json:"id"`
	RetCode ErrorType `json:"retcode"`
	Uuid    int64     `json:"uuid"`
	Body    string    `json:"body"`
}

type SettledMsg struct { // 结算完成协议
	RoomID  int64  `json:"roomid"`
	Expect  string `json:"expect"`
	Expinfo string `json:"expinfo"`
	Time    int64  `json:"id"`
}

type RecordList struct {
	Maxrecordnum int64       `json:"maxrecordnum"` // 最大记录数
	Maxpagenum   int64       `json:"maxpagenum"`   // 最大页码数
	Data         interface{} `json:"data"`         // 数据
}

/**********************
* 一中多信息
**********************/
type WinMulti struct {
	ItemID int64   // 中奖项
	Count  int64   // 中奖注数
	Odds   float64 // 中奖赔率
	Profit float64 // 中奖金额
}
type Betinfo struct {
	Id      int64   `json:"id"`      // 结果ID
	Money   float64 `json:"money"`   // 下注金额
	Num     string  `json:"num"`     // 连码下注号
	Expinfo string  `json:"expinfo"` // 扩展信息
}

//官方玩法扩展信息结构
type Expinfo struct {
	MonetaryUnit float64 `json:"monetary_unit"` //货币单位
	Multiple     int64   `json:"multiple"`      //倍率
	Ordernum     int64   `json:"ordernum"`      //注单个数
	Officeflag   int64   `json:"officeflag"`    //官方玩法标志
}

//五福临门扩展信息结构
type WflmExpinfo struct {
	LuckyFlag bool
	ExpInfo   string
}

type SetnumCfg struct {
	Id             int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"`
	SettleType     string  `json:"settletype"`    // 玩法类型
	ExpectPer      float64 `json:"expectper"`     // 杀期率
	RetwagerPerMin float64 `json:"profitper_min"` // 返奖率最低值
	RetwagerPerMax float64 `json:"profitper_max"` // 返奖率最高值
	RandPerMax     float64 `json:"randper_max"`   // 随机最大返奖值
}

type GetCodeInfo struct {
	Id         int64 `gorm:"primary_key;not null;AUTO_INCREMENT"`
	RoomID     int64
	Lottery    string `gorm:"type:varchar(40);not null"`
	Expect     string `gorm:"type:varchar(40);not null"`
	WishTime   int64
	OpenCode   string
	ArriveTime int64
	ArriveFlag int64 // 0 等待获取 1 获取成功  2 获取超时  3 提前到达
	ExpInfo    string
}

type LostExpect struct {
	Id         int64  `gorm:"primary_key;not null;AUTO_INCREMENT"`
	Lottery    string `gorm:"type:varchar(40);not null"`
	Expect     string `gorm:"type:varchar(40);not null"`
	Wishtime   int64
	Optime     int64
	Opencode   string
	Repairflag int64
	Time       int64
}

type OrderExpinfo struct {
	Count int64  `json:"count"` // 注数
	Multi int64  `json:"multi"` // 倍数
	Unit  string `json:"unit"`  // 单位
}

type GameNameInfo struct {
	GameNameENG string `json:"gamenameeng"` // 游戏名ENG
	GameNameCN  string `json:"gamenamecn"`  // 游戏名CN
}

// 订单的占成信息
type AgentZhanchenginfo struct {
	Uuid               int64   `json:"uuid"`               // 用户ID
	Account            string  `json:"account"`            // 账号
	RoleType           string  `json:"roletype"`           // 账号类型
	BaseZhancheng      float64 `json:"basezhancheng"`      // 初始占成
	TransferAmount     int64   `json:"transferamount"`     // 触发额度
	ZhanchengPer       float64 `json:"zhanchengper"`       // 我的实际占成
	ZhanchengAmount    float64 `json:"zhanchengamount"`    // 我的实占额
	SufZhanchengPer    float64 `json:"sufzhanchengper"`    // 下级实际占成
	SufZhanchengAmount float64 `json:"sufzhanchengamount"` // 下级实占额
	Tuishui            float64 `json:"pretuishui"`         // 上级给水
	SufTuishui         float64 `json:"suftuishui"`         // 给出退水
	ProfitTuishui      float64 `json:"profittuishui"`      // 代理赚水
	PreBuhuoset        int64   `json:"prebuhuoset"`        // 上级补货设置 0 关闭 1 打开
	MyBuhuoset         int64   `json:"mybuhuoset"`         // 自己补货设置 0 关闭 1 打开
	// ShouhuoPer         float64 `json:"shouhuoper"`         // 收货比例
	// ShouhuoAmount      float64 `json:"shouhuoamount"`      // 收货额
	// BuchuPer           float64 `json:"buchuper"`           // 补出比例
	// BuchuAmount        float64 `json:"buchuamount"`        // 补出额
}

type SessionValue struct {
	Uuid      int64  `json:"uuid"`      // uuid
	SID       string `json:"sid"`       // sid
	Account   string `json:"account"`   // account
	RoleType  string `json:"roletype"`  // roletype
	Expiry    int64  `json:"expiry"`    // sid过期时间
	IP        string `json:"ip"`        // 登录IP
	IPPlace   string `json:"ipplace"`   // IP归属地
	IsMobile  bool   `json:"ismobile"`  // 客户端类型
	Logintime int64  `json:"logintime"` // 登录时间
	Ordertime int64  `json:"ordertime"` // 下注时间
}

type Orderdata struct {
	AgentID       int64   `json:"uuid"`          // 用户ID
	Key           int64   `json:"id"`            // itemid /portid
	Info          string  `json:"info"`          // 中文解析
	ShizhanAmount float64 `json:"shizhanamount"` // 实占金额
	Amount        float64 `json:"amount"`        // 全部金额
}

// 下发给用户的开奖时间
type ToCli_Opentime struct {
	Nowkithe      string // 当前期号
	Closetime     int64  // 距离封盘的时间
	TemaClosetime int64  `json:"temaclosetime"` // 特码封盘时间
	Opentime      int64  // 开盘时间
	Totaltime     int64  // 一期的所有时长 用来做进度条总长度的
	Bantime       int64  // 距离开奖时间
	State         int64  // 房间状态
}

// 下发给用户的开奖时间
type ToCliKaipantime struct {
	Nowkithe   string `json:"nowkithe"`   // 当前期号
	Julikaipan int64  `json:"julikaipan"` // 距离开盘时间
	Kaipantime int64  `json:"kaipantime"` // 开盘时间
	State      int64  `json:"state"`      // 状态
}

// 下发给用户的开奖号码
type ToCli_Wininfo struct {
	Closetime  int64
	Lastkithe  string
	Nowkithe   string
	Opentime   int64
	Resultinfo string
	Resultnum  string
	Totaltime  int64
}

// 注单返回
type WarningOrder struct {
	Orderid         int64   `json:"orderid"`         // 自动生成ID
	Uuid            int64   `json:"uuid"`            // 下单用户ID
	Account         string  `json:"account"`         // 下单账户
	Roomid          int64   `json:"roomid"`          // 房间ID
	RoomENG         string  `json:"roomeng"`         // 房间英文名
	RoomCN          string  `json:"roomcn"`          // 房间中文名
	RoleType        string  `json:"roletype"`        // 房间中文名
	Expect          string  `json:"expect"`          // 期号
	LotteryDalei    string  `json:"platform"`        // 平台
	SettleType      string  `json:"settletype"`      // 结算类型
	GameDalei       string  `json:"gamedalei"`       // 玩法大类
	GameXiaolei     string  `json:"gamexiaolei"`     // 玩法小类
	Iteminfo        string  `json:"iteminfo"`        // 玩法名称
	Pan             string  `json:"pan"`             // 盘口
	PortID          int64   `json:"portid"`          // 下注种类
	ItemID          int64   `json:"itemid"`          // 下单结果ID
	Amount          float64 `json:"amount"`          // 投注金额
	IsZidongbuhuo   int64   `json:"iszidongbuhuo"`   // 自动补货标识
	IsShoudongbuhuo int64   `json:"isshoudongbuhuo"` // 手动补货标识
	WarningFlag     int64   `json:"warningflag"`     // 警示标志(1警示)
	ZhanchengInfo   string  `json:"zhanchenginfo"`   // 代理占成信息
	BaseZhancheng   float64 `json:"basezhancheng"`   // 代理的基础占成
	BuchuPer        float64 `json:"buchuper"`        // 代理补出占成
	ActualZhancheng float64 `json:"actualzhancheng"` // 代理的实占比例
	Orderinfo       string  `json:"orderinfo"`       // 投注内容(盘口, 投注金额, 该单注数, 上级给赔, 上级给水, 变赔率)
	TotalZhancheng  float64 `json:"totalzhancheng"`  // 代理的代理累计实占
	OrderNum        string  `json:"ordernum"`        // 下注号码
	Optime          int64   `json:"optime"`          // 投注时间
	SettleTime      int64   `json:"settletime"`      // 结算时间
	IsWin           int64   `json:"iswin"`           // 中奖列表
	State           int64   `json:"state"`           // 结算状态
	Opencode        string  `json:"opencode"`        // 开奖结果
	Winodds         float64 `json:"winodds"`         // 赔付赔率
	Wager           float64 `json:"wager"`           // 派彩金额
	Tuishui         float64 `json:"tuishui"`         // 退水
	Shuying         float64 `json:"shuying"`         // 输赢
	AgentAmount     float64 `json:"agentamount"`     // 本级占成(当前查看的代理)
	AgentShuying    float64 `json:"agentshuying"`    // 本级输赢(当前查看的代理)
	IsBuhuo         int64   `json:"isbuhuo"`         // 是否为补货单子
	Touzhufangshi   string  `json:"touzhufangshi"`   // 投注方式
}

// 计算类开奖的号码源
type RetHaomares struct {
	Expect        string // 期号
	Nameeng       string `json:"nameeng"`
	Opencode      string // 号码
	Opentime      string // 时间
	Opentimestamp int64  //
}

// 连码项 各盘口 统计信息
type LianmaPanStats struct {
	Pan           string  `json:"pan"`           // 盘口
	Amount        float64 `json:"amount"`        // 投注金额
	SzAmount      float64 `json:"szamount"`      // 实占金额
	Zidongbuchu   float64 `json:"zidongbuchu"`   // 自动补出
	Shoudongbuchu float64 `json:"shoudongbuchu"` // 手动补出
	Winodds       float64 `json:"winodds"`       // 赔付赔率
	Tuishui       float64 `json:"tuishui"`       // 退水
}

// 代理的 即时注单 连码项 统计信息
type AgentLianmastats struct {
	Roomid       int64            `json:"roomid"`      // 房间ID
	RoomENG      string           `json:"roomeng"`     // 房间英文名
	RoomCN       string           `json:"roomcn"`      // 房间中文名
	Expect       string           `json:"expect"`      // 期号
	LotteryDalei string           `json:"platform"`    // 平台
	SettleType   string           `json:"settletype"`  // 结算类型
	GameDalei    string           `json:"gamedalei"`   // 玩法大类
	GameXiaolei  string           `json:"gamexiaolei"` // 玩法小类
	Iteminfo     string           `json:"iteminfo"`    // 玩法名称
	PortID       int64            `json:"portid"`      // 下注种类
	ItemID       int64            `json:"itemid"`      // 下单结果ID
	OrderNum     string           `json:"ordernum"`    // 下注号码
	Optime       int64            `json:"optime"`      // 投注时间
	State        int64            `json:"state"`       // 结算状态
	RiskAmount   float64          `json:"riskamount"`  // 风险值
	PaninfoS     []LianmaPanStats `json:"paninfos"`    // 各盘详细信息
	// Amount        float64          `json:"amount"`        // 投注金额
	// SzAmount      float64          `json:"szamount"`      // 实占金额
	// Zidongbuchu   float64          `json:"zidongbuchu"`   // 自动补出
	// Shoudongbuchu float64          `json:"shoudongbuchu"` // 手动补出
	// Winodds    float64          `json:"winodds"`    // 赔付赔率
	// Tuishui    float64          `json:"tuishui"`    // 退水
}

type DaohangUrl struct {
	UrlType string `json:"urltype"` // 线路类型
	UrlCN   string `json:"urlcn"`   // 线路中文
	UrlENG  string `json:"urleng"`  // 线路域名
}

type XianluUrl struct {
	PcUrlCN   string `json:"pcurlcn"`   // PC线路中文
	PcUrlENG  string `json:"pcurleng"`  // PC线路域名
	MbdUrlCN  string `json:"mbdurlcn"`  // MBD线路中文
	MbdUrlENG string `json:"mbdurleng"` // MBD线路域名
	AdmUrlCN  string `json:"admurlcn"`  // ADM线路中文
	AdmUrlENG string `json:"admurleng"` // ADM线路域名
}

// 公司的初始化配置
type RetCompanyConfig struct {
	Gongsiming   string `json:"gongsiming"`  // 平台名称
	QiantaiLogo  string `json:"qiantailogo"` // 前台Logo
	HoutaiLogo   string `json:"houtailogo"`  // 前台Ico
	QiantaiIco   string `json:"qiantaiico"`  // 后台Logo
	HoutaiIco    string `json:"houtaiico"`   // 后台Ico
	QiantaiTitle string `json:"qiantaititle"`
	HoutaiTitle  string `json:"houtaititle"`
	PcLoginType  int64  `json:"pclogintype"`  // PC登录页
	PcGuoduType  int64  `json:"pcguodutype"`  // PC过渡页
	PcGameType   int64  `json:"pcgametype"`   // PC游戏
	MbdLoginType int64  `json:"mbdlogintype"` // MBD登录页
	MbdGameType  int64  `json:"mbdgametype"`  // MBD游戏
	AdmLoginType int64  `json:"admlogintype"` // ADM登录页
	AdmGameType  int64  `json:"admgametype"`  // ADM模板
	SkinType     int64  `json:"skintype"`     // 皮肤类型
	IsChatroom   int64  `json:"ischatroom"`   // 聊天室状态
	PcShiwan     int64  `json:"pcshiwan"`     // 试玩开关 0隐藏/1显示/2关闭
	Zhudanyuming string `json:"zhudanyuming"` // 注单域名
	Expcodeurl1  string `json:"expcodeurl1"`  // 外部链接1
	Expcodeurl2  string `json:"expcodeurl2"`  // 外部链接2
	Expcodeurl3  string `json:"expcodeurl3"`  // 外部链接3
	Expcodeurl4  string `json:"expcodeurl4"`  // 外部链接4
	Expcodeurl5  string `json:"expcodeurl5"`  // 外部链接5
	Expcodeurl6  string `json:"expcodeurl6"`  // 外部链接6
}

type RetFengkongSet struct {
	CompanyID    int64   `json:"companyid"`
	RoomID       int64   `json:"roomid"`
	QiuhaoInuse  int64   `json:"qiuhaoinuse"`
	QiuhaoAmount float64 `json:"qiuhaoamount"`
	QiuhaoNum    int64   `json:"qiuhaonum"`
}

type RetFengkongLog struct {
	Uuid                int64   `json:"uuid"`                // 下单用户ID
	Account             string  `json:"account"`             // 会员账号
	Nickname            string  `json:"nickname"`            // 会员昵称
	Realname            string  `json:"realname"`            // 真实姓名
	RoleType            string  `json:"roletype"`            // 角色类型
	WalletType          string  `json:"wallettype"`          // 钱包类型
	Level               int64   `json:"level"`               // 层级
	Cashtype            string  `json:"cashtype"`            // 币别
	CreateTime          int64   `json:"createtime"`          // 创建时间
	WalletAmount        float64 `json:"walletamount"`        // 钱包余额
	Yinglv              float64 `json:"yinglv"`              // 赢率
	Ordernum            int64   `json:"ordernum"`            // 下注笔数
	Jine                float64 `json:"touzhujine"`          // 下注金额
	RevokeAmount        float64 `json:"revokeamount"`        // 撤单金额
	SettledJine         float64 `json:"settledjine"`         // 已结算金额
	ValidJine           float64 `json:"validjine"`           // 有效金额
	ProfitWager         float64 `json:"profitwager"`         // 派奖
	Shuying             float64 `json:"shuying"`             // 输赢
	Tuishui             float64 `json:"tuishui"`             // 退水
	Limited             int64   `json:"limited"`             // 停押状态
	PlayerYingkuijieguo float64 `json:"playeryingkuijieguo"` // 会员盈亏结果
	ZhanchengPer        float64 `json:"zhanchengper"`        // 本级占成
	Online              int     `json:"online"`              // 在线状态
	FengkongDate        int64   `json:"fengkongdate"`        // 风控日期
	FengkongCount       int64   `json:"fengkongcount"`       // 风控次数
	FengkongNameCN      string  `json:"fengkongnamecn"`      // 风控游戏
	FengkongExpect      string  `json:"fengkongexpect"`      // 风控期号
	FengkongType        string  `json:"fengkongtype"`        // 风控类型
	FengkongInfo        string  `json:"fengkonginfo"`        // 风控信息
	DayYingkuijieguo    float64 `json:"dayyingkuijieguo"`    // 会员日盈亏结果
	Frequency           string  `json:"frequency"`           // 开奖频率
}
