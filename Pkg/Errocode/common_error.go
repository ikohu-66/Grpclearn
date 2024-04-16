package Errocode

// 可根据请求服务端api自定义错误
var (
	Sucess        = NewError("0", "成功")
	Fail          = NewError("10000000", "内部错误")
	InvalidParams = NewError("10000001", "无效参数")
	Unauthorized  = NewError("10000002", "认证错误")
	NotFound      = NewError("10000003", "没有找到")
)
