package controllers

import (
	"app/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) Save() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		beego.Error(err)
	}
	imagePath, originalFilename, uuid := c.saveFormFileImageToS3("avatar")
	if imagePath != "" {
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
	}
	oldUser, err := models.FindUserById(o, id)
	if err != nil {
		err := errors.New("User doesnt exist")
		beego.Error(err)
		c.Redirect("/404", 404)
	}

	// if image is not uploaded - use old one
	if imagePath == "" {
		imagePath = oldUser.Avatar
	}
	u := models.User{
		Id:            id,
		Email:         c.GetString("email"),
		FirstName:     c.GetString("first_name"),
		LastName:      c.GetString("last_name"),
		Password:      c.GetString("password"),
		About:         c.GetString("about"),
		PinterestLink: c.GetString("pinterest_link"),
		InstagramLink: c.GetString("instagram_link"),
		FacebookLink:  c.GetString("facebook_link"),
		Phone:         c.GetString("phone"),
		Skype:         c.GetString("skype"),
		Telegram:      c.GetString("telegram"),
		Avatar:        imagePath,
	}
	userModelValidation := models.ValidateUserModelOnUpdate(o, u)
	if userModelValidation.HasErrors() {
		c.Data["ValidationErrors"] = userModelValidation.Errors
		err = o.Rollback()
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.UpdateUser(o, u)
		if err != nil {
			beego.Error(err)
			err = o.Rollback()
			if err != nil {
				beego.Error(err)
			}
		} else {
			err = o.Commit()
			if err != nil {
				beego.Error(err)
			}
		}
	}
	c.Redirect("/profile/update?id="+strconv.Itoa(id), 302)
}

func (c *ProfileController) Update() {
	o := orm.NewOrm()
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
	}
	user, err := models.FindUserById(o, id)
	fmt.Printf("%+v", err)
	fmt.Printf("%+v", user)
	if err == orm.ErrNoRows {
		err := errors.New("User doesnt exist")
		beego.Error(err)
		c.Redirect("/404", 404)
	}
	title := fmt.Sprintf("Edit Profile: %s %s", user.FirstName, user.LastName)
	c.setRenderData(title, "profile/update")
	c.Data["User"] = user

}
