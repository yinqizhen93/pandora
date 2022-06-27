package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/singleflight"
	"pandora/service"
)

var sf singleflight.Group // 多个goroutine使用同一个singleflight

type cacheWriter struct {
	key   string
	cache []byte
	gin.ResponseWriter
}

func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			url := c.Request.URL.RequestURI()
			if data, ok := service.Cache.Get(url); ok {
				fmt.Println("find cache...")
				ReplyFromCache(c, data)
			} else {
				writer := cacheWriter{key: url, ResponseWriter: c.Writer}
				data, err := sf.Do(url, func() (interface{}, error) {
					fmt.Println("create cache...")
					c.Writer = &writer // 修改c.Writer为自定义writer, 添加缓存写入
					c.Next()           // 确保先请求，cw.cache被赋值
					return writer.cache, nil
				})
				// todo err != nil
				if err == nil {
					if writer.cache == nil { // singleflight 阻塞的其他请求
						ReplyFromCache(c, data)
					}
				}
			}
		}
		//c.Next()  c.Next()仅在需要返回后，继续逻辑处理的情况下
		// do something
	}
}

func (cw *cacheWriter) Write(data []byte) (int, error) {
	cw.cache = data
	service.Cache.Set(cw.key, data, 1)
	service.Cache.Wait()
	return cw.ResponseWriter.Write(data) // 返回请求结果
}

func ReplyFromCache(c *gin.Context, val interface{}) {
	c.Writer.WriteHeader(200)
	c.Header("Content-Type", "application/json; charset=utf-8")
	if _, err := c.Writer.Write(val.([]byte)); err != nil {
		panic(err)
	}
	c.Abort()
}
