package commonstruct

const (
	MD5KEY    = "fgdforewrewrqewq"         // 游戏跳转验签KEY
	MianmiKEY = "FDSREWORWOBObvvoeworewer" // 免密KEY
)

const (
	// 系统运行的基础配置表
	WY_gm_config_portclass          = "wy_gm_config_portclass"          // 盘口设置分类
	WY_gm_config_room               = "wy_gm_config_room"               // 房间基础信息表
	WY_gm_config_roomgame           = "wy_gm_config_roomgame"           // 房间挂靠游戏分类
	WY_gm_config_item               = "wy_gm_config_item"               // 玩法分类
	WY_gm_config_wallet             = "wy_gm_config_wallet"             // 钱包配置表
	WY_gm_config_bknavi             = "wy_gm_config_bknavi"             // 公司后台总目录表
	WY_gm_config_ballinfo           = "wy_gm_config_ballinfo"           // 球信息表
	WY_gm_config_level              = "wy_gm_config_level"              //
	WY_gm_config_recharge_way       = "wy_gm_config_recharge_way"       // 支付方式列表
	WY_gm_config_expgame            = "wy_gm_config_expgame"            // 外接游戏配置
	WY_gm_config_defaultodds        = "wy_gm_config_defaultodds"        // 系统默认最大赔率
	WY_gm_config_defaultcompanyinfo = "wy_gm_config_defaultcompanyinfo" // 系统默认的公司基本信息
	WY_gm_config_promotion          = "wy_gm_config_promotion"          // 所有活动列表
	WY_gm_config_bkset              = "wy_gm_config_bkset"              // 所有后台权限
	WY_gm_config_moneytype          = "wy_gm_config_moneytype"          // 账变分类配置表
	WY_gm_config_serverid           = "wy_gm_config_serverid"           // 服务器ID
	WY_gm_config_ip                 = "wy_gm_config_ip"                 // ip信息库
	WY_gm_config_ip_elide           = "wy_gm_config_ip_elide"           //
	WY_gm_config_shengxiaoball      = "wy_gm_config_shengxiaoball"      // 生肖库
	WY_gm_config_setnum             = "wy_gm_config_setnum"             // 自开配置信息
	WY_gm_config_serverkey          = "wy_gm_config_serverkey"          // 服务Key设置
	WY_gm_config_wxurl              = "wy_gm_config_wxurl"              // 微信配置表
	WY_gm_config_closevalue         = "wy_gm_config_closevalue"         // 提前封盘值
	WY_gm_config_serverset          = "wy_gm_config_serverset"          // 特殊的服务器设置
	WY_lotterytime_                 = "wy_lotterytime_"                 // 开奖时间表
	WY_num_                         = "wy_num_"                         // 开奖号码

	// 公司的基础信息表 (新系统搭建可以选择性转移)
	WY_company_dyncitem         = "wy_company_dyncitem"         // 公司的动态变赔表
	WY_company_rechargeway      = "wy_company_rechargeway"      // 公司的支付通道
	WY_company_selfrechargeway  = "wy_company_selfrechargeway"  // 公司入款通道
	WY_company_notice           = "wy_company_notice"           // 公司公告
	WY_company_bkset            = "wy_company_bkset"            // 公司后台权限设置
	WY_company_base             = "wy_company_base"             // 公司基本信息
	WY_company_promotion        = "wy_company_promotion"        // 公司活动列表
	WY_company_maxodds          = "wy_company_maxodds"          // 公司的最大赔率
	WY_company_game             = "wy_company_game"             // 公司的游戏表
	WY_company_portclass        = "wy_company_portclass"        // 公司的玩法种类信息表
	WY_company_portclass_jiashi = "wy_company_portclass_jiashi" // 公司的玩法种类信息表
	WY_company_limitinfo        = "wy_company_limitinfo"        // 公司限制的用户表
	WY_company_setnum           = "wy_company_setnum"           // 公司自行配置的自开配置信息
	WY_company_changlongodds    = "wy_company_changlongodds"    // 公司长龙变赔设置

	// 用户基础信息表 (新系统搭建可以选择性转移)
	WY_user_base         = "wy_user_base"         // 用户基本信息表
	WY_user_branch       = "wy_user_branch"       // 用户上下级关系表
	WY_user_money        = "wy_user_money"        // 用户资金表
	WY_user_navi_        = "wy_user_navi_"        // 获取用户的目录列表
	WY_user_odds         = "wy_user_odds"         // 代理的占成比
	WY_user_gameset      = "wy_user_gameset"      // 会员的游戏设置
	WY_user_portclassset = "wy_user_portclassset" // 会员的玩法设置
	WY_user_settletype   = "wy_user_settletype"   // 营销的赔率阈值表

	// 临时数据表
	WY_tmp_user_daystatistic     = "wy_tmp_user_daystatistic"     // 会员日统计
	WY_tmp_company_level         = "wy_tmp_company_level"         // 升级规则
	WY_tmp_user_retmsg           = "wy_tmp_user_retmsg"           // 反馈会员信息
	WY_tmp_user_teamstats        = "wy_tmp_user_teamstats"        // 公司统计
	WY_tmp_wait_setnum           = "wy_tmp_wait_setnum"           // 自开信息
	WY_tmp_order_stats           = "wy_tmp_order_stats"           // 下注统计
	WY_tmp_user_tag              = "wy_tmp_user_tag"              // 用户添加标签
	WY_tmp_inout_statistic       = "wy_tmp_inout_statistic"       // 用户存取统计
	WY_tmp_log_moneyinout        = "wy_tmp_log_moneyinout"        // 用户资金存取记录
	WY_tmp_user_zoufei           = "wy_tmp_user_zoufei"           // 走飞设置
	WY_tmp_user_datares          = "wy_tmp_user_datares"          // 真人数据源账号表
	WY_tmp_user_withdraw         = "wy_tmp_user_withdraw"         // 用户提现记录
	WY_tmp_user_op_statistic     = "wy_tmp_user_op_statistic"     // 用户行为统计
	WY_tmp_user_lotdata_data     = "wy_tmp_user_lotdata_data"     // 彩票数据 日统计 (交收报表)
	WY_tmp_user_lotdata_item     = "wy_tmp_user_lotdata_item"     // 彩票数据 item统计 (分类账报表、自动补货、警示注单)
	WY_tmp_user_lotdata_tema     = "wy_tmp_user_lotdata_tema"     // 特码统计数据(特码A、特码B合并)
	WY_tmp_tool_news             = "wy_tmp_tool_news"             // 新闻表
	WY_tmp_tool_changlong        = "wy_tmp_tool_changlong"        // 长龙统计表
	WY_tmp_user_dyncitem         = "wy_tmp_user_dyncitem"         //
	WY_tmp_lotteryorder          = "wy_tmp_lotteryorder"          // 游戏注单表
	WY_tmp_wait_settle           = "wy_tmp_wait_settle"           // 等待结算表
	WY_tmp_settleresult_expect   = "wy_tmp_settleresult_"         // 结算结果表 彩种
	WY_tmp_log_retcash           = "wy_tmp_log_retcash"           // 返现日志
	WY_tmp_log_settle            = "wy_tmp_log_settle"            // 结算日志
	WY_tmp_log_revoke            = "wy_tmp_log_revoke"            // 结算日志
	WY_tmp_wait_retcash          = "wy_tmp_wait_retcash"          // 等待返现表
	WY_tmp_wait_promotion        = "wy_tmp_wait_promotion"        // 等待发放活动表
	WY_tmp_log_money             = "wy_tmp_log_money"             // 用户资金日志表
	WY_tmp_log_port              = "wy_tmp_log_port"              // 盘口修改日志表
	WY_tmp_log_userop            = "wy_tmp_log_userop"            // 用户操作日志表
	WY_tmp_log_promotion         = "wy_tmp_log_promotion"         // 活动领取记录
	WY_tmp_log_warning           = "wy_tmp_log_warning"           // 活动领取记录
	WY_tmp_log_gentou            = "wy_tmp_log_gentou"            // 跟投日志
	WY_tmp_user_logininfo        = "wy_tmp_user_logininfo"        // 用户登录信息表
	WY_tmp_user_dyncchanglong    = "wy_tmp_user_dyncchanglong"    // 长龙变赔表
	WY_tmp_user_dyncamount       = "wy_tmp_user_dyncamount"       // 货量变赔表
	WY_tmp_user_ordertransfer    = "wy_tmp_user_ordertransfer"    // 货量走飞表
	WY_tmp_user_recharge         = "wy_tmp_user_recharge"         // 用户充值记录
	WY_tmp_user_wallettransfer   = "wy_tmp_user_wallettransfer"   // 钱包间转换记录
	WY_tmp_wait_getcode          = "wy_tmp_wait_getcode"          //
	WY_tmp_session               = "wy_tmp_session"               //
	WY_tmp_company_fengkongset   = "wy_tmp_company_fengkongset"   // 公司风控设置表
	WY_tmp_company_fengkonglog   = "wy_tmp_company_fengkonglog"   // 公司风控日志表
	WY_tmp_user_fengkongmark_log = "wy_tmp_user_fengkongmark_log" // 用户被风控标记日志
	WY_tmp_user_gentouplan       = "wy_tmp_user_gentouplan"       // 跟投设置表
	WY_tmp_log_huoliang          = "wy_tmp_log_huoliang"          // 货量记录
	WY_DaohangUrl                = "wy_daohang_url"               // 导航key=>url
	Daohang_key_url              = "daohang_key_url"              // 导航key=>url
)

// 彩种信息
type CRoomInfo struct {
	ID             int64  `gorm:"primary_key;AUTO_INCREMENT;not null"` // 用户ID
	NameCN         string // 房间名字
	NameENG        string // 房间名字 作为账号
	Walletid       int64  // 钱包ID
	GameDalei      string // 游戏平台
	GameType       string // 游戏种类
	SettleType     string // 结算类型
	SettleTypeCN   string // 结算类型中文
	SettleTypeSort int64  // 结算类型中文
	ClientShowtype string // 客户端显示类型
	Frequency      string // 开奖频率
	CodeType       string // 号码来源
	ResRoomeng     string // 计算类彩种号码来源
	CreateTime     int64  // 房间创建时间
	InValidTime    int64  // 房间过期时间
	LimitNum       int64  // 限制人数
	Opentime       int64  // 提前开盘量
	Closevalue     int64  // 系统基础封盘量
	TemaClosevalue int64  //
	Description    string // 游戏说明
}

// 公司基本信息
type CompanyBase struct {
	Uuid            int64   `json:"uuid"`            // 公司ID
	Account         string  `json:"account"`         // 公司账号
	NameCN          string  `json:"namecn"`          // 公司名称
	MinAmount       float64 `json:"minamount"`       // 单注提现下注金额
	IntervalTime    float64 `json:"intervaltime"`    // 提现间隔
	DayCount        int64   `json:"daycount"`        // 每日提现次数
	MaxOnlinenum    int64   `json:"maxonlinenum"`    // 最大在线人数
	Baseoddsper     float64 `json:"baseoddsper"`     // 赔率模式
	WithdrawSet     int64   `json:"withdrawset"`     // 提现额度流水倍率
	WalletSort      string  `json:"walletsort"`      // 钱包排序
	InitCash        float64 `json:"initcash"`        // 公司的初始金额
	InitPwd         string  `json:"initpwd"`         // 公司的初始密码
	Shiwanid        int64   `json:"shiwanid"`        // 试玩代理ID
	FixedamountUrl  string  `json:"fixedamount_url"` // 查询固定额度
	RechargeUrl     string  `json:"recharge_url"`    // 支付接口
	CallbackUrl     string  `json:"callback_url"`    // 充值完成后的回调接口
	SignType        string  `json:"sign_type"`       // 校验类型
	ThirdAccount    string  `json:"thirdaccount"`    // 第三方支付平台账号
	ThirdKey        string  `json:"thirdkey"`        // 支付验证码
	FrontUrls       string  `json:"fronturls"`       // 前端可用域名
	BackUrls        string  `json:"backurls"`        // 后端可用域名
	ShiwanCreateurl string  `json:"shiwancreateurl"` // 创建试玩账号地址
	Issetwhite      int64   `json:"issetwhite"`      // 是否启用白名单
	WhiteList       string  `json:"whitelist"`       // 白名单列表
	ExpgameList     string  `json:"expgamelist"`     // 外接游戏列表
	OperateState    int64   `json:"operatestate"`    // 运营状态
	DaohangKey      string  `json:"daohangkey"`      // 导航key
	DaohangUrl      string  `json:"daohangurl"`      // 导航url
	UpdateZhancheng int64
	UpdateBuhuo     int64
	Changlongcount  int64 `json:"changlongcount"` // 长龙期数
	LiankaiWarning  int64
	UpdateXianhong  int64
	MaxPaicai       int64  `json:"maxpaicai"`    // 最大派彩
	QiantaiLogo     string `json:"qiantailogo"`  // 前台Logo
	HoutaiLogo      string `json:"houtailogo"`   // 前台Ico
	QiantaiIco      string `json:"qiantaiico"`   // 后台Logo
	HoutaiIco       string `json:"houtaiico"`    // 后台Ico
	QiantaiTitle    string `json:"qiantaititle"` // 前台title
	HoutaiTitle     string `json:"houtaititle"`  // 后台title
	PcLoginType     int64  `json:"pclogintype"`  // PC登录页
	PcGuoduType     int64  `json:"pcguodutype"`  // PC过渡页
	PcGameType      int64  `json:"pcgametype"`   // PC游戏
	MbdLoginType    int64  `json:"mbdlogintype"` // MBD登录页
	MbdGameType     int64  `json:"mbdgametype"`  // MBD游戏
	AdmLoginType    int64  `json:"admlogintype"` // ADM登录页
	AdmGameType     int64  `json:"admgametype"`  // ADM模板
	SkinType        int64  `json:"skintype"`     // 皮肤类型
	IsChatroom      int64  `json:"ischatroom"`   // 聊天室状态
	PcShiwan        int64  `json:"pcshiwan"`     // 试玩开关 0隐藏/1显示/2关闭
	Peishuiquan     int64  `json:"peishuiquan"`  // 赔水权
	Zhudanyuming    string `json:"zhudanyuming"` // 注单外链域名
	KfzYuming       string `json:"kfzyuming"`    // 开发组域名
	Expcodeurl1     string `json:"expcodeurl1"`  // 外接开奖号码url
	Expcodeurl2     string `json:"expcodeurl2"`  // 外接开奖号码url
	Expcodeurl3     string `json:"expcodeurl3"`  // 外接开奖号码url
	Expcodeurl4     string `json:"expcodeurl4"`  // 外接开奖号码url
	Expcodeurl5     string `json:"expcodeurl5"`  // 外接开奖号码url
	Expcodeurl6     string `json:"expcodeurl6"`  // 外接开奖号码url

}

// 公告
type NewsInfo struct {
	ID   int64  `gorm:"primary_key;AUTO_INCREMENT;not null"` // 玩法ID
	Time int64  // 大类ID
	Info string `gorm:"type:text;not null"` // 公告内容
}

type BallInfo struct {
	ID       int64  `gorm:"primary_key;AUTO_INCREMENT;not null"` // 玩法ID
	Balltype string // 分类
	Balllist string `gorm:"type:text;not null"` // 球列表
}

type ChanglongInfo struct {
	ID        int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	GameDalei string `json:"gamedalei"` // 游戏平台
	GameName  string `json:"nameeng"`   // 游戏名
	ItemID    int64  `json:"itemid"`    // 下注ID
	Num       int64  `json:"num"`       // 连开次数
	Info      string `json:"info"`      // 下注项中文匹配
}

// 赔率设置分类
type CGameItem struct {
	ID            int64  `gorm:"primary_key;AUTO_INCREMENT;not null"` // 玩法ID
	LotteryDalei  string // 彩种大类
	GameType      string // 彩种类型
	GameDalei     string // 盘口大类
	GameXiaolei   string // 盘口小类
	GameMode      string // 具体玩法
	GameItem      string // 玩法具体项
	RelatedID     int64  // 关联ID （特码A/B关联）
	IsXiazhuitem  int64  // 下注项标识
	Islianma      int64  // 连码标识
	Ischanglong   int64  // 是否计入长龙列表标识 (0不计入 1 计入)
	RiskGroupid   int64  // 风险分组ID
	JszdGroupName string // 即时注单分组名
}

// 房间游戏分类
type CRoomGame struct {
	ID          int64  `gorm:"primary_key;AUTO_INCREMENT;not null"` // 用户ID
	GameName    string // 游戏分类英文名
	TableSuf    string // 表名后缀
	SettleType  string // 结算类型
	PercentType string // 分成分类
}

// 分类玩法赔率
type PortClass struct {
	ID           int64   // 用户ID
	LotteryDalei string  // 彩种大类
	SettleType   string  // 结算类型
	GameDalei    string  // 盘口大类
	GameXiaolei  string  // 盘口小类
	DefaultOdds  float64 // 默认赔率
	SysKaiguan   int64   // 系统玩法开关
}

// 营销玩法设置
type SalerBaseodds struct {
	Uuid         int64   // 用户ID
	Portid       int64   // 用户ID
	LotteryDalei string  // 彩种大类
	SettleType   string  // 结算类型
	GameDalei    string  // 盘口大类
	GameXiaolei  string  // 盘口小类
	DefaultOdds  float64 // 默认赔率
	DesOdds      float64 // 基准降赔
	// MaxOdds      float64 // 默认赔率
}

// 赔率的详细信息
type OddsInfo struct {
	GameXiaolei int64   `json:"gameid"` // 游戏小类ID
	PreOdds     float64 `json:"po"`     // 上级给赔
	SufOdds     float64 `json:"so"`     // 下级赔
	MinOdds     float64 `json:"mo"`     // 最小赔
	PreShui     float64 `json:"ps"`     // 上级水
	SufShui     float64 `json:"ss"`     // 下级水
	MinShui     float64 `json:"ms"`     // 最小水
}

// 赔率的详细信息
type LimitamountInfo struct {
	GameXiaolei     int64   `json:"gameid"` // 游戏小类ID
	PreMinamount    float64 `json:"pmina"`  // 上级最低限额
	SufMinamount    float64 `json:"smina"`  // 下级最低限额
	PreMaxamount    float64 `json:"pmaxa"`  // 上级最大额
	SufMaxamount    float64 `json:"smaxa"`  // 下级最大额
	PreMaxamountOne float64 `json:"pmaxas"` // 上级单注最大额 single
	SufMaxamountOne float64 `json:"smaxas"` // 下级单注最大额
}

// 盘口设置
type PortInfo struct {
	Uuid              int64 // 用户ID
	Gameid            int
	Port              string `gorm:"type:varchar(3)"`
	OddsInfoS         string `gorm:"type:text"` // 盘口详细水赔
	LimitamountInfoS  string `gorm:"type:text"` // 盘口详细限额
	ValidExpect_begin int64  // 起效期数
	ValidExpect_end   int64  // 失效期数
}

// 手动降赔(即时注单降赔)
type DynimicOddsInfo struct {
	Uuid        int64   `json:"uuid"`        // 用户ID
	RoomID      int64   `json:"roomid"`      // 房间ID
	Expect      string  `json:"expect"`      // 期数
	ItemID      int64   `json:"itemid"`      // 下注项
	DynimicOdds float64 `json:"dynimicodds"` // 动态赔率
	ValidTime   int64   `json:"validtime"`   // 生效时间
}

// 两面(长龙)降赔
type ChanglongOdds struct {
	Uuid      int64   `json:"uuid"`      // 用户ID
	RoomID    int64   `json:"roomid"`    // 房间ID
	Liankai1  float64 `json:"liankai1"`  //
	Liankai2  float64 `json:"Liankai2"`  //
	Liankai3  float64 `json:"Liankai3"`  //
	Liankai4  float64 `json:"Liankai4"`  //
	Liankai5  float64 `json:"Liankai5"`  //
	Liankai6  float64 `json:"liankai6"`  //
	Liankai7  float64 `json:"liankai7"`  //
	Liankai8  float64 `json:"liankai8"`  //
	Liankai9  float64 `json:"liankai9"`  //
	Liankai10 float64 `json:"liankai10"` //
	Liankai11 float64 `json:"liankai11"` //
	Liankai12 float64 `json:"liankai12"` //
	Liankai13 float64 `json:"liankai13"` //
	Liankai14 float64 `json:"liankai14"` //
	Liankai15 float64 `json:"liankai15"` //
	Liankai16 float64 `json:"liankai16"` //
	Liankai17 float64 `json:"liankai17"` //
	Liankai18 float64 `json:"liankai18"` //
	Liankai19 float64 `json:"liankai19"` //
	Liankai20 float64 `json:"liankai20"` //
	Weikai1   float64 `json:"Weikai1"`   //
	Weikai2   float64 `json:"Weikai2"`   //
	Weikai3   float64 `json:"Weikai3"`   //
	Weikai4   float64 `json:"Weikai4"`   //
	Weikai5   float64 `json:"Weikai5"`   //
	Weikai6   float64 `json:"Weikai6"`   //
	Weikai7   float64 `json:"Weikai7"`   //
	Weikai8   float64 `json:"Weikai8"`   //
	Weikai9   float64 `json:"Weikai9"`   //
	Weikai10  float64 `json:"Weikai10"`  //
	Weikai11  float64 `json:"Weikai11"`  //
	Weikai12  float64 `json:"Weikai12"`  //
	Weikai13  float64 `json:"Weikai13"`  //
	Weikai14  float64 `json:"Weikai14"`  //
	Weikai15  float64 `json:"Weikai15"`  //
	Weikai16  float64 `json:"Weikai16"`  //
	Weikai17  float64 `json:"Weikai17"`  //
	Weikai18  float64 `json:"Weikai18"`  //
	Weikai19  float64 `json:"Weikai19"`  //
	Weikai20  float64 `json:"Weikai20"`  //
}

// 长龙变赔表
type DyncChanglong struct {
	Uuid     int64  `gorm:"not null" json:"uuid"`      // 用户ID
	RoomID   int64  `json:"roomid"`                    // 房间
	DyncInfo string `gorm:"type:text" json:"dyncinfo"` // 变赔信息
}

// 货量变赔表
type DyncAmount struct {
	Uuid     int64  `gorm:"not null" json:"uuid"`      // 用户ID
	RoomID   int64  `json:"roomid"`                    // 房间
	DyncInfo string `gorm:"type:text" json:"dyncinfo"` // 变赔信息
}

