package task

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/api"
	"pandora/db"
	"pandora/ent/task"
	"pandora/logs"
	"runtime/debug"
	"strconv"
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

	ctx := context.Background()
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
		logs.Logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(3002, "查询失败"))
		return
	}
	stocks, err := stockQuery.Offset(offset).Limit(pageSize).Select().All(ctx)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
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
