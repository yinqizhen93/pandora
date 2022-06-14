package router

import (
	"pandora/api/stock"
	"pandora/middleware"
)

func initStockRouter() {
	//r.GET("/", api.LandingPage)
	r := Router.Group("/stocks")
	r.Use(middleware.JWTAuthMiddleware())
	{
		r.GET("/daily", stock.GetStock)
		r.POST("/daily/upload", stock.UploadStock)
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
