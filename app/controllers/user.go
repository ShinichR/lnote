package controllers

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/utils"
	"github.com/ShinichR/lnote/app/models"
    "github.com/ShinichR/lnote/app/libs"
	//"strings"
	"time"
)

type UserController struct {
	BaseController
}


func (this *UserController) Signup() {
    beego.ReadFromRequest(&this.Controller)
    beego.Debug("Signup in")
	if this.isPost() {
        flash := beego.NewFlash()
        password1 := this.GetString("password")
		password2 := this.GetString("password2")
        beego.Debug("pw1,pw2",password1,password2)
        if password1 != "" {
            if len(password1) < 6 {
				flash.Error("密码长度必须大于6位")
				flash.Store(&this.Controller)
                
				this.redirect(beego.UrlFor("MainController.Login"))
			} else if password2 != password1 {
				flash.Error("两次输入的密码不一致")
				flash.Store(&this.Controller)
				this.redirect(beego.UrlFor("MainController.Login"))
            }else{
                user := new(models.User)
                user.Id = models.MaxUserid() + 1 
                user.LastIp = this.getClientIp()
                user.UserName =this.GetString("username")
                user.Password = this.GetString("Password")
                user.LastLogin = time.Now().Unix()
                user.Salt = string(utils.RandomCreateBytes(10))
                user.Password = libs.Md5([]byte(password1 + user.Salt))
                
                 beego.Debug("add user:",user.UserName)
                 models.UserAdd(user)
                 flash.Success("注册成功！")
	             flash.Store(&this.Controller)
        
                 this.Data["user"] = user
                 this.redirect(beego.UrlFor("MainController.Login"))
            }
        }
        
        flash.Success("注册失败！")
	    flash.Store(&this.Controller)
      //  this.ajaxMsg("", MSG_OK)
	}
   	
	this.Data["pageTitle"] = "注册用户"
    this.TplNames = "main/login.html"
}

func (this *UserController) Edit() {

	this.Data["pageTitle"] = "编辑信息"
	this.display()
    
}