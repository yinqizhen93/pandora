package middleware

import (
	"github.com/gin-gonic/gin"
	"pandora/service/logger"
	"time"
)

func (mdw *Middleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		mdw.logger.Info(path,
			logger.Pair{K: "status", V: c.Writer.Status()},
			logger.Pair{K: "method", V: c.Request.Method},
			logger.Pair{K: "path", V: path},
			logger.Pair{K: "query", V: query},
			logger.Pair{K: "ip", V: c.ClientIP()},
			logger.Pair{K: "user-agent", V: c.Request.UserAgent()},
			logger.Pair{K: "errors", V: c.Errors.ByType(gin.ErrorTypePrivate).String()},
			logger.Pair{K: "cost", V: cost},
		)
	}
}
