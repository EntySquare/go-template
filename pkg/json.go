package pkg

import (
	"github.com/sirupsen/logrus"
	"go-template/log"
)

const (
	CodeOk       = 0  // 成功
	CodeErr      = -1 // 失败
	CodeErrToken = -2 // token相关的异常
)

// JSONResponse represents an HTTP response which contains a JSON body.
type JSONResponse struct {
	// HTTP status code.
	Code int `json:"code"`
	// JSON represents the JSON that should be serialized and sent to the client
	Data interface{} `json:"data"`
}

func SuccessResponse(data interface{}) JSONResponse {
	return JSONResponse{
		Code: 0,
		Data: data,
	}
}

// MessageResponse returns a JSONResponse with a 'message' key containing the given text.
func MessageResponse(code int, msg, msgZh string) JSONResponse {
	log.Log.WithFields(logrus.Fields{
		"code":   code,
		"msg_zh": msgZh,
	}).Warnf(msg)
	return JSONResponse{
		Code: code,
		Data: struct {
			Message   string `json:"message"`
			MessageZh string `json:"message_zh"`
		}{msg, msgZh},
	}
}
