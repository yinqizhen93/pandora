package router

import (
	"pandora/api/task"
	"pandora/middleware"
)

func initTaskRouter() {
	//r.GET("/", api.LandingPage)
	r := Router.Group("/tasks")
	r.Use(middleware.JWTAuthMiddleware())
	{
		r.GET("/list", task.GetTask)
		r.POST("/once", task.UploadStockOnce)

		//user := baseUrl.Group("/user")
		//{
		//	user.GET("/", api.GetUser)
		//	user.POST("/", api.CreateUser)
		//	user.PUT("/:id", api.UpdateUser)
		//	user.DELETE("/:id", api.DeleteUser)
		//}
		//stock := baseUrl.Group("/stock")
		//{
		//	stock.GET("/", api.GetStock)
		//}

	}
}
