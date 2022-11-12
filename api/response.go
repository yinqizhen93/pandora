package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func FailResponse(code int32, msg string, err ...error) gin.H {
	var m string
	if len(err) > 0 {
		m = fmt.Sprintf("%s: %v", msg, err)
	} else {
		m = msg
	}
	return gin.H{
		"success": false,
		"errCode": code,
		"errMsg":  m,
	}
}
