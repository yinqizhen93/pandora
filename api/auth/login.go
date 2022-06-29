package auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"pandora/api"
	"pandora/db"
	"pandora/ent/user"
	"pandora/service"
	"runtime/debug"
)

type UserInfo struct {
	Username string
	Password string
}

func Login(c *gin.Context) {
	// 用户发送用户名和密码过来
	var userInf UserInfo
	err := c.ShouldBind(&userInf)
	if err != nil {
		c.JSON(http.StatusOK, api.FailResponse(2001, "无效的参数"))
		return
	}
	// 校验用户名和密码是否正确
	if userInf.Username == "admin" && userInf.Password == "123" {
		// 生成Token
		ctx := c.Request.Context()
		userId, err := getUserIdByName(ctx, userInf.Username)
		if err != nil {
			// todo 记录日志
			service.Logger.Error(fmt.Sprintf("获取用户失败：%s; \n %s", err, debug.Stack()))
			c.JSON(http.StatusOK, api.FailResponse(2009, "登录失败"))
			return
		}
		tokenString, err := service.CreateToken(userId)
		if err != nil {
			panic(err)
		}
		// 保存refreshToken
		refreshToken := service.CreateRefreshToken()
		service.SaveRefreshToken(ctx, userId, refreshToken)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, api.FailResponse(2002, "用户名或密码错误"))
	return
}

func getUserIdByName(ctx context.Context, name string) (int, error) {
	userId, err := db.Client.User.Query().Where(user.UsernameEQ(name)).Select("id").Int(ctx)
	//db.DB.Where("username = ?", name).First(&user)
	if err != nil {
		return 0, errors.Wrap(err, "getUserIdByName failed")
	}
	return userId, nil
}
