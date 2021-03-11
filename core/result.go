package core

import (
	"net/http"
	"time"

	"github.com/anthonyzero/go-quick-api/public"
	"github.com/gin-gonic/gin"
)

//Result 统一返回类型
type Result struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

//Success 返回统一的函数
func Success(g *gin.Context, data interface{}) {
	time := time.Now().Unix()
	g.JSON(http.StatusOK, Result{
		Code:      public.OK.Code,
		Message:   public.OK.Message,
		Data:      data,
		Timestamp: time,
	})
}

//Error 错误返回 参数为error或者codeno
func Error(g *gin.Context, err error) {
	time := time.Now().Unix()
	// 解析 err或者codeno
	code, message := public.DecodeErr(err)

	// 返回json
	g.JSON(http.StatusOK, Result{
		Code:      code,
		Message:   message,
		Data:      nil,
		Timestamp: time,
	})
}
