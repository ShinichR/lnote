package models

import (
	"fmt"
	"time"
    "github.com/astaxie/beego"
    "github.com/go-mgo/mgo/bson"
)

type Note struct {
	Id           int
    User         string
	Words        string
	CreateTime   time.Time    
}

func (t *Note) Update(fields ...string) error {
	
	return nil
}

func NoteAdd(note *Note) (int64, error) {
	if note.Words == "" {
		return 0, fmt.Errorf("内容不能为空")
	}
	mConn := Conn()
    defer mConn.Close()
    var code int64
    c := mConn.DB(Dbname()).C(Dbname())
	err := c.Insert(note)
    if err != nil {
		code = 1
	} else {
		code = 0
	}
    beego.Error("note add :",code,err)
    return code ,err
}

func  NoteGetList(u string,page, pageSize int) ([]*Note, int64) {
	notes := make([]*Note, 0)
	mConn := Conn()
    defer mConn.Close()
    c := mConn.DB(Dbname()).C(Dbname())
    err := c.Find(&bson.M{"user":u}).Sort("createtime").All(&notes)
    var code int64
    if err != nil {
		code = 1
        beego.Error("note GetList :",notes,err)
	} else {
		code = 0
	}
   
    for i, m := range notes {
        beego.Debug("note:",m.User, i, m.Words)
    }
	return notes, code
}

func  NoteGetOne(id int) (Note, int64) {
	var note Note
	mConn := Conn()
    defer mConn.Close()
    c := mConn.DB(Dbname()).C(Dbname())
    err := c.Find(&bson.M{"id":id}).One(&note)
    var code int64
    if err != nil {
		code = 1
        beego.Error("note NoteGetOne :",note,err)
	} else {
		code = 0
	}
   
	return note, code
}
func NoteEdit(id int,words string) error {
    mConn := Conn()
    defer mConn.Close()
    c := mConn.DB(Dbname()).C(Dbname())
    err := c.Update(bson.M{"id": id}, bson.M{"$set": bson.M{"words":words }})
    if err != nil{
         beego.Error("NoteEdit err,",err)
    }
	return err
}

func NoteDel(id int,name string) error {
    mConn := Conn()
    defer mConn.Close()
    c := mConn.DB(Dbname()).C(Dbname())
    err := c.Remove(&bson.M{"id": id,"user":name})
    if err != nil{
         beego.Error("notedel err,",err)
    }
	return err
}
