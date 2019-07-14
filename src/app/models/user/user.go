package user

import (
	"app/models"
	"github.com/astaxie/beego/orm"
)

func FindByEmail(email string, or orm.Ormer) (*models.User, orm.Ormer) {
	var user models.User
	or.QueryTable("user").Filter("email", email).One(&user)

	return &user, or
}

func OneById(id int64, or orm.Ormer) (*models.User, orm.Ormer) {
	var user models.User
	or.QueryTable("user").Filter("id", id).One(&user)

	return &user, or
}
