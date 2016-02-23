package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/go-mgo/mgo"
    	_"github.com/go-mgo/mgo/bson"
    

)
const (
	MONGOSERVER = "shinichr.eicp.net:28190"
)

var session *mgo.Session
var dbname  string
var userdbname  string
var id      int

func Conn() *mgo.Session {
	return session.Copy()
}

func Dbname() string {
    if dbname == ""{
        dbname =  beego.AppConfig.String("db.name") 
    }
	return dbname
}

func Userdbname() string {
    if userdbname == ""{
        userdbname =  "user" 
    }
	return userdbname
}

func Maxid() int{
    var note Note
    c := Conn().DB(Dbname()).C(Dbname())
    num,_ := c.Count()
    if num == 0 {
	return 0	
    }
    c.Find(nil).Sort("-id").One(&note)
    beego.Debug("maxid: ",note.Id)

    return note.Id	
}
func MaxUserid() int{
    var user User
    c := Conn().DB(Dbname()).C(Userdbname())
    num,_ := c.Count()
    if num == 0 {
	   return 0	
    }
    c.Find(nil).Skip(num - 1).One(&user)
    beego.Debug("MaxUserid: ",user.Id)

    return user.Id	
}

func init() {
	//dbhost := beego.AppConfig.String("db.host")
	//dbport := beego.AppConfig.String("db.port")
	//dbuser := beego.AppConfig.String("db.user")
	//dbpassword := beego.AppConfig.String("db.password")
	//dbname := beego.AppConfig.String("db.name")
	//timezone := beego.AppConfig.String("db.timezone")
	/*if dbport == "" {
		dbport = "27017"
	}*/
	
	sess, err := mgo.Dial(MONGOSERVER) 
	if err != nil {
        beego.Error("err in mongo dial")
		panic(err)
	}
    session = sess
	session.SetMode(mgo.Monotonic, true)
	//db := this.session.DB("note") 
	
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func Finish() {
	

}
