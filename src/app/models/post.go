package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type PostOnView struct {
	Id          int
	Title       string
	Description string
	Image       string
	User        User
	CreatedAt   string
}

type Post struct {
	Id        int `orm:"auto"`
	Title     string
	Content   string
	UserId    int
	CreatedAt int
	Publish   int
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

func FindAllPosts(o orm.Ormer) (posts []Post, err error) {
	_, err = o.QueryTable("post").
		OrderBy("-id").
		All(&posts)
	if err != nil {
		beego.Error(err)
	}
	return
}

func CountPosts(o orm.Ormer) int64 {
	var maps []orm.Params
	num, err := o.Raw("SELECT count(id) as count FROM post").Values(&maps)
	if err == nil && num > 0 {
		i, err := strconv.ParseInt(maps[0]["count"].(string), 10, 64)
		if err != nil {
			beego.Error(err)
		}
		return i
	}
	return 0
}

func ListPostsByOffsetAndLimit(o orm.Ormer, offset, limit int) (posts []Post, err error) {
	_, err = o.QueryTable("post").
		Filter("Publish", 1).
		Offset(offset).
		Limit(limit).
		OrderBy("-id").
		All(&posts)
	if err != nil {
		beego.Error(err)
	}
	return
}

func init() {
	orm.RegisterModel(new(Post))
}
