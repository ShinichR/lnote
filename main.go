package main

import (
	"github.com/astaxie/beego"
	"github.com/ShinichR/lnote/app/controllers"
	"html/template"
	"net/http"
)

const VERSION = "1.0.0"

func main() {
	// 设置默认404页面
	beego.Errorhandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	// 生产环境不输出debug日志
	if beego.AppConfig.String("runmode") == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}
	beego.AppConfig.Set("version", VERSION)

	// 路由设置
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
    beego.Router("/Signup", &controllers.UserController{}, "*:Signup")
   // beego.Router("/Signup", &controllers.UserController{}, "*:Signup")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.AutoRouter(&controllers.NoteController{})
   // beego.AutoRouter(&controllers.UserController{})
	beego.SessionOn = true
	beego.Run()
}
