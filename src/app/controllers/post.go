package controllers

import (
	"app/models"
	"app/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type PostController struct {
	BaseController
}

func (c *PostController) EditList() {
	c.setResponseData("Posts", "post/edit-list")
	o := orm.NewOrm()
	posts, _ := models.FindAllPosts(o)
	c.Data["Posts"] = utils.GetPostOnViewList(o, posts)
}

func (c *PostController) Create() {
	c.setResponseData("Create New Post", "post/create")
}

func (c *PostController) Save() {
	o := orm.NewOrm()

	userId, err := strconv.Atoi(c.GetString("user_id"))
	if err != nil {
		beego.Error(err)
	}
	user, err := models.FindUserById(o, userId)
	if err == orm.ErrNoRows {
		beego.Info("User not found")
		c.Redirect("/404", 302)
	}

	post := models.Post{
		Title:     c.GetString("title"),
		Content:   c.GetString("content"),
		UserId:    user.Id,
		CreatedAt: int(time.Now().Unix()),
		Publish:   1,
	}
	err = models.InsertPost(o, post)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/", 302)
}

func (c *PostController) Update() {
	o := orm.NewOrm()

	userId, err := strconv.Atoi(c.GetString("user_id"))
	if err != nil {
		beego.Error(err)
	}
	_, err = models.FindUserById(o, userId)
	if err == orm.ErrNoRows {
		beego.Info("User not found")
		c.Redirect("/404", 302)
	}

	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		beego.Error(err)
	}

	post, err := models.FindPostById(o, id)
	if err == orm.ErrNoRows {
		beego.Info("Post not found")
		c.Redirect("/404", 302)
	}
	post.Title = c.GetString("title")
	post.Content = c.GetString("content")

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
	c.setResponseData(title, "post/edit")
	c.Data["Post"] = post
}

func (c *PostController) View() {
	id, err := strconv.Atoi(c.GetString("id"))
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()
	post, err := models.FindPostById(o, id)
	if err == orm.ErrNoRows {
		c.Redirect("/404", 302)
	}
	title := post.Title
	c.setResponseData(title, "post/view")
	c.Data["Post"] = post
}
