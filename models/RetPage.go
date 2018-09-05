package models

type RetPage struct {
	Count       int         `json:"count"`
	TotalPages  int         `json:"totalPages"`
	Pagesize    int         `json:"pagesize"`
	CurrentPage int         `json:"currentPage"`
	Data        interface{} `json:"data"`
}
