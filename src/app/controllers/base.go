package controllers

import (
	"app/models/auth"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) setAuthData() {
	isLoggedIn, token := auth.ValidateAuth(c.Ctx)
	c.Data["IsLoggedIn"] = isLoggedIn
	if token.JWT != nil {
		c.Data["UserId"] = token.JWT.ID
	} else {
		c.Data["UserId"] = ""
	}
}

func (c *BaseController) setRenderData(title, templateName string) {
	c.Data["title"] = title
	c.Layout = "layout.html"
	c.TplName = templateName + ".html"
}
