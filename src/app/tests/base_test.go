package test

import (
	"app/models/post"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"math/rand"
	"path/filepath"
	"runtime"
	"time"
)

var Cfg = beego.AppConfig

func initOrm() {
	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	format := "%s:%s@tcp(%s:%s)/%s?charset=utf8"
	dbLink := fmt.Sprintf(format, dbUser, dbPass, dbHost, dbPort, dbName)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink)
}

func initTestBeegoApp() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func cleanDb() {
	o := orm.NewOrm()
	tables := []string{"user", "post"}
	for _, element := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE `%s`;", element)
		o.Raw(query).Exec()
	}
}

func init() {
	initOrm()
	initTestBeegoApp()
	cleanDb()
	uploadPostFixtures()
}

func uploadPostFixtures() {
	or := orm.NewOrm()
	post.Create("post_1_title", "post_1_content", or)
	post.Create("post_2_title", "post_2_content", or)
}

func getRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func rollbackTransaction(or orm.Ormer) {
	err := or.Rollback()
	if err != nil {
		log.Fatal(err)
	}
}

func beginTransaction() orm.Ormer {
	or := orm.NewOrm()
	err := or.Begin()
	if err != nil {
		log.Fatal(err)
	}

	return or
}
