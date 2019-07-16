package controllers

import (
	"app/auth"
	"app/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Login() {
	c.setResponseData("Login", "auth/login")
	o := orm.NewOrm()
	if c.Ctx.Input.IsPost() {
		loginModel := auth.Login{
			Email:    c.GetString("email"),
			Password: c.GetString("password"),
		}
		v := auth.ValidateLoginModel(loginModel)
		if v.HasErrors() {
			c.Data["ValidationErrors"] = v.Errors
		} else {
			user, _ := models.FindUserByEmail(o, loginModel.Email)
			auth.LoginHandler(user, c.Ctx)
			c.Redirect("/", 302)
		}
	}
}

func (c *AuthController) Logout() {
	auth.Logout(c.Ctx)
	c.Redirect("/", 302)
}

func (c *AuthController) Register() {
	c.setResponseData("Register", "auth/register")
	if c.Ctx.Input.IsPost() {
		m := models.User{
			Email:     c.GetString("email"),
			Password:  c.GetString("password"),
			FirstName: c.GetString("first_name"),
			LastName:  c.GetString("last_name"),
			Role:      models.RoleUser,
		}
		o := orm.NewOrm()
		userModelValidation := models.ValidateUserModelOnRegister(o, m)
		if userModelValidation.HasErrors() {
			c.Data["ValidationErrors"] = userModelValidation.Errors
		} else {
			err := models.InsertUser(o, m)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/auth/login", 302)
		}
	}
}
