package middleware

import "github.com/gin-gonic/gin"

func SSEHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Header("Content-Type", "text/event-stream")
		//c.Header("Cache-Control", "no-cache")
		//c.Header("Connection", "keep-alive")
		//c.Header("Transfer-Encoding", "chunked")
		// 添加跨域头
		c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:8000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Authorization")
		c.Header("Access-Control-Allow-Methods", "GET")
		//c.Next()
	}
}
