// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import "fmt"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]struct{}

	// Inbound messages from the clients.
	Message chan []byte // 桥接作用，实现一个服务端发送，多个客户端接收

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

var WSHub = NewHub()

func NewHub() (hub *Hub) {
	hub = &Hub{
		clients:    make(map[*Client]struct{}),
		Message:    make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go hub.listen()
	return
}

func (h *Hub) listen() {
	for {
		select {
		case client := <-h.register:
			fmt.Println("registering new client")
			h.clients[client] = struct{}{}
		case client := <-h.unregister:
			// todo 这里的OK 判断是否必须？
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.sendStream)
			}
		case msg := <-h.Message: // 接受到服务端写入数据，广播给所有客户端连接
			for client := range h.clients {
				select {
				case client.sendStream <- msg:
				default:
					// client.stream被堵塞 说明client 没有接收，则客户端连接可能已经断开
					close(client.sendStream)
					delete(h.clients, client)
				}
			}
		}
	}
}
