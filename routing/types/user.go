package types

type SengMsg struct {
	Msg   string `json:"msg,omitempty"` //内容
	Phone string `json:"phone"`         //手机号
	Area  string `json:"area"`          //地区
}
