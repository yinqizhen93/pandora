package router

import (
	"pandora/api/auth"
	"pandora/middleware"
)

func initStockRouter() {
	//r.GET("/", api.LandingPage)
	r := Router.Group("/stock")
	r.Use(middleware.Authorize())
	{
		r.GET("/daily", auth.Login)
		r.GET("/day", auth.Login)
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
