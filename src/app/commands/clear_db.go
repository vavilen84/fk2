package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func clearDb() {
	o := orm.NewOrm()
	_, err := o.Raw("TRUNCATE TABLE user").Exec()
	if err != nil {
		beego.Error(err)
	}
	_, err = o.Raw("TRUNCATE TABLE post").Exec()
	if err != nil {
		beego.Error(err)
	}
}

func init() {
	Cmd.Run(clearDb)
}
