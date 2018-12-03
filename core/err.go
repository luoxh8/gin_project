package core

type ServerError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

var (
	/* 0 - 500 请求错误*/
	ErrSysBadRequest   = &ServerError{400, "Bad Request"}
	ErrSysPageNotFound = &ServerError{404, "找不到该路由"}
	ErrSysServer       = &ServerError{500, "服务器内部错误"}
	ErrSysSocket       = &ServerError{600, "Socket连接错误"}

	/* 400-500 Base错误 */
	ErrBaseParams            = &ServerError{401, "参数错误"}
	ErrBaseInputData         = &ServerError{402, "数据输入错误"}
	ErrBaseDatabase          = &ServerError{403, "数据库错误"}
	ErrBaseOpenFile          = &ServerError{405, "服务器内部错误"}
	ErrBaseWriteFile         = &ServerError{406, "服务器内部错误"}
	ErrBaseIllegalPermission = &ServerError{407, "非法权限"}
	ErrBaseParmsFormat       = &ServerError{408, "参数传递不符合规范"}
	ErrBaseDataNotFound      = &ServerError{409, "找不到该数据"}
	ErrBaseFrequentRequests  = &ServerError{410, "频率超过了限制，请休息片刻在请求"}
	ErrBaseUnknown           = &ServerError{411, "未知错误"}

	/* 500 -600 用户错误 */
	ErrUserNotLogin         = &ServerError{501, "没有登录"}
	ErrUserAlreadyExists    = &ServerError{502, "用户已存在"}
	ErrUserDoesNotExist     = &ServerError{503, "用户不存在"}
	ErrUserTokenCreate      = &ServerError{504, "生成token失败"}
	ErrUserNoToken          = &ServerError{505, "非法token"}
	ErrUserTokenExpired     = &ServerError{506, "token过期"}
	ErrUserTokenNotValidYet = &ServerError{507, "token无法通过验证"}
	ErrUserTokenMalformed   = &ServerError{507, "token格式错误"}
	ErrUserTokenInvalid     = &ServerError{508, "token无效"}
	ErrUserPassword         = &ServerError{509, "密码错误"}
	ErrUserBalance          = &ServerError{510, "余额不足"}
	ErrUserUnusual          = &ServerError{511, "账户异常"}

	// 商城对应的错误码
	ErrGoodsUnusual = &ServerError{601, "商品当前价格更新异常"}
	ErrGoodsBid     = &ServerError{602, "正在出价，请勿重复操作"}
)
