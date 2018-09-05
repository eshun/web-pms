package models

import (
	"path/filepath"
	"io/ioutil"
	"os"
	"github.com/BurntSushi/toml"
	"github.com/kataras/iris"
)

var AppConfig = appConfig{
	AppPort:    8080,
	DriverName: "mysql",
	DBName:     "web_pms",
	DBServer:   "127.0.0.1",
	DBUser:     "root",
	DBPassword: "root",
	DBCharset:  "utf8",
	DBPort:     3306,
	ShowSql:    false,
	Prefix:     "tb_",
}

func InitAppConfig(app *iris.Application) {
	dbPath, err := filepath.Abs("./conf/app.conf")
	if err != nil {
		app.Logger().Fatalf("AppConfig not find conf: %v", err)
		return
	}
	data, err := ioutil.ReadFile(dbPath)
	if err != nil {
		os.Exit(-1)
		app.Logger().Fatalf("AppConfig ReadFile failed: %v", err)
		return
	}
	if _, err := toml.Decode(string(data), &AppConfig); err != nil {
		app.Logger().Fatalf("AppConfig Decode failed: %v", err)
		return
	}
}

type appConfig struct {
	AppPort    int
	DriverName string
	DBName     string
	DBServer   string
	DBUser     string
	DBPassword string
	DBCharset  string
	DBPort     int
	ShowSql    bool
	Prefix  string
}