package tablestruct

type RetMsg struct {
	Status     bool
	Error_code int64
	Msg        string
}

//
type UserInfo struct {
	Uuid     int64  `json:"uuid"`     // 分配ID
	Account  string `json:"account"`  // 用户名登录
	Password string `json:"password"` // 用户名登录
	Nickname string `json:"nickname"` // 占成信息
}

type CallLog struct {
	ID          int64  `json:"id"`          // 分配ID
	SrcAccount  string `json:"srcaccount"`  // 呼出账号
	CallTime    int64  `json:"calltime"`    // 呼出时间
	DestAccount string `json:"destaccount"` // 被呼账号
	RespTime    int64  `json:"resqtime"`    // 被呼处理时间
	Reqinfo     string `json:"reqinfo"`     // 呼出内容
	Respinfo    string `json:"respinfo"`    // 回复内容
}
