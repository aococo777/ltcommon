package commonstruct

const (
	SYSTEMID          = 10000 // 系统ID
	TOURISTID_BEGIN   = 1000000
	TOURISTID_END     = 9800000
	CSYY_SYSITEMID    = 20000    //创世云游系统id
	SessionExpiryTime = 120 * 60 // session过期时间
)

const (
	GGExp = "ggexp"
)

type OgGameType int

var (
	OGGameMap = map[string]OgGameType{
		"oglive": 1,
		"mg":     4,
		"pp":     5,
		"cmd":    6,
		"bbin":   38,
		"saba":   46,
	}
)

type ErrorType int

// http status 数字文本映射
var (
	BBINERROR = map[string]string{
		"0":      "成功",
		"10000":  "未授权登入",
		"10300":  "查寻许可失败",
		"10301":  "验证授权失败",
		"11000":  "系统维修中",
		"12000":  "内部服务器错误",
		"13000":  "数据库错误",
		"14000":  "方式未执行",
		"15000":  "频繁询问",
		"15001 ": "测试字符频繁询问",
		"15002":  "操作码频繁询问",
	}

	ErrorMap = map[ErrorType]string{
		1:  "注册站长出错",
		2:  "注册会员出错",
		3:  "新注单失败",
		4:  "登录失败",
		5:  "已经登录",
		6:  "退出登录失败",
		7:  "参数错误",
		8:  "已经封盘啦",
		9:  "注册会员查询代理资金类型出错",
		10: "新增用户基础信息",
		11: "查询上级代理的代理树出错",
		12: "解析上级代理的代理树",
		13: "新增用户代理树",
		14: "新增用户资金",
		15: "提交用户注册事务",
		16: "账号已存在",
		17: "用户在退出后查询信息",
		18: "新注单扣款",
		19: "资金状态异常",
		20: "用户资金不足",
		21: "日志生成失败",
		22: "站长下单",
		23: "已经返过现了",
		24: "登录已过期",
		25: "Sid错误",
		26: "新增用户统计",
		27: "Websocket 登录错误",
		28: "原密码错误",
		29: "修改密码失败",
		30: "手机修改失败",
		31: "修改交易密码",
		32: "修改银行卡",
		33: "查询记录失败",
		34: "注单ID异常",
		35: "设置开奖号码",
		36: "进入房间失败",
		37: "目标房间已满",
		38: "退出房间失败",
		39: "游戏类型查询房间失败",
		40: "用户ID查询房间失败",
		41: "创建充值记录失败",
		42: "获取失败",
		43: "解析失败",
		44: "IP异常",
		45: "提现失败",
		46: "PostData失败",
		47: "获取上级ID失败",
		48: "获取站长ID失败",
		49: "获取支付方式出错",
		50: "赔率内容有误",
		51: "房间ID有误",
		52: "微信号不存在",
		53: "上级查询订单失败",
		54: "没有账号",
		55: "上级代理信息异常",
		56: "新增用户赔率",
		57: "房间列表为空",
		58: "用户没有下级",
		59: "用户上传异常调水值",
		60: "新增用户赔率错误",
		61: "没有权限",
		62: "修改权限失败",
		63: "修改基准赔率失败",
		64: "修改水赔比失败",
		65: "动态调赔失败",
		66: "长龙调赔失败",
		67: "货量调赔失败",
		68: "查询货量变赔信息失败",
		69: "查询长龙变赔信息失败",
		70: "团队列表为空",
		71: "用户登录限制",
		72: "转换钱包失败",
		73: "使用激活码失败",
		74: "增加使用天数失败",
		75: "账号已过期",
		88: "下注金额超出范围",
		89: "账号格式错误",
		91: "超出单期限额",
		92: "水赔异常",
	}

	GAMEMAP = map[string]int{
		"wylottery": EXP_WYLOTTERY,
		"bbinlive":  EXP_BBINLIVE,
		"bbinslots": EXP_BBINSLOTS,
		"bbinsport": EXP_BBINSPORT,
		"sabalive":  EXP_SABALIVE,
		"sabaslots": EXP_SABASLOTS,
		"sabasport": EXP_SABASPORT,
	}

	QYMAP = map[int64]string{
		1:   "契约一级",
		2:   "契约二级",
		3:   "契约三级",
		4:   "契约四级",
		5:   "契约五级",
		6:   "契约六级",
		7:   "契约七级",
		100: "常规账户",
	}
)

// 平台、游戏 分类
var (
	// 平台信息
	PLATFORMMAP = map[string]string{
		"lottery":   "彩票",
		"actual":    "真人",
		"slots":     "电子",
		"uclottery": "UC彩票",
	}

	// 平台信息
	PLATFORMTYPEMAP = map[string]string{
		"lottery":   "GG",
		"actual":    "GG",
		"slots":     "GG",
		"uclottery": "UC",
	}

	// 种类信息
	GAMETYPEMAP = map[string]string{
		"hk6":    "六合类",
		"11x5":   "11选5类",
		"k3":     "快3类",
		"klsf8":  "快乐十分8球类",
		"klsf5":  "快乐十分5球类",
		"ssc":    "时时彩类",
		"pk10":   "pk10类",
		"fc3d":   "福彩3D类",
		"pl3":    "排列三类",
		"kl8":    "快乐8类",
		"pcdd":   "PC蛋蛋类",
		"baijia": "百家类",
		"bzbm":   "奔驰宝马类",
		"k3dw":   "快三定位类",
		"brnn":   "百人牛牛类",
		"wflm":   "老虎机类",
		"lhdz":   "龙虎大战类",
	}
)

