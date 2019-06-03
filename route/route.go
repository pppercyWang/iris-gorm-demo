package route

import (
	"../controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func InitRouter(app *iris.Application) {
	bathPath := "/api/v1"
	//app.Party（）分组路由
	mvc.New(app.Party(bathPath+"/book")).Handle(controllers.NewBookController())
}

