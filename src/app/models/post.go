package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id      int `orm:"auto"`
	Title   string
	Content string
}

func FindPostById(o orm.Ormer, id int) (post Post, err error) {
	err = o.QueryTable("post").Filter("id", id).One(&post)
	if err != nil {
		beego.Error(err)
	}
	return
}

func DeleletePost(o orm.Ormer, id int) (err error) {
	_, err = o.QueryTable("post").Filter("id", id).Delete()
	if err != nil {
		beego.Error(err)
	}
	return
}

func InsertPost(o orm.Ormer, post Post) (err error) {
	_, err = o.Insert(&post)
	if err != nil {
		beego.Error(err)
	}
	return
}

func UpdatePost(o orm.Ormer, post Post) (err error) {
	_, err = o.Update(&post)
	if err != nil {
		beego.Error(err)
	}
	return
}

func FindAllPosts(or orm.Ormer) (posts []Post, err error) {
	_, err = or.QueryTable("post").All(&posts)
	if err != nil {
		beego.Error(err)
	}
	return
}

func init() {
	orm.RegisterModel(new(Post))
}
