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
