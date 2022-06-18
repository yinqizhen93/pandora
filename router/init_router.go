package router

import (
	"github.com/gin-gonic/gin"
	"pandora/api/auth"
	"pandora/api/stock"
	"pandora/api/task"
	"pandora/middleware"
)

var Router = gin.Default()

//var Router = gin.New()

func InitRouter() {
	addLoginRouter()
	addAuthRouter()
	addStockRouter()
	addTaskRouter()
}

func addLoginRouter() {
	Router.POST("/login", auth.Login)
}

func addAuthRouter() {
	r := Router.Group("/auth", middleware.JWTAuthMiddleware())
	{
		r.GET("/currentUser", auth.GetCurrentUser)
		//r.POST("/login", auth.Login)
		// 重定向
		r.POST("/regist", func(c *gin.Context) {
			c.Request.URL.Path = "/auth/users"
			Router.HandleContext(c)
		})
		user := r.Group("/users")
		{
			user.GET("/", auth.GetUser)
			user.POST("/", auth.CreateUser)
			user.PUT("/:id", auth.UpdateUser)
			user.DELETE("/:id", auth.DeleteUser)
		}
	}
}

func addStockRouter() {
	r := Router.Group("/stocks", middleware.JWTAuthMiddleware())
	{
		r.GET("/daily", stock.GetStock)
		r.POST("/daily/upload", stock.UploadStock)
		r.POST("/daily/download", stock.DownloadStock)
	}
}

func addTaskRouter() {
	r := Router.Group("/tasks", middleware.JWTAuthMiddleware())
	{
		r.GET("/list", task.GetTask)
		r.POST("/once", task.UploadStockOnce)
		r.GET("/sse", task.SSE)
	}
}
