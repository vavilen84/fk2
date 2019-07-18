package controllers

import (
	"app/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type GalleryController struct {
	BaseController
}

func (c *GalleryController) Student() {
	c.setResponseData("Ученики", "gallery/users")
	o := orm.NewOrm()
	list, err := models.FindUserListByType(o, models.TypeStudent)
	if err != nil {
		beego.Error(err)
	}
	c.Data["UserList"] = list
}

func (c *GalleryController) Graduate() {
	c.setResponseData("Выпускники", "gallery/users")
	o := orm.NewOrm()
	list, err := models.FindUserListByType(o, models.TypeGraduate)
	if err != nil {
		beego.Error(err)
	}
	c.Data["UserList"] = list
}