// 货量变赔表
type Transfer struct {
	Uuid         int64  `gorm:"not null" json:"uuid"`          // 用户ID
	RoomID       int64  `json:"roomid"`                        // 房间
	TransferInfo string `gorm:"type:text" json:"transferinfo"` // 变赔信息
}

//注单  (分表) 分表规则？
type OrderS struct {
	Orderid         int64   `gorm:"primary_key;not null;AUTO_INCREMENT"` // 自动生成ID
	ExpOrderid      string  `json:"exporderid"`                          // 第三方订单号
	ParentID        int64   `json:"parentid"`                            // 父级关联单号
	Uuid            int64   // 下单用户ID
	Account         string  // 账户
	RoleType        string  // 角色类型
	Roomid          int64   // 房间ID
	Roomeng         string  // 房间名
	Roomcn          string  // 房间中文名
	Pan             string  // 盘口
	Level           int64   // 用户层级
	PreID           int64   // 上级代理ID
	MasterID        int64   // 站长ID
	ResID           int64   // 结算对象
	Expect          string  // 期号
	LotteryDalei    string  // 平台  LotteryDalei = lottery_dalei
	SettleType      string  // 结算类型
	GameDalei       string  // 玩法大类
	GameXiaolei     string  // 玩法小类
	PortID          int64   // 下注种类
	ItemID          int64   // 下单结果ID
	Iteminfo        string  // 下注项信息
	Official        int64   // 官方玩法标志
	Amount          float64 // 投注金额
	IsZidongbuhuo   int64   // 自动补货标识
	IsShoudongbuhuo int64   // 手动补货标识
	Touzhufangshi   string  // 投注方式
	WarningFlag     int64   // 警示标志(1警示)
	ZhanchengInfo   string  // 该单上级占成信息
	Numberinfo      string  // 号码信息
	BaseOdds        float64 // 基础赔率
	DynimicOdds     float64 // 货量变赔
	ChanglongOdds   float64 // 长龙变赔
	ManagerOdds     float64 // 手动变赔
	OrderNum        string  // 投注号码
	WinItems        string  // 投注号码
	Orderinfo       string  // 投注内容(盘口, 投注金额, 该单注数, 上级给赔, 上级给水, 变赔率)
	Optime          int64   // 投注时间
	State           int64   // 订单状态(0 未结算 1 撤单)
	SettleTime      int64   // 结算时间
	Opencode        string  // 开奖结果
	IsWin           int64   // 是否中奖
	WinList         string  // 中奖信息
	Winodds         float64 // 赔付赔率(baseodds - dynimicodds - changlongodds - managerodds)
	Wager           float64 // 派彩金额
	ShuiPer         float64 // 用户退水比例
	ShuiAmount      float64 // 用户退水
	IsStats         int64   // 是否统计订单
	Expinfo         string  `gorm:"type:text"` // 扩展信息(前端版本等等)
}

type OrderInfo struct {
	Pan         string  `json:"pan"`         // 盘口
	Num         string  `json:"Num"`         // 注单号码
	Count       int64   `json:"Count"`       // 该单注数
	Multiple    int64   `json:"Multiple"`    // 倍率
	PreOdds     float64 `json:"PreOdds"`     // 上级给赔
	PreShui     float64 `json:"PreShui"`     // 上级给水
	DynimicOdds float64 `json:"DynimicOdds"` // 变赔率
	Expinfo     string  `json:"Expinfo"`     // 扩展信息
}

// 订单可中奖项信息
type WinIteminfo struct {
	ItemID      int64   `json:"itemid"`      // itemid
	PreOdds     float64 `json:"preodds"`     // 上级给赔
	DynimicOdds float64 `json:"dynimicodds"` // 变赔率
}

// 下注号码信息
type NumberInfo struct {
	Num          string        `json:"num"`          // 注单号码
	WinIteminfoS []WinIteminfo `json:"winiteminfos"` // 可中奖项列表
}

type OrderInfoOf struct {
	Port         string  // 盘口
	Num          string  // 注单号码
	Count        int64   // 该单注数
	Multiple     int64   //倍率
	MonetaryUnit float64 //货币单位
	PreOdds      float64 // 上级给赔
	PreShui      float64 // 上级给水
	DynimicOdds  float64 // 变赔率
	Expinfo      string  // 扩展信息
}

//结算结果 - 按彩种分
type SettleresultExpect struct {
	Uuid              int64   `gorm:"not null" json:"uuid"` // 下单用户ID
	Date              int64   `gorm:"not null" json:"date"`
	Roomid            int64   `gorm:"not null" josn:"roomid"`
	Expect            string  `gorm:"type:varchar(40);not null" json:"expect"`
	Touzhujine        float64 `json:"touzhujine"`        // 投注金额
	Ordernum          int64   `json:"ordernum"`          // 投注数
	Settlednum        int64   `json:"settlednum"`        // 已结算数量
	Youxiaojine       float64 `json:"youxiaojine"`       // 有效金额
	Shizhanjine       float64 `json:"shizhanjine"`       // 实占金额中
	Shizhanjieguo     float64 `json:"shizhanjieguo"`     // 实占结果
	Shizhantuishui    float64 `json:"shizhantuishui"`    // 实占退水
	ProfitShui        float64 `json:"profitshui"`        // 赚水
	ProfitShuiDiff    float64 `json:"profitshuidiff"`    // 吃水
	ProfitWager       float64 `json:"profitwager"`       // 派奖
	ProfitWagerDiff   float64 `json:"profitwagerdiff"`   // 代理赚赔
	Yingkuijieguo     float64 `json:"yingkuijieguo"`     // 盈亏结果
	UpShizhanjine     float64 `json:"upshizhanjine"`     // 上级实占金额中，自己的贡献部分
	UpShizhanjieguo   float64 `json:"upshizhanjieguo"`   // 上级实占结果中，自己的贡献部分
	UpShizhantuishui  float64 `json:"upshizhantuishui"`  // 上级实占退水中，自己的贡献部分
	UpProfitShui      float64 `json:"profitshui"`        // 赚水
	UpProfitShuiDiff  float64 `json:"profitshuidiff"`    // 上级赚水中，自己的贡献部分
	UpProfitWager     float64 `json:"profitwager"`       // 派奖
	UpProfitWagerDiff float64 `json:"profitwagerdiff"`   // 上级赚赔中，自己的贡献部分
	UpYingkuijieguo   float64 `json:"upyingkuijieguo"`   // 上级盈亏结果中，自己的贡献部分
	Yingshouhuoliang  float64 `json:"yingshouxiaxian"`   // 应收下线
	Shangjiaohuoliang float64 `json:"shangjiaohuoliang"` // 上交货量
	Shangjiaojieguo   float64 `json:"shangjiaojieguo"`   // 上交结果
	SettleTime        int64   `json:"settletime"`        // 是否结算
}

//结算结果 - 按用户分
type SettleresultUuid struct {
	Uuid              int64   `gorm:"not null" json:"uuid"`   // 下单用户ID
	Roomid            int64   `gorm:"not null" json:"roomid"` // 房间号
	Date              int64   `gorm:"not null" json:"date"`   // 日期
	Ordernum          int64   `json:"ordernum"`               // 下注数
	Jine              float64 `json:"touzhujine"`             // 投注金额
	Settlednum        int64   `json:"settlednum"`             // 已结算注数
	SettledJine       float64 `json:"settledjine"`            // 已结算金额 (不知道这个算不算有效金额)
	ProfitWager       float64 `json:"profitwager"`            // 派奖
	ProfitWagerDiff   float64 `json:"profitwagerdiff"`        // 代理赚赔
	ProfitShui        float64 `json:"profitshui"`             // 会员赚水
	ProfitShuiDiff    float64 `json:"profitshuidiff"`         // 会员赚水
	Yingkui           float64 `json:"yingkui"`                // 盈亏结果
	UpShizhanjine     float64 `json:"upshizhanjine"`          // 上级实占金额中，自己的贡献部分
	UpShizhanjieguo   float64 `json:"upshizhanjieguo"`        // 上级实占结果中，自己的贡献部分
	UpShizhantuishui  float64 `json:"upshizhantuishui"`       // 上级实占退水中，自己的贡献部分
	UpProfitShui      float64 `json:"profitshui"`             // 赚水
	UpProfitShuiDiff  float64 `json:"profitshuidiff"`         // 上级赚水中，自己的贡献部分
	UpProfitWager     float64 `json:"profitwager"`            // 派奖
	UpProfitWagerDiff float64 `json:"profitwagerdiff"`        // 上级赚赔中，自己的贡献部分
	UpYingkuijieguo   float64 `json:"upyingkuijieguo"`        // 上级盈亏结果中，自己的贡献部分
	Yingshouhuoliang  float64 `json:"yingshouxiaxian"`        // 应收下线
	Shangjiaohuoliang float64 `json:"shangjiaohuoliang"`      // 上交货量
	Shangjiaojieguo   float64 `json:"shangjiaojieguo"`        // 上交结果
}

