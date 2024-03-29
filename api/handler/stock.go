package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"net/http"
	"pandora/api"
	"pandora/ent"
	"pandora/ent/stock"
	"pandora/utils"
	"runtime/debug"
	"time"
)

// time_format:"2006-01-02" validate 只在tag form 里面起作用，tag json里不起作用
type stockQuery struct {
	Page      int       `form:"page" binding:"required,gte=1"`
	PageSize  int       `form:"pageSize" binding:"required"`
	StartDate time.Time `form:"startDate" binding:"required,ltefield=EndDate" time_format:"2006-01-02"`
	EndDate   time.Time `form:"endDate" binding:"required" time_format:"2006-01-02"`
	SearchVal string    `form:"searchVal"`
}

// GetStock 查询Stock接口
// @Summary 查询Stock接口
// @Tags 查询Stock接口
// @Accept application/json
// @Produce application/json
// @Param object query stockQuery true "查询参数"
// @Security ApiKeyAuth
// @Router /stocks/daily [get]
func (h Handler) GetStock(c *gin.Context) {
	var sq stockQuery
	if err := c.ShouldBindQuery(&sq); err != nil {
		c.JSON(http.StatusOK, api.FailResponse(1200, err.Error()))
		return
	}
	ctx := c.Request.Context()
	offset := (sq.Page - 1) * sq.PageSize
	stockQuery := h.db.Stock.Query().Where(stock.And(
		stock.DateGTE(sq.StartDate),
		stock.DateLTE(sq.EndDate),
	))
	if sq.SearchVal != "" {
		stockQuery = stockQuery.Where(
			stock.Or(
				stock.CodeContains(sq.SearchVal),
				stock.NameContains(sq.SearchVal),
			))
	}
	total, err := stockQuery.Count(ctx)
	if err != nil {
		h.logger.Error(fmt.Sprintf("查询失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(3002, "查询失败"))
		return
	}
	stocks, err := stockQuery.Offset(offset).Limit(sq.PageSize).Select().All(ctx)
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

func (h Handler) UploadStock(c *gin.Context) {
	formFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, api.FailResponse(3004, "参数错误"))
		return
	}
	file, err := formFile.Open()
	if err != nil {
		c.JSON(200, api.FailResponse(3004, "打开上传文件失败"))
		return
	}
	// todo 变成任务形式，可以看到成功还是失败
	go h.uploadStockFromFile(file)
	resp := gin.H{
		"code": "success",
	}
	c.JSON(200, resp)
}

func (h Handler) uploadStockFromFile(file multipart.File) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		// todo handle error
		return err
	}
	rows, err := f.GetRows("Sheet1")
	header := rows[0]
	bulk := make([]*ent.StockCreate, len(rows)-1)
	for ri, r := range rows[1:] {
		sc := h.db.Stock.Create()
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
	//ctx := c.Request.Context()
	// 此处不能使用c.Request.Context()，因为请求结束后context就reset了，除非copy 一个c
	ctx := context.Background()
	_, err = h.db.Stock.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		// todo handle error
		fmt.Println(err)
		h.logger.Error(fmt.Sprintf("插入数据失败: %s", err))
		return err
	}
	return nil
}

//func strToDate(val string) time.Time {
//	// todo Excelizer 读取Excel日期类型数据待研究
//	v, err := time.Parse("2006/1/2", val)
//	if err != nil {
//		panic(err)
//	}
//	return v
//}
//
//func strToFloat32(val string) float32 {
//	v, err := strconv.ParseFloat(val, 32)
//	if err != nil {
//		panic(err)
//	}
//	return float32(v)
//}
//
//func strToInt32(val string) int32 {
//	v, err := strconv.Atoi(val)
//	if err != nil {
//		panic(err)
//	}
//	return int32(v)
//}

// json.Unmarshal 反序列化类似"2020-01-01"格式日期 , 会出错, 只能识别RFC3339格式日期
// 使用 utils.LocalTime 替代
// 另外一种方法是，将日期字段定义为string ，使用binding:datetime 检验格式， 在使用时，将string类型转为time.Time类型
type userDownload struct {
	SearchVal string          `json:"searchVal"`
	StartDate utils.LocalTime `json:"startDate" binding:"required"`
	EndDate   utils.LocalTime `json:"endDate" binding:"required"`
}

func (h Handler) DownloadStock(c *gin.Context) {
	var ud userDownload
	err := c.ShouldBindJSON(&ud)
	if err != nil {
		h.logger.Error(fmt.Sprintf("参数错误：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(3002, "参数错误"))
		return
	}
	stockQuery := h.db.Stock.Query().Where(stock.And(
		stock.DateGTE(time.Time(ud.StartDate)),
		stock.DateLTE(time.Time(ud.EndDate)),
	))
	if ud.SearchVal != "" {
		stockQuery = stockQuery.Where(
			stock.Or(
				stock.CodeContains(ud.SearchVal),
				stock.NameContains(ud.SearchVal),
			))
	}
	ctx := c.Request.Context()
	stocks, err := stockQuery.Select().All(ctx)
	if err != nil {
		c.JSON(200, api.FailResponse(3002, "查询数据失败"))
	}
	stks := make([]any, len(stocks))
	for i, s := range stocks {
		stks[i] = s
	}
	file := excelize.NewFile()
	if xs, err := utils.NewXlsxStorage(file, stks); err != nil {
		h.logger.Error(fmt.Sprintf("生成XlsxStorage失败：%+v; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(3002, "生成XlsxStorage失败"))
		return
	} else {
		if err := xs.WriteXlsx(); err != nil {
			h.logger.Error(fmt.Sprintf("生成XlsxStorage失败：%+v; \n %s", err, debug.Stack()))
			c.JSON(200, api.FailResponse(3002, "生成XlsxStorage失败"))
			return
		}
	}
	//streamWriter, err := file.NewStreamWriter("Sheet1")
	//if err != nil {
	//	c.JSON(200, api.FailResponse(3002, "生成Excel失败"))
	//}
	//buff, err := file.WriteToBuffer()

	fileName := "stock.xlsx"
	//c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	//c.Header("response-type", "blob")
	//file.Write(c.Writer)

	//c.Writer.Write(buff.Bytes())
	//c.Data(h   ttp.StatusOK, contentType, buff.Bytes())
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", "attachment;filename="+fileName)
	c.Writer.Header().Add("Content-Transfer-Encoding", "binary")
	_ = file.Write(c.Writer)
}
