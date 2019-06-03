package main

import (
	"./conf"
	"./route"
	"flag"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	flag.Parse()
	app := newApp()
	route.InitRouter(app)
	err := app.Run(iris.Addr(":"+conf.Sysconfig.Port), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	//配置允许跨域请求
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodOptions)
	return app
}
