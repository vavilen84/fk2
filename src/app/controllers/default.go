package controllers

import (
	"app/models"
	"app/settings"
	"app/utils"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
)

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	c.setResponseData("Home", "index")
	o := orm.NewOrm()

	postsPerPage := settings.PostsPerPage
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, models.CountPosts(o))

	posts, _ := models.ListPostsByOffsetAndLimit(o, paginator.Offset(), postsPerPage)
	c.Data["Posts"] = utils.GetPostOnViewList(o, posts)

}

func (c *MainController) PageNotFound() {
	c.setResponseData("404 - Page Not Found :(", "404")
}
