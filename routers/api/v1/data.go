package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Tags API
// @Summary 延时返回 [get|post]
// @Produce  json
// @Param delay query int false "延时时长(默认5s)"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/data/delay [get]
func DataDelay(c *gin.Context) {
	appG := app.Gin{C: c}
	delay := c.DefaultQuery("delay", "5")
	d, err := strconv.Atoi(delay)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	time.Sleep(time.Second * time.Duration(d))
	appG.Response(http.StatusOK, e.SUCCESS, fmt.Sprintf("delay %ds", d))
}

var data0k = "Hello"
var data1k = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`
var data10k = data1k + data1k + data1k + data1k + data1k + data1k + data1k + data1k + data1k + data1k

// @Tags API
// @Summary 获取一定量的数据 [get|post]
// @Produce  json
// @Param size query int false "数据量(k)默认0k"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/data/normal [get]
func DataNormal(c *gin.Context) {
	appG := app.Gin{C: c}
	size := c.DefaultQuery("size", "0")
	d, err := strconv.Atoi(size)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	if d > 0 {
		var data string
		for i := 0; i < d; i++ {
			data += data1k
		}
		c.String(200, data)
		return
	}
	c.String(200, data0k)
}

// @Tags API
// @Summary 获取一定量的数据(并发) [get|post]
// @Produce  json
// @Param size query int false "数据量(0,1,10)默认0"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/data/p [get]
func DataP(c *gin.Context) {
	size := c.DefaultQuery("size", "0")
	var data string
	switch size {
	case "1":
		data = data1k
	case "10":
		data = data10k
	default:
		data = data0k
	}
	c.String(200, data)
}
