package controllers

import (
	"app/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type PostController struct {
	BaseController
}

func (c *PostController) EditList() {
	c.setResponseData("Posts", "post/edit-list")
	o := orm.NewOrm()
	posts, _ := models.FindAllPosts(o)
	c.Data["Posts"] = posts
}

func (c *PostController) Create() {
	c.setResponseData("Create New Post", "post/create")
}

func (c *PostController) Save() {
	o := orm.NewOrm()
	post := models.Post{
		Title:   c.GetString("title"),
		Content: c.GetString("content"),
	}
	err := models.InsertPost(o, post)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)
}

func (c *PostController) Update() {
	o := orm.NewOrm()
	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		beego.Error(err)
	}
	post := models.Post{
		Id:      id,
		Title:   c.GetString("title"),
		Content: c.GetString("content"),
	}
	fmt.Printf("%+v", post)
	err = models.UpdatePost(o, post)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)
}

func (c *PostController) Delete() {
	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()
	err = models.DeleletePost(o, id)
	c.Redirect("/", 302)
}

func (c *PostController) Edit() {
	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()
	post, err := models.FindPostById(o, id)
	if err == orm.ErrNoRows {
		c.Redirect("/404", 302)
	}
	title := fmt.Sprintf("Edit Post #%s", c.GetString("id"))
	c.setRenderData(title, "post/edit")
	c.Data["Post"] = post
}
