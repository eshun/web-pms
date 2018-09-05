package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/iris-contrib/middleware/cors"

	"strconv"
	"web-pms/controllers"

	"io/ioutil"
	"fmt"
)

// Route 路由
func InitRoute(app *iris.Application) {
	// Method:    GET
	// Resource:  http://localhost:8080

	app.StaticWeb("/", "./public/view")

	app.Get("/h.js", func(c iris.Context) {
		dat, err := ioutil.ReadFile("./public/js/pms.min.js")
		if(err!=nil){
			c.StatusCode(iris.StatusInternalServerError)
			fmt.Println(err)
			c.HTML("Error h.js")
		}
		//apikey
		//dm
		c.Write(dat)
	})

	app.OnAnyErrorCode(func(c iris.Context) {
		statusCode := c.GetStatusCode()

		c.JSON(iris.Map{"status": false, "message": strconv.Itoa(statusCode)})
	})

	app.WrapRouter(cors.WrapNext(cors.Options{
	 	AllowedOrigins:   []string{"*"},
	 	AllowCredentials: true,
	 }))
	//crs := cors.NewPartyMiddleware(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowCredentials: true,
	//})
	//v1 := app.Party("/api/v1")
	//v1.ConfigureParty(crs)
	//{
	//	v1.Get("/home", func(ctx iris.Context) {
	//		ctx.WriteString("Hello from /home")
	//	})
	//}

	mvc.New(app.Party( "/javascript/event")).Handle(new(controllers.EventController))
}

func CORSMiddleware(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if ctx.Method() == "OPTIONS" {
		ctx.StatusCode(204)
		return
	}

	ctx.Next()
}