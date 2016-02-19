package controllers

import (
	"time"
)

type MainController struct {
	BaseController
}


func (this *MainController) Index() {
	this.Data["pageTitle"] = "note"
	this.Layout = "layout/layout.html"
	this.TplNames = "main/index.html"
}

// 获取系统时间
func (this *MainController) GetTime() {
	out := make(map[string]interface{})
	out["time"] = time.Now().UnixNano() / int64(time.Millisecond)
	this.jsonResult(out)
}

