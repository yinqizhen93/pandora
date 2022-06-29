package router

import (
	"github.com/gin-gonic/gin"
	"pandora/api/auth"
	"pandora/api/stock"
	"pandora/api/task"
	mdw "pandora/middleware"
	"pandora/service"
	ws "pandora/service/websocket"
)

var Router = gin.Default()

//var Router = gin.New()

func InitRouter() {
	addLoginRouter()
	addAuthRouter()
	addStockRouter()
	addTaskRouter()
	addSSERouter()
	addWSRouter()
}

func addLoginRouter() {
	Router.POST("/login", auth.Login)
}

func addAuthRouter() {
	r := Router.Group("/auth", mdw.JWTAuthMiddleware())
	{
		r.GET("/currentUser", auth.GetCurrentUser)
		//r.POST("/login", auth.Login)
		// 重定向
		r.POST("/register", func(c *gin.Context) {
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
	r := Router.Group("/stocks", mdw.JWTAuthMiddleware())
	{
		r.GET("/daily", mdw.TimeOut(2000), mdw.RateLimit(), mdw.AccessControl(), mdw.CacheHandler(), stock.GetStock)
		r.POST("/daily/upload", stock.UploadStock)
		r.POST("/daily/download", stock.DownloadStock)
	}
}

func addTaskRouter() {
	r := Router.Group("/tasks", mdw.JWTAuthMiddleware())
	{
		r.GET("/list", task.GetTask)
		r.POST("/once", task.UploadStockOnce)
	}
}

func addSSERouter() {
	r := Router.Group("/sse", mdw.JWTAuth(), service.Stream.SSEHandler(), mdw.SSEHeaderMiddleware()) // JWTAuth授权
	{
		r.GET("/task", task.StartTaskSSE)
	}
}

func addWSRouter() {
	r := Router.Group("/ws", mdw.JWTAuth(), ws.WebSocketHandler()) // JWTAuth授权
	{
		r.GET("/task", task.StartTaskWS)
	}
}
