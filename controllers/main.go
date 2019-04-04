package controllers

import (
	"github.com/astaxie/beego"
)

// MainController operations for Main
type MainController struct {
	beego.Controller
}

//Response 结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//Response 结构体
type ErrResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type Page struct {
	NowPage  int `json:"page"`
	PageSize int `json:"size"`
	Total    int `json:"total"`
}

type Result struct {
	Item     interface{} `json:"item"`
	PageItem Page        `json:"page_item"`
}
