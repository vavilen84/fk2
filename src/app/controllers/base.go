package controllers

import (
	"app/auth"
	"app/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) setAuthData() {
	token := auth.GetToken(c.Ctx)
	fmt.Printf("%+v", token)
	c.Data["IsLoggedIn"] = token.IsLoggedIn
	if token.User != "" {
		var user models.User
		err := json.Unmarshal([]byte(token.User), &user)
		if err != nil {
			beego.Error(err)
		}
		c.Data["User"] = user
		if user.Role == models.RoleAdmin {
			c.Data["IsAdmin"] = true
		} else {
			c.Data["IsAdmin"] = false
		}
	}
}

func (c *BaseController) setResponseData(title, templateName string) {
	c.setRenderData(title, templateName)
	c.setAuthData()
}

func (c *BaseController) setRenderData(title, templateName string) {
	c.Data["title"] = title
	c.Layout = "layout.html"
	c.TplName = templateName + ".html"
}
