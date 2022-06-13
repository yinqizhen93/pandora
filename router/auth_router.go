package router

import (
	"github.com/gin-gonic/gin"
	"pandora/api/auth"
	"pandora/middleware"
)

func initAuthRouter() {
	//r.GET("/", api.LandingPage)
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

		//role := r.Group("/roles")
		//{
		//	role.GET("/", auth.GetRole)
		//	role.POST("/", auth.CreateRole)
		//	role.PUT("/:id", auth.UpdateRole)
		//	role.DELETE("/:id", auth.DeleteRole)
		//}
	}
}
