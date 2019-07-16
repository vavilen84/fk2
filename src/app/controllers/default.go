package controllers

import (
	"app/models"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	c.setResponseData("Home", "index")
	o := orm.NewOrm()
	posts, _ := models.FindAllPosts(o)
	c.Data["Posts"] = posts
}

func (c *MainController) PageNotFound() {
	c.setResponseData("Resource Not Found :(", "404")
	c.Ctx.Output.Status = 404
}
