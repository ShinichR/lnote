package main

import (
	"fmt"
	"github.com/go-mgo/mgo"
	_"gopkg.in/mgo.v2/bson"
	_"github.com/go-mgo/mgo/bson"
	"time"
)

const (
	MONGOSERVER = "127.0.0.1:10000"
)
type Note struct {
	Id           int
    	User         string
	Words        string
	CreateTime   time.Time    
}

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
func main() {

	session, err := mgo.Dial(MONGOSERVER) //l½Óý	if
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	db := session.DB("note")  
	c := db.C("note")
	countNum, err := c.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("Things objects count: ", countNum)
	//var user User
	//notes := make([]*Note, 0)
	var note Note
	//c.Find(nil).Skip(countNum - 2).One(&notes)
	//c.Find(&bson.M{"user":"lei"}).One(&notes)
	c.Find(nil).Sort("-id").One(&note)
	//for _, m := range notes {
      	fmt.Println("note:",note.User, note.Id)
  	//}
        //fmt.Println("note:",user.UserName, user.Password, user.LastLogin,user.Id)

	
	

}
