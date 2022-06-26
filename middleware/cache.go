package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			url := c.Request.URL.RequestURI()
			fmt.Println(url)
			c.Abort()
		}

		//c.Next()  c.Next()仅在需要返回后，继续逻辑处理的情况下
		// do something
	}
}
