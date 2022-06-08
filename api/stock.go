package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/db"
	"pandora/models"
)

func GetStock(c *gin.Context) {
	//var stock models.Stock
	var data []models.Stock
	db.DB.Find(&data).Scan(&data)
	fmt.Println(data)
	resp := gin.H{
		"code": "success",
		"data": data,
	}
	c.JSON(200, resp)
}
