package service

import (
	"github.com/gin-gonic/gin"
	"log"
)

var Stream = NewSSEvent()

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

//func main() {
//	router := gin.Default()
//
//	// Initialize new streaming server
//	stream := NewSSEvent()
//	router.Use(stream.serveHTTP())
//
//	// Basic Authentication
//	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
//		"admin": "admin123", // username : admin, password : admin123
//	}))
//
//	// Authorized client can stream the event
//	authorized.GET("/stream", func(c *gin.Context) {
//		// We are streaming current time to clients in the interval 10 seconds
//		go func() {
//			for {
//				time.Sleep(time.Second * 10)
//				now := time.Now().Format("2006-01-02 15:04:05")
//				currentTime := fmt.Sprintf("The Current Time Is %v", now)
//
//				// Send current time to clients message channel
//				stream.Message <- currentTime
//			}
//		}()
//
//		c.Stream(func(w io.Writer) bool {
//			// Stream message to client from message channel
//			if msg, ok := <-stream.Message; ok {
//				c.SSEvent("message", msg)
//				return true
//			}
//			return false
//		})
//	})
//
//	//Parse Static files
//	router.StaticFile("/", "./public/index.html")
//
//	router.Run(":8085")
//}

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
	for {
		select {
		// Add new available client
		case client := <-sse.NewClient:
			sse.Clients[client] = struct{}{}
			log.Printf("Client added. %d registered clients", len(sse.Clients))

		// Remove closed client
		case client := <-sse.CloseClient:
			delete(sse.Clients, client)
			log.Printf("Removed client. %d registered clients", len(sse.Clients))

		// Broadcast message to client
		case eventMsg := <-sse.Message:
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
		// Send new connection to event server
		sse.NewClient <- clientChan
		defer func() {
			// Send closed connection to event server
			sse.CloseClient <- clientChan
		}()
		go func() {
			// Send connection that is closed by client to event server
			<-c.Done()
			sse.CloseClient <- clientChan
		}()
		c.Next()
	}
}
