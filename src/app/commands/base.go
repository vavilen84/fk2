package main

import (
	"os"

	"github.com/CastAI/database-package/castdatabase"
	"github.com/astaxie/beego"

	"auth-service/config"
)

var Cmd Command

type Command struct {
	Func func()
}

func initDB(mode string) error {
	dbConf := castdatabase.Config{
		Mode:          mode,
		Driver:        config.Get("db_driver"),
		Source:        config.Get("db_data_source"),
		SourceTestApi: config.Get("db_data_source_test_api"),
	}

	return castdatabase.Init(dbConf)
}

func (c *Command) Run(f func()) {
	c.Func = f
}

func main() {
	mode := os.Getenv("APP_ENV")
	if mode == "" {
		mode = "cmd"
	}

	err := config.Init(mode)
	if err != nil {
		beego.Error(err)
	}

	err = initDB(mode)
	if err != nil {
		beego.Error(err)
	}

	if Cmd.Func != nil {
		Cmd.Func()
	}
}