package task

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"pandora/api"
	"pandora/db"
	"pandora/ent"
	"pandora/ent/task"
	"pandora/service"
	"strconv"
	"time"
)

func UploadStockOnce(c *gin.Context) {
	formFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, api.FailResponse(3004, "参数错误"))
		return
	}
	upload := func() {
		fmt.Println("uploading...")
		file, err := formFile.Open()
		if err != nil {
			c.JSON(200, api.FailResponse(3004, "打开上传文件失败"))
			return
		}
		f, err := excelize.OpenReader(file)
		if err != nil {
			c.JSON(200, api.FailResponse(3004, "读取上传文件失败"))
			return
		}
		rows, err := f.GetRows("Sheet1")
		header := rows[0]
		bulk := make([]*ent.StockCreate, len(rows)-1)
		for ri, r := range rows[1:] {
			sc := db.Client.Stock.Create()
			for i, c := range header {
				switch c {
				case "market":
					sc.SetMarket(r[i])
				case "code":
					sc.SetCode(r[i])
				case "name":
					sc.SetName(r[i])
				case "date":
					sc.SetDate(strToDate(r[i]))
				case "open":
					sc.SetOpen(strToFloat32(r[i]))
				case "close":
					sc.SetClose(strToFloat32(r[i]))
				case "high":
					sc.SetHigh(strToFloat32(r[i]))
				case "low":
					sc.SetLow(strToFloat32(r[i]))
				case "volume":
					sc.SetVolume(strToInt32(r[i]))
				case "outstanding_share":
					sc.SetOutstandingShare(strToInt32(r[i]))
				case "turnover":
					sc.SetTurnover(strToFloat32(r[i]))
				}
			}
			bulk[ri] = sc
		}
		ctx := context.Background()
		_, err = db.Client.Stock.CreateBulk(bulk...).Save(ctx)
		if err != nil {
			fmt.Println(err)
			service.Logger.Error(fmt.Sprintf("插入数据失败: %s", err))
		}
	}

	tk := NewTask("stock upload", task.Type("once"), "上传每日股票数据", upload)
	go tk.Start(c)
	//fmt.Println(stocks)
	resp := gin.H{
		"code": "success",
		//"data": stocks,
	}
	c.JSON(200, resp)
}

func strToDate(val string) time.Time {
	// todo Excelizer 读取Excel日期类型数据待研究
	v, err := time.Parse("2006/1/2", val)
	if err != nil {
		panic(err)
	}
	return v
}

func strToFloat32(val string) float32 {
	v, err := strconv.ParseFloat(val, 32)
	if err != nil {
		panic(err)
	}
	return float32(v)
}

func strToInt32(val string) int32 {
	v, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return int32(v)
}