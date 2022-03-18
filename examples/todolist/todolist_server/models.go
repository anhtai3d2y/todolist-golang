package main

import "time"

type TodoInfo struct {
	Id      int64  `orm:"auto"`
	Time    string `orm:"size(10)"`
	Action  string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
