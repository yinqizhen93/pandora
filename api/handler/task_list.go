package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"pandora/api"
	"pandora/ent/task"
	"pandora/service"
	ws "pandora/service/websocket"
	"runtime/debug"
	"time"
)

type taskQuery struct {
	Page     int `form:"page" binding:"required,gte=1"`
	PageSize int `form:"pageSize" binding:"required"`
	//StartDate time.Time `form:"startDate" binding:"required,ltefield=EndDate" time_format:"2006-01-02"`
	//EndDate   time.Time `form:"endDate" binding:"required" time_format:"2006-01-02"`
	//SearchVal string    `form:"searchVal"`
	Status []int8 `form:"taskStatus"`
}

func (h *Handler) GetTask(c *gin.Context) {
	//var req UserQueryParams
	var tq taskQuery
	if err := c.ShouldBindQuery(&tq); err != nil {
		c.JSON(http.StatusOK, api.FailResponse(1200, err.Error()))
		return
	}

	ctx := c.Request.Context()
	offset := (tq.Page - 1) * tq.PageSize
	stockQuery := h.db.Task.Query()
	if tq.Status != nil {
		stockQuery = stockQuery.Where(task.StatusIn(tq.Status...))
	}
	total, err := stockQuery.Count(ctx)
	if err != nil {
		h.logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(3002, "查询失败"))
		return
	}
	stocks, err := stockQuery.Offset(offset).Limit(tq.PageSize).Select().All(ctx)
	if err != nil {
		h.logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(3002, "查询失败"))
		return
	}

	//fmt.Println(stocks)
	resp := gin.H{
		"code":  "success",
		"data":  stocks,
		"total": total,
	}
	c.JSON(200, resp)
}

// StartTaskSSE an example of Server Send Event
func (h *Handler) StartTaskSSE(c *gin.Context) {
	go func() {
		for {
			time.Sleep(time.Second * 1)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			// Send current time to clients message channel
			service.Stream.Message <- service.Message{Pipeline: "message2", Data: currentTime}
		}
	}()

	cliStream, ok := c.Get(service.SSEClient)
	if !ok {
		panic("no sseClient find")
	}
	c.Stream(func(w io.Writer) bool {
		// Stream message to client from message channel
		if msg, ok := <-cliStream.(service.ClientChan); ok {
			fmt.Println(msg)
			c.SSEvent(msg.Pipeline, msg.Data)
			return true
		}
		return false
	})
}

// StartTaskWS an example of WebSocket
func (h *Handler) StartTaskWS(c *gin.Context) {
	cli, ok := c.Get(ws.WebSocketClient)
	if !ok {
		panic("no WebSocketClient find")
	}

	for {
		msg, ok := <-cli.(*ws.Client).ReceiveStream
		if !ok {
			return
		}
		// Send current time to clients message channel
		now := time.Now().Format("2006-01-02 15:04:05")
		currentTime := fmt.Sprintf("The Current Time Is %v with msg %v", now, msg)
		ws.WSHub.Message <- []byte(currentTime)
	}

	//go func() {
	//	for {
	//		time.Sleep(time.Second * 1)
	//		now := time.Now().Format("2006-01-02 15:04:05")
	//		currentTime := fmt.Sprintf("The Current Time Is %v", now)
	//
	//		// Send current time to clients message channel
	//		ws.WSHub.Message <- []byte(currentTime)
	//	}
	//}()

}
