package route

import (
	"../controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"net/http"
)

func InitRouter(app *iris.Application) {
	//app.Use(CrossAccess)
	bathPath := "/api/v1"
	mvc.New(app.Party(bathPath+"/book")).Handle(controllers.NewBookController())
}

func CrossAccess11(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
func CrossAccess(ctx iris.Context) {
	ctx.ResponseWriter().Header().Add("Access-Control-Allow-Origin", "*")
}
