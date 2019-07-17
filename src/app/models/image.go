package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Image struct {
	Uuid             string
	OriginalFilename string
	Filepath         string
}

func InsertImage(o orm.Ormer, i Image) (err error) {
	_, err = o.Insert(&i)
	if err != nil {
		beego.Error(err)
	}
	return
}
