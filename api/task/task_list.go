package task

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"pandora/api"
	"pandora/db"
	"pandora/ent/task"
	"pandora/service"
	"pandora/service/logger"
	ws "pandora/service/websocket"
	"runtime/debug"
	"strconv"
	"time"
)

func GetTask(c *gin.Context) {
	//var req UserQueryParams
	strPage, ok := c.GetQuery("page")
	if !ok {
		c.JSON(200, api.FailResponse(3002, "page参数缺失"))
		return
	}
	page, err := strconv.Atoi(strPage)
	if err != nil {
		c.JSON(200, api.FailResponse(3002, "page参数错误"))
		return
	}

	strPageSize, ok := c.GetQuery("pageSize")
	if !ok {
		c.JSON(200, api.FailResponse(3002, "pageSize参数缺失"))
		return
	}
	pageSize, err := strconv.Atoi(strPageSize)
	if err != nil {
		c.JSON(200, api.FailResponse(3002, "pageSize参数错误"))
		return
	}

	ctx := c.Request.Context()
	offset := (page - 1) * pageSize
	stockQuery := db.Client.Task.Query()
	strTaskStatus, ok := c.GetQueryArray("taskStatus")
	if ok {
		fmt.Println(strTaskStatus)
		taskStatus := make([]int8, len(strTaskStatus))
		for i, s := range strTaskStatus {
			a, err := strconv.Atoi(s)
			if err != nil {
				c.JSON(200, api.FailResponse(3002, "taskStatus参数错误"))
				return
			}
			taskStatus[i] = int8(a)
		}
		stockQuery = stockQuery.Where(task.StatusIn(taskStatus...))
	}

	total, err := stockQuery.Count(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(3002, "查询失败"))
		return
	}
	stocks, err := stockQuery.Offset(offset).Limit(pageSize).Select().All(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
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
func StartTaskSSE(c *gin.Context) {
	go func() {
		for {
			time.Sleep(time.Second * 1)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			// Send current time to clients message channel
			service.Stream.Message <- currentTime
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
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}

// StartTaskWS an example of WebSocket
func StartTaskWS(c *gin.Context) {
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