// 用户基础信息表
type Users struct {
	Uuid            int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"uuid"` // 分配ID
	Account         string  `gorm:"type:varchar(40)" json:"account"`                 // 用户名登录
	AccountOgactual string  `gorm:"type:varchar(40)" json:"account_ogactual"`        // OG视讯账户
	AccountBbin     string  `gorm:"type:varchar(40)" json:"account_bbin"`            // BBIN 账户
	AccountSaba     string  `gorm:"type:varchar(40)" json:"account_saba"`            // SABA 账户
	AccountMgactual string  `gorm:"type:varchar(40)" json:"account_mgactual"`        // 用户名登录
	Wechat          string  `gorm:"type:varchar(400)" json:"wechat"`                 // 微信登录
	IsShiwan        int64   `json:"ishiwan"`                                         // 试玩账号标识
	Email           string  `gorm:"type:varchar(40)" json:"email"`                   // 邮箱登录
	WalletType      string  `gorm:"type:varchar(40)" json:"wallettype"`              // 会员钱包类型
	Iphone          string  `gorm:"type:varchar(40)" json:"iphone"`                  // 手机登录
	Realname        string  `gorm:"type:varchar(40)" json:"realname"`                // 真实姓名
	QQ              string  `gorm:"type:varchar(40)" json:"qq"`                      // QQ登录
	Exp1            string  `gorm:"type:varchar(40)" json:"exp1"`                    // Exp1登录
	Exp2            string  `gorm:"type:varchar(40)" json:"exp2"`                    // Exp2登录
	Exp3            string  `gorm:"type:varchar(40)" json:"exp3"`                    // Exp3登录
	Exp4            string  `gorm:"type:varchar(40)" json:"exp4"`                    // Exp4登录
	Nickname        string  `gorm:"type:varchar(40)" json:"nickname"`                // 昵称
	Userbrand       string  `gorm:"type:varchar(20)" json:"userbrand"`               // 一代品牌名
	Headpic         string  `gorm:"type:varchar(40)" json:"headpic"`                 // 头像名
	Sex             int64   `gorm:"type:varchar(40);DEFAULT 0" json:"sex"`           // 性别
	ExtCode         string  `gorm:"type:varchar(40)" json:"extcode"`                 // 推广码
	RoleType        string  `gorm:"type:varchar(40)" json:"roletype"`                // 角色类型
	LeaderID        int64   `json:"leaderid"`                                        // 领导ID(子站号专用)
	MasterID        int64   `json:"masterid"`                                        // 站长ID
	PreID           int64   `json:"preid"`                                           // 上级ID
	Limited         int64   `json:"limited"`                                         // 登陆限制
	IsGeneral       int64   `json:"is_general"`                                      // 是否为会员
	QyType          int64   `json:"qy_type"`                                         // 契约等级
	IsResetpwd      bool    `json:"isresetpwd"`                                      // 是否重置过密码
	OpTime          int64   `json:"optime"`                                          // 注册时间
	LimitTime       int64   `json:"limittime"`                                       // 限制登录时间
	LastLogintime   int64   `json:"lastlogintime"`                                   // 最新登录时间
	LastOrdertime   int64   `json:"lastordertime"`                                   // 最后下注时间
	Validamount     float64 `json:"validamount"`                                     // 有效金额
	Paicai          float64 `json:"paicai"`                                          // 派彩金额
	Tuishui         float64 `json:"tuishui"`                                         // 退水金额
	MoneyType       string  `json:"moneytype"`                                       // 资金类型
	VipLevel        int64   `json:"viplevel"`                                        // 用户会员等级
	WhiteIps        string  `gorm:"type:varchar(300)" json:"whiteips"`               // 白名单列表
	NaviidList      string  `gorm:"type:varchar(300)" json:"naviidlist"`             // 权限列表
	UpTime          int64   `json:"uptime"`                                          // 升级时间
	AgentType       string  `gorm:"type:varchar(40)" json:"agenttype"`               // 代理类型
	WxExpinfo       string  `gorm:"type:varchar(300)" json:"wxexpinfo"`              // 微信扩展信息
	ErrPwdCount     int64   `json:"errpwdcount"`                                     // 错误密码次数
	FengkongCount   int64   `json:"fengkongcount"`                                   // 标记次数
	Expinfo         []byte  `gorm:"type:text"  json:"expinfo"`                       // 扩展信息
}

// 会员管理信息
type UserManager struct {
	Uuid          int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"uuid"` // 分配ID
	Account       string  `gorm:"type:varchar(40)" json:"account"`                 // 用户名登录
	Cash          float64 // 现金
	CashLimit     float64 // 现金限额
	PercentInfo   string  `json:"percentinfo"`   // 占成信息
	RegTime       int64   `json:"regtime"`       // 注册时间
	LastLogintime int64   `json:"lastlogintime"` // 注册时间
	PreAccount    string  `json:"pre_account"`   // 上级账号
	Odds          float64 `json:"odds"`          // 彩票赔率
	IsGeneral     int64   `json:"isgeneral"`     // 代理/会员
	ShuiLottery   float64 `json:"shui_lottery"`  // 彩票返点
	ShuiActual    float64 `json:"shui_actual"`   // 真人返点
	ShuiSport     float64 `json:"shui_sport"`    // 体育返点
	ShuiSlots     float64 `json:"shui_slots"`    // 电子返点
	ShuiCard      float64 `json:"shui_card"`     // 棋牌返点
	Tag           string  `json:"tag"`           // 用户标签
	HaveSuf       int64   `json:"havesuf"`       // 是否拥有下级
	Online        bool    `json:"online"`        // 在线标志位
}

type UserExpinfo struct {
	Password string
	//	Realname   string
	Nickname   string
	Birthday   string
	Addr       string
	Bankname   string
	Bankaddr   string // 支行地址
	Bankuser   string
	Banknum    string
	Dealpwd    string
	Alipayuser string
	Alipaynum  string
	Info       string // 备注信息
}

//用户资金
type UserMoney struct {
	Uuid                int64   `gorm:"primary_key;not null"` // 分配ID
	Withdraw            float64 // 可提现金额
	Lastgame            string  // 用户上次所玩游戏
	Cash                float64 // 现金
	Gaopin              float64 // 高频钱包
	GaopinLimit         float64 // 高频钱包限额
	Tcp3                float64 // 体彩P3
	Tcp3Limit           float64 // 体彩P3
	Fc3d                float64 // 福彩3D
	Fc3dLimit           float64 // 福彩3D
	Taiwanliuhe         float64 // 台湾六合
	TaiwanliuheLimit    float64 // 台湾六合
	Taiwandaletou       float64 // 台湾大乐透
	TaiwandaletouLimit  float64 // 台湾大乐透
	Xianggangliuhe      float64 // 香港六合
	XianggangliuheLimit float64 // 香港六合
	Aomenliuhe          float64 // 澳门六合
	AomenliuheLimit     float64 // 澳门六合
	Ogactual            float64 // og视讯
	Mgactual            float64 // mg视讯
	Bbinlive            float64 // bbin视讯
	Bbinslots           float64 // bbin电子
	Sabasport           float64 // bbin电子
	Cmdsport            float64 // cmd体育
	AG                  float64 // AG
	GG                  float64 // GG
	// Credit_Liuhe   float64 // 信用现金-六合
}

// 分成的详细信息
type PercentInfo struct {
	UpPercent float64 `json:"up"` // 上级占本条线的成数
	SfPercent float64 `json:"sf"` // 上级给本级的占成
}

type PercentInfo_Type struct {
	HighRate PercentInfo `json:"gp"` // 高频占成
	LhRate   PercentInfo `json:"lh"` // 六合占成
	LowRate  PercentInfo `json:"dp"` // 低频占成
}

// 用户上下级关系表
type BranchUsers struct {
	Uuid      int64  `gorm:"primary_key;not null"` // 用户id
	MasterID  int64  // 公司ID
	PreID     int64  // 上级代理ID
	PreList   []byte `gorm:"type:text"` // 上级列表
	SalerID   int64  // 上级销售ID
	SalerList []byte `gorm:"type:text"` // 销售列表
	RoleType  string `json:"roletype"`
}

// 盘口修改日志
type GamePortLog struct {
	ID           int64  `gorm:"primary_key;not null;AUTO_INCREMENT"`
	Uuid         int64  `gorm:"not null"` // 用户ID
	Lottery      string `gorm:"type:varchar(40)"`
	Time         int64  // 修改时间
	NewPortinfoS string `gorm:"type:text"` // 新的盘口设置
}

// 定制操作记录
type UserOpLog struct {
	ID           int64  `gorm:"primary_key;not null;AUTO_INCREMENT"`
	Uuid         int64  `gorm:"not null"` // 用户ID
	Account      string // 操作者账户
	RoleType     string // 操作者账户
	Nickname     string // 昵称
	Realname     string // 真实姓名
	DestUuid     int64  // 被操作者
	DestAccount  string // 被操作者账户
	DestNickname string // 昵称
	DestRealname string // 真实姓名
	Time         int64  `gorm:"not null"` // 修改时间
	Mode         string
	Yuanshizhi   string
	Biangengzhi  string
	OpType       UserOpType
	IP           string
	IPPlace      string
	Url          string
	Desc         string `gorm:"type:text"` // 新的盘口设置
}

type MoneyUpdateLog struct { // 资金变动日志
	ID         int64           `gorm:"primary_key;not null;AUTO_INCREMENT"`
	Uuid       int64           `gorm:"not null"`   // 用户ID
	Time       int64           `gorm:"not null"`   // 修改时间
	WalletName string          `json:"walletname"` // 钱包名称
	OpType     MoneyUpdateType // 操作类型
	OpGold     float64         // 操作金额
	OldGold    float64         // 操作前资金
	NewGold    float64         // 操作后资金
	Expinfo    string          // 用户操作备注
	Opinfo     string          // 其他备注信息
}

// 活动领取记录
type PromotionLog struct {
	ID            int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"`
	Uuid          int64   `json:"uuid"`          // 用户ID
	Account       string  `json:"account"`       // 账户
	PreID         int64   `json:"preid"`         //
	MasterID      int64   `json:"masterid"`      //
	Optime        int64   `json:"optime"`        // 修改时间
	PromotionID   int64   `json:"promotionid"`   // 活动ID
	PromotionName string  `json:"promotionname"` // 活动名
	Amount        float64 `json:"amount"`        // 奖励金额
	Desc          string  `json:"desc"`          // 说明
}

// 警示记录
type WarningLog struct {
	ID           int64   `json:"id"`
	Uuid         int64   `json:"uuid"`         // 用户ID
	Account      string  `json:"account"`      // 账户
	RoomID       int64   `json:"roomid"`       //
	Expect       string  `json:"expect"`       //
	ItemID       int64   `json:"itemid"`       //
	Iteminfo     string  `json:"iteminfo"`     //
	WarningCount int64   `json:"warningcount"` //
	Amount       float64 `json:"amount"`       // 修改时间
	WarningTime  int64   `json:"warningtime"`  // 活动ID
}

type SettleWait struct { // 等待结算表
	ID            int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"ID"`
	RoomID        int64  `json:"roomid"`                            // 房间ID
	Expect        string `json:"expect"`                            // 期号
	Wishtime      int64  `json:"wishtime"`                          // 预计开奖时间
	Code          string `gorm:"type:varchar(100)" json:"code"`     // 开奖号码
	Arrivetime    int64  `json:"arrivetime"`                        // 实际获取时间
	Issettled     int64  `gorm:"DEFAULT 0" json:"issettled"`        // 正式服结算标志
	SettledTime   int64  `json:"settledtime"`                       // 正式服结算时间
	Isretcash     int64  `gorm:"DEFAULT 0" json:"isretcash"`        // 正式服返现标志
	RetcashTime   int64  `json:"retcashtime"`                       // 正式服返现时间
	Isstatistic   int64  `gorm:"DEFAULT 0"`                         // 正式服结算标志
	StatisticTime int64  `json:"statistictime"`                     // 正式服结算时间
	Expinfo       string `gorm:"type:varchar(1000)" json:"expinfo"` // 日志记录
}

type PromotionWait struct { // 等待结算表
	ID            int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"`
	CompanyID     int64  `json:"companyid"`                      // 发起结算ID
	PromotionID   int64  `json:"promotionid"`                    //
	PromotionName string `json:"promotionname"`                  //
	BeginTime     int64  `json:"begintime"`                      //
	OpTime        int64  `json:"optime"`                         //
	OpState       int64  `json:"opstate"`                        //
	Desc          string `gorm:"type:varchar(1000)" json:"desc"` // 日志记录
}

type SetWait struct { // 等待生成号码表
	ID          int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"ID"`
	RoomID      int64   `json:"roomid"`      // 房间ID
	RoomName    string  `json:"roomname"`    // 房间名
	Expect      string  `json:"expect"`      // 期号
	WishTime    int64   `json:"wishtime"`    // 预期开奖时间
	Opencode    string  `json:"opencode"`    // 开奖号码
	SetTime     int64   `json:"settime"`     // 开奖时间
	Wager       float64 `json:"wager"`       // 下注额
	ProfitWager float64 `json:"profitwager"` // 派彩金额
	Profit      float64 `json:"profit"`      // 盈利
	Expinfo     string  `json:"expinfo"`     // 日志记录
}

type RetCashWait struct { // 等待返现表
	ID       int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"ID"`
	MasterID int64  `json:"masterid"`                          // 站长ID
	Date     int64  `json:"date"`                              // 日期
	Platform string `json:"platform"`                          // 平台
	Gametype string `json:"gametype"`                          // 游戏
	Optime   int64  `json:"optime"`                            // 返佣时间
	Result   int64  `json:"result"`                            // 返佣标志
	Expinfo  string `gorm:"type:varchar(1000)" json:"expinfo"` // 日志记录
}

type SettleLog struct { // 结算日志
	ID        int64  `gorm:"primary_key;not null;AUTO_INCREMENT"`
	ResID     int64  `json:"resid"`     // 站长ID
	BeginTime int64  `json:"begintime"` // 进入时间
	EndTime   int64  `json:"endtime"`   // 完成时间
	RoomID    int64  `json:"roomid"`    // 房间ID
	Expect    string `json:"expect"`    // 期号
	ExpInfo   string `json:"expinfo"`   // 备注信息
}

type RetCashLog struct { // 返现日志
	ID        int64  `gorm:"primary_key;not null;AUTO_INCREMENT"`
	ResID     int64  // 发起返现ID
	MasterID  int64  // 站长ID
	BeginTime int64  // 进入时间
	EndTime   int64  // 完成时间
	RoomID    int64  // 房间ID
	Expect    string // 期号
	RetType   int    // 是否返现  (返现，无效退款，打和退款)
	ExpInfo   string // 备注信息
}

type LotteryOpentime struct { // 返现日志
	ID            int64  `gorm:"primary_key;not null;AUTO_INCREMENT"`
	Expect        string // 发起返现ID
	Opentimestamp int64
}

type OpentimeS struct {
	Id            int64  `gorm:"primary_key;not null;AUTO_INCREMENT"`
	Expect        string `gorm:"type:varchar(40);not null"`
	Opentimestamp int64
	IsShoudong    int64 // 手动设置标识
}

type LostWininfo struct { // 丢失的开奖信息
	ID         int64  `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"`
	RoomENG    string // 房间名
	Expect     string // 期号
	Opentime   int64  // 预计开奖时间
	RepairTime int64  // 修复时间
	RepairWay  string // 修复方式
}

type WYManager struct { // 文远管理员信息
	ID          int64  `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"`
	Account     string `json:"account"`     // 账号
	Password    string `json:"password"`    // 密码
	NickName    string `json:"nickname"`    // 昵称
	Phone       string `json:"phone"`       // 电话
	RoleType    int64  `json:"roletype"`    // 角色类型
	RegTime     int64  `json:"regtime"`     // 注册时间
	LoginTime   int64  `json:"logintime"`   // 上次登录时间
	Companylist string `json:"companylist"` // 可监控公司列表
	Roomlist    string `json:"roomlist"`    // 可查房间列表
}

type Closevalue struct { // 代理的提前关闭值
	Uuid   int64 `gorm:"not null" json:"uuid"`
	Roomid int64 `json:"roomid"`
	Value  int64 `json:"value"`
}

