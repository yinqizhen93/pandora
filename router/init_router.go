package router

import (
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func InitRouter() {
	initAuthRouter()
	initStockRouter()
	initDoctorRouter()
	initDepartmentRouter()
	//err := r.Run(":8080")// listen and serve on 0.0.0.0:8080
	//if err != nil {
	//	return
	//}
}
