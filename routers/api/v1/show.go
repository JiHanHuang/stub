package v1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Tags API
// @Summary 请求该接口，打印出请求相关信息 [get|post]
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/show/url [get]
func ShowUrl(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make([]string, 0, 4)
	body, _ := ioutil.ReadAll(c.Request.Body)
	for k, v := range c.Request.Header {
		data = append(data, fmt.Sprintf("%s:%v", k, v))
	}
	data = append(data, "body:"+string(body))
	data = append(data, "url:"+appG.C.Request.RequestURI)
	data = append(data, "content_len:"+strconv.FormatInt(c.Request.ContentLength, 10))
	appG.Response(http.StatusOK, e.SUCCESS, &data)
}

// @Tags API
// @Summary 请求该接口，打印出请求相关信息，更友好显示 [get|post]
// @Success 200 string string
// @Router /api/v1/show/base [get]
func ShowBase(c *gin.Context) {
	appG := app.Gin{C: c}
	var data strings.Builder
	body, _ := ioutil.ReadAll(c.Request.Body)
	data.WriteString(fmt.Sprintf("%s %s %s\n",
		appG.C.Request.Method, appG.C.Request.RequestURI, appG.C.Request.Proto))
	data.WriteString(fmt.Sprintf("<br> Host:%s\n", appG.C.Request.Host))
	for k, v := range c.Request.Header {
		data.WriteString(fmt.Sprintf("<br> %s:%v\n", k, v))
	}
	data.WriteString(fmt.Sprintf("<br> Content-Len:%d\n", c.Request.ContentLength))
	data.WriteString(fmt.Sprintf("<br><br> %s\n", string(body)))

	appG.C.Header("Content-Type", "text/html; charset=utf-8")
	appG.C.String(http.StatusOK, "<h4> %s </h4>", data.String())
}
