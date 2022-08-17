package middleware

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"sync"
	"time"
)

type timeoutWriter struct {
	gin.ResponseWriter
	h           http.Header
	wbuf        bytes.Buffer
	code        int
	timedOut    bool
	wroteHeader bool // 超时时， WriteHeader可能还没调用， code就没被赋值
	mu          sync.Mutex
}

// 参考 https://github.com/golang/go/blob/003dbc4cda6a1418fc419461799320521d64f4e5/src/net/http/server.go#L3209
// ServeHttp(w ResponseWriter, r *Request), w实际传入的是一个response结构体，实现了ResponseWriter
// 对于gin的中间件而言，最终（注意是最终）返回的内容一定要写入其传入的writer中，也就是调用最原始的c.Writer.Write方法，从而最终调用response.Write()

// TimeOut 应该放在其他中间件的前面，避免writer被修改，
// 每次请求不管是否超时都会多出一个timeoutWriter结构占用内存，只能等待gc清除
func (mdw *Middleware) TimeOut(t int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(t)*time.Millisecond)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		tw := timeoutWriter{ResponseWriter: c.Writer, h: make(http.Header)}
		c.Writer = &tw
		done := make(chan struct{})
		// 带一个缓冲，防止超时后panicChan没有读取方，c.Next 发生异常，defer里面panicChan <- err被阻塞
		panicChan := make(chan interface{}, 1)
		go func() {
			defer func() {
				// recovery错误，避免让其他goroutine panic
				if err := recover(); err != nil {
					// 通知主协程panic发生
					panicChan <- fmt.Sprintf("程序异常:%s; %s", err, debug.Stack()) // todo err添加堆栈信息
				}
			}()
			c.Next() // 这里面的Write方法不会写入到最原始的Writer中
			done <- struct{}{}
		}()
		select {
		case p := <-panicChan:
			panic(p)
		case <-done:
			tw.mu.Lock()
			defer tw.mu.Unlock()
			// tw.code, tw.wroteHeader, tw.h都存在并发读写
			if !tw.wroteHeader {
				tw.code = http.StatusOK
			}
			tw.ResponseWriter.WriteHeader(tw.code)
			for k, v := range tw.h {
				tw.ResponseWriter.Header()[k] = v
			}
			tw.ResponseWriter.Write(tw.wbuf.Bytes())
		case <-ctx.Done():
			tw.mu.Lock()
			defer tw.mu.Unlock()
			// tw.timedOut, tw.h都存在并发读写
			tw.timedOut = true
			tw.ResponseWriter.WriteHeader(http.StatusGatewayTimeout)
			for k, v := range tw.h {
				tw.ResponseWriter.Header()[k] = v
			}
		}
	}
}

//var ErrHandlerTimeout = errors.New("http: Handler timeout")

// Write 完全覆盖，没有调用原始Write方法，避免在c.Next()和主goroutine写入一个writer
func (tw *timeoutWriter) Write(p []byte) (int, error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	// tw.timedOut 可能存在并发读写，所以加锁
	if tw.timedOut {
		return 0, nil
	}
	if !tw.wroteHeader { // 后面的middleware 可能覆盖了WriteHeader方法， 导致 tw.writeHeader未调用
		tw.writeHeader(http.StatusOK)
	}
	return tw.wbuf.Write(p)
}

// WriteHeader 覆盖WriteHeader方法
func (tw *timeoutWriter) WriteHeader(statusCode int) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	// tw.code 可能存在并发读写，所以加锁
	// 已经超时了就不写tw.code，减少内存占用， 后面的middleware可能会覆盖方法， 导致tw.Write 可能在tw.WriteHeader之前调用
	// tw.wroteHeader 就会为true
	if tw.timedOut || tw.wroteHeader {
		return
	}
	tw.writeHeader(statusCode)
}

func (tw *timeoutWriter) writeHeader(code int) {
	tw.code = code
	tw.wroteHeader = true
}

// Header 覆盖Header方法， 让后面的方法实际在tw.h操作
func (tw *timeoutWriter) Header() http.Header {
	return tw.h
}