// http status 错误类型枚举
const (
	NoError                    ErrorType = iota
	Error_RegManager                     // 注册站长出错  1
	Error_RegUser                        // 注册会员出错  2
	Error_NewOrder                       // 新注单失败  3
	Error_Login                          // 登录失败  4
	Error_AlreadyLogin                   // 已经登录  5
	Error_Exit                           // 退出登录失败  6
	Error_ParamsWrong                    // 参数错误  7
	Error_Ban                            // 已经封盘啦  8
	Error_RegUser_PreIDInfo              // 注册会员查询代理资金类型出错  9
	Error_RegUser_NewUserbase            // 新增用户基础信息  10
	Error_RegUser_SelPreBranch           // 查询上级代理的代理树出错  11
	Error_RegUser_UmPreBranch            // 解析上级代理的代理树 12
	Error_RegUser_NewBranch              // 新增用户代理树  13
	Error_RegUser_NewMoney               // 新增用户资金  14
	Error_RegUser_Commit                 // 提交用户注册事务  15
	Error_RegUser_Exist                  // 账号已存在  16

	Error_Query_Notlogin        // 用户在退出后查询信息 17
	Error_NewOrder_UptMoney     // 新注单扣款  18
	Error_Money_Unusual         // 资金状态异常  19
	Error_NewOrder_NotEnough    // 用户资金不足  20
	Error_NewOrder_Addlog       // 日志生成失败  21
	Error_NewOrder_Master       // 站长下单 22
	Error_RetCash_Done          // 已经返过现了  23
	Error_LoginTimeOut          // 登录已过期  24
	Error_SidWrong              // Sid错误  25
	Error_RegUser_NewResult     // 新增用户统计  26
	Error_Login_WebSocket       // Websocket 登录错误  27
	Error_Resetpwd_Oldpwd       // 原密码错误  28
	Error_Resetpwd_Update       // 修改密码失败  29
	Error_Resetphone_Update     // 手机修改失败  30
	Error_Resetdealpwd_Update   // 修改交易密码  31
	Error_ResetBankcard_Update  // 修改银行卡  32
	Error_Query_NotFound        // 查询记录失败  33
	Error_NewOrder_ResultID     // 注单ID异常 34
	Error_TEST_SETOPENCODE      // 设置开奖号码  35
	Error_Room_Into             // 进入房间失败 36
	Error_Room_Full             // 目标房间已满  37
	Error_Room_Out              // 退出房间失败  38
	Error_Room_NoGametype       // 游戏类型查询房间失败  39
	Error_Room_NoUuid           // 用户ID查询房间失败  40
	Error_NewRecharge           // 创建充值记录失败  41
	Error_Fixedamount_get       // 获取失败  42
	Error_Fixedamount_unmarshal // 解析失败  43
	Error_WrongIP               // IP异常  44
	Error_Withdraw              // 提现失败  45
	Error_PostData              // PostData失败  46
	Error_GetPreID              // 获取上级ID失败  47
	Error_GetMasterID           // 获取站长ID 失败  48
	Error_GetRechargewayList    // 获取支付方式出错  49
	Error_Portvalue             // 赔率内容有误  50
	Error_RoomID                // 房间ID有误  51
	Error_NoWxid                // 微信号不存在  52
	Error_GetOrderByPreid       // 微信号不存在  53
	Error_Login_NoAccount       // 没有账号  54
	Error_NoPreID               // 上级代理信息异常  55
	Error_RegUser_NewAgentodds  // 新增用户赔率  56
	Error_RoomlistIsNull        // 房间列表为空  57
	Error_UuidNoSuf             // 用户没有下级  58
	Error_DynamicShui           // 用户上传异常调水值  59
	Error_RegUser_NewAuthority  // 新增用户赔率  60
	Error_NoAuthority           // 没有权限 61
	Error_UpdateAuthority       // 修改权限失败 62
	Error_UpdateBaseodds        // 修改基准赔率失败 63
	Error_UpdateOddspercent     // 修改水赔比失败 64
	Error_UpdateDync_item       // 动态调赔失败 65
	Error_UpdateDync_changlong  // 长龙调赔失败 66
	Error_UpdateDync_amount     // 货量调赔失败 67
	Error_QueryDync_amount      // 货量调赔失败 68
	Error_QueryDync_changlong   // 货量调赔失败 69
	Error_TeamIsNull            // 团队列表为空 70
	Error_UpdateUserLimited     // 用户登录限制 71
	Error_TransferWallet        // 转换钱包失败 72
	Error_UseActcode            // 使用激活码失败 73
	Error_AddValiddate          // 增加使用天数失败 74
	Error_NotValidtime          // 账号已过期 75
	Error_Rechargewayinfo       // 通道信息错误 76
	Error_OGReigster            // OG注册失败 77
	Error_QRCode                // 获取二维码失败 78
	Error_OGTransfer            // OG钱包转换失败 79
	Error_OGPlay                // OG开始游戏 80
	Error_IGRESETPWD            // IG重置密码 81
	Error_IGLOGIN               // IG登录游戏 82
	Error_IGGETBALANCE          // IG查询余额 83
	Error_IGCHECKREF            // IG查询订单 84
	Error_BBINCREATEMEMBER      // BBIN注册失败 85
	Error_BBINLOGIN             // BBIN登录失败 86
	Error_NODATA                // 没有数据 87
	Error_AMOUNTNOTRANGE        // 下注金额超出范围  88
	Error_WrongForm             // 格式错误  89
	Error_BINDWECHAT            // 绑定微信失败 90
	Error_NOTEXPECTLIMIT        // 超出单期限额  91
	Error_WRONGODDS             // 水赔异常 92
	Error_REGTYPE               // regtype异常 93
	Error_ISGENERAL             // 代理会员类型异常 94
	Error_NOTPRE                // 非上级操作 95
	Error_ORDERID               // 错误的OrderID 96
	Error_SQLEXEC               // 数据库操作错误 97
	Error_RegCoThirdRechargeway // 新增用户赔率  98
)

