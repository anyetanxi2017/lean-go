package utils

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{}
}
func (r *Result) SetCode(code int) *Result {
	r.Code = code
	return r
}

func (r *Result) SetMsg(msg string) *Result {
	r.Msg = msg
	return r
}
func (r *Result) SetData(data interface{}) *Result {
	r.Data = data
	return r
}
