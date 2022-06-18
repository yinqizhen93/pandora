package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"pandora/service"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "请求头中auth为空",
			})
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "请求头中auth格式有误",
			})
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := service.ParseToken(parts[1])
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "无效的Token",
			})
			return
		}
		// 将当前请求的userId信息保存到请求的上下文c上, c每次请求都会被初始化，所以每次要保存
		c.Set("userId", mc.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authInfo := c.Query("token")
		if authInfo == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "未授权的请求",
			})
		}
		// 按空格分割
		parts := strings.SplitN(authInfo, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "token格式有误",
			})
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := service.ParseToken(parts[1])
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "无效的Token",
			})
		}
		// 将当前请求的userId信息保存到请求的上下文c上
		c.Set("userId", mc.UserId)
		c.Next()
	}
}
