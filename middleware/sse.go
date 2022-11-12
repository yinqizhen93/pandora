package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/service/sse"
)

// SSE act as a middleware
func (mdw *Middleware) SSE() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Authorization")
		c.Header("Access-Control-Allow-Methods", "GET")
		// Initialize client channel
		clientChan := make(chan sse.Message)
		c.Set(sse.SSEClient, clientChan)
		// Send new connection to event server
		fmt.Println("create new sse client...")
		mdw.sse.NewClient <- clientChan
		fmt.Println("new sse client created...")
		defer func() {
			// Send closed connection to event server
			fmt.Println("sse client is closing...")
			mdw.sse.CloseClient <- clientChan
		}()
		go func() {
			// Send connection that is closed by client to event server
			<-c.Done()
			fmt.Println("request context is closed...")
			mdw.sse.CloseClient <- clientChan
		}()
		c.Next()
	}
}
