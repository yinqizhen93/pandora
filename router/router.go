package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"os"
	"pandora/api/handler"
	_ "pandora/docs" // 以上导入for Swagger
	mdw "pandora/middleware"
	ws "pandora/service/websocket"
)

//var Router = gin.Default()

//var Router = gin.New()

var ProviderSet = wire.NewSet(NewAppRouter)

type AppRouter struct {
	router  *gin.Engine
	handler *handler.Handler
	mdw     *mdw.Middleware
}

func (ar *AppRouter) NewEngine() *gin.Engine {
	if os.Getenv("PANDORA") == "production" {
		e := gin.New()
		e.Use(ar.mdw.Logger(), ar.mdw.Recovery(true))
		return e
	}
	return gin.Default()
}

func NewAppRouter(handler *handler.Handler, mdw *mdw.Middleware) *AppRouter {
	ar := &AppRouter{
		handler: handler,
		mdw:     mdw,
	}
	ar.router = ar.NewEngine()
	return ar
}

func (ar *AppRouter) Run(addr ...string) error {
	ar.InitRouter()
	return ar.router.Run(addr...)
}

func (ar *AppRouter) InitRouter() {
	ar.addSwaggerRouter()
	ar.addLoginRouter()
	ar.addAuthRouter()
	ar.addStockRouter()
	ar.addTaskRouter()
	ar.addSSERouter()
	ar.addWSRouter()
	ar.addMaterialRouter()
	ar.addDepartmentRouter()
}

// see api docs on http://localhost:5001/swagger/index.html
func (ar *AppRouter) addSwaggerRouter() {
	ar.router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
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

func (ar *AppRouter) addMaterialRouter() {
	r := ar.router.Group("/materials")
	{
		r.GET("", ar.handler.GetMaterial)
		r.GET("/edit", ws.WebSocketHandler(), ar.handler.EditMaterial)
		r.PUT("/:id", ar.handler.UpdateMaterial)
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
	r := ar.router.Group("/sse", ar.mdw.JWTAuth(), ar.mdw.SSE()) // JWTAuth授权
	{
		r.GET("/task", ar.handler.StartTaskSSE)
	}
}

func (ar *AppRouter) addWSRouter() {
	r := ar.router.Group("/ws", ar.mdw.JWTAuth(), ws.WebSocketHandler()) // JWTAuth授权
	{
		r.GET("/task", ar.handler.StartTaskWS)
	}
}

func (ar *AppRouter) addDepartmentRouter() {
	r := ar.router.Group("/departments", ar.mdw.JWTAuthMiddleware())
	{
		r.POST("/", ar.handler.CreateDepartment)
	}
}
