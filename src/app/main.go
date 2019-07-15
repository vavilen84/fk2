package main

import (
	_ "app/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var Cfg = beego.AppConfig

func initOrm() {
	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")

	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink)
}

func main() {
	initOrm()
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
