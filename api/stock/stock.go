package stock

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/db"
	"pandora/models"
	"pandora/utils"
)

func GetStock(c *gin.Context) {
	//var stock models.Stock
	var data []models.Stock
	ctx := context.Background()
	err := db.Client.Stock.Query().Select().Scan(ctx, data)
	if err != nil {
		c.JSON(http.StatusOK, utils.FailResponse(3002, "查询失败"))
	}
	fmt.Println(data)
	resp := gin.H{
		"code": "success",
		"data": data,
	}
	c.JSON(200, resp)
}
