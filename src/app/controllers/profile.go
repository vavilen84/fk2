package controllers

import (
	_ "github.com/go-sql-driver/mysql"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) Save() {
	c.setAuthData()
	//title := c.GetString("title")
	//content := c.GetString("content")
	//or := orm.NewOrm()
	//post.Create(title, content, or)
	//c.Redirect("/", 302)
}

func (c *ProfileController) Update() {
	c.setAuthData()
	//id, e := c.GetInt(":id")
	//if e != nil {
	//	log.Fatal(e)
	//}
	//c.Data["title"] = "Edit Profile "
	//or := orm.NewOrm()
	//user, _ := user.OneById(int64(id), or)
	//c.Data["User"] = user
	//c.Layout = "layout.html"
	//c.TplName = "post/edit.html"
}
