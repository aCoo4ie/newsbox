package controller

type RespCode int

const (
	CodeSuccess RespCode = 1000 + iota
	CodeUserExists
	CodeUserNotExists
	CodeInvalidParams
	CodeInvalidPassword
	CodeServerBusy
)

var RespMsg = map[RespCode]string{
	CodeSuccess:         "成功",
	CodeUserExists:      "用户已存在",
	CodeUserNotExists:   "用户不存在",
	CodeInvalidParams:   "参数错误",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器内部错误",
}

// Msg return the message corresponding to the response code
func (c RespCode) Msg() string {
	msg, ok := RespMsg[c]
	if !ok {
		return RespMsg[CodeServerBusy]
	}
	return msg
}
