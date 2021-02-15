// Code generated by goctl. DO NOT EDIT.
package types

type ListReq struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type InfoReq struct {
	Id int64 `json:"id"`
}

type InfoRes struct {
	Item interface{} `json:"item"`
	List interface{} `json:"list"`
}

type ListRes struct {
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	Page      int         `json:"page"`
	TotalPage int         `json:"totalPage"`
	Limit     int         `json:"limit"`
}
