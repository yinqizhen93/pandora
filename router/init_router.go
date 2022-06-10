package router

import (
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

//var Router = gin.New()

func InitRouter() {
	initAuthRouter()
	initStockRouter()
	initDoctorRouter()
	initDepartmentRouter()
	initLoginRouter()
	//err := r.Run(":8080")// listen and serve on 0.0.0.0:8080
	//if err != nil {
	//	return
	//}
}

func init() {
	// 使用自定义logger 和 recovery
	//Router.Use(middleware.GinLogger(logs.Logger), middleware.GinRecovery(logs.Logger, true))
}
