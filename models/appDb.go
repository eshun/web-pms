package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"errors"
	"strconv"
	"github.com/kataras/iris"
	"github.com/go-xorm/core"
)

var Orm *xorm.Engine

//数据库连接
func InitDB(app *iris.Application) {
	var dataSourceName string
	var err error
	driverName := AppConfig.DriverName
	if driverName == "sqlite3" {

	} else if driverName == "mysql" {
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBServer, strconv.Itoa(AppConfig.DBPort), AppConfig.DBName, AppConfig.DBCharset)
	}
	Orm, err = xorm.NewEngine(AppConfig.DriverName, dataSourceName)
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
		return
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, AppConfig.Prefix)
	Orm.SetTableMapper(tbMapper)

	Orm.ShowSQL(AppConfig.ShowSql)
	Orm.Logger().SetLevel(core.LOG_INFO)
	//Orm.SetMaxIdleConns(setting.MaxIdle)
	//Orm.SetMaxOpenConns(setting.MaxOpen)

	Syncdb(app)
}

func Syncdb(app *iris.Application) {
	err := Orm.Sync2(new(Event),new(Resource),new(PageTimes),new(Target),new(Req),new(Res))
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized table: %v", err)
		return
	}

	//Orm.CreateTables()
	//Orm.IsTableExist();DropTables(),IsTableEmpty()
}

var (
	ErrNotExist = errors.New("not exist")
)