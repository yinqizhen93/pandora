package ws

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//	// 解决跨域问题
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}

const WebSocketClient = "_WebSocketClient"

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	ReceiveStream chan []byte // 缓存接收的消息
	SendStream    chan []byte // 缓存发送的消息
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
//
func (c *Client) KeepReceive() {
	defer func() {
		// 退出，删除client, 关闭连接
		c.Hub.Unregister <- c
		fmt.Println("ws client is closing...")
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.ReceiveStream <- message // 接收到的消息放入receiveStream chan
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
//
func (c *Client) KeepSend() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.SendStream: // 接收sendStream消息
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.SendStream)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.SendStream)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

//// WebSocketHandler handles websocket requests from the peer.
//func WebSocketHandler() gin.HandlerFunc {
//
//	return func(c *gin.Context) {
//		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		client := &Client{
//			hub:           hub,
//			conn:          conn,
//			ReceiveStream: make(chan []byte, 256),
//			sendStream:    make(chan []byte, 256),
//		}
//		fmt.Println("find new ws client")
//		WSHub.register <- client
//		// Allow collection of memory referenced by the caller by doing all work in
//		// new goroutines.
//		go client.keepReceive()
//		go client.keepSend()
//		c.Set(WebSocketClient, client)
//	}
//}
