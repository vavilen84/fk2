package models

import "github.com/astaxie/beego/orm"

type Collection struct {
	Id     int
	Name   string
	Status int
}

func init() {
	orm.RegisterModel(new(Collection))
}
