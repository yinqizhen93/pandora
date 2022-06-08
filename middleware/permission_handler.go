package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/service"
)

func Authorize() gin.HandlerFunc {

	return func(c *gin.Context) {

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		fmt.Println(obj)
		//获取请求方法
		act := c.Request.Method
		fmt.Println(act)
		//todo 获取用户的角色
		sub := "admin"

		//判断策略中是否存在
		if ok, _ := service.Enforcer.Enforce(sub, obj, act); ok {
			fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			fmt.Println("很遗憾,权限验证没有通过")
			c.Abort()
		}
	}
}
