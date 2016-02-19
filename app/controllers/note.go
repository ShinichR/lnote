package controllers

import (
	//"github.com/astaxie/beego"
	"github.com/shinichr/lnote/app/models"
	//"strconv"
	//"strings"
	//"time"
)

type NoteController struct {
	BaseController
}

func (this *NoteController) List() {
	result, count := models.NoteGetList(1, 10)
	this.Data["notes"] = result
	this.display()
}


func (this *NoteController) Add() {

	if this.isPost() {
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加笔记"
	this.display()
}
