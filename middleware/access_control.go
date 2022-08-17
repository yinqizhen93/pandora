package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"pandora/ent"
	"pandora/service/access"
	"strings"
)

func (mdw *Middleware) AccessControl() gin.HandlerFunc {
	if mdw.accessCtrl == nil {
		mdw.accessCtrl = access.NewRBAC(mdw.db)
	}
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
			panic("user do not exists")
		}
		subs, err := mdw.getRolesByUserId(ctx, id.(int))
		if err != nil {
			// todo 需要用fmt.Sprintf包裹一层吗
			panic(fmt.Sprintf("getRolesByUserId error: %s", err)) // todo 过滤context超时错误
		}
		//判断策略中是否存在
		for _, sub := range subs {
			if ok := mdw.accessCtrl.HasAccess(sub.Name, baseUrl, act); ok {
				return
			}
		}
		//c.Status(http.StatusForbidden)
		c.JSON(http.StatusForbidden, gin.H{"msg": "未授权操作"})
		c.Abort()
	}
}

func (mdw *Middleware) getRolesByUserId(ctx context.Context, id int) ([]*ent.Role, error) {
	user, err := mdw.db.User.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "查询id=%d的用户失败", id)
	}
	roles, err := user.QueryRoles().All(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "查询用户id=%d的角色失败", id)
	}
	return roles, nil
}
