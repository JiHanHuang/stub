package v1

import (
	"io/ioutil"
	"net/http"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Tags API
// @Summary 获取自定义返回数据 [get|post]
// @Produce  json
// @Param name query string false "自定义返回(可选)" default(set_response)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/define/resp [get]
func DefineResp(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	if name == "" {
		appG.Response(http.StatusOK, e.SUCCESS, string(body))
		return
	}
	appG.ResponseExt(name)
}

// @Tags API
// @Summary 获取自定义返回数据 [get|post]
// @Produce  json
// @Param name query string false "自定义返回(可选)" default(file_response)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/define/file [get]
func DefineRespFile(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	if name == "" {
		appG.Response(http.StatusOK, e.SUCCESS, string(body))
		return
	}
	appG.ResponseFile(name)
}
