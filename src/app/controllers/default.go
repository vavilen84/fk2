package controllers

import (
	"app/models/auth"
	"app/models/post"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["IsLoggedIn"] = auth.ValidateAuth(c.Ctx)
	or := orm.NewOrm()
	posts, _ := post.FindAll(or)
	c.Data["Posts"] = posts
	c.Layout = "layout.html"
	c.TplName = "index.html"
}
