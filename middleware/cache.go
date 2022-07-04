package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"pandora/service/cache"
	"pandora/service/logger"
	"time"
)

var sf singleflight.Group // 多个goroutine使用同一个singleflight

// cacheWriter 继承gin.ResponseWriter，可以替换c.writer
type cacheWriter struct {
	key   string
	cache []byte
	gin.ResponseWriter
}

type cacheOption struct {
	timeout int
	forget  int
	schemas []string
}

type OptionFunc func(*cacheOption)

var defaultCacheOption = cacheOption{
	timeout: 2000,
	//forget:  100,
}

func WithTimeOut(t int) OptionFunc {
	return func(co *cacheOption) {
		co.timeout = t
	}
}

func WithSchemas(schemas ...string) OptionFunc {
	return func(co *cacheOption) {
		co.schemas = schemas
	}
}

//func WithForget(t int) OptionFunc {
//	return func(co *cacheOption) {
//		co.forget = t
//	}
//}

// CacheWithOpt 带超时限制，不建议使用
func CacheWithOpt(opts ...OptionFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			opt := defaultCacheOption
			for _, o := range opts {
				o(&opt)
			}
			url := c.Request.URL.RequestURI()
			if data, ok := cache.Cache.Get(url); ok {
				replyFromData(c, data)
			} else {
				writer := cacheWriter{key: url, ResponseWriter: c.Writer}
				ch := sf.DoChan(url, fnWrapper(c, opt.schemas, url, &writer))
				if opt.timeout <= 0 {
					select {
					case result := <-ch:
						if result.Err == nil {
							// singleflight 阻塞的其他请求, 其cache未初始化，直接返回从singleflight获取的返回值
							if writer.cache == nil {
								replyFromData(c, result.Val)
							}
						} else {
							// todo err != nil 的场景待添加
							replyFromData(c, []byte("请求失败"))
						}
					}
				} else {
					select {
					case result := <-ch:
						if result.Err == nil {
							// singleflight 阻塞的其他请求, 其cache未初始化，直接返回从singleflight获取的返回值
							if writer.cache == nil {
								replyFromData(c, result.Val)
							}
						} else {
							// todo err != nil 的场景待添加
							replyFromData(c, []byte("请求失败"))
						}
					case <-time.After(time.Duration(opt.timeout) * time.Millisecond):
						replyFromData(c, []byte("请求超时"))
					}
				}

			}
		}
	}
}

// CacheHandler 不带超时限制, 仅缓存GET请求和json返回
func CacheHandler(schema ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			url := c.Request.URL.RequestURI()
			if data, ok := cache.Cache.Get(url); ok {
				replyFromData(c, data)
			} else {
				writer := cacheWriter{key: url, ResponseWriter: c.Writer}
				// 一个goroutine panic, 其他goroutine也panic， 然后都被recovery middleware 捕获
				data, err, _ := sf.Do(url, fnWrapper(c, schema, url, &writer))
				if err == nil {
					// singleflight 阻塞的其他请求, 其cache未初始化，直接返回从singleflight获取的返回值
					if writer.cache == nil {
						replyFromData(c, data)
					}
				} else {
					logger.Error(fmt.Sprintf("cache singleflight 返回error:%+v", err))
					// todo err != nil 的场景待添加
					replyFromData(c, []byte("请求失败"))
				}
			}
		}
	}
}

// fnWrapper 返回singleflight参数fn， 负责执行请求和更新缓存
func fnWrapper(c *gin.Context, schema []string, url string, writer *cacheWriter) func() (interface{}, error) {
	return func() (interface{}, error) {
		// 为避免一个网络波动等造成查询的goroutine失败，而使得依赖该goroutine的其他goroutine都失败，
		// 每100ms， 就会有一个新的goroutine去请求
		go func() {
			time.Sleep(100 * time.Millisecond)
			sf.Forget(url)
		}()
		c.Writer = writer // 修改c.Writer为自定义writer, 添加缓存写入
		c.Next()          // 确保先请求，cw.cache被赋值
		if c.Writer.Status() == 200 {
			rs := struct {
				Success bool `json:"success"`
			}{}
			err := json.Unmarshal(writer.cache, &rs)
			if err != nil {
				panic(fmt.Sprintf("Unmarshal writer.cache fail: %s", err))
			}
			if rs.Success { // 请求返回success，更新缓存
				cache.Cache.Set(schema, url, writer.cache, 1)
			}
		}
		return writer.cache, nil
	}
}

// Write 覆盖Write方法
func (cw *cacheWriter) Write(data []byte) (int, error) {
	cw.cache = data
	return cw.ResponseWriter.Write(data) // 返回请求结果
}

func replyFromData(c *gin.Context, val interface{}) {
	c.Writer.WriteHeader(200)
	c.Header("Content-Type", "application/json; charset=utf-8")
	if _, err := c.Writer.Write(val.([]byte)); err != nil {
		panic(err)
	}
	c.Abort()
}

//// updateCache 更新缓存
//func updateCache(key string, data interface{}) {
//	service.Cache.Set(key, data, 1)
//	service.Cache.Wait()
//}
