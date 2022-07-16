package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	//code = 1000 --- 1 开头 用户模块错误
	ERROR_USERNAME_DUPLICATED = 1001 //用户名重复
	ERROE_USERPASSWORD_WRONG  = 1002 //用户密码错误
	ERROR_USER_NOT_EXIST      = 1003 //用户不存在
	ERROR_TOKEN_NOT_EXIST     = 1004
	ERROR_TOKEN_RUNTIME       = 1005
	ERROR_TOKEN_RUN           = 1006 //虚假token
	ERROR_TOKEN_TYPEWRONG     = 1007
	//code = 2000 ----2 开头 文章模块错误
	//code = 3000 ----3 开头 分类模块错误
)

//错误状态码-错误信息
var CodeMsg = map[int]string{
	SUCCESS:                   "OK",
	ERROR:                     "FAIL",
	ERROR_USERNAME_DUPLICATED: "用户名已存在！",
	ERROE_USERPASSWORD_WRONG:  "密码错误",
	ERROR_USER_NOT_EXIST:      "用户不存在",
	ERROR_TOKEN_NOT_EXIST:     "Token不存在",
	ERROR_TOKEN_RUNTIME:       "Token过期",
	ERROR_TOKEN_RUN:           "Token错误",
	ERROR_TOKEN_TYPEWRONG:     "Token格式错误",
}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}