type MsgID int

// ws 错误类型枚举
const (
	NoError_WS             ErrorType = iota // 无错
	Error_WS_NoAuthority                    // 没有权限 1
	Error_WS_Body                           // 上传Body错误 2
	Error_WS_RoomState                      // 房间状态异常 3
	Error_WS_WRONGACCOUNT                   // 账号密码错误 4
	Error_WS_ALREADYONLINE                  // 已经在线 5
	Error_WS_NOROOM                         // 没有房间使用 6
	Error_WS_NOPREID                        // 代理信息异常 7
	Error_WS_SETTLED                        // 信彩结算失败 8
	Error_WS_RETCASH_SYS                    // 系统返现失败 9
	Error_WS_RETCASH_USER                   // 用户返现失败 10
	Error_WS_CHANGLONG                      // 长龙保存异常 11
	Error_WS_NoUuid                         // 用户ID为空 12
	Error_WS_SWAPROOM                       // 切换房间错误 13
	Error_WS_SETTLE                         // 结算错误 14
	Error_WS_LOGINOUT                       // 登录过期 15
	Error_WS_UNSETTLE                       // 有未结算订单 16
)

// ws 百家乐 心跳类型
const (
	// 通用配置
	MSG_ID_XINTIAO MsgID = iota + 2000 // 心跳包

	// 用户登录
	MSG_ID_IN_CLI_LOGIN  //2001
	MSG_ID_OUT_CLI_LOGIN //2002

	// 用户退出
	MSG_ID_IN_CLI_EXIT  //2003
	MSG_ID_OUT_CLI_EXIT //2004

	// 用户切换房间
	MSG_ID_IN_CLI_SWAPROOM  //2005
	MSG_ID_OUT_CLI_SWAPROOM //2006

	// 荷官登录
	MSG_ID_IN_DEALER_BIND  //2007
	MSG_ID_OUT_DEALER_BIND //2008

	// 荷官开靴
	MSG_ID_IN_LEADER_NEWBOOT  //2009
	MSG_ID_OUT_LEADER_NEWBOOT //2010

	// 荷官开局
	MSG_ID_IN_LEADER_NEWROUND  //2011
	MSG_ID_OUT_LEADER_NEWROUND //2012

	// 荷官开牌
	MSG_ID_IN_LEADER_OPENCODE  //2013
	MSG_ID_OUT_LEADER_OPENCODE //2014

	// 荷官重置牌靴
	MSG_ID_IN_LEADER_RESET  //2015
	MSG_ID_OUT_LEADER_RESET //2016

	// 荷官延长下注时间
	MSG_ID_IN_LEADER_DELAY  //2017
	MSG_ID_OUT_LEADER_DELAY //2018

	// 结算完成
	MSG_ID_OUT_SETTLE_FINISH //2019

	// 用户发言
	MSG_ID_IN_CLI_TALK //2020

	// 用户资金
	MSG_ID_OUT_MONEY //2021

	// 开奖号码
	MSG_ID_OUT_WINNUM //2022

	// 开奖时间
	MSG_ID_OUT_WINTIME //2023

	// 当前服务器人数
	MSG_ID_OUT_ONLINE_ALL //2024

	// 真人在线人数
	MSG_ID_OUT_ACTUAL_ONLINE //2025

	// 当前房间人数统计
	MSG_ID_OUT_ROOM_ONLINE //2026

	// 结算结果
	MSG_ID_OUT_SETTLERESULT //2027

	// 下注统计
	MSG_ID_OUT_ORDERINFO //2028

	// 百家在线人数
	MSG_ID_OUT_BAIJIA_ONLINE //2029

	// 百家所有房间下注统计
	MSG_ID_OUT_BAIJIA_ORDERINFO //2030

	// 荷官封盘
	MSG_ID_IN_LEADER_BAN  //2031
	MSG_ID_OUT_LEADER_BAN //2032

	// 聊天代投
	MSG_ID_IN_TALK_ORDER_AGENT  //2033
	MSG_ID_OUT_TALK_ORDER_AGENT //2034

	// 长龙信息
	MSG_ID_OUT_CHANGLONG //2035

	// 房间状态
	MSG_ID_OUT_ROOMSTATE //2036

	// 向用户广播发言
	MSG_ID_OUT_CLI_TALK //2037

	// 监控信息
	MSG_ID_OUT_CLI_MONITERINFO //2038

	// 视频源
	MSG_ID_OUT_CLI_VIDEOURL //2039

	// 今日盈亏
	MSG_ID_OUT_CLI_TODAYRESULT //2040

	// 动态赔率
	MSG_ID_OUT_CLI_DYNIMICODDS //2041

	// 站位的
	MSG_ID_OUT_CLI_ZHANWEI //2042

	// 具体牌号
	MSG_ID_IN_LEADER_OPENITEM  //2043
	MSG_ID_OUT_LEADER_OPENITEM //2044

	// 剩余牌数
	MSG_ID_IN_REMAIN_CARDS  //2045
	MSG_ID_OUT_REMAIN_CARDS //2046

	// 进入房间的用户信息
	MSG_ID_OUT_ENTERUSER_INFO //2047

	// 用户入座
	MSG_ID_IN_INTOSEAT  //2048 入座
	MSG_ID_OUT_INTOSEAT //2049 入座返回
	MSG_ID_IN_OUTSEAT   //2050 离座
	MSG_ID_OUT_OUTSEAT  //2051 离座返回
	MSG_ID_OUT_SEATINFO //2052 当前座位信息

	MSG_ID_OUT_SUPERWINNER //2053 大赢家信息
	MSG_ID_OUT_SEATRESULT  //2054  座位输赢结果

	// 奔驰宝马
	MSG_ID_IN_BZBM_NEWROUND  //2055 开局
	MSG_ID_OUT_BZBM_NEWROUND //2056 开局

	MSG_ID_IN_BZBM_BAN  //2057 封盘
	MSG_ID_OUT_BZBM_BAN //2058 封盘

	MSG_ID_IN_BZBM_OPENCODE  //2059 开奖结果
	MSG_ID_OUT_BZBM_OPENCODE //2060 开奖结果

	MSG_ID_OUT_BZBM_USERORDER  //2061 客户下注
	MSG_ID_OUT_BZBM_ORDERSTATS //2062 下注统计

	MSG_ID_IN_BZBM_CLILOGIN  //2063 客户登录
	MSG_ID_OUT_BZBM_CLILOGIN //2064 客户登录

	MSG_ID_IN_BZBM_CLIEXIT  //2065 客户退出
	MSG_ID_OUT_BZBM_CLIEXIT //2066 客户退出

	MSG_ID_IN_BZBM_INTOBANK  //2067 客户上庄
	MSG_ID_OUT_BZBM_INTOBANK //2068 客户上庄

	MSG_ID_IN_BZBM_OUTBANK  //2069 客户下庄
	MSG_ID_OUT_BZBM_OUTBANK //2070 客户下庄

	MSG_ID_OUT_BZBM_SETTLEFINISH //2071 结算完成
	MSG_ID_OUT_BZBM_SETTLERESULT //2072 结算结果

	MSG_ID_OUT_BZBM_RANKINGLIST  //2073  盈利榜
	MSG_ID_OUT_BZBM_WINTIME      //2074  开奖时间
	MSG_ID_OUT_BZBM_ONLINE       //2075  当前房间人数统计
	MSG_ID_OUT_BZBM_CHANGEBANKER //2076  切换庄家

	// 百人牛牛
	MSG_ID_IN_BRNN_NEWROUND  //2077 开局
	MSG_ID_OUT_BRNN_NEWROUND //2078 开局

	MSG_ID_IN_BRNN_BAN  //2079 封盘
	MSG_ID_OUT_BRNN_BAN //2080 封盘

	MSG_ID_IN_BRNN_OPENCODE  //2081 开奖结果
	MSG_ID_OUT_BRNN_OPENCODE //2082 开奖结果

	MSG_ID_OUT_BRNN_USERORDER  //2083 客户下注
	MSG_ID_OUT_BRNN_ORDERSTATS //2084 下注统计

	MSG_ID_IN_BRNN_CLILOGIN  //2085 客户登录
	MSG_ID_OUT_BRNN_CLILOGIN //2086 客户登录

	MSG_ID_IN_BRNN_CLIEXIT  //2087 客户退出
	MSG_ID_OUT_BRNN_CLIEXIT //2088 客户退出

	MSG_ID_IN_BRNN_INTOBANK  //2089 客户上庄
	MSG_ID_OUT_BRNN_INTOBANK //2090 客户上庄

	MSG_ID_IN_BRNN_OUTBANK  //2091 客户下庄
	MSG_ID_OUT_BRNN_OUTBANK //2092 客户下庄

	MSG_ID_OUT_BRNN_SETTLEFINISH //2093 结算完成
	MSG_ID_OUT_BRNN_SETTLERESULT //2094 结算结果(个人、座位、玩家列表)

	MSG_ID_OUT_BRNN_RANKINGLIST  //2095  盈利榜
	MSG_ID_OUT_BRNN_WINTIME      //2096  开奖时间
	MSG_ID_OUT_BRNN_ONLINE       //2097  当前房间人数统计
	MSG_ID_OUT_BRNN_CHANGEBANKER //2098  切换庄家
	MSG_ID_OUT_BRNN_ORDERINFO    //2099  奖池统计

	MSG_ID_IN_BRNN_INTOSEAT  //2100 入座
	MSG_ID_OUT_BRNN_INTOSEAT //2101 入座返回
	MSG_ID_IN_BRNN_OUTSEAT   //2102 离座
	MSG_ID_OUT_BRNN_OUTSEAT  //2103 离座返回
	MSG_ID_OUT_BRNN_SEATINFO //2104 当前座位信息

	MSG_ID_OUT_BRNN_SUPERWINNER //2105 大赢家信息

	MSG_ID_OUT_BRNN_BANKERLIST //2106 庄家列表

	// 重庆时时彩
	MSG_ID_IN_CQSSC_NEWROUND  //2107 开局
	MSG_ID_OUT_CQSSC_NEWROUND //2108 开局

	MSG_ID_IN_CQSSC_BAN  //2109 封盘
	MSG_ID_OUT_CQSSC_BAN //2110 封盘

	MSG_ID_IN_CQSSC_OPENCODE  //2111 开奖结果
	MSG_ID_OUT_CQSSC_OPENCODE //2112 开奖结果

	MSG_ID_OUT_SLOTSBAIJIA_XianHong // 2113	slotsbaijia限红改变
)

