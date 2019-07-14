package controllers

import (
	"app/models"
	"app/models/auth"
	"app/models/user"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	o := orm.NewOrm()
	c.Data["title"] = "Login"
	c.Layout = "layout.html"
	c.TplName = "auth/login.html"
	isLoggedIn, _ := auth.ValidateAuth(c.Ctx)
	c.Data["IsLoggedIn"] = isLoggedIn
	if c.Ctx.Input.IsPost() {
		email := c.GetString("email")
		password := c.GetString("password")
		m := models.Login{Email: email, Password: password}
		loginModelValidation := auth.ValidateLoginModel(&m)
		if loginModelValidation.HasErrors() {
			c.Data["ValidationErrors"] = loginModelValidation.Errors
		} else {
			u, _ := user.FindByEmail(m.Email, o)
			auth.LoginHandler(u, c.Ctx)
			c.Redirect("/", 302)
		}
	}
}

func (c *AuthController) Logout() {
	auth.Logout(c.Ctx)
	c.Redirect("/", 302)
}

func (c *AuthController) Register() {
	c.Data["title"] = "Register"
	c.Layout = "layout.html"
	c.TplName = "auth/register.html"
	isLoggedIn, _ := auth.ValidateAuth(c.Ctx)
	c.Data["IsLoggedIn"] = isLoggedIn
	c.Data["ValidationErrors"] = make([]*validation.Error, 0)
	if c.Ctx.Input.IsPost() {
		email := c.GetString("email")
		password := c.GetString("password")
		m := models.User{Email: email, Password: password}
		userModelValidation := auth.ValidateUserModel(&m)
		userModelValidation = auth.ValidateUserModelOnRegister(&m, userModelValidation)
		if userModelValidation.HasErrors() {
			c.Data["ValidationErrors"] = userModelValidation.Errors
		} else {
			or := orm.NewOrm()
			auth.CreateUser(&m, or)
			c.Redirect("/auth/login", 302)
		}
	}
}
