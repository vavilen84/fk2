package controllers

import (
	"app/models/post"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	c.setAuthData()
	or := orm.NewOrm()
	posts, _ := post.FindAll(or)
	c.Data["Posts"] = posts
	c.Layout = "layout.html"
	c.TplName = "index.html"
}
