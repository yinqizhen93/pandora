package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"pandora/api"
	"sort"
	"strings"
)

const (
	//token = "wechat4go"
	token = "pandora"
)

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return hex.EncodeToString(s.Sum(nil))
}

type RequestInfo struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	Echostr   string `form:"echostr"`
}

func (h *Handler) ValidWechat(c *gin.Context) {
	var ri RequestInfo
	if err := c.ShouldBindQuery(&ri); err != nil {
		c.JSON(http.StatusOK, api.FailResponse(1200, err.Error()))
		return
	}
	fmt.Println(ri)
	sig := makeSignature(ri.Timestamp, ri.Nonce)
	fmt.Println(sig)
	if sig == ri.Signature {
		c.String(200, ri.Echostr)
		return
	}
	c.JSON(200, api.FailResponse(3004, "参数错误"))
}
