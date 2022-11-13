package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type XueQiuRealtimeReplay struct {
	Data []struct {
		Symbol             string      `json:"symbol"`
		Current            float64     `json:"current"`
		Percent            float64     `json:"percent"`
		Chg                float64     `json:"chg"`
		Timestamp          int64       `json:"timestamp"`
		Volume             int         `json:"volume"`
		Amount             float64     `json:"amount"`
		MarketCapital      float64     `json:"market_capital"`
		FloatMarketCapital float64     `json:"float_market_capital"`
		TurnoverRate       float64     `json:"turnover_rate"`
		Amplitude          float64     `json:"amplitude"`
		Open               float64     `json:"open"`
		LastClose          float64     `json:"last_close"`
		High               float64     `json:"high"`
		Low                float64     `json:"low"`
		AvgPrice           float64     `json:"avg_price"`
		TradeVolume        interface{} `json:"trade_volume"`
		Side               interface{} `json:"side"`
		IsTrade            bool        `json:"is_trade"`
		Level              int         `json:"level"`
		TradeSession       interface{} `json:"trade_session"`
		TradeType          interface{} `json:"trade_type"`
		CurrentYearPercent float64     `json:"current_year_percent"`
		TradeUniqueID      interface{} `json:"trade_unique_id"`
		Type               int         `json:"type"`
		BidApplSeqNum      interface{} `json:"bid_appl_seq_num"`
		OfferApplSeqNum    interface{} `json:"offer_appl_seq_num"`
		VolumeExt          interface{} `json:"volume_ext"`
		TradedAmountExt    interface{} `json:"traded_amount_ext"`
		TradeTypeV2        interface{} `json:"trade_type_v2"`
	} `json:"data"`
	ErrorCode        int         `json:"error_code"`
	ErrorDescription interface{} `json:"error_description"`
}

func requestXueQiu(code, market string) (float64, error) {
	fmt.Println("requesting xueqiu api..")
	market = strings.ToUpper(market)
	url := fmt.Sprintf("https://stock.xueqiu.com/v5/stock/realtime/quotec.json?symbol=%s%s", market, code)
	resp, err := http.Get(url)
	if err != nil {
		return 0, errors.Wrap(err, "http.Get请求出错")
	}
	defer resp.Body.Close()
	var r XueQiuRealtimeReplay
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return 0, errors.Wrap(err, "解码resp.Body出错")
	}
	fmt.Println("value", r.Data[0].Current)
	return r.Data[0].Current, nil
}

func requestQt(code, market string) (float64, error) {
	fmt.Println("requesting qt api..")
	market = strings.ToLower(market)
	url := fmt.Sprintf("http://qt.gtimg.cn/q=%s%s", market, code)
	resp, err := http.Get(url)
	if err != nil {
		return 0, errors.Wrap(err, "http.Get请求出错")
	}
	defer resp.Body.Close()
	// fmt.Println(resp.Header) 编码方式为gbk, 需要转为gbk 解码
	utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	r, err := io.ReadAll(utf8Reader)
	if err != nil {
		return 0, errors.Wrap(err, "读取resp.Body出错")
	}
	v := strings.Split(string(r), "~")[3]
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, errors.Wrap(err, "value转float64出错")
	}
	fmt.Println("value", f)
	return f, nil
}

func RealtimePrice(code, market string) float64 {
	rand.Seed(time.Now().Unix())
	choices := []func(string, string) (float64, error){requestQt, requestXueQiu}
	tick := time.NewTicker(time.Second * 3)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			v, err := choices[rand.Intn(len(choices))](code, market)
			if err != nil {
				panic(err)
			}
			fmt.Println(v)
		}
	}
}

func main() {
	RealtimePrice("600000", "sh")
}
