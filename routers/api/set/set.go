package set

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
)

type SetResponseForm struct {
	Code        int    `form:"code" example:"200"`
	ContentType string `form:"type" example:"json"`
	Data        string `form:"data" example:"your response data" valid:"Required"`
}

// @Tags Set
// @Summary 设置自定义返回
// @Produce  json
// @Param setResponse body SetResponseForm false "设自定义返回结构"
// @Param name query string true "自定义返回名" default(set_response)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/set/response [post]
func SetResponse(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = SetResponseForm{
			Code:        200,
			ContentType: "application/json",
		}
	)
	sectionName := c.Query("name")
	if sectionName == "" {
		appG.Response(http.StatusBadRequest, e.ERROR, "name parameter needed.")
		return
	}
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	if err := app.SetResponseExtData(&form, sectionName); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Tags Set
// @Summary 自定义返回列表
// @Produce  json
// @Success 200 {object} app.Response
// @Router /api/set/list [get]
func GetResponse(c *gin.Context) {
	appG := app.Gin{C: c}
	bodys := app.ListResponseExtData()
	appG.Response(http.StatusOK, e.SUCCESS, bodys)
	return
}
