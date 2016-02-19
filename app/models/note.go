package models

import (
	"fmt"
	"time"
)

type Note struct {
	Id           int
	Words      string
	CreateTime   int64
	//base        Base        
}

func (t *Note) Update(fields ...string) error {
	
	return nil
}

func NoteAdd(note *Note) (int64, error) {
	if note.Words == "" {
		return 0, fmt.Errorf("内容不能为空")
	}
	if note.CreateTime == 0 {
		note.CreateTime = time.Now().Unix()
	}
	return 1,nil
	//return orm.NewOrm().Insert(note)
}

func NoteGetList(page, pageSize int) ([]*Note, int64) {
	notes := make([]*Note, 0)
	
	return notes, 0
}


func NoteDel(id int) error {
	return nil
}
