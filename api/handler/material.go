package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/api"
	ws "pandora/service/websocket"
	"runtime/debug"
	"strconv"
	"time"
)

// time_format:"2006-01-02" validate 只在tag form 里面起作用，tag json里不起作用
type materialQuery struct {
	Page     int `form:"page" binding:"required,gte=1"`
	PageSize int `form:"pageSize" binding:"required"`
}

func (h Handler) GetMaterial(c *gin.Context) {
	var mq materialQuery
	if err := c.ShouldBindQuery(&mq); err != nil {
		c.JSON(http.StatusOK, api.FailResponse(1200, err.Error()))
		return
	}
	ctx := c.Request.Context()
	offset := (mq.Page - 1) * mq.PageSize
	stockQuery := h.db.Material.Query()
	total, err := stockQuery.Count(ctx)
	if err != nil {
		h.logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(3002, "查询失败"))
		return
	}
	stocks, err := stockQuery.Offset(offset).Limit(mq.PageSize).Select().All(ctx)
	if err != nil {
		h.logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(3002, "查询失败"))
		return
	}
	resp := gin.H{
		"success": true,
		"data":    stocks,
		"total":   total,
	}
	c.JSON(200, resp)
}

func (h Handler) UpdateMaterial(c *gin.Context) {
	strId := c.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(200, api.FailResponse(1002, "参数错误"))
		h.logger.Error(fmt.Sprintf("参数错误:%s; %s", err, string(debug.Stack())))
		return
	}
	um, err := api.ParseJsonFormInputMap(c)
	fmt.Println("um", um)
	if err != nil {
		h.logger.Error(fmt.Sprintf("请求参数解析失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(1001, "请求参数解析失败"))
		return
	}
	upd := h.db.Material.UpdateOneID(intId)
	if name, ok := um["name"]; ok {
		upd.SetName(name.(string))
	}
	if code, ok := um["code"]; ok {
		upd.SetCode(code.(string))
	}
	if describe, ok := um["describe"]; ok {
		upd.SetDescribe(describe.(string))
	}
	if price, ok := um["price"]; ok {
		upd.SetPrice(price.(float64))
	}
	if startDate, ok := um["buyDate"]; ok {
		upd.SetBuyDate(startDate.(time.Time))
	}
	if _, err := upd.Save(c.Request.Context()); err != nil {
		h.logger.Error(fmt.Sprintf("更新保存失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(1005, "更新保存失败"))
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"msg":     "更新成功",
	})
}

// EditMaterial an example of WebSocket
func (h *Handler) EditMaterial(c *gin.Context) {
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

func (h Handler) updateDbMaterial(c *gin.Context) {
	um, err := api.ParseJsonFormInputMap(c)
	strId, ok := um["id"]
	if !ok {
		c.JSON(200, api.FailResponse(1002, "参数错误"))
		h.logger.Error(fmt.Sprintf("参数错误:%s; %s", err, string(debug.Stack())))
		return
	}
	intId, ok := strId.(int)
	if !ok {
		c.JSON(200, api.FailResponse(1002, "参数错误"))
		h.logger.Error(fmt.Sprintf("参数错误:%s; %s", err, string(debug.Stack())))
		return
	}
	fmt.Println("um", um)
	if err != nil {
		h.logger.Error(fmt.Sprintf("请求参数解析失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(1001, "请求参数解析失败"))
		return
	}
	upd := h.db.Material.UpdateOneID(intId)
	if name, ok := um["name"]; ok {
		upd.SetName(name.(string))
	}
	if code, ok := um["code"]; ok {
		upd.SetCode(code.(string))
	}
	if describe, ok := um["describe"]; ok {
		upd.SetDescribe(describe.(string))
	}
	if price, ok := um["price"]; ok {
		upd.SetPrice(price.(float64))
	}
	if startDate, ok := um["buyDate"]; ok {
		upd.SetBuyDate(startDate.(time.Time))
	}
	if _, err := upd.Save(c.Request.Context()); err != nil {
		h.logger.Error(fmt.Sprintf("更新保存失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(1005, "更新保存失败"))
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"msg":     "更新成功",
	})
}
