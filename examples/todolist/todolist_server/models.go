package main

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	pb "google.golang.org/grpc/examples/todolist/todolist"
)

type TodoInfo struct {
	Id           int64
	Time         string
	Action       string
	DetailAction string    `orm:"type(text)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}

func ConvertPbTodo2TodoInfo(pbTodo *pb.Todo) *TodoInfo {
	return &TodoInfo{
		Id:           pbTodo.Id,
		Time:         pbTodo.Time,
		Action:       pbTodo.Action,
		DetailAction: pbTodo.DetailAction,
	}
}

func ConvertTodoInfo2PbTodo(ti *TodoInfo) *pb.Todo {
	return &pb.Todo{
		Id:           ti.Id,
		Time:         ti.Time,
		Action:       ti.Action,
		DetailAction: ti.DetailAction,
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

func GetAll() (*TodoInfo, error) {
	o := orm.NewOrm()
	ti := &TodoInfo{}
	err := o.Read(ti)
	if err != nil {
		log.Printf("Get todo %+v err %v\n", ti, err)
		return nil, err
	}

	return ti, nil
}

func Get(id int64) (*TodoInfo, error) {
	o := orm.NewOrm()
	ti := &TodoInfo{
		Id: id,
	}
	err := o.Read(ti)
	if err != nil {
		log.Printf("Get todo %+v err %v\n", ti, err)
		return nil, err
	}

	return ti, nil
}

func (t *TodoInfo) Update() error {
	o := orm.NewOrm()

	num, err := o.Update(t, "Time", "Action", "detailAction")
	if err != nil {
		log.Printf("Update %+v err %v\n", t, err)
		return err
	}
	log.Printf("update todo %+v, affect %d row\n", t, num)
	return nil
}
