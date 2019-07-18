package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Image struct {
	Uuid             string `orm:"pk"`
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

func init() {
	orm.RegisterModel(new(Image))
}
