package handler

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) Register(c *gin.Context) {
	h.CreateUser(c)
}