// ws 红黑大战 心跳类型
const (
	MSG_ID_IN_HHDZ_NEWROUND  MsgID = iota + 2120 //2120 开局
	MSG_ID_OUT_HHDZ_NEWROUND                     //2121 开局

	MSG_ID_IN_HHDZ_BAN  //2122 封盘
	MSG_ID_OUT_HHDZ_BAN //2123 封盘

	MSG_ID_IN_HHDZ_OPENCODE  //2124 开奖结果
	MSG_ID_OUT_HHDZ_OPENCODE //2125 开奖结果

	MSG_ID_OUT_HHDZ_USERORDER  //2126 客户下注
	MSG_ID_OUT_HHDZ_ORDERSTATS //2127 下注统计

	MSG_ID_IN_HHDZ_CLILOGIN  //2128 客户登录
	MSG_ID_OUT_HHDZ_CLILOGIN //2129 客户登录

	MSG_ID_IN_HHDZ_CLIEXIT  //2130 客户退出
	MSG_ID_OUT_HHDZ_CLIEXIT //2131 客户退出

	MSG_ID_IN_HHDZ_INTOBANK  //2132 客户上庄
	MSG_ID_OUT_HHDZ_INTOBANK //2133 客户上庄

	MSG_ID_IN_HHDZ_OUTBANK  //2134 客户下庄
	MSG_ID_OUT_HHDZ_OUTBANK //2135 客户下庄

	MSG_ID_OUT_HHDZ_SETTLEFINISH //2136 结算完成
	MSG_ID_OUT_HHDZ_SETTLERESULT //2137 结算结果

	MSG_ID_OUT_HHDZ_RANKINGLIST  //2138  盈利榜
	MSG_ID_OUT_HHDZ_WINTIME      //2139  开奖时间
	MSG_ID_OUT_HHDZ_ONLINE       //2140  当前房间人数统计
	MSG_ID_OUT_HHDZ_CHANGEBANKER //2141  切换庄家

	MSG_ID_IN_HHDZ_INTOSEAT  //2142 入座
	MSG_ID_OUT_HHDZ_INTOSEAT //2143 入座返回
	MSG_ID_IN_HHDZ_OUTSEAT   //2144 离座
	MSG_ID_OUT_HHDZ_OUTSEAT  //2145 离座返回
	MSG_ID_OUT_HHDZ_SEATINFO //2146 当前座位信息
)

