package v1

import (
	"encoding/base64"
	"net/http"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Tags API
// @Summary 校验账号密码，若密码为账号的base64，则返回yes，否则返回no。支持自定义返回 [get|post]
// @Produce  json
// @Param yes query string false "匹配自定义返回(可选)"
// @Param no query string false "不匹配自定义返回(可选)"
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/check/pass [get]
func CheckPass(c *gin.Context) {
	appG := app.Gin{C: c}
	yes := c.Query("yes")
	no := c.Query("no")
	name := c.Query("name")
	password := c.Query("password")
	if name == "" || password == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	if yes == "" {
		yes = "yes"
	}
	if no == "" {
		no = "no"
	}
	var key string
	if base64.StdEncoding.EncodeToString([]byte(name)) == password {
		key = yes
	} else {
		key = no
	}
	appG.C.String(http.StatusOK, "%s", key)
}
