package models

type RetMsg struct {
	ErrNo  int         `json:"errno"`
	ErrMsg interface{} `json:"errmsg"`
	Data   interface{} `json:"data"`
}