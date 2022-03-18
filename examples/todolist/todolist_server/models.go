package main

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	pb "google.golang.org/grpc/examples/todolist/todolist"
)

type TodoInfo struct {
	Id           int64 `orm:"auto"`
	Time         string
	Action       string
	DetailAction string    `orm:"type(text)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}

func ConvertPbTodo2TodoInfo(pbTodo *pb.Todo) *TodoInfo {
	return &TodoInfo{
		Time:         pbTodo.Time,
		Action:       pbTodo.Action,
		DetailAction: pbTodo.DetailAction,
	}
}

func (t *TodoInfo) Insert() error {
	o := orm.NewOrm()

	_, err := o.Insert(t)
	if err != nil {
		log.Printf("Insert todo %+v err %v\n", t, err)
		return err
	}

	log.Printf("Insert %+v successfully\n", t)
	return nil
}
