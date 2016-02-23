package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ShinichR/lnote/app/models"
	"strconv"
	"strings"
	"time"
)

type NoteController struct {
	BaseController
}

func (this *NoteController) List() {
    uname := this.Ctx.GetCookie("user")
     if uname == ""{
        beego.Debug("redirect to login")
        this.redirect(beego.UrlFor("MainController.Login"))
    }
    beego.Debug("get user list:",uname)
	result ,_:= models.NoteGetList(uname,1, 10)
	this.Data["notes"] = result

	this.display()
}


func (this *NoteController) Add() {

	if this.isPost() {
        note := new(models.Note)
        note.User = this.Ctx.GetCookie("user")
	    note.Id = models.Maxid() + 1 
        note.CreateTime = time.Now()
        note.Words =  strings.TrimSpace(this.GetString("words"))
        beego.Debug("add note:",note.Words)
        models.NoteAdd(note)
        this.ajaxMsg("", MSG_OK)
	}
   
	this.Data["pageTitle"] = "添加笔记"
	this.display()
}
func (this *NoteController) Del() {
    id ,_ := strconv.ParseInt(this.Input().Get("id"),10,64)
    uname := this.Ctx.GetCookie("user")
    beego.Debug("del note",id,uname)
    models.NoteDel(int(id),uname)
    refer := this.Ctx.Request.Referer()
	if refer == "" {
		refer = beego.UrlFor("NoteController.List")
	}
	this.redirect(refer)
    
}
func (this *NoteController) Edit() {
  	id,_ := strconv.ParseInt(this.Input().Get("id"),10,64)
	if this.isPost() {
	  
        words :=  strings.TrimSpace(this.GetString("words"))
        beego.Debug("edit note:",words)
        models.NoteEdit(int(id),words)
        this.ajaxMsg("", MSG_OK)
	}
	note,_ := models.NoteGetOne(int(id))
	this.Data["note"] = note
	this.Data["pageTitle"] = "编辑笔记"
	this.display()
    
}