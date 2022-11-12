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

const WebSocketClient = "_WebSocketClient"

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *ClientHub
	// The websocket connection.
	Conn *websocket.Conn
	// Buffered channel of outbound messages.
	ReceiveStream chan []byte // 缓存接收的消息
	SendStream    chan []byte // 缓存发送的消息
}

// KeepReceive 从c.Conn读取数据放入ReceiveStream
func (c *Client) KeepReceive() {
	defer func() {
		// 退出，删除client, 关闭连接
		c.Hub.ClientUnregister <- c
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

// KeepSend 从c.SendStream读取数据并写入c.Conn
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
			// 读取所有缓存的数据，换行分隔
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.SendStream)
			}
			// w.Close() 回将消息写入网络
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			// 定时ping客户端， 以判断连接是否存活
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
