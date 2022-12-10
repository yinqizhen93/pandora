package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"pandora/api"
	"pandora/ent/user"
	"pandora/service"
	"runtime/debug"
)

type UserInfo struct {
	Username string
	Password string
}

func (h Handler) Login(c *gin.Context) {
	// 用户发送用户名和密码过来
	var userInf UserInfo
	err := c.ShouldBind(&userInf)
	if err != nil {
		c.JSON(http.StatusOK, api.FailResponse(2001, "无效的参数"))
		return
	}
	// 校验用户名和密码是否正确
	//if userInf.Username == "admin" && userInf.Password == "123" {
	ctx := c.Request.Context()
	// todo handle err
	if ok, err := h.validUserAndPasswd(ctx, userInf.Username, userInf.Password); err == nil && ok {
		// 生成Token
		userId, err := h.getUserIdByName(ctx, userInf.Username)
		if err != nil {
			// todo 记录日志
			h.logger.Error(fmt.Sprintf("获取用户失败：%s; \n %s", err, debug.Stack()))
			c.JSON(http.StatusOK, api.FailResponse(2009, "登录失败"))
			return
		}
		tokenString, err := service.CreateToken(userId)
		if err != nil {
			panic(err)
		}
		// 保存refreshToken
		refreshToken := service.CreateRefreshToken()
		service.SaveRefreshToken(ctx, h.db, userId, refreshToken)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, api.FailResponse(2002, "用户名或密码错误"))
	return
}

func (h Handler) getUserIdByName(ctx context.Context, name string) (int, error) {
	userId, err := h.db.User.Query().Where(user.UsernameEQ(name)).Select("id").Int(ctx)
	//db.DB.Where("username = ?", name).First(&user)
	if err != nil {
		return 0, errors.Wrap(err, "getUserIdByName failed")
	}
	return userId, nil
}

func (h Handler) validUserAndPasswd(ctx context.Context, username, passwd string) (bool, error) {
	pwdHash, err := h.db.User.Query().Where(user.UsernameEQ(username)).Select("password").String(ctx)
	if err != nil {
		return false, errors.Wrap(err, "查询用户失败")
	}
	err = bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(passwd))
	if err != nil {
		return false, nil
	}
	return true, nil
}
