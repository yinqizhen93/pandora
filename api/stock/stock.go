package stock

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/api"
	"pandora/db"
	"pandora/models"
)

func GetStock(c *gin.Context) {
	//var stock models.Stock
	var data []models.Stock
	ctx := context.Background()
	err := db.Client.Stock.Query().Select().Scan(ctx, data)
	if err != nil {
		c.JSON(http.StatusOK, api.FailResponse(3002, "查询失败"))
	}
	fmt.Println(data)
	resp := gin.H{
		"code": "success",
		"data": data,
	}
	c.JSON(200, resp)
}
