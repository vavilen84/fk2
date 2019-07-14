package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id      int64 `orm:"auto"`
	Title   string
	Content string
}

type Login struct {
	Email    string
	Password string
}

type User struct {
	Id       int64  `orm:"auto"`
	Email    string `orm:"unique"`
	Password string
	Salt     string
}

func init() {
	orm.RegisterModel(new(Post))
	orm.RegisterModel(new(User))
}
