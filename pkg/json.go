package pkg

import (
	"github.com/sirupsen/logrus"
	"go-template/log"
)

const (
	CodeOk       = 0  // 成功
	CodeErr      = -1 // 失败
	CodeErrToken = -2 // token相关的异常
	//CodeReject   = "2" // 拒绝
	//CodeTimeout  = "3" // 超时
)

// JSONResponse represents an HTTP response which contains a JSON body.
type JSONResponse struct {
	// HTTP status code.
	Code int `json:"code"`
	// JSON represents the JSON that should be serialized and sent to the client
	JSON interface{} `json:"json"`
}

func SuccessResponse(data interface{}) JSONResponse {
	return JSONResponse{
		Code: 0,
		JSON: data,
	}
}

// MessageResponse returns a JSONResponse with a 'message' key containing the given text.
func MessageResponse(code int, msg, msgZh string) JSONResponse {
	log.Log.WithFields(logrus.Fields{
		"code":   code,
		"msg_zh": msgZh,
	}).Warnf(msg)
	//Log.Warnf("12312")
	return JSONResponse{
		Code: code,
		JSON: struct {
			Message   string `json:"message"`
			MessageZh string `json:"message_zh"`
		}{msg, msgZh},
	}
}
