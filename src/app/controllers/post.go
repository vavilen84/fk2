package controllers

import (
	"app/models/auth"
	"app/models/post"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type PostController struct {
	beego.Controller
}

func (c *PostController) Create() {
	c.Data["title"] = "Create New Post"
	c.Layout = "layout.html"
	c.TplName = "post/create.html"
	isLoggedIn, _ := auth.ValidateAuth(c.Ctx)
	c.Data["IsLoggedIn"] = isLoggedIn
}

func (c *PostController) Save() {
	title := c.GetString("title")
	content := c.GetString("content")
	or := orm.NewOrm()
	post.Create(title, content, or)
	c.Redirect("/", 302)
}

func (c *PostController) Update() {
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
	id, e := c.GetInt("id")
	if e != nil {
		log.Fatal(e)
	}
	or := orm.NewOrm()
	post.Del(int64(id), or)
	c.Redirect("/", 302)
}

func (c *PostController) Edit() {
	id, e := c.GetInt(":id")
	if e != nil {
		log.Fatal(e)
	}
	c.Data["title"] = "Edit Post #"
	or := orm.NewOrm()
	post, _ := post.OneById(int64(id), or)
	c.Data["Post"] = post
	c.Layout = "layout.html"
	c.TplName = "post/edit.html"
	isLoggedIn, _ := auth.ValidateAuth(c.Ctx)
	c.Data["IsLoggedIn"] = isLoggedIn
}
