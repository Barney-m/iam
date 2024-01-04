package res

type CommonRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok(data interface{}) *CommonRes {
	return &CommonRes{Code: StatusOk, Msg: "Success", Data: data}
}

func NotOk(errMsg string) *CommonRes {
	return &CommonRes{Code: StatusNotOk, Msg: errMsg, Data: new(map[string]interface{})}
}
