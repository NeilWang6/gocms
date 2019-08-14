package enums

type JsonResultCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode200 = 200 //成功
	JRCode201 = 201 //没有对应结果
	JRCode302 = 302 //跳转至地址
	JRCode400 = 400 //请求参数有误
	JRCode401 = 401 //未授权访问
	JRCode402 = 402 //token超时
	JRCode501 = 501 //系统出错
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)