type Wininfo_actual struct { // 百家的开奖号码
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

type WininfoFFC struct { // 分分彩开奖号码
	RoomID   int64  `json:"roomid"`   // 房间号
	Expect   string `json:"expect"`   // 期号
	ResID    int64  `json:"resid"`    // 庄家ID
	Opencode string `json:"opencode"` // 中奖Item
	Openitem string `json:"openitem"`
	Opentime int64  `json:"opentime"` // 开奖时间
}

// 公司的游戏表
type CompanyGame struct {
	CompanyID       int64  `gorm:"not null" json:"companyid"` // 公司ID
	Port            string `gorm:"type:varchar(3);not null" json:"port"`
	RoomID          int64  `json:"roomid"`
	Platform        string `json:"platform"`
	OpTime          int64  `json:"optime"`
	ValidTime       int64  `json:"validtime"`
	InUse           int64  `json:"inuse"`
	PreKaiguan      int64  `json:"prekaiguan"`
	SortID          int64  `json:"sortid"`
	Openvalue       int64  `json:"openvalue"`       // 六合的开盘时间
	Closevalue      int64  `json:"closevalue"`      // 公司自行增加的提前封盘量
	TemaClosevalue  int64  `json:"temaclosevalue"`  // 六合的特码封盘时间
	JsClosetime     int64  `json:"jsclosetime"`     // 加时封盘
	JsTemaClosetime int64  `json:"jstemaclosetime"` // 加时特码封盘
	IsHot           int64  `json:"ishot"`           // 热门标识
	IsChedan        int64  `json:"ischedan"`        // 允许撤单标识
}

// 公司的port玩法赔率表
type CompanyPortinfo struct {
	CompanyID         int64   `json:"companyid"` // 公司ID
	RoomID            int64   `json:"roomid"`
	PortID            int64   `json:"portid"`
	OddsA             float64 `json:"oddsa"`             // 赔率A
	TuishuiA          float64 `json:"tuishuia"`          // 退水A
	OddsB             float64 `json:"oddsb"`             // 赔率B
	TuishuiB          float64 `json:"tuishuib"`          // 退水B
	OddsC             float64 `json:"oddsc"`             // 赔率C
	TuishuiC          float64 `json:"tuishuic"`          // 退水C
	OddsD             float64 `json:"oddsd"`             // 赔率D
	TuishuiD          float64 `json:"tuishuid"`          // 退水D
	DefaultOdds       float64 `json:"defaultodds"`       // 最大赔率
	MinAmount         float64 `json:"minamount"`         // 单注最低金额
	MaxAmount         float64 `json:"maxamount"`         // 单注最大金额
	MaxAmountExpect   float64 `json:"expectamount"`      // 当期最大金额
	TransferAmount    int64   `json:"transferamount"`    // 走飞限额
	IsTransfer        int64   `json:"istransfer"`        // 走飞启用
	WarningAmount     int64   `json:"warningamount"`     // 警示金额
	WarningLoopamount int64   `json:"warningloopamount"` // 警示循环金额
	PortSwitch        int64   `json:"portswitch"`        // 玩法下注开关
	Kaiguan           int64   `json:"kaiguan"`           // 玩法开关
	// UserOdds        float64 `json:"userodds"`
}

// 公司的加时port玩法赔率表
type CompanyJiashiPortinfo struct {
	CompanyID       int64   `json:"companyid"` // 公司ID
	RoomID          int64   `json:"roomid"`
	PortID          int64   `json:"portid"`
	OddsDesA        float64 `json:"oddsdesa"`     // 赔率A
	OddsDesB        float64 `json:"oddsdesb"`     // 赔率B
	OddsDesC        float64 `json:"oddsdesc"`     // 赔率C
	OddsDesD        float64 `json:"oddsdesd"`     // 赔率D
	MinAmount       float64 `json:"minamount"`    // 单注最低金额
	MaxAmount       float64 `json:"maxamount"`    // 单注最大金额
	MaxAmountExpect float64 `json:"expectamount"` // 当期最大金额
}

// 会员的游戏设置
type UserGameset struct {
	Uuid                int64   `json:"uuid"`                // 用户ID
	Account             string  `json:"account"`             //
	Platform            string  `json:"platform"`            //
	RoomID              int64   `json:"roomid"`              // 房间ID
	InUse               int64   `json:"inuse"`               // 上级设置的游戏启用标识
	Zhancheng           float64 `json:"zhancheng"`           // 上级想在我这里的占成
	ZhanchengNext       float64 `json:"zhanchengnext"`       // 上级想在我这里的占成(下次维护值)
	SufMinZhancheng     float64 `json:"sufminzhancheng"`     // 上级给我的最低占成
	SufMinZhanchengNext float64 `json:"sufminzhanchengnext"` // 上级给我的最低占成(下次维护值)
	SufMaxZhancheng     float64 `json:"sufmaxzhancheng"`     // 上级能给我的最高占成
	SufMaxZhanchengNext float64 `json:"sufmaxzhanchengnext"` // 上级能给我的最高占成(下次维护值)
	Pans                string  `json:"pans"`                // 开通盘口
	BuhuoFlag           int64   `json:"buhuoflag"`           // 补货启用标识
}

// 会员的玩法设置
type UserPortclassset struct {
	Uuid              int64   `json:"uuid"`              // 会员ID
	RoomID            int64   `json:"roomid"`            // 房间ID
	PortID            int64   `json:"portid"`            // 玩法ID
	OddsA             float64 `json:"oddsa"`             // 赔率A
	TuishuiA          float64 `json:"tuishuia"`          // 退水A
	OddsB             float64 `json:"oddsb"`             // 赔率B
	TuishuiB          float64 `json:"tuishuib"`          // 退水B
	OddsC             float64 `json:"oddsc"`             // 赔率C
	TuishuiC          float64 `json:"tuishuic"`          // 退水C
	OddsD             float64 `json:"oddsd"`             // 赔率D
	TuishuiD          float64 `json:"tuishuid"`          // 退水D
	MinAmount         float64 `json:"minamount"`         // 单注最低
	MaxAmount         float64 `json:"maxamount"`         // 单注最高
	MaxAmountExpect   float64 `json:"maxamountexpect"`   // 单期最高
	TransferAmount    int64   `json:"transferamount"`    // 走飞限额
	IsTransfer        int64   `json:"istransfer"`        // 走飞启用
	WarningAmount     int64   `json:"warningamount"`     // 警示金额
	WarningLoopamount int64   `json:"warningloopamount"` // 警示循环金额
}

type BaijiaOdds struct { // 公司基准赔率
	ID   int64   `gorm:"not null" json:"id"`
	Odds float64 `json:"odds"`
}

type BzbmOdds struct { // 奔驰宝马公司基准赔率(飞禽走兽也用此结构)
	ID             int64   `gorm:"not null" json:"id"`
	Odds           float64 `json:"odds"`
	MinProbability int64   `json:"minprobability"`
	MaxProbability int64   `json:"maxprobability"`
}

type WFLMOdds struct { // 五福临门赔率
	ID    int64   `json:"id"`    // 牌ID
	Three float64 `json:"three"` // 中三赔率
	Four  float64 `json:"Four"`  // 中四赔率
	Five  float64 `json:"Five"`  // 中五赔率
}

type AgentOdds struct { // 代理的水赔
	Uuid     int64   `gorm:"not null" json:"uuid"`      // 代理ID
	Preid    int64   `gorm:"default 0" json:"preid"`    // 上级代理ID
	Masterid int64   `gorm:"default 0" json:"masterid"` // 站长ID
	Port     string  `json:"port"`                      // 盘口
	InUse    int     `json:"inuse"`                     // 是否使用中
	OpTime   int64   `json:"optime"`                    // 操作时间
	PreOdds  float64 `json:"preodds"`                   // 基准赔率
	// SufMinodds    float64 `json:"sufminodds"`                // 下线的最低赔率
	// RangeLottery float64 `json:"range_lottery"` // 彩票返点
	// RangeActual   float64 `json:"range_actual"`              // 真人返点
	// RangeElectric float64 `json:"range_electric"`            // 电子返点
	// RangeCard     float64 `json:"range_card"`                // 棋牌返点
	// RangeSport    float64 `json:"range_sport"`               // 体育返点
	// OddsLottery   float64 `json:"shui_lottery"`              // 彩票返点
	// ShuiActual    float64 `json:"shui_actual"`               // 真人返点
	// ShuiElectric  float64 `json:"shui_electric"`             // 电子返点
	// ShuiCard      float64 `json:"shui_card"`                 // 棋牌返点
	// ShuiSport     float64 `json:"shui_sport"`                // 体育返点
}

type DefaultOdds struct { // 系统的默认赔率信息
	RoomID      int64   `json:"roomid"`
	PortID      int64   `json:"portid"`
	DefaultOdds float64 `json:"defaultodds"`
	// UserOdds        float64 `json:"userodds"`
	// MinAmount       float64 `json:"minamount"`
	// MaxAmount       float64 `json:"maxamount"`
	// MaxAmountExpect float64 `json:"expectamount"`
	// DesAmount       float64 `json:"desamount"`      // 种类降赔金额
	// DesOdds         float64 `json:"desodds"`        // 种类降赔赔率
	// DesMinodds      float64 `json:"desminodds"`     // 种类最低赔率
	// TransferAmount  float64 `json:"transferamount"` // 走飞限额
	// WarningAmount   float64 `json:"warningamount"`
	// PortSwitch      int64   `json:"portswitch"` // 种类开关
}

type DefaultOddsJson struct { // json格式 系统默认赔率
	SettleType  string `json:"settletype"`                   // 结算类型
	MaxoddsInfo string `gorm:"type:text" json:"maxoddsinfo"` // 赔率信息
}

type UserzoufeiSet struct {
	Uuid   int64   `json:"uuid"`   //
	RoomID int64   `json:"roomid"` //
	PortID int64   `json:"portid"` // 盘口
	InUse  int64   `json:"inuse"`  // 是否使用中
	Value  float64 `json:"value"`  //
}

type DayProfit struct { // 日盈亏报表
	ExpGameID       int64   `gorm:"not null" json:"expgameid"` // 游戏分类
	Uuid            int64   `gorm:"not null"`                  // 用户ID
	Date            int64   `gorm:"not null"`                  // 日期
	Wager           float64 `json:"wager"`                     // 赌金
	OrderNum        int64   `json:"ordernum"`                  // 下注数
	SettledWager    float64 `json:"settledwager"`              // 结算数量
	SettledNum      int64   `json:"settlednum"`                // 结算注数
	ValidAmount     float64 `json:"validamount"`               // 有效金额
	ProfitWager     float64 `json:"profit_wager"`              // 赌金利润
	ProfitWagerDiff float64 `json:"profit_wagerdiff"`          // 赔差
	ProfitShui      float64 `json:"profit_shui"`               // 下注退水
	ProfitShuiDiff  float64 `json:"profit_shuidiff"`           // 水差
	ProfitTotal     float64 `json:"profit_total"`              // 总盈利
}

// 彩票统计数据(日+游戏级)
type LotdataStatistic struct {
	Uuid                       int64   `gorm:"not null"`                   // 分配ID
	PreID                      int64   `json:"preid"`                      //
	MasterID                   int64   `json:"masterid"`                   //
	RoleType                   string  `json:"roletype"`                   //
	Date                       int64   `gorm:"not null"`                   // 日期
	RoomID                     int64   `gorm:"not null"`                   // 房间号
	RoomENG                    string  `json:"roomeng"`                    // 房间号
	RoomCN                     string  `json:"roomcn"`                     // 房间号
	OrderNum                   int64   `json:"ordernum"`                   // 下注数
	OrderAmount                float64 `json:"orderamount"`                // 赌金
	SettledNum                 int64   `json:"settlednum"`                 // 结算注数
	SettledAmount              float64 `json:"settledamount"`              // 结算数量
	ValidAmount                float64 `json:"validamount"`                // 有效金额
	RevokeAmount               float64 `json:"revokeamount"`               // 撤单金额
	Wager                      float64 `json:"profit_wager"`               // 派彩金额
	Tuishui                    float64 `json:"profit_shui"`                // 退水金额
	Yingshouxiaxian            float64 `json:"yingshouxiaxian"`            // 应收下线
	Shizhanhuoliang            float64 `json:"shizhanhuoliang"`            // 实占货量
	Shizhanshuying             float64 `json:"shizhanshuying"`             // 实占输赢
	Shizhanjieguo              float64 `json:"shizhanjieguo"`              // 实占结果
	Shizhantuishui             float64 `json:"shizhantuishui"`             //
	ProfitTuishui              float64 `json:"profittuishui"`              // 水差
	Shizhanpeicha              float64 `json:"shizhanpeicha"`              // 实占赔差
	ProfitWager                float64 `json:"profitwager"`                // 赔差
	Yingkuijieguo              float64 `json:"yingkuijieguo"`              // 代理的盈亏结果
	Shangjiaohuoliang          float64 `json:"shangjiaohuoliang"`          // 上交货量
	Shangjijiaoshou            float64 `json:"shangjijiaoshou"`            // 上级交收
	Shouhuo                    float64 `json:"shouhuo"`                    // 收货
	Shouhuoyingkui             float64 `json:"shouhuoyingkui"`             // 收货盈亏
	Shoudongbuchu              float64 `json:"shoudongbuchu"`              // 手动补出
	Buchu                      float64 `json:"buchu"`                      // 补出
	Buchuyingkui               float64 `json:"buchuyingkui"`               // 补出盈亏
	SettleShoudongbuchu        float64 `json:"settleshoudongbuchu"`        // 结算后的手动补货
	SettleShoudongbuchuyingkui float64 `json:"settleshoudongbuchuyingkui"` // 结算后的手动补货盈亏
	Shizhanbuhuo               float64 `json:"shizhanbuhuo"`               // 实占补货
	Shizhanbuhuoyingkui        float64 `json:"shizhanbuhuoyingkui"`        // 实占补货盈亏
}

// 彩票统计数据(日+游戏级)
type LotdataUseritem struct {
	Uuid                       int64   `gorm:"not null"`                   // 分配ID
	PreID                      int64   `json:"preid"`                      //
	MasterID                   int64   `json:"masterid"`                   //
	RoleType                   string  `json:"roletype"`                   //
	Date                       int64   `gorm:"not null"`                   // 日期
	RoomID                     int64   `gorm:"not null"`                   // 房间号
	RoomENG                    string  `json:"roomeng"`                    // 房间号
	RoomCN                     string  `json:"roomcn"`                     // 房间号
	Expect                     string  `json:"expect"`                     // 期号
	Pan                        string  `json:"pan"`                        // 盘口
	PortID                     int64   `json:"portid"`                     // portid
	ItemID                     int64   `json:"itemid"`                     // itemid
	OrderNum                   int64   `json:"ordernum"`                   // 下注数
	OrderAmount                float64 `json:"orderamount"`                // 赌金
	Zhanchenghuoliang          float64 `json:"zhanchenghuoliang"`          // 占成货量
	Shizhanhuoliang            float64 `json:"shizhanhuoliang"`            // 实占货量
	Xiajibuhuo                 float64 `json:"xiajibuhuo"`                 // 下级补货
	Zidongbuchu                float64 `json:"zidongbuchu"`                // 自动补出
	Shoudongbuchu              float64 `json:"shoudongbuchu"`              // 手动补出
	SettledNum                 int64   `json:"settlednum"`                 // 结算注数
	SettledAmount              float64 `json:"settledamount"`              // 结算数量
	ValidAmount                float64 `json:"validamount"`                // 有效金额
	RevokeAmount               float64 `json:"revokeamount"`               // 撤单金额
	Wager                      float64 `json:"profit_wager"`               // 派彩金额
	Tuishui                    float64 `json:"profit_shui"`                // 退水金额
	Yingshouxiaxian            float64 `json:"yingshouxiaxian"`            // 应收下线
	SettleShizhanhuoliang      float64 `json:"settleshizhanhuoliang"`      // 结算实占货量
	Shizhanshuying             float64 `json:"shizhanshuying"`             // 实占输赢
	Shizhanjieguo              float64 `json:"shizhanjieguo"`              // 实占结果
	Shizhantuishui             float64 `json:"shizhantuishui"`             // 实占退水
	ProfitWager                float64 `json:"profit_wagerdiff"`           // 赔差
	ProfitTuishui              float64 `json:"profit_shuidiff"`            // 水差
	Shizhanpeicha              float64 `json:"shizhanpeicha"`              // 实占赔差
	Yingkuijieguo              float64 `json:"yingkuijieguo"`              // 代理的盈亏结果
	Shangjiaohuoliang          float64 `json:"shangjiaohuoliang"`          // 上交货量
	Shangjijiaoshou            float64 `json:"shangjijiaoshou"`            // 上级交收
	Shouhuo                    float64 `json:"shouhuo"`                    // 收货
	Shouhuoyingkui             float64 `json:"shouhuoyingkui"`             // 收货盈亏
	Buchu                      float64 `json:"buchu"`                      // 补出
	Buchuyingkui               float64 `json:"buchuyingkui"`               // 补出盈亏
	SettleShoudongbuchu        float64 `json:"settleshoudongbuchu"`        // 结算后的手动补货
	SettleShoudongbuchuyingkui float64 `json:"settleshoudongbuchuyingkui"` // 结算后的手动补货盈亏
	Shizhanbuhuo               float64 `json:"shizhanbuhuo"`               // 实占补货
	Shizhanbuhuoyingkui        float64 `json:"shizhanbuhuoyingkui"`        // 实占补货盈亏
}

// 用户特码统计数据(合并特码A、B)
type LotdataUsertema struct {
	Uuid        int64   `gorm:"not null"`    // 分配ID
	Account     string  `json:"account"`     //
	PreID       int64   `json:"preid"`       //
	MasterID    int64   `json:"masterid"`    //
	RoleType    string  `json:"roletype"`    //
	Date        int64   `gorm:"not null"`    // 日期
	RoomID      int64   `gorm:"not null"`    // 房间号
	RoomENG     string  `json:"roomeng"`     // 房间号
	RoomCN      string  `json:"roomcn"`      // 房间号
	Expect      string  `json:"expect"`      // 期号
	PortID      int64   `json:"portid"`      // portid
	ItemID      int64   `json:"itemid"`      // itemid
	OrderNum    int64   `json:"ordernum"`    // 下注数
	OrderAmount float64 `json:"orderamount"` // 下注金额
}

// 单期统计数据
type UserItemstatistic struct {
	Uuid                 int64   // 用户ID
	Pan                  string  // 盘口
	Expect               string  // 期号
	PortID               int64   //
	ItemID               int64   //
	Settlenum            int64   // 结算单数
	Settledamount        float64 // 结算金额
	Validamount          float64 // 有效金额
	Wager                float64 // 派彩金额
	Shuiamount           float64 // 退水金额
	Yingshouxiaxian      float64 // 应收下线
	Shizhanhuoliang      float64 // 实占货量
	Shizhanshuying       float64 // 实占输赢
	Shizhanjieguo        float64 // 实占结果
	Shizhantuishui       float64 // 实占退水
	Shizhanpeicha        float64 // 实占赔差
	ProfitWager          float64 // 赚取赔差
	Porfittuishui        float64 // 赚取退水
	Yingkuijieguo        float64 // 代理的盈亏结果
	Shangjiaohuoliang    float64 // 上交货量
	Shangjijiaoshou      float64 // 上级交收
	Shouhuo              float64 // 收货金额
	Shouhuoyingkui       float64 // 收货盈亏
	Buchu                float64 // 补出金额
	Buchuyingkui         float64 // 补出盈亏
	Shoudongbuchu        float64 // 手动补出
	Shoudongbuchuyingkui float64 // 手动补出盈亏
	Shizhanbuhuo         float64 // 实占补货
	Shizhanbuhuoyingkui  float64 // 实占补货盈亏
}

// 未结算订单需要回清的数据
type UnsettleRevokedata struct {
	Uuid              int64   // 用户ID
	Pan               string  // 盘口
	Expect            string  // 期号
	PortID            int64   //
	ItemID            int64   //
	OrderNum          int64   //
	OrderAmount       float64 //
	Zhanchenghuoliang float64 //
	Shizhanhuoliang   float64 //
	Xiajibuhuo        float64 //
	Zidongbuchu       float64 //
	Shoudongbuchu     float64 //
}

type GamemodeStats struct { // 玩法统计
	Uuid            int64   `gorm:"not null"`         // 分配ID
	PreID           int64   `json:"preid"`            //
	MasterID        int64   `json:"masterid"`         //
	Date            int64   `gorm:"not null"`         // 日期
	RoomID          int64   `gorm:"not null"`         // 房间号
	Expect          string  `gorm:"type:varchar(40)"` // 期号
	ModeID          int64   `gorm:"not null"`         // 玩法ID
	Wager           float64 `json:"wager"`            // 下注金额
	OrderNum        int64   `json:"ordernum"`         // 下注数
	SettledWager    float64 `json:"settledwager"`     // 结算数量
	SettledNum      int64   `json:"settlednum"`       // 结算注数
	ValidAmount     float64 `json:"validamount"`      // 有效金额
	ProfitWager     float64 `json:"profit_wager"`     // 赌金利润
	ProfitWagerDiff float64 `json:"profit_wagerdiff"` // 赔差
	ProfitShui      float64 `json:"profit_shui"`      // 下注退水
	ProfitShuiDiff  float64 `json:"profit_shuidiff"`  // 水差
	ProfitTotal     float64 `json:"profit_total"`     // 总盈利
}

type ExpYingkuiDay struct { // 日盈亏报表
	Uuid            int64   `gorm:"not null"`         // 分配ID
	Date            int64   `gorm:"not null"`         // 日期
	Wager           float64 `json:"wager"`            // 赌金
	OrderNum        int64   `json:"ordernum"`         // 下注数
	SettledWager    float64 `json:"settledwager"`     // 结算数量
	SettledNum      int64   `json:"settlednum"`       // 结算注数
	ValidAmount     float64 `json:"validamount"`      // 有效金额
	ProfitWager     float64 `json:"profit_wager"`     // 赌金利润
	ProfitWagerDiff float64 `json:"profit_wagerdiff"` // 赔差
	ProfitShui      float64 `json:"profit_shui"`      // 下注退水
	ProfitShuiDiff  float64 `json:"profit_shuidiff"`  // 水差
	ProfitTotal     float64 `json:"profit_total"`     // 总盈利
}

type ExpGame struct { // 外接游戏里诶包
	ID         int64  `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"` // 订单ID
	GameName   string `json:"gamename"`                                      // 游戏名
	InUse      string `json:"inuse"`                                         // 启用状态
	NameCN     string `json:"name_cn"`                                       // 钱包名
	GamenameCN string `json:"gamename_cn"`                                   // 游戏名
	Url        string `json:"url"`                                           // 跳转地址
}

type ExpLimit struct { // 外接游戏限额
	ID          int64   `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"` // 订单ID
	CompanyID   int64   `json:"companyid"`                                     // 公司ID
	ExpName     string  `gorm:"type:varchar(20)" json:"expname"`               // 外接游戏名
	LimitAmount float64 `json:"limitamount"`                                   // 限额
}

type UserTag struct { // 用户标签信息
	ResID  int64  `json:"resid"`  // 写标签人
	DestID int64  `json:"destid"` // 被标签人
	Tag    string `json:"tag"`    // 标签信息
}

type IPInfo struct { // IP信息
	ID   int64  `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"` // 订单ID
	IP   string `gorm:"type:varchar(40)" json:"ip"`                    // 公司ID
	Info string `gorm:"type:varchar(100)" json:"info"`                 // 外接游戏名
}

type ExpGameInfo struct { // IP信息
	ID       int64  `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id"` //
	GameName string `gorm:"type:varchar(40)" json:"gamename"`              //
	InUse    int64  `json:"inuse"`                                         //
}

// 用户扩展信息
type CUserbaseExp struct {
	Uuid        int64  `gorm:"primary_key;not null" json:"uuid"`      // 用户ID
	Mobilemodel string `gorm:"type:text;not null" json:"mobilemodel"` // 机型
}

// 查询用户下注统计地址
type OrderstatsUrl struct {
	ID  int64  `gorm:"primary_key;not null;AUTO_INCREMENT"` // 自动生成ID
	Url string // 查询地址
}

type OfOrderinfo struct {
	Multiple     int64           //官方玩法下注倍数
	MonetaryUnit float64         //官方玩法下注金额单位
	Ordernum     int64           //官方玩法下单注数
	Preodds      float64         //官方玩法下注项赔率
	Betlist      []int           //官方玩法下注号码解析
	Betmap       map[int64][]int //官方玩法下注号码解析
	Num          string          //官方玩法下注号码
}

// 占成信息
type CShare struct {
	Uuid     int64   `gorm:"primary_key;not null" json:"id"` //
	Preid    int64   `json:"preid"`                          //
	Masterid int64   `json:"masterid"`                       //
	Share    float64 `json:"minquota"`                       //
}

// 真人百家操作记录
type ACTDeallog struct {
	ID       int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Roomid   int64  // 房间号
	Time     int64  // 时间
	DealerID int64  // 荷官ID
	Optype   int64  // 操作类型
	Boot     int64  // 靴号
	Round    int64  // 局号
	Expect   string // 期号
	Desc     string // 操作备注
}

// 电子百家操作记录
type DZBJDeallog struct {
	ID     int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Roomid int64  // 房间号
	Time   int64  // 时间
	Optype int64  // 操作类型
	Boot   int64  // 靴号
	Round  int64  // 局号
	Expect string // 期号
	Desc   string // 操作备注
}

//  奔驰宝马操作记录
type BZBMDeallog struct {
	ID     int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Roomid int64  // 房间号
	Time   int64  // 时间
	Optype int64  // 操作类型
	Expect string // 期号
	Desc   string // 操作备注
}

//  奔驰宝马操作记录
type BRNNDeallog struct {
	ID     int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Roomid int64  // 房间号
	Time   int64  // 时间
	Optype int64  // 操作类型
	Expect string // 期号
	Desc   string // 操作备注
}

//  红黑大战操作记录
type HHDZDeallog struct {
	ID     int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Roomid int64  // 房间号
	Time   int64  // 时间
	Optype int64  // 操作类型
	Expect string // 期号
	Desc   string // 操作备注
}

//  龙虎大战操作记录
type LHDZDeallog struct {
	ID     int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Roomid int64  // 房间号
	Time   int64  // 时间
	Optype int64  // 操作类型
	Expect string // 期号
	Desc   string // 操作备注
}

//  飞禽走兽操作记录
type FQZSDeallog struct {
	ID     int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Roomid int64  // 房间号
	Time   int64  // 时间
	Optype int64  // 操作类型
	Expect string // 期号
	Desc   string // 操作备注
}

//  飞禽走兽操作记录
type WFLMDeallog struct {
	ID     int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Roomid int64  // 房间号
	Time   int64  // 时间
	Optype int64  // 操作类型
	Expect string // 期号
	Desc   string // 操作备注
}

// 重庆时时彩房间信息
type CCQSSCRoomInfo struct {
	ID           int64   `gorm:"primary_key;not null"` //
	NameCN       string  // 房间名字
	NameENG      string  // 房间名字 作为账号
	GameDalei    string  // 游戏大类
	TableSuf     string  // 表名后缀
	SettleType   string  // 结算类型
	GameName     string  // 挂靠游戏名
	Masterid     int64   // 公司
	CreateTime   int64   // 房间创建时间
	InValidTime  int64   // 房间过期时间
	EnterAmount  float64 // 进入要求
	MinAmount    float64 // 房间最低限红
	MaxAmount    float64 // 房间最高限红
	LimitNum     int64   // 限制人数
	Jackpot      float64 // 奖池
	Commission   float64 // 佣金提成
	Opencodetime int64   //开奖间隔
	Bantime      int64   //封盘时间
	Scrollinfo   string  `gorm:"type:text;not null"` // 滚筒设置

}

// 分分彩房间信息
type CFFCRoomInfo struct {
	ID             int64   `gorm:"primary_key;not null"` //
	NameCN         string  // 房间名字
	NameENG        string  // 房间名字 作为账号
	GameDalei      string  // 游戏大类
	TableSuf       string  // 表名后缀
	SettleType     string  // 结算类型
	GameName       string  // 挂靠游戏名
	Masterid       int64   // 公司
	CreateTime     int64   // 房间创建时间
	InValidTime    int64   // 房间过期时间
	EnterAmount    float64 // 进入要求
	MinAmount      float64 // 房间最低限红
	MaxAmount      float64 // 房间最高限红
	LimitNum       int64   // 限制人数
	Jackpot        float64 // 奖池
	Expect         int64   // 期号
	Opentime       int64   // 可下注时间
	Bantime        int64   // 封盘时间
	Flashtime      int64   //  动画时间
	Waittime       int64   // 新开局等待时间
	Commission     float64 // 佣金提成
	MinimumBalance float64 //用户下注所需的最低余额
}

// 系统默认活动设置
type CSysDefaultPromotion struct {
	ID            int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // ID
	PromotionName string `json:"promotionname"`                                 //
	PromotionType string `json:"promotiontype"`                                 //
	BeginTime     int64  `json:"begintime"`                                     //
	EndTime       int64  `json:"endtime"`                                       //
	SetInfo       string `json:"setinfo"`                                       //
}

// 公司活动设置
type CCoPromotion struct {
	ID            int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // ID
	CompanyID     int64  `json:"companyid"`                                     //
	PromotionID   int64  `json:"promotionid"`                                   //
	PromotionName string `json:"promotionname"`                                 //
	PromotionType string `json:"promotiontype"`                                 //
	InUse         int64  `json:"inuse"`                                         // 启用标志
	BeginTime     int64  `json:"begintime"`                                     //
	EndTime       int64  `json:"endtime"`                                       //
	SetInfo       string `json:"setinfo"`                                       //
}

// 注册送金
type CPromotionReg struct {
	Regamount float64 `json:"regamount"` // 注册送金
	Count     int64   `json:"count"`     // 可领取人数
}

// 登录送金
type CPromotionLogin struct {
	Loginamount float64 `json:"loginamount"` // 登录送金
	Count       int64   `json:"count"`       // 可领取人数
}

// 周业绩活动
type CPromotionWeekvalidamount struct {
	Validamount float64 `json:"validamount"` // 业绩
	Percentage  float64 `json:"percentage"`  // 提成比例(/万)
}

// 排线业绩
type CPromotionPaixianvalidamount struct {
	Paixian     []CPromotionWeekvalidamount `json:"paixian"` // 排线提成
	BaodiAmount float64                     `json:"baodi"`   // 保底提成
}

// 充值送金
type CPromotionRechargebenifit struct {
	Recharge float64 `json:"recharge"` // 充值
	Benifit  float64 `json:"benifit"`  // 优惠
}

// 推广码信息
type CExtInfo struct {
	ID     int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // ID
	Uuid   int64   `json:"uuid"`                                          // 用户ID
	LtOdds float64 `json:"lotteryodds"`                                   // 彩票赔率
	LtShui float64 `json:"lotteryshui"`                                   // 彩票退税
}

type ExpgameSetByThird struct {
	ID            int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 自动生成ID
	CompanyID     int64   `json:"companyid"`                                     // 公司ID
	ExpGameid     int64   `json:"expgameid"`                                     // 游戏ID
	ExpName       string  `gorm:"type:varchar(40)" json:"expname"`               // 游戏名
	InUse         int64   `json:"inuse"`                                         // 启用标识
	SortID        int64   `json:"sortid"`                                        // 排序ID
	LimitAmount   float64 `json:"limitamount"`                                   // 资金限额
	RequestUrl    string  `json:"requesturl"`                                    // 请求域名
	CagentAccount string  `json:"cagentaccount"`                                 // 代理账号
	Md5Key        string  `json:"md5key"`                                        // MD5 Key
}

type ExpgameSetForThird struct {
	ID             int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 自动生成ID
	CompanyID      int64   `json:"companyid"`                                     // 公司ID
	CompanyAccount string  `gorm:"type:varchar(40)" json:"coaccount"`             // 第三方公司账户
	InUse          int64   `json:"inuse"`                                         // 启用标识
	LimitAmount    float64 `json:"limitamount"`                                   // 资金限额
	PreStr         string  `gorm:"type:varchar(10)" json:"prestr"`                // 账号前缀
	Md5Key         string  `gorm:"type:varchar(10)" json:"md5key"`                // MD5验签
}

// UC彩票订单
type UCLtOrderInfo struct {
	ID         int64   `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 本地订单ID
	Uuid       int64   `json:"uuid"`                                          //
	PreID      int64   `json:"preid"`                                         //
	MasterID   int64   `json:"masterid"`                                      //
	Account    string  `json:"account"`                                       //
	CreateTime int64   `json:"createtime"`                                    //
	ExpOrderid string  `json:"exporderid"`                                    //
	GameType   string  `json:"gametype"`                                      //
	GameName   string  `json:"gamename"`                                      //
	Expect     string  `json:"expect"`                                        //
	Lotname    string  `json:"lotname"`                                       //
	Depiction  string  `json:"depiction"`                                     //
	State      int64   `json:"state"`                                         //
	Cash       float64 `json:"cash"`                                          //
	Rate       float64 `json:"rate"`                                          //
	Result     float64 `json:"result"`                                        //
	Desc       string  `gorm:"type:varchar(1000)" json:"desc"`                //
}

// 排线基础信息
type PaixianBase struct {
	Uuid         int64   `json:"uuid"`         // 用户ID
	Account      string  `json:"account"`      // 账户
	PreID        int64   `json:"preid"`        // 上级ID
	MasterID     int64   `json:"masterid"`     // 公司ID
	RegType      string  `json:"regtype"`      // 注册类型
	RegTime      int64   `json:"regtime"`      // 注册时间
	BaodiAmount  float64 `json:"baodiamount"`  // 保底金额
	PaixianCode  int64   `json:"paixiancode"`  // 排线码
	PaixianLevel int64   `json:"paixianlevel"` // 排线层级
	PaixianList  string  `json:"paixianlist"`  // 排线列表
}

// 排线设置
type PaixianSet struct {
	ID          int64  `gorm:"primary_key;not null;AUTO_INCREMENT" json:"id"` // 排线ID
	Uuid        int64  `json:"uuid"`                                          //
	Account     string `json:"account"`                                       //
	PreID       int64  `json:"preid"`                                         // 上级ID
	MasterID    int64  `json:"masterid"`                                      // 公司ID
	PaixianName string `json:"paixianname"`                                   // 排线名
	CreateTime  int64  `json:"createtime"`                                    // 排线创建时间
	NowNum      int64  `json:"nownum"`                                        // 当前人数
	LimitNum    int64  `json:"limitnum"`                                      // 限制人数
	Url         string `json:"url"`                                           // 排线推广地址
}

// 房间庄信息记录
type RoomBankerRecord struct {
	Companyid int64
	Roomid    int64
	Bankerid  int64
	Time      int64
}

type UserTeamstats struct { // 用户团队统计
	Uuid             int64   `json:"uuid"`             // 公司ID
	Date             int64   `json:"date"`             // 日期
	RegNum           int64   `json:"regnum"`           // 注册数
	HighestNum       int64   `json:"highestnum"`       // 最高在线
	OrderNum         int64   `json:"ordernum"`         // 下注人数
	WalletAmount     float64 `json:"walletamount"`     // 钱包余额(日末)
	UnsettleAmount   float64 `json:"unsettleamount"`   // 未结算金额(日末)
	UnwithdrawAmount float64 `json:"unwithdrawamount"` // 未提现金额(日末)
	RechargeAmount   float64 `json:"rechargeamount"`   // 系统充值
	MgrRemit         float64 `json:"mgrremit"`         // 后台确认
	MgrIncrease      float64 `json:"mgrincrease"`      // 管理员上分
	MgrDecrease      float64 `json:"mgrdecrease"`      // 管理员下分
	WithdrawAmount   float64 `json:"withdrawamount"`   // 提现总额
	SysRevoke        float64 `json:"sysrevoke"`        // 系统撤单
	SysBack          float64 `json:"sysback"`          // 系统退款(提现打回)
	OrderAmount      float64 `json:"orderamount"`      // 下注金额
	RevokeAmount     float64 `json:"revokeamount"`     // 用户撤单
	SettleAmount     float64 `json:"settleamount"`     // 有效金额
	WagerAmount      float64 `json:"wageramount"`      // 派彩金额
	GrossProfit      float64 `json:"grossprofit"`      // 毛利润
	ShuiAmount       float64 `json:"shuiamount"`       // 公司返佣
	PureProfit       float64 `json:"pureprofit"`       // 纯利润
}

type ZoufeiSet struct { // 用户的走飞设置
	Uuid   int64 `json:"uuid"`
	RoomID int64 `json:"roomid"`
	PortID int64 `json:"portid"`
	InUse  int64 `json:"inuse"`
	Value  int64 `json:"value"`
}

// 服务器设置
type ServerSet struct {
	SetKey string `json:"setkey"` //
	Value  string `json:"value"`  //
}

type Walletinfo struct { // 钱包信息
	ID            int64  `json:"id"`
	WalletNameCN  string `json:"walletnamecn"`
	WalletNameENG string `json:"walletnameeng"`
	WalletType    string `json:"wallettype"`
}

type KeyJumpurl struct { //
	CompanyAccount string `json:"companyaccount"`
	CompanyName    string `json:"companyname"`
	DaohangKey     string `json:"daohangkey"`
	DaohangUrl     string `json:"daohangurl"`
	XianluUrl      string `json:"xianluurl"`
}

type FengkongSet struct {
	CompanyID    int64   `json:"companyid"`
	RoomID       int64   `json:"roomid"`
	QiuhaoInuse  int64   `json:"qiuhaoinuse"`
	QiuhaoAmount float64 `json:"qiuhaoamount"`
	QiuhaoNum    int64   `json:"qiuhaonum"`
}

type FengkongLog struct {
	ID           int64  `json:"id"`
	Uuid         int64  `json:"uuid"`         // 下单用户ID
	Account      string `json:"account"`      // 会员账号
	CompanyID    int64  `json:"companyid"`    // 下单用户ID
	FengkongDate int64  `json:"fengkongdate"` // 风控日期
	RoomID       int64  `json:"roomid"`       // 游戏ID
	RoomCN       string `json:"roomcn"`       // 游戏名
	Expect       string `json:"expect"`       // 期号
	FengkongType string `json:"fengkongtype"` // 风控类型
	FengkongInfo string `json:"fengkonginfo"` // 风控信息
}

// 跟投计划
type GentouPlan struct {
	PlanID             int64   `json:"planid"`
	PlanName           string  `json:"planname"`
	Uuid               int64   `json:"uuid"`
	CreateTime         int64   `json:"createtime"`
	InUse              int64   `json:"inuse"`
	GentouUuid         int64   `json:"gentouuuid"`
	GentouAccount      string  `json:"gentouaccount"`
	GentouGameList     string  `json:"gentougamelist"`
	GentouType         int64   `json:"gentoutype"`
	GentouPercent      float64 `json:"gentoupercent"`
	TouzhuAccountinfo1 string  `json:"touzhuaccountinfo1"`
	TouzhuAccountinfo2 string  `json:"touzhuaccountinfo2"`
	TouzhuAccountinfo3 string  `json:"touzhuaccountinfo3"`
	TouzhuAccountinfo4 string  `json:"touzhuaccountinfo4"`
	TouzhuAccountinfo5 string  `json:"touzhuaccountinfo5"`
	TouzhuAccountinfo6 string  `json:"touzhuaccountinfo6"`
}

// 跟投账号信息
type TouzhuAccountinfo struct {
	WebName  string  `json:"webname"`  // 网站名
	SafeKey  string  `json:"safekey"`  // 安全码
	Account  string  `json:"account"`  // 账号
	Password string  `json:"password"` // 密码
	Percent  float64 `json:"percent"`  // 分配比例
}

// 跟投日志
type GentouLog struct {
	ID         int64   `json:"id"`
	Uuid       int64   `json:"uuid"`
	Account    string  `json:"account"`
	RoomID     int64   `json:"roomid"`
	Expect     string  `json:"expect"`
	ItemID     int64   `json:"itemid"`
	Iteminfo   string  `json:"iteminfo"`
	Num        string  `json:"num"`
	Amount     float64 `json:"amount"`
	GentouTime int64   `json:"gentoutime"`
	PlanID     int64   `json:"planid"`
	OrderID    int64   `json:"orderid"`
	Result     int64   `json:"result"`
}

// 风标日志
type FengkongmarkLog struct {
	ID         int64  `json:"id"`
	Uuid       int64  `json:"uuid"`       // 下单用户ID
	Account    string `json:"account"`    // 会员账号
	MarkDate   int64  `json:"markdate"`   // 风控日期
	SrcUuid    int64  `json:"srcuuid"`    //
	SrcAccount string `json:"srcaccount"` //
}

// 货量变动日志
type HuoliangLog struct {
	ID             int64 `json:"id"`
	OrderID        int64
	Uuid           int64 // 补货代理ID
	Roomid         int64
	Pan            string
	Itemid         int64
	Transferamount int64   // 补货额度
	Prehuoliang    float64 // 目前DB已收
	Bendanhuoliang float64 // 本单基础占成
	Sufhuoliang    float64 // 本单基础占成
}

type RevokeLog struct { // 结算日志
	ID        int64 `gorm:"primary_key;not null;AUTO_INCREMENT"`
	OrderID   int64 `json:"orderid"`   // 订单ID
	Uuid      int64 `json:"uuid"`      // 撤单人
	OpTime    int64 `json:"optime"`    // 撤单时间
	StatsFlag int64 `json:"statsflag"` // 货量修改标识
}
