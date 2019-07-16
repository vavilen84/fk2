package controllers

import (
	"app/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) Save() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
	}
	u := models.User{
		Id:        id,
		Email:     c.GetString("email"),
		FirstName: c.GetString("first_name"),
		LastName:  c.GetString("last_name"),
		Password:  c.GetString("password"),
	}
	o := orm.NewOrm()
	userModelValidation := models.ValidateUserModelOnUpdate(o, u)
	if userModelValidation.HasErrors() {
		c.Data["ValidationErrors"] = userModelValidation.Errors
	} else {
		err := models.UpdateUser(o, u)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Redirect("/", 302)
}

func (c *ProfileController) Update() {
	o := orm.NewOrm()
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
	}
	user, err := models.FindUserById(o, id)
	if err == orm.ErrNoRows {
		err := errors.New("User not exists")
		beego.Error(err)
		c.Redirect("/", 302)
	}
	title := fmt.Sprintf("Edit Profile: %s %s", user.FirstName, user.LastName)
	c.setRenderData(title, "profile/update")
	c.Data["User"] = user
}