// ws 飞禽走兽 心跳类型
const (
	MSG_ID_IN_FQZS_NEWROUND  MsgID = iota + 2280 //2280 开局
	MSG_ID_OUT_FQZS_NEWROUND                     //2281 开局

	MSG_ID_IN_FQZS_BAN  //2282 封盘
	MSG_ID_OUT_FQZS_BAN //2283 封盘

	MSG_ID_IN_FQZS_OPENCODE  //2284 开奖结果
	MSG_ID_OUT_FQZS_OPENCODE //2285 开奖结果

	MSG_ID_OUT_FQZS_USERORDER  //2286 客户下注
	MSG_ID_OUT_FQZS_ORDERSTATS //2287 下注统计

	MSG_ID_IN_FQZS_CLILOGIN  //2288 客户登录
	MSG_ID_OUT_FQZS_CLILOGIN //2289 客户登录

	MSG_ID_IN_FQZS_CLIEXIT  //2290 客户退出
	MSG_ID_OUT_FQZS_CLIEXIT //2291 客户退出

	MSG_ID_IN_FQZS_INTOBANK  //2292 客户上庄
	MSG_ID_OUT_FQZS_INTOBANK //2293 客户上庄

	MSG_ID_IN_FQZS_OUTBANK  //2294 客户下庄
	MSG_ID_OUT_FQZS_OUTBANK //2295 客户下庄

	MSG_ID_OUT_FQZS_SETTLEFINISH //2296 结算完成
	MSG_ID_OUT_FQZS_SETTLERESULT //2297 结算结果

	MSG_ID_OUT_FQZS_RANKINGLIST  //2298  盈利榜
	MSG_ID_OUT_FQZS_WINTIME      //2299  开奖时间
	MSG_ID_OUT_FQZS_ONLINE       //2300  当前房间人数统计
	MSG_ID_OUT_FQZS_CHANGEBANKER //2301  切换庄家  成功返回新的庄家信息/失败返回错误
	MSG_ID_OUT_FQZS_BANKERLIST   //2302  排庄列表
)

