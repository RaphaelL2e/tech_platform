package model

type ListModel struct {
	PageNum int `json:"page_num" default:"1"`
	PageSize int `json:"page_size" default:"10"`
}