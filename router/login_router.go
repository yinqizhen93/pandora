package router

import "pandora/api/auth"

func initLoginRouter() {
	//r.GET("/", api.LandingPage)
	Router.POST("/login", auth.Login)
}
