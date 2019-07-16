package main

import (
	"app/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)

func createAdminUser() {
	beego.Info("Start create admin ...")
	admin := models.User{
		FirstName: os.Getenv("ADMIN_FIRST_NAME"),
		LastName:  os.Getenv("ADMIN_LAST_NAME"),
		Email:     os.Getenv("ADMIN_EMAIL"),
		Password:  os.Getenv("ADMIN_PASS"),
		Role:      models.RoleAdmin,
	}
	o := orm.NewOrm()
	err := models.InsertUser(o, admin)
	if err != nil {
		beego.Error(err)
	}
	beego.Info("Create admin success!")
}

func init() {
	Cmd.Run(createAdminUser)
}