// ws 龙虎大战 心跳类型
const (
	MSG_ID_IN_LHDZ_NEWROUND  MsgID = iota + 2350 //2350 开局
	MSG_ID_OUT_LHDZ_NEWROUND                     //2351 开局

	MSG_ID_IN_LHDZ_BAN  //2352 封盘
	MSG_ID_OUT_LHDZ_BAN //2353 封盘

	MSG_ID_IN_LHDZ_OPENCODE  //2354 开奖结果
	MSG_ID_OUT_LHDZ_OPENCODE //2355 开奖结果

	MSG_ID_OUT_LHDZ_USERORDER  //2356 客户下注
	MSG_ID_OUT_LHDZ_ORDERSTATS //2357 下注统计

	MSG_ID_IN_LHDZ_CLILOGIN  //2358 客户登录
	MSG_ID_OUT_LHDZ_CLILOGIN //2359 客户登录

	MSG_ID_IN_LHDZ_CLIEXIT  //2360 客户退出
	MSG_ID_OUT_LHDZ_CLIEXIT //2361 客户退出

	MSG_ID_IN_LHDZ_INTOBANK  //2362 客户上庄
	MSG_ID_OUT_LHDZ_INTOBANK //2363 客户上庄

	MSG_ID_IN_LHDZ_OUTBANK  //2364 客户下庄
	MSG_ID_OUT_LHDZ_OUTBANK //2365 客户下庄

	MSG_ID_OUT_LHDZ_SETTLEFINISH //2366 结算完成
	MSG_ID_OUT_LHDZ_SETTLERESULT //2367 结算结果

	MSG_ID_OUT_LHDZ_RANKINGLIST  //2368  盈利榜
	MSG_ID_OUT_LHDZ_WINTIME      //2369  开奖时间
	MSG_ID_OUT_LHDZ_ONLINE       //2370  当前房间人数统计
	MSG_ID_OUT_LHDZ_CHANGEBANKER //2371  切换庄家

	MSG_ID_IN_LHDZ_INTOSEAT  //2372 入座
	MSG_ID_OUT_LHDZ_INTOSEAT //2373 入座返回
	MSG_ID_IN_LHDZ_OUTSEAT   //2374 离座
	MSG_ID_OUT_LHDZ_OUTSEAT  //2375 离座返回
	MSG_ID_OUT_LHDZ_SEATINFO //2376 当前座位信息
)

// ws 分分彩 心跳类型
const (
	MSG_ID_IN_FFC_NEWROUND  = iota + 2377 //2377 开局
	MSG_ID_OUT_FFC_NEWROUND               //2378 开局

	MSG_ID_IN_FFC_OPENCODE  //2379 开奖结果
	MSG_ID_OUT_FFC_OPENCODE //2380 开奖结果

	MSG_ID_OUT_FFC_SETTLEFINISH //2381 结算完成
	MSG_ID_OUT_FFC_SETTLERESULT //2382 结算结果

	MSG_ID_OUT_FFC_RANKINGLIST //2383  盈利榜
	MSG_ID_OUT_FFC_WINTIME     //2384  开奖时间
	MSG_ID_OUT_FFC_ONLINE      //2385  当前房间人数统计

	MSG_ID_IN_FFC_BAN  //2386 封盘
	MSG_ID_OUT_FFC_BAN //2387 封盘

	MSG_ID_IN_FFC_CLILOGIN  //2388 客户登录
	MSG_ID_OUT_FFC_CLILOGIN //2389 客户登录

	MSG_ID_IN_FFC_CLIEXIT  //2390 客户退出
	MSG_ID_OUT_FFC_CLIEXIT //2391 客户退出

	MSG_ID_OUT_FFC_USERORDER  //2392 客户下注
	MSG_ID_OUT_FFC_ORDERSTATS //2393 下注统计
)

// 服务器内部辨识ID段
const (
	MSG_ID_SEVER_ONE        MsgID = iota + 1001 // 服务器辨识ID = 单人推送
	MSG_ID_SEVER_ROOM                           // 服务器辨识ID = 房间推送
	MSG_ID_SEVER_ALL                            // 服务器辨识ID = 所有在线推送给
	MSG_ID_SEVER_SETTLETYPE                     // 服务器辨识ID = 结算类型
	MSG_ID_SEVER_DATARES                        // 服务器辨识ID = 数据来源推送
)

// 服务间通信
const (
	MSG_ID_RPC_XINTIAO    MsgID = iota + 5000
	MSG_ID_RPC_SETTLE_REQ       // 百家结算请求
	MSG_ID_RPC_SETTLE_RES       // 百家结算回复
	MSG_ID_RPC_SETTLED_XC       // 推送信彩结算完成
	MSG_ID_RPC_SETTLED_GC       // 推送官彩结算完成
	MSG_ID_RPC_RETCASH          // 返现完成
	MSG_ID_RPC_CHANGLONG        // 推送长龙信息
)

