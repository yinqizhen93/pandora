package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/service"
	"strings"
)

func AccessControl() gin.HandlerFunc {

	return func(c *gin.Context) {
		//获取请求的URI
		url := c.Request.URL.RequestURI()
		baseUrl := strings.Split(url, "?")[0]
		//获取请求方法
		act := c.Request.Method
		//todo 获取用户的角色
		sub := "yinqizhen"

		//判断策略中是否存在
		if ok, _ := service.Enforcer.Enforce(sub, baseUrl, act); ok {
			c.Next()
		} else {
			c.Status(http.StatusForbidden)
			c.Abort()
		}
	}
}
