package controllers

import (
	"app/models/auth"
	"app/models/post"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	isLoggedIn, token := auth.ValidateAuth(c.Ctx)
	fmt.Printf("%+v", token.JWT.ID)

	c.Data["IsLoggedIn"] = isLoggedIn
	c.Data["UserId"] = token.JWT.ID
	or := orm.NewOrm()
	posts, _ := post.FindAll(or)
	c.Data["Posts"] = posts
	c.Layout = "layout.html"
	c.TplName = "index.html"
}