const (
	Normal = iota + 1 // 普通注单
	ZouFei            // 走飞注单
)

// 游戏种类区分
const (
	Xincai  = iota + 1 // lottery
	Guancai            // guanfang
	Zhenren            // 真人
	Jingcai            // 竞彩
)

const (
	Long = iota + 1
	Hu
	He
	Baozi
	Shunzi
	Duizi
	Banshun
	Zaliu
)

type UserOpType int

// 后台操作分类
const (
	OpType_Zhuce             UserOpType = iota // 注册 0
	OpType_Login                               // 登录 1
	OpType_Exit                                // 退出 2
	OpType_ReCharge                            // 充值 3
	OpType_NewOrder                            // 下注 4
	OpType_Query                               // 查询 5
	OpType_ReSettle                            // 重新结算 6
	OpType_Userbaseinfo                        // 用户基本信息 7
	OpType_Companybaseinfo                     // 公司基本信息 11
	OpType_SetConfig                           // 设置基本配置 14
	OpType_SetPromotion                        // 设置活动 16
	OpType_CompanyGonggao                      // 公司公告 16
	OpType_DeleteAccount                       // 删除 17
	OpType_Pankoushezhi                        // 盘口设置 18
	OpType_Jizhunpeilv                         // 基准赔率 19
	OpType_Quanxian                            // 权限
	OpType_Youxikaiguan                        // 游戏开关
	OpType_Qiangzhijiangshui                   // 玩法设置
	OpType_Shoudongjiangpei                    // 手动降赔
	OpType_Xianzhidenglu                       // 限制登录
	OpType_Buhuoshezhi                         // 补货设置
	OpType_Jingshizhudan                       // 警示注单设置
	OpType_Youxishuoming                       // 游戏说明
	OpType_Peilvkaobei                         // 赔率拷贝
	OpType_GongsiGameset                       // 游戏设置
	OpType_SetSettlecode                       // 修改开奖号码

)

type MoneyType int

const (
	MoneyType_Cash    MoneyType = iota + 1 // 现金
	MoneyType_Virtual                      // 虚拟
	MoneyType_Credit                       // 信用
)

type SettleType string

// 两面彩结算类型分类
const (
	GameType_ssc    SettleType = "ssc" // 重庆时时彩
	GameType_hk6    SettleType = "hk6" // 香港彩
	GameType_11x5   SettleType = "11x5"
	GameType_k3     SettleType = "k3"
	GameType_klsf8  SettleType = "klsf8"
	GameType_klsf5  SettleType = "klsf5"
	GameType_kl8    SettleType = "kl8"
	GameType_pk10   SettleType = "pk10"
	GameType_cakeno SettleType = "cakeno"
	GameType_mlaft  SettleType = "mlaft"
	GameType_fc3d   SettleType = "fc3d"
	GameType_pl3    SettleType = "pl3"
	GameType_hall   SettleType = "hall" // 大厅
	GameType_baijia SettleType = "baijia"
	GameType_toubao SettleType = "toubao"
	GameType_lunpan SettleType = "lunpan"
	GameType_longhu SettleType = "longhu"
)

type ResetType int

// 个人信息重置分类
const (
	Resetpwd      ResetType = iota + 1 // 重置密码 1
	Resetphone                         // 重置电话号码 2
	Resetbankinfo                      // 重置银行信息 3
	Resetdealpwd                       // 重置交易密码 4
	Resetbase                          // 重置基本信息 5
	ResetRealname                      // 充值真实姓名
)

// 下注订单结算状态
const (
	Ret_UnSettle = iota // 未结算   0
	Ret_System          // 系统结算 1
	Ret_Manual          // 手动结算 2
	Ret_UnValid         // 系统退单 3
	Ret_Revoke          // 用户撤单 4
)

// 充值订单状态
const (
	WAIT    int64 = iota // 等待充值返回
	SUCCESS              // 充值成功
	OUTTIME              // 充值超时
)

const (
	Order_Normal int64 = iota + 6100 // 普通注单 6100
	Order_Agent                      // 代投  1 6101
	Order_Chat                       // 聊天投注6102
)

const (
	WebRegister int64 = iota + 6200 // 网站注册 6200
	WXRegister                      // 微信注册 6201
	WebAgent                        // 代理注册 6202
	UrlExt                          // 推广链接注册 6203
	SupRegister                     // 辅助账号注册 6204
)

const (
	LABEL_LEADER  int = iota + 6300 // 用户标签 = 荷官
	LABEL_MANAGER                   // 用户标签 = 管理
	LABEL_VIP                       // 用户标签 = VIP
)

type MoneyUpdateType int

