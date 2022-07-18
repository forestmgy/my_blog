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
	ERROR_TOKEN_WRONG         = 1006 //虚假token
	ERROR_TOKEN_TYPEWRONG     = 1007
	ERROR_USER_NO_RIGHT       = 1008
	//code = 2000 ----2 开头 分类模块错误
	ERROR_CATENAME_DULPICATED = 2001
	ERROR_CATENAME_NOT_EXIST  = 2002
	//code = 3000 ----3 开头 文章模块错误
	ERROR_ART_NOT_EXIST = 3001
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
	ERROR_TOKEN_WRONG:         "Token错误",
	ERROR_TOKEN_TYPEWRONG:     "Token格式错误",
	ERROR_CATENAME_DULPICATED: "分类已存在！",
	ERROR_CATENAME_NOT_EXIST:  "分类不存在",
	ERROR_ART_NOT_EXIST:       "文章不存在",
	ERROR_USER_NO_RIGHT:       "该用户不是管理员",
}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}
