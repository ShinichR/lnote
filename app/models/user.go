package models

import (
	"fmt"
    "github.com/astaxie/beego"
    "github.com/go-mgo/mgo/bson"
    "errors"
)

type User struct {
	Id        int
	UserName  string
	Password  string
	Salt      string
	Email     string
	LastLogin int64
	LastIp    string
	Status    int
}

func UserAdd(user *User) (int64, error) {
	if user.Password == "" {
		return 0, fmt.Errorf("密码不能为空")
	}
    _,geterr := UserGetByName(user.UserName)
    if geterr == 0{
         beego.Error("UserAdd err,has same name user :",user.UserName)
         return 1 ,errors.New("has same name user")
    }
	mConn := Conn()
    defer mConn.Close()
    var code int64
    c := mConn.DB(Dbname()).C(Userdbname())
	err := c.Insert(user)
    if err != nil {
		code = 1
	} else {
		code = 0
	}
    beego.Error("UserAdd add :",user,err)
    return code ,err
}

func UserGetById(id int) (*User, int64) {
	user := new(User)
	mConn := Conn()
    defer mConn.Close()
    c := mConn.DB(Dbname()).C(Userdbname())
    err := c.Find(&bson.M{"id":id}).One(user)
    var code int64
    if err != nil {
		code = 1
        beego.Error("note UserGetById :",user,err)
	} else {
		code = 0
	}
	return user, code
}

func UserGetByName(userName string) (*User, int64) {
    user := new(User)
	mConn := Conn()
    defer mConn.Close()
    c := mConn.DB(Dbname()).C(Userdbname())
    err := c.Find(&bson.M{"username":userName}).One(user)
    var code int64
    if err != nil {
		code = 1
        beego.Error("note UserGetByName :",user,err)
	} else {
		code = 0
	}
	return user, code
}

func UserUpdate(user *User, fields ...string) error {

	return nil
}
