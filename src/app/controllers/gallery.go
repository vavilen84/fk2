package controllers

import (
	"app/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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

func (c *GalleryController) User() {
	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()
	user, err := models.FindUserById(o, id)
	if err == orm.ErrNoRows {
		c.Redirect("/404", 302)
	}
	images, err := models.FindImageListByUser(o, user)
	if err != nil {
		beego.Error(err)
	}
	title := user.FirstName + " " + user.LastName
	c.setResponseData(title, "gallery/user")
	c.Data["User"] = user
	c.Data["ImageList"] = images
}

func (c *GalleryController) UpdatePortfolio() {
	id, err := strconv.Atoi(c.GetString("userId"))
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()
	user, err := models.FindUserById(o, id)
	if err == orm.ErrNoRows {
		c.Redirect("/404", 302)
	}
	images, err := models.FindImageListByUser(o, user)
	if err != nil {
		beego.Error(err)
	}
	title := user.FirstName + " " + user.LastName
	c.setResponseData(title, "gallery/update-portfolio")
	c.Data["User"] = user
	c.Data["ImageList"] = images
}
