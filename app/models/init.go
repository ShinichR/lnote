package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/go-mgo/mgo"
	"github.com/go-mgo/mgo/bson"
	"net/url"
)
type Base struct {
	session *mgo.Session
}
const (
	MONGOSERVER = "127.0.0.1:27017"
)

func (this *Base )Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "27017"
	}
	
	this.session, err := mgo.Dial(MONGOSERVER) //连接数据库
	this.session.SetMode(mgo.Monotonic, true)
	db := this.session.DB("note")  //数据库名称
	
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
func (this *Base )Finish() {
	defer func() {
        if this.session != nil {
           // mongo.CloseSession(this.UserId, this.session)
            this.session = nil
        }
    }()

}