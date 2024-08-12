package controller

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var RespErrASNInvalid = Resp{
	Code: -1,
	Msg:  "invalid ASN",
}
