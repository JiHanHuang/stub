package v1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/JiHanHuang/stub/pkg/file"
)

var filesPath = "./runtime/files/"

// @Tags API
// @Summary 下载文件(需要文件上传接口上传) [get]
// @Param filename query string true "file name"
// @Success 200 {object} gin.Context
// @Router /api/v1/file/download [get]
func FileDown(c *gin.Context) {
	filename := c.DefaultQuery("filename", "")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filesPath + filename)
}

// @Tags API
// @Summary 下载文件(不可靠)(需要文件上传接口上传) [get]
// @Param filename query string true "file name"
// @Router /api/v1/file/download2 [get]
func FileDown2(c *gin.Context) {
	appG := app.Gin{C: c}
	filename := c.Query("filename")
	if filename == "" {
		fstr := strings.SplitN(c.Request.RequestURI, "/api/v1/download", 2)
		if len(fstr) < 2 || fstr[1] == "/" {
			appG.Response(http.StatusBadRequest, e.ERROR_INVALID_JSON, nil)
			return
		}
		filename = fstr[1][1:]
	}
	file, err := os.Open(filesPath + filename) //Create a file
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	defer file.Close()
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
}

// @Tags API
// @Summary 上传文件 [post]
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/file/upload [post]
func FileUp(c *gin.Context) {
	appG := app.Gin{C: c}

	f, header, err := c.Request.FormFile("file")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	filename := header.Filename

	if err := file.IsNotExistMkDir(filesPath); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	newFile, err := os.Create(filesPath + filename)
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
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Tags API
// @Summary 删除上传文件 [delete]
// @Param filename query string true "file name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/file/del [delete]
func FileDel(c *gin.Context) {
	appG := app.Gin{C: c}

	filename := c.Query("filename")
	if filename == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "Please input file name")
		return
	}
	err := os.Remove(filesPath + filename) //Create a file
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
