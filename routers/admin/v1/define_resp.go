package set

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/JiHanHuang/stub/pkg/file"
	"github.com/JiHanHuang/stub/pkg/logging"
)

type SetResponseForm struct {
	Code   int    `form:"code" example:"200"`
	Header string `form:"header" example:"{\"content-type\":\"application/json\"}"`
	Body   string `form:"Body" example:"{\"data\":\"your response data\"}"`
}

// @Tags Admin
// @Summary 设置自定义返回
// @Produce  json
// @Param setResponse body SetResponseForm false "设自定义返回结构"
// @Param name query string true "自定义返回名" default(set_response)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/v1/define/resp [post]
func SetResponse(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = SetResponseForm{}
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

	var m map[string]string

	if err := json.Unmarshal([]byte(form.Header), &m); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, "header must be a map[string]string, err:"+err.Error())
		return
	}

	logging.Debug("Input code:", form.Code, "Header:", form.Header, "Body:", form.Body)

	if err := app.SetResponseExtData(&form, sectionName); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type SetResponseFileData struct {
	Code     int    `form:"code" example:"200"`
	FileName string `form:"file_name" example:"file_response.html"`
}

// @Tags Admin
// @Summary 设置自定义文件返回
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Param name query string true "自定义返回名" default(file_response)
// @Param code query int true "自定义http status" default(200)
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/v1/define/file [post]
func SetResponseFile(c *gin.Context) {
	appG := app.Gin{C: c}

	f, header, err := c.Request.FormFile("file")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	filename := header.Filename
	sectionName := c.Query("name")
	if sectionName == "" {
		appG.Response(http.StatusBadRequest, e.ERROR, "name parameter needed.")
		return
	}
	code, err := strconv.Atoi(c.Query("code"))
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, err.Error())
		return
	}

	if err := file.IsNotExistMkDir(app.SaveFilesPath); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	newFile, err := os.Create(app.SaveFilesPath + filename)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, f)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	form := SetResponseFileData{
		Code:     code,
		FileName: filename,
	}

	if err := app.SetResponseExtData(&form, sectionName); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Tags Admin
// @Summary 自定义返回列表
// @Produce  json
// @Success 200 {object} app.Response
// @Router /admin/v1/define/resp [get]
func ListResponse(c *gin.Context) {
	appG := app.Gin{C: c}
	bodys := app.ListResponseExtData()
	appG.Response(http.StatusOK, e.SUCCESS, bodys)
	return
}

// @Tags Admin
// @Summary 删除自定义返回列表
// @Produce  json
// @Param name query string true "自定义返回名" default(set_response)
// @Success 200 {object} app.Response
// @Router /admin/v1/define/resp [delete]
func DelResponse(c *gin.Context) {
	appG := app.Gin{C: c}
	sectionName := c.Query("name")
	if sectionName == "" {
		appG.Response(http.StatusBadRequest, e.ERROR, "name parameter needed.")
		return
	}
	section, err := app.DelResponseExtData(sectionName)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, err.Error())
		return
	}
	if section.HasKey("FileName") {
		filename := section.Key("FileName").String()
		if !file.CheckNotExist(app.SaveFilesPath + filename) {
			err = os.Remove(app.SaveFilesPath + filename)
			if err != nil {
				appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
				return
			}
		}
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}
