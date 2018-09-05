package main

import (
	"github.com/kataras/iris"
	"web-pms/routers"
	"web-pms/models"
	"strconv"
	"os"
	"fmt"
)

func main() {
	fmt.Println("Starting....")
	app := iris.New()

	initialize(app)
	iris.RegisterOnInterrupt(func() {
		defer models.Orm.Close()
	})

	routers.InitRoute(app)
	// Start the server using a network address.
	app.Run(iris.Addr(":" + strconv.Itoa(models.AppConfig.AppPort)))
}

func initialize(app *iris.Application)  {
	models.InitAppConfig(app)
	models.InitDB(app)
	initArgs(app)
	//defer func(){
	//	if err:=recover();err!=nil{
	//	}
	//}()
}

func initArgs(app *iris.Application) {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb(app)
			os.Exit(0)
		} else if v == "-appPort" {
			models.AppConfig.AppPort, _ = strconv.Atoi(os.Args[1])
		}
	}
}