package consts

const (
	SUCCECC = 200
	ERROR   = 0

	ERROR_CODE_PARAMETER_ILLEGAL = 1000   // 参数错误
	ERROR_DES_PARAMETER_ILLEGAL  = "参数错误" // 参数错误

	ERROR_CODE_REQUEST = 1001   // 请求出错
	ERROR_DES_REQUEST  = "请求出错" // 请求出错

	ERROR_CODE_REFRESH_TOKEN = 1002                 // 请使用新的accessToken请求
	ERROR_DES_REFRESH_TOKEN  = "请使用新的accessToken请求" // 请求出错
)