// 资金变动分类
const (
	UptType_NewOrder      MoneyUpdateType = iota + 6600 // 用户注单 6600
	UptType_ReCharge                                    // 在线充值 6601
	UptType_Withdraw                                    // 用户取款 6602
	UptType_Paijiang                                    // 系统派奖 6603
	UptType_Chipei                                      // 代理返点 6604
	UptType_Gongzi                                      // 代理返佣 6605
	UptType_RemitIn                                     // 管理员上分 6606
	UptType_RemitOut                                    // 管理员下分 6607
	UptType_CONTRACT                                    // 契约履行 6608
	UptType_Benifit                                     // 活动优惠 6609
	UptType_PreIn                                       // 上级转账 6610
	UptType_ThirdRepair                                 // 第三方补单 6611
	UptType_OutSuf                                      // 转入下级 6612
	UptType_ExpTransfer                                 // 钱包转换 6613
	UptType_AgentWX                                     // 公司入款 6614
	UptType_DecSuf                                      // 回收下级 6615
	UptType_PreDec                                      // 被上级回收 6616
	UptType_Revoke                                      // 用户撤单 6617
	UptType_Loseback                                    // 会员输返 6618
	UptType_CoRecharge                                  // 公司线下收款 6619
	UptType_SysRet                                      // 系统退款 6620
	UptType_ThirdOut                                    // 第三方平台转出 6621
	UptType_ThirdIn                                     // 第三方平台转入 6622
	UptType_WalletLimit                                 // 修改钱包额度 6623
	UptType_Recoveramount                               // 钱包额度恢复 6624
)

const (
	GENERAL int = iota + 6700 // 普通会员
	AGENT                     // 代理
	TOURIST                   // 游客
)

const (
	Role_CashAgent  = iota // 现金代理 0
	Role_CashPlayer        // 现金会员 1
)

const (
	Role_Player = iota + 6700 // 信用会员 6700
	Role_Agent                // 信用代理 6701
)

const (
	Role_Tourist   = iota + 6702 // 游客 6702
	Role_Company                 // 公司 6703
	Role_Assistant               // 子账号 6704
	Role_Dealer                  // 荷官 6705
	Role_Super                   // 超管 6706
	Role_Middleman               // 中间人  6707
	Role_Financer                // 财务  6708
)

const (
	EXP_WYLOTTERY int = iota + 7000 // wy彩票
	EXP_BBINLIVE                    // BBIN视讯
	EXP_BBINSLOTS                   // BBIN电子
	EXP_BBINSPORT                   // BBIN体育
	EXP_SABALIVE                    // SABA视讯
	EXP_SABASLOTS                   // SABA电子
	EXP_SABASPORT                   // SABA体育
)

const (
	Way_Alipay = iota + 7100 // 微信支付 7100
	Way_Wechat               // 支付宝支付 7101
	Way_Bank                 // 银行卡转入 7102
)

// 充值状态
const (
	State_Wait    = iota // 等待付款 0
	State_Success        // 充值完成 1
	State_Timeout        // 充值超时 2
)

// 提款状态
const (
	Withdraw_Wait    = iota + 7200 // 等待审核 7200
	Withdraw_Success               // 提现完成 7201
	Withdraw_Failed                // 提现失败 7202
	Recharge_BK                    // 后台确认充值 7203
)

// 游戏结算状态
const (
	SaveResult_Success = iota + 7300 // 保存结算完成 7300
	SaveResult_Failed                // 保存结算失败 7301
	Paicai_Success                   // 派彩成功 7302
	Paicai_Failed                    // 派彩失败 7303
	Fanyong_Success                  // 返佣成功 7304
	Fanyong_Failed                   // 返佣失败 7305
	SettleBegin                      // 开始结算 7306
	CodePre                          // 提前开奖 7307
	CodeError                        // 号码异常 7308
)

const (
	Recharge1 = iota + 7500 // 普通 7500
	Recharge2               // 关闭 7501
	Recharge3               // 维护 7502
	Recharge4               // 推荐 7503
	Recharge5               // 未开通 7504
)

const (
	RoomState_Close        = iota // 关闭 0
	RoomState_Open                // 开通 1 (正常)
	RoomState_Repair              // 维护 2
	RoomState_NotSet              // 未开通 3
	RoomState_Hide                // 隐藏 4
	RoomState_ChangeDealer        // 缓荷官维护 5
	RoomState_LostConn            // 网络连接失败 6
)

const ( // 公司入款分类
	CoBank   = iota + 7600 // 银联   7600
	CoAli                  // 支付宝  7601
	CoWechat               // 微信  7602
	CoQQ                   // QQ  7603
	CoJD                   // 京东  7604
)

const ( // 百家角色分类
	BJ_Banker = iota + 7610 // 庄家   7610
	BJ_Player               // 闲家  7611
)

const ( // 公告类型
	SHOWTYPE_SYSTEM = iota + 7700 // 系统讯息   7700
	SHOWTYPE_WINDOW               // 弹窗   7701
	SHOWTYPE_MSG                  // 通告   7702
)

const ( // 公告类型
	LIMIT_LOGIN_ONE  = iota + 1 // 用户限制登录   1
	LIMIT_LOGIN_TEAM            // 团队限制登录   2
	LIMIT_ORDER_HAND            // 停押 3
	LIMIT_ORDER_SYS             // 系统临时停押 4
)

/*
cash
gaopin	高频
gaopinlimit 	高频额度
xianggangliuhe	香港六合
xianggangliuhelimit	香港六合额度
tcp3	体彩P3
tcp3limit	体彩P3额度
fc3d	福彩3D
fc3dlimit	福彩3D额度
taiwandaletou	台湾大乐透
taiwandaletoulimit	台湾大乐透额度
aomenliuhe	澳门六合
aomenliuhelimit	澳门六合额度
taiwanliuhe	台湾六合
taiwanliuhelimit	台湾六合额度

*/
