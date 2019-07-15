package controllers

import (
	"app/models"
	"app/models/auth"
	"app/models/user"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) Save() {
	c.setAuthData()
	email := c.GetString("email")
	password := c.GetString("password")
	firstName := c.GetString("first_name")
	lastName := c.GetString("last_name")
	id, err := c.GetInt("id")
	if err != nil {
		beego.Error(err)
	}
	u := models.User{
		Id:        int64(id),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	}
	o := orm.NewOrm()
	userModelValidation := auth.ValidateUserModelOnUpdate(&u, o)
	if userModelValidation.HasErrors() {
		c.Data["ValidationErrors"] = userModelValidation.Errors
	} else {
		or := orm.NewOrm()
		_, err := auth.UpdateUser(&u, or)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Redirect("/", 302)
}

func (c *ProfileController) Update() {
	c.setAuthData()
	or := orm.NewOrm()
	idParam := c.GetString("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		beego.Error(err)
	}
	u, _ := user.OneById(int64(id), or)
	if u.Email == "" {
		err := errors.New("User not exists")
		beego.Error(err)
		c.Redirect("/", 302)
	}
	c.Data["User"] = u
	c.Data["title"] = "Edit Profile: " + u.FirstName + " " + u.LastName
	c.Layout = "layout.html"
	c.TplName = "profile/update.html"
}
