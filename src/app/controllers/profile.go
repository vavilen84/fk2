package controllers

import (
	"app/models/auth"
	"app/models/post"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type ProfileController struct {
	beego.Controller
}

func (c *ProfileController) Save() {
	title := c.GetString("title")
	content := c.GetString("content")
	or := orm.NewOrm()
	post.Create(title, content, or)
	c.Redirect("/", 302)
}

func (c *ProfileController) Update() {
	id, e := c.GetInt(":id")
	if e != nil {
		log.Fatal(e)
	}
	c.Data["title"] = "Edit Profile #"
	or := orm.NewOrm()
	post, _ := post.OneById(int64(id), or)
	c.Data["Post"] = post
	c.Layout = "layout.html"
	c.TplName = "post/edit.html"
	isLoggedIn, _ := auth.ValidateAuth(c.Ctx)
	c.Data["IsLoggedIn"] = isLoggedIn
}
