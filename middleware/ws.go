package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	ws "pandora/service/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket handles websocket requests from the peer.
func (mdw *Middleware) WebSocket() gin.HandlerFunc {

	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &ws.Client{
			Hub:           mdw.ws,
			Conn:          conn,
			ReceiveStream: make(chan []byte, 256),
			SendStream:    make(chan []byte, 256),
		}
		fmt.Println("find new ws client")
		client.Hub.Register <- client
		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.KeepReceive()
		go client.KeepSend()
		c.Set(ws.WebSocketClient, client)
	}
}
