package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ImageToUser struct {
	ImageUuid string `orm:"pk"`
	UserId    int
}

func InsertImageToUser(o orm.Ormer, i ImageToUser) (err error) {
	_, err = o.Insert(&i)
	if err != nil {
		beego.Error(err)
	}
	return
}

func init() {
	orm.RegisterModel(new(ImageToUser))
}
