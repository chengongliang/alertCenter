package model

// ReqBody 通用请求体
type ReqBody struct {
	Mode  string                 `json:"mode"`
	Query map[string]interface{} `json:"query"`
}
