package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"pandora/api/handler"
	"pandora/api/task"
	mdw "pandora/middleware"

	"pandora/service"
	ws "pandora/service/websocket"
)

//var Router = gin.Default()

//var Router = gin.New()

var ProviderSet = wire.NewSet(NewEngine, NewAppRouter)

type AppRouter struct {
	router  *gin.Engine
	handler *handler.Handler
	mdw     *mdw.Middleware
}

func NewEngine() *gin.Engine {
	return gin.Default()
}

func NewAppRouter(router *gin.Engine, handler *handler.Handler, mdw *mdw.Middleware) *AppRouter {
	return &AppRouter{
		router:  router,
		handler: handler,
		mdw:     mdw,
	}
}

func (ar *AppRouter) Run(addr ...string) error {
	ar.InitRouter()
	return ar.router.Run(addr...)
}

func (ar *AppRouter) InitRouter() {
	ar.addLoginRouter()
	ar.addAuthRouter()
	ar.addStockRouter()
	ar.addTaskRouter()
	ar.addSSERouter()
	ar.addWSRouter()
}

func (ar *AppRouter) addLoginRouter() {
	ar.router.POST("/login", ar.handler.Login)
}

func (ar *AppRouter) addAuthRouter() {
	r := ar.router.Group("/auth", ar.mdw.JWTAuthMiddleware())
	{
		r.GET("/currentUser", ar.handler.GetCurrentUser)
		//r.POST("/login", auth.Login)
		// 重定向
		r.POST("/register", func(c *gin.Context) {
			c.Request.URL.Path = "/auth/users"
			ar.router.HandleContext(c)
		})
		user := r.Group("/users", ar.mdw.CacheHandler("User"))
		{
			user.GET("/", ar.handler.GetUser)
			user.POST("/", ar.handler.CreateUser)
			user.PUT("/:id", ar.handler.UpdateUser)
			user.DELETE("/:id", ar.handler.DeleteUser)
		}
	}
}

func (ar *AppRouter) addStockRouter() {
	r := ar.router.Group("/stocks", ar.mdw.JWTAuthMiddleware())
	{
		r.GET("/daily", ar.mdw.TimeOut(2000), ar.mdw.RateLimit(), ar.mdw.AccessControl(), ar.mdw.CacheHandler("Stock"), ar.handler.GetStock)
		r.POST("/daily/upload", ar.handler.UploadStock)
		r.POST("/daily/download", ar.handler.DownloadStock)
	}
}

func (ar *AppRouter) addTaskRouter() {
	r := ar.router.Group("/tasks", ar.mdw.JWTAuthMiddleware())
	{
		r.GET("/list", ar.handler.GetTask)
		r.POST("/once", ar.handler.UploadStockOnce)
	}
}

func (ar *AppRouter) addSSERouter() {
	r := ar.router.Group("/sse", ar.mdw.JWTAuth(), service.Stream.SSEHandler(), ar.mdw.SSEHeaderMiddleware()) // JWTAuth授权
	{
		r.GET("/task", task.StartTaskSSE)
	}
}

func (ar *AppRouter) addWSRouter() {
	r := ar.router.Group("/ws", ar.mdw.JWTAuth(), ws.WebSocketHandler()) // JWTAuth授权
	{
		r.GET("/task", ar.handler.StartTaskWS)
	}
}
