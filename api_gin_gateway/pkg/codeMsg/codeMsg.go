package codeMsg

type CodeMessage struct {
	StatusCode    uint64
	StatusMessage string
}

const (
	SUCCESS = 11
	Failed  = 12
	TRUE    = 21
	FALSE   = 22

	/* 关于用户基本信息的错误 */

	ErrorUserExist        uint64 = 10001
	ErrorUserNotExist     uint64 = 10002
	ErrorEncryptionFailed uint64 = 10003
	ErrorDecryptionFailed uint64 = 10004
	ErrorRegisterFailed   uint64 = 10005
	ErrorLoginFailed      uint64 = 10006
	ErrorInvalidParameter uint64 = 10007
	ErrorInvalidToken     uint64 = 10008
)

var ErrorMessageFlags = map[uint64]string{
	ErrorUserExist:        "用户已存在",
	ErrorUserNotExist:     "用户不存在",
	ErrorEncryptionFailed: "加密失败",
	ErrorDecryptionFailed: "解密失败",
	ErrorRegisterFailed:   "注册失败",
	ErrorLoginFailed:      "登录失败",
	ErrorInvalidParameter: "无效参数",
	ErrorInvalidToken:     "无效token",
}

// GetErrorMsg 获取error状态码对应信息
func GetErrorMsg(code uint64) string {
	msg, ok := ErrorMessageFlags[code]
	if !ok {
		return "error状态码信息获取失败"
	}
	return msg
}
