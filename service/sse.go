package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var Stream = NewSSEvent()

const SSEClient = "sseClient"

// SSEvent  keeps a list of clients those are currently attached
//and broadcasting events to those clients.
type SSEvent struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message chan string
	// New client connections
	NewClient chan chan string
	// Closed client connections
	CloseClient chan chan string
	// Total client connections
	Clients map[chan string]struct{}
}

// ClientChan New event messages are broadcast to all registered client connection channels
type ClientChan chan string

// NewSSEvent Initialize event and Start procnteessing requests
func NewSSEvent() (event *SSEvent) {
	event = &SSEvent{
		Message:     make(chan string),
		NewClient:   make(chan chan string),
		CloseClient: make(chan chan string),
		Clients:     make(map[chan string]struct{}),
	}
	go event.listen()
	return
}

//It Listens all incoming requests from clients.
//Handles addition and removal of clients and broadcast messages to clients.
func (sse *SSEvent) listen() {
	fmt.Println("sse is listening...")
	for {
		select {
		// Add new available client
		case client := <-sse.NewClient:
			fmt.Println("find new client...")
			sse.Clients[client] = struct{}{}
			log.Printf("Client added. %d registered clients", len(sse.Clients))

		// Remove closed client
		case client := <-sse.CloseClient:
			fmt.Println("find client closed...")
			delete(sse.Clients, client)
			log.Printf("Removed client. %d registered clients", len(sse.Clients))

		// Broadcast message to client
		case eventMsg := <-sse.Message:
			fmt.Println("find message...")
			for clientMessageChan := range sse.Clients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (sse *SSEvent) SSEHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize client channel
		clientChan := make(ClientChan)
		c.Set("sseClient", clientChan)
		// Send new connection to event server
		fmt.Println("create new client...")
		sse.NewClient <- clientChan
		fmt.Println("new client created...")
		defer func() {
			// Send closed connection to event server
			fmt.Println("client is closing...")
			sse.CloseClient <- clientChan
		}()
		go func() {
			// Send connection that is closed by client to event server
			<-c.Done()
			fmt.Println("request context is closed...")
			sse.CloseClient <- clientChan
		}()
		c.Next()
	}
}
