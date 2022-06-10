package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/db"
	"pandora/models"
	"pandora/service"
	"pandora/utils"
)

type UserInfo struct {
	Username string
	Password string
}

func Login(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, utils.FailResponse(2001, "无效的参数"))
		return
	}
	// 校验用户名和密码是否正确
	if user.Username == "admin" && user.Password == "123" {
		// 生成Token
		userId := getUserIdByName(user.Username)
		tokenString, err := service.CreateToken(userId)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, utils.FailResponse(2002, "用户名或密码错误"))
	return
}

func getUserIdByName(name string) int32 {
	var user models.User
	db.DB.Where("username = ?", name).First(&user)
	return user.Id
}
