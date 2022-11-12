package ws

import (
	"fmt"
	"github.com/google/wire"
)

// ClientHub maintains the set of active clients and broadcasts messages to the
// clients.
type ClientHub struct {
	// Registered clients.
	clients map[*Client]struct{}
	// Inbound messages from the clients.
	MessageBroadcaster chan []byte // 桥接作用，将服务端写入的数据，广播给所有客户端
	// Register requests from the clients.
	ClientRegister chan *Client
	// Unregister requests from clients.
	ClientUnregister chan *Client
}

var ProviderSet = wire.NewSet(NewHub)

func NewHub() (hub *ClientHub) {
	hub = &ClientHub{
		clients:            make(map[*Client]struct{}),
		MessageBroadcaster: make(chan []byte),
		ClientRegister:     make(chan *Client),
		ClientUnregister:   make(chan *Client),
	}
	go hub.listen()
	return
}

func (h *ClientHub) listen() {
	for {
		select {
		// 新的客户端连接
		case client := <-h.ClientRegister:
			fmt.Println("registering new ws client")
			h.clients[client] = struct{}{}
		// 客户端断开连接
		case client := <-h.ClientUnregister:
			// todo 这里的OK 判断是否必须？
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.SendStream)
			}
		// 接受到服务端写入数据，广播给所有客户端连接
		case msg := <-h.MessageBroadcaster: // 接受到服务端写入数据，广播给所有客户端连接
			for client := range h.clients {
				select {
				case client.SendStream <- msg:
				default:
					// client.sendStream被堵塞 说明client 没有接收，则客户端连接可能已经断开
					close(client.SendStream)
					delete(h.clients, client)
				}
			}
		}
	}
}
