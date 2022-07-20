package middleware

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"pandora/service/logger"
	"runtime/debug"
	"strings"
)

// Recovery recover掉项目可能出现的panic
func (mdw *Middleware) Recovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					mdw.logger.Error(c.Request.URL.Path,
						logger.Pair{K: "error", V: err},
						logger.Pair{K: "request", V: string(httpRequest)},
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					mdw.logger.Error("[Recovery from panic]",
						logger.Pair{K: "error", V: err},
						logger.Pair{K: "request", V: string(httpRequest)},
						logger.Pair{K: "stack", V: string(debug.Stack())},
					)
				} else {
					mdw.logger.Error("[Recovery from panic]",
						logger.Pair{K: "error", V: err},
						logger.Pair{K: "request", V: string(httpRequest)},
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
