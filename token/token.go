package token

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"ltback/src/ltcommon/commonfunc"

	"ltback/src/tt.com/component/uuid"
)

type Token struct {
	uidGen uuid.UuidGen // 唯一计数器
}

func (this *Token) Init() {
	this.uidGen.SetGenID(19070)
	fmt.Sprintf("Token enter!!!!")
}

func (this *Token) GetToken(username string) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s%d", username, commonfunc.BeijingTime().UnixNano()))) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	id := this.uidGen.GenID()
	if id > 999995 {
		this.uidGen.SetGenID(100)
	}
	encodestr := fmt.Sprintf("%s%07d", hex.EncodeToString([]byte(cipherStr)), this.uidGen.GenID())
	return encodestr
}
