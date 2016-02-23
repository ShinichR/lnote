package controllers

import (
	"time"
    "github.com/astaxie/beego"
    "github.com/ShinichR/lnote/app/models"
    "github.com/ShinichR/lnote/app/libs"
    "strings"
    "strconv"
)

type MainController struct {
	BaseController
}


func (this *MainController) Index() {
	this.Data["pageTitle"] = "note"
    uname := this.Ctx.GetCookie("user")
    if uname == ""{
        beego.Debug("redirect to login")
        this.redirect(beego.UrlFor("MainController.Login"))
    }
    beego.Debug("get user list:",uname)
    
    result ,_:= models.NoteGetList(uname,1, 10)
	this.Data["notes"] = result
    beego.Error("note list result:",result)
	this.Layout = "layout/layout.html"
	this.TplNames = "main/index.html"
}

// 获取系统时间
func (this *MainController) GetTime() {
	out := make(map[string]interface{})
	out["time"] = time.Now().UnixNano() / int64(time.Millisecond)
	this.jsonResult(out)
}
func (this *MainController) Login() {
	if this.userId > 0 {
		this.redirect("/")
	}
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		flash := beego.NewFlash()

		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
        beego.Debug("login:",username)
		remember := this.GetString("remember")
		if username != "" && password != "" {
			user, err := models.UserGetByName(username)
			errorMsg := ""
			if err != 0 || user.Password != libs.Md5([]byte(password+user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = this.getClientIp()
				user.LastLogin = time.Now().Unix()
				models.UserUpdate(user)

				authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))
				if remember == "yes" {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
				} else {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey)
                    
				}
                this.Ctx.SetCookie("user", user.UserName)

				this.redirect(beego.UrlFor("NoteController.List"))
			}
			flash.Error(errorMsg)
			flash.Store(&this.Controller)
            beego.Debug("errorMsg:",errorMsg)
			this.redirect(beego.UrlFor("MainController.Login"))
		}
	}

	this.TplNames = "main/login.html"
}

// 退出登录
func (this *MainController) Logout() {
	this.Ctx.SetCookie("auth", "")
    this.Ctx.SetCookie("user", "")
	this.redirect(beego.UrlFor("MainController.Login"))
}

