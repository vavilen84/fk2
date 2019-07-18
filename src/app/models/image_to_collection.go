package models

import "github.com/astaxie/beego/orm"

type ImageToCollection struct {
	ImageUuid    string `orm:"pk"`
	CollectionId int
}

func init() {
	orm.RegisterModel(new(ImageToCollection))
}
