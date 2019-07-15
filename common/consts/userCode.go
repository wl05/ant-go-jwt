package consts

const (
	ERROR_CODE_USER_EXIST = 2000     // 用户名已存在
	ERROR_DES_USER_EXIST  = "用户名已存在" // 用户名已存在

	ERROR_CODE_USER_NOT_EXIST = 2001     // 用户名不存在
	ERROR_DES_USER_NOT_EXIST  = "用户不存在" // 用户名不存在

	ERROR_CODE_USERNAME_OR_PASSWORD_ERROR = 2002        // 用户名或者密码错误
	ERROR_DES_USERNAME_OR_PASSWORD_ERROR  = "用户名或者密码错误" // 用户名或者密码错误

	ERROR_CODE_LOGIN_ERROR = 2003    // 请重新登录
	ERROR_DES_LOGIN_ERROR  = "请重新登录" // 请重新登录

	ERROR_CODE_LOGIN_EXPIRED_ERROR = 2004   // 登录过期
	ERROR_DES_LOGIN_EXPIRED_ERROR  = "登录过期" // 登录过期

	ERROR_CODE_EMAIL_EXIST = 2005    // 邮箱已存在
	ERROR_DES_EMAIL_EXIST  = "邮箱已存在" // 邮箱已存在
)
