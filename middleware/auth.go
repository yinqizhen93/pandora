package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"pandora/ent"
	"pandora/service"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func (mdw *Middleware) JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "请求头中auth为空",
			})
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "请求头中auth格式有误",
			})
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		tk, mc, err := service.ParseToken(parts[1])
		if err != nil {
			log.Println(err)
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors == jwt.ValidationErrorExpired { // token过期
					if claims, ok := tk.Claims.(*service.Claims); ok {
						can, err := mdw.tokenCanRefresh(c.Request.Context(), claims.UserId)
						if err != nil {
							// same as
							//nfe := &ent.NotFoundError{}
							//if errors.As(err, &nfe) { // 必须写nfe地址
							if _, ok2 := err.(*ent.NotFoundError); ok2 {
								goto fail
							} else {
								panic(err)
							}
						}
						if can {
							newToken, err := service.CreateToken(claims.UserId)
							if err != nil {
								panic(err)
							}
							c.Header("x-refreshed-token", newToken)
							c.Set("userId", claims.UserId) // token 过期也要设置userID
							return
						}
					} else {
						panic("token claims 有错误")
					}
				}
			}
		fail:
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "无效的Token",
			})
			return
		}
		// 将当前请求的userId信息保存到请求的上下文c上, c每次请求都会被初始化，所以每次要保存
		c.Set("userId", mc.UserId)
		//c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func (mdw *Middleware) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authInfo := c.Query("token")
		if authInfo == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "未授权的请求",
			})
		}
		// 按空格分割
		parts := strings.SplitN(authInfo, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "token格式有误",
			})
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		tk, mc, err := service.ParseToken(parts[1])
		if err != nil {
			log.Println(err)
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors == jwt.ValidationErrorExpired { // token过期
					if claims, ok := tk.Claims.(*service.Claims); ok {
						can, err := mdw.tokenCanRefresh(c.Request.Context(), claims.UserId)
						if err != nil {
							panic(err)
						}
						if can {
							newToken, err := service.CreateToken(claims.UserId)
							if err != nil {
								panic(err)
							}
							c.Header("x-refreshed-token", newToken)
							c.Set("userId", claims.UserId) // token 过期也要设置userID
							return
						}
					} else {
						panic("token claims 有错误")
					}
				}
			}
			c.AbortWithStatusJSON(401, gin.H{
				"msg": "无效的Token",
			})
			return
		}
		// 将当前请求的userId信息保存到请求的上下文c上
		c.Set("userId", mc.UserId)
		//c.Next()
	}
}

func (mdw *Middleware) tokenCanRefresh(ctx context.Context, id int) (bool, error) {
	user, err := mdw.db.User.Get(ctx, id)
	if err != nil {
		return false, err
	}
	ok, err := service.RefreshTokenExpired(user.RefreshToken)
	if err != nil {
		return ok, err
	}
	return !ok, nil
}
