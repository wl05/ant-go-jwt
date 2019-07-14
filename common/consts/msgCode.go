package consts

const (
	SUCCECC = 200
	ERROR   = 0

	ERROR_CODE_PARAMETER_ILLEGAL = 1101   // 参数错误
	ERROR_DES_PARAMETER_ILLEGAL  = "参数错误" // 参数错误

	ERROR_CODE_REQUEST = 1102   // 请求出错
	ERROR_DES_REQUEST  = "请求出错" // 请求出错

	ERROR_CODE_USER_EXIST = 1103     // 用户名已存在
	ERROR_DES_USER_EXIST  = "用户名已存在" // 用户名已存在

	ERROR_CODE_USER_NOT_EXIST = 1104     // 用户名不存在
	ERROR_DES_USER_NOT_EXIST  = "用户名不存在" // 用户名不存在

	ERROR_CODE_USERNAME_OR_PASSWORD_ERROR = 1105        // 用户名或者密码错误
	ERROR_DES_USERNAME_OR_PASSWORD_ERROR  = "用户名或者密码错误" // 用户名或者密码错误

	ERROR_CODE_LOGIN_ERROR = 1106    // 请重新登录
	ERROR_DES_LOGIN_ERROR  = "请重新登录" // 请重新登录

	ERROR_CODE_LOGIN_EXPIRED_ERROR = 1107   // 登录过期
	ERROR_DES_LOGIN_EXPIRED_ERROR  = "登录过期" // 登录过期

	ERROR_CODE_TAG_EXIST = 1108     // 标签名已存在
	ERROR_DES_TAG_EXIST  = "标签名已存在" // 标签名已存在

	ERROR_CODE_CATEGORY_EXIST = 1109    // 分类已存在
	ERROR_DES_CATEGORY_EXIST  = "分类已存在" // 分类已存在

)
