package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"pandora/service"
)

var limiter *rate.Limiter

func init() {
	limiter = rate.NewLimiter(100, 10000)
}

func (mdw *Middleware) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		if err := limiter.Wait(ctx); err != nil {

			c.Status(http.StatusGatewayTimeout)
			fmt.Println(err)
			//if errors.Is(err, ctx.Err()) {
			//	fmt.Println("chaoshi ....")
			//	c.Abort()
			//} else {
			//	panic(err)
			//}
		}
	}
}

// 自己生成的rate limiter
var bucket *service.Bucket

func init() {
	bucket = service.NewBucket(100, 10000)
}

func RateLimit2() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !bucket.Pick() {
			c.Status(http.StatusTooManyRequests)
			c.Abort()
		}
	}
}
