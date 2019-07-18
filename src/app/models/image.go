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

func FindImageListByUser(o orm.Ormer, user User) (list []Image, err error) {
	sql := "SELECT image.* FROM image " +
		"INNER JOIN image_to_user ON image.uuid = image_to_user.image_uuid " +
		"WHERE image_to_user.user_id = ?"
	_, err = o.Raw(sql, user.Id).QueryRows(&list)
	return
}

func init() {
	orm.RegisterModel(new(Image))
}
