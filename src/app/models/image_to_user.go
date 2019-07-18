package models

import "github.com/astaxie/beego/orm"

type ImageToUser struct {
	ImageUuid string `orm:"pk"`
	UserId    int
}

func init() {
	orm.RegisterModel(new(ImageToUser))
}
