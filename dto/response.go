package dto

type ResponseObject struct {
	Code  uint64      `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Total int         `json:"total,omitempty"`
}
