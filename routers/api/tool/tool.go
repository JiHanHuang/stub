package tool

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/JiHanHuang/stub/pkg/logging"
	"github.com/gin-gonic/gin"
)

// @Tags Tool
// @Summary 获取数据
// @Produce  json
// @Param app_key query string true "appkey"
// @Param data body string true "data" default({"app_id":"xxxxx",...})
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/tool/fingerprint [post]
func FingerPrint(c *gin.Context) {
	appG := app.Gin{C: c}
	appKey := c.Query("app_key")
	if appKey == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "Please input app_key")
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(body), &mapResult); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_INVALID_JSON, nil)
		return
	}
	if _, ok := mapResult["app_id"]; !ok {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "Json struct must include app_id")
		return
	}
	inputFP := make(map[string]string)
	for k, v := range mapResult {
		inputFP[k] = Strval(v)
	}
	inputFP["app_key"] = appKey
	logging.Debug("input map:", inputFP)
	appG.Response(http.StatusOK, e.SUCCESS, fingerPrintEncode(inputFP))
}

func fingerPrintEncode(args map[string]string) string {
	var keySlice []string
	var dataSlice []string
	for key, _ := range args {
		keySlice = append(keySlice, key)
	}
	sort.Strings(keySlice)
	for _, key := range keySlice {
		dataSlice = append(dataSlice, key+"="+args[key])
	}
	oneStr := strings.Join(dataSlice, "&")

	//SHA1
	h := sha1.New()
	h.Write([]byte(oneStr))
	bs := h.Sum(nil)
	//base64
	return base64.StdEncoding.EncodeToString(bs)
}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
