package utils

import "github.com/gin-gonic/gin"

func FailResponse(code int32, msg string) gin.H {
	return gin.H{
		"success": false,
		"errCode": code,
		"errMsg":  msg,
	}
}
