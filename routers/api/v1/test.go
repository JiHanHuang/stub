package v1

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/JiHanHuang/stub/pkg/file"
)

var filesPath = "./runtime/files/"

// @Tags Test
// @Summary 获取数据
// @Produce  json
// @Param name query string false "自定义返回(可选)"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/get [get]
func Tget(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	if name == "" {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
		return
	}
	appG.ResponseExt(name)
}

// @Tags Test
// @Summary 上传数据
// @Produce  json
// @Param post body string false "post" default({"data":"helllo"})
// @Param name query string false "自定义返回(可选)"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/post [post]
func Tpost(c *gin.Context) {
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

// @Tags Test
// @Summary 上传数据
// @Produce  json
// @Param post body string false "post" default({"data":"helllo"})
// @Param yes query string false "自定义返回(可选)"
// @Param no query string false "自定义返回(可选)"
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/check [get]
// @Router /api/v1/check [post]
func Tcheck(c *gin.Context) {
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

// @Tags Test
// @Summary get url信息获取
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/geturl [get]
func TgetUrl(c *gin.Context) {
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

// @Tags Test
// @Summary post url信息获取
// @Produce  json
// @Param data body string false "Data" default({"data":"helllo"})
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/posturl [post]
func TpostUrl(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make([]string, 0, 4)
	body, _ := ioutil.ReadAll(c.Request.Body)
	for k, v := range c.Request.Header {
		data = append(data, fmt.Sprintf("%s:%v", k, v))
	}
	if !isJSON(string(body)) {
		appG.Response(http.StatusBadRequest, e.ERROR_INVALID_JSON, nil)
		return
	}
	data = append(data, "body:"+string(body))
	data = append(data, "content_len:"+strconv.FormatInt(c.Request.ContentLength, 10))
	appG.Response(http.StatusOK, e.SUCCESS, &data)
}

// @Tags Test
// @Summary get url信息获取
// @Success 200 string string
// @Router /api/v1/show [get]
// @Router /api/v1/show [post]
func Show(c *gin.Context) {
	appG := app.Gin{C: c}
	var data strings.Builder
	body, _ := ioutil.ReadAll(c.Request.Body)
	data.WriteString(fmt.Sprintf("%s %s %s",
		appG.C.Request.Method, appG.C.Request.RequestURI, appG.C.Request.Proto))
	data.WriteString(fmt.Sprintf("<br> Host:%s", appG.C.Request.Host))
	for k, v := range c.Request.Header {
		data.WriteString(fmt.Sprintf("<br> %s:%v", k, v))
	}
	data.WriteString(fmt.Sprintf("<br> Content-Len:%d", c.Request.ContentLength))
	data.WriteString(fmt.Sprintf("<br><br> %s", string(body)))

	appG.C.Header("Content-Type", "text/html; charset=utf-8")
	appG.C.String(http.StatusOK, "<h4> %s </h4>", data.String())
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

// @Tags Test
// @Summary 下载文件
// @Param filename query string true "file name"
// @Success 200 {object} gin.Context
// @Router /api/v1/download [get]
func DownFile(c *gin.Context) {
	filename := c.DefaultQuery("filename", "")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filesPath + filename)
}

// @Tags Test
// @Summary 下载文件(不可靠)
// @Param filename query string true "file name"
// @Router /api/v1/download2 [get]
func DownFile2(c *gin.Context) {
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

// @Tags Test
// @Summary 上传文件
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/upload [post]
func UpFile(c *gin.Context) {
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

// @Tags Test
// @Summary 延时返回
// @Produce  json
// @Param delay query int false "延时时长(默认5s)"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/delay [get]
func Delay(c *gin.Context) {
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

// @Tags Test
// @Summary 获取一定量的数据
// @Produce  json
// @Param size query int false "数据量(k)默认0k"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/data [get]
func Data(c *gin.Context) {
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

// @Tags Test
// @Summary 获取一定量的数据(并发)
// @Produce  json
// @Param size query int false "数据量(0,1,10)默认0"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/pdata [get]
func PData(c *gin.Context) {
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
