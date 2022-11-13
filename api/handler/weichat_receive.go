package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/api"
	"time"
)

type Message struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        int64  `xml:"MsgId"`
	MsgDataId    string `xml:"MsgDataId"`
	Idx          int    `xml:"Idx"`
}

var AccessToken string

var AppId = "wx285e860ef821d3d0"

var AppSecret = "adf6d2947489627d91f124df0a941ede"

var EncodingAESKey = "4IFFW477xAiIgaNgFVmKtsa4dkB0IcaQEc8NG7YX5Xu"

func GetAccessToken() {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", AppId, AppSecret)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var m map[string]any
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		panic(err)
	}
	AccessToken = m["access_token"].(string)
}

type xml struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}

func (h *Handler) WechatReceive(c *gin.Context) {
	var msg Message
	if err := c.ShouldBindXML(&msg); err != nil {
		c.JSON(http.StatusOK, api.FailResponse(1200, err.Error()))
		return
	}
	fmt.Println(msg)
	if AccessToken == "" {
		GetAccessToken()
	}
	c.XML(200, xml{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      "this a test replay",
	})
}
