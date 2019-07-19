package controllers

import (
	"app/models"
	"errors"
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
	id, err := c.GetInt("id")
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

func (c *GalleryController) AddImageToPortfolio() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
	}

	o := orm.NewOrm()

	_, err = models.FindUserById(o, id)
	if err != nil {
		err := errors.New("User doesnt exist")
		beego.Error(err)
		c.Redirect("/404", 404)
	}

	imagePath, originalFilename, uuid := c.saveFormFileImageToS3("image")
	if imagePath == "" {
		c.Redirect("/gallery/update-portfolio?id="+strconv.Itoa(id), 302)
	}

	err = o.Begin()
	if err != nil {
		beego.Error(err)
	}

	m := models.Image{
		Uuid:             uuid,
		OriginalFilename: originalFilename,
		Filepath:         imagePath,
	}
	err = models.InsertImage(o, m)
	if err != nil {
		beego.Error(err)
		err = o.Rollback()
		if err != nil {
			beego.Error(err)
		}
	}
	imageToUser := models.ImageToUser{
		ImageUuid: uuid,
		UserId:    id,
	}
	err = models.InsertImageToUser(o, imageToUser)
	if err != nil {
		beego.Error(err)
		err = o.Rollback()
		if err != nil {
			beego.Error(err)
		}
	}

	err = o.Commit()
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/gallery/update-portfolio?id="+strconv.Itoa(id), 302)
}
