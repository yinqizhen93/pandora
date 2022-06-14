package stock

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
	"pandora/api"
	"pandora/db"
	"pandora/ent"
	"pandora/ent/stock"
	"pandora/logs"
	"runtime/debug"
	"strconv"
	"time"
)

//type UserQueryParams struct {
//	Page      int    `json:"page" binding:"required"`
//	PageSize  int    `json:"pageSize" binding:"required"`
//	SearchVal string `json:"searchVal" binding:"required"`
//	StartDate string `json:"startDate" binding:"required"`
//	EndDate   string `json:"endDate" binding:"required"`
//}

func GetStock(c *gin.Context) {
	//var req UserQueryParams
	strPage, ok := c.GetQuery("page")
	if !ok {
		c.JSON(http.StatusOK, api.FailResponse(3002, "page参数缺失"))
		return
	}
	page, err := strconv.Atoi(strPage)
	if err != nil {
		c.JSON(http.StatusOK, api.FailResponse(3002, "page参数错误"))
		return
	}

	strPageSize, ok := c.GetQuery("pageSize")
	if !ok {
		c.JSON(http.StatusOK, api.FailResponse(3002, "pageSize参数缺失"))
		return
	}
	pageSize, err := strconv.Atoi(strPageSize)
	if err != nil {
		c.JSON(http.StatusOK, api.FailResponse(3002, "pageSize参数错误"))
		return
	}

	searchVal, ok := c.GetQuery("searchVal")
	if !ok {
		c.JSON(http.StatusOK, api.FailResponse(3002, "searchVal参数缺失"))
		return
	}

	strStartDate, ok := c.GetQuery("startDate")
	if !ok {
		c.JSON(http.StatusOK, api.FailResponse(3002, "startDate参数缺失"))
		return
	}
	startDate, err := time.Parse("2006-01-02", strStartDate)
	if !ok {
		c.JSON(http.StatusOK, api.FailResponse(3002, "searchVal参数缺失"))
		return
	}

	strEndDate, ok := c.GetQuery("endDate")
	if !ok {
		c.JSON(http.StatusOK, api.FailResponse(3002, "endDate参数缺失"))
		return
	}
	endDate, err := time.Parse("2006-01-02", strEndDate)
	if !ok {
		c.JSON(http.StatusOK, api.FailResponse(3002, "searchVal参数缺失"))
		return
	}

	ctx := context.Background()
	offset := (page - 1) * pageSize
	stockQuery := db.Client.Stock.Query().Where(stock.And(
		stock.DateGTE(startDate),
		stock.DateLTE(endDate),
	))
	if searchVal != "" {
		stockQuery = stockQuery.Where(
			stock.Or(
				stock.CodeContains(searchVal),
				stock.NameContains(searchVal),
			))
	}
	total, err := stockQuery.Count(ctx)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
		c.JSON(http.StatusOK, api.FailResponse(3002, "查询失败"))
		return
	}
	stocks, err := stockQuery.Offset(offset).Limit(pageSize).Select().All(ctx)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
		c.JSON(http.StatusOK, api.FailResponse(3002, "查询失败"))
		return
	}

	fmt.Println(stocks)
	resp := gin.H{
		"code":  "success",
		"data":  stocks,
		"total": total,
	}
	c.JSON(200, resp)
}

func UploadStock(c *gin.Context) {
	formFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, api.FailResponse(3004, "参数错误"))
		return
	}
	file, err := formFile.Open()
	if err != nil {
		c.JSON(http.StatusOK, api.FailResponse(3004, "打开上传文件失败"))
		return
	}
	f, err := excelize.OpenReader(file)
	if err != nil {
		c.JSON(http.StatusOK, api.FailResponse(3004, "读取上传文件失败"))
		return
	}
	go func() {
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
			logs.Logger.Error(fmt.Sprintf("插入数据失败: %s", err))
		}
	}()

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
