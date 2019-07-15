package controllers

import (
	"app/models/post"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type PostController struct {
	BaseController
}

func (c *PostController) Create() {
	c.setAuthData()
	c.Data["title"] = "Create New Post"
	c.Layout = "layout.html"
	c.TplName = "post/create.html"
}

func (c *PostController) Save() {
	c.setAuthData()
	title := c.GetString("title")
	content := c.GetString("content")
	or := orm.NewOrm()
	_, err := post.Create(title, content, or)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)
}

func (c *PostController) Update() {
	c.setAuthData()
	id, e := c.GetInt("id")
	if e != nil {
		log.Fatal(e)
	}
	title := c.GetString("title")
	content := c.GetString("content")
	or := orm.NewOrm()
	post.Update(int64(id), title, content, or)
	c.Redirect("/", 302)
}

func (c *PostController) Delete() {
	c.setAuthData()
	id, e := c.GetInt("id")
	if e != nil {
		log.Fatal(e)
	}
	or := orm.NewOrm()
	post.Del(int64(id), or)
	c.Redirect("/", 302)
}

func (c *PostController) Edit() {
	c.setAuthData()
	id, e := c.GetInt(":id")
	if e != nil {
		log.Fatal(e)
	}
	c.Data["title"] = "Edit Post #"
	or := orm.NewOrm()
	p, _ := post.OneById(int64(id), or)
	c.Data["Post"] = p
	c.Layout = "layout.html"
	c.TplName = "post/edit.html"
}
