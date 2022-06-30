package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/db"
	"pandora/ent"
	"pandora/service"
	"strings"
)

func AccessControl() gin.HandlerFunc {

	return func(c *gin.Context) {
		//获取请求的URI
		url := c.Request.URL.RequestURI()
		baseUrl := strings.Split(url, "?")[0]
		//获取请求方法
		act := c.Request.Method
		//todo 获取用户的角色, 避免每次请求数据库, 可将用户角色数据放入缓存
		ctx := c.Request.Context()
		id, ok := c.Get("userId")
		if !ok {
			panic("use do not exists")
		}
		subs := getRolesByUserId(ctx, id.(int))
		//判断策略中是否存在
		for _, sub := range subs {
			if ok, _ := service.Enforcer.Enforce(sub.Name, baseUrl, act); ok {
				return
			}
		}
		c.Status(http.StatusForbidden)
		c.Abort()
	}
}

func getRolesByUserId(ctx context.Context, id int) []*ent.Role {
	user, err := db.Client.User.Get(ctx, id)
	if err != nil {
		panic(err) // todo 过滤超时错误
	}
	roles, err := user.QueryRoles().All(ctx)
	if err != nil {
		panic(err) // todo 过滤超时错误
	}
	return roles
}
