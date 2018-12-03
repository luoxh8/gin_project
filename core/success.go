package core

var SuccessCode = 200

type ServerSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

var (
	SuccessCodeSend = &ServerSuccess{200, "送出成功"}
)
