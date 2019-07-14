package post

import (
	"app/models"
	"github.com/astaxie/beego/orm"
	"log"
)

func OneById(id int64, or orm.Ormer) (*models.Post, orm.Ormer) {
	var post models.Post
	or.QueryTable("post").Filter("id", id).One(&post)

	return &post, or
}

func Del(id int64, or orm.Ormer) orm.Ormer {
	_, e := or.QueryTable("post").Filter("id", id).Delete()
	if e != nil {
		log.Fatal(e)
	}

	return or
}

func Create(title string, content string, or orm.Ormer) (orm.Ormer, error) {
	post := &models.Post{Content: content, Title: title}
	_, e := or.Insert(post)
	if e != nil {
		log.Fatal(e)
	}

	return or, e
}

func Update(id int64, title string, content string, or orm.Ormer) orm.Ormer {
	post := &models.Post{Id: id, Content: content, Title: title}
	_, e := or.Update(post)
	if e != nil {
		log.Fatal(e)
	}

	return or
}

func FindAll(or orm.Ormer) ([]*models.Post, orm.Ormer) {
	var posts []*models.Post
	_, e := or.QueryTable("post").All(&posts)
	if e != nil {
		log.Fatal(e)
	}

	return posts, or
}
