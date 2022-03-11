package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/JiHanHuang/stub/pkg/file"
	"github.com/JiHanHuang/stub/pkg/logging"
	"gopkg.in/ini.v1"
)

var responseFileName = "set_response.ini"
var responseFilePath = "runtime/app/"

var fini *ini.File

func SetResponseExtData(data interface{}, sectionName string) error {

	f, er := file.MustOpen(responseFileName, responseFilePath)
	if er != nil {
		return er
	}
	f.Close()
	var err error
	if fini == nil {
		fini, err = ini.Load(responseFilePath + responseFileName)
		if err != nil {
			return err
		}
	}
	section := fini.Section(sectionName)
	err = section.ReflectFrom(data)
	if err != nil {
		return err
	}
	err = fini.SaveToIndent(responseFilePath+responseFileName, "\t")
	if err != nil {
		return err
	}
	return nil
}

// Response setting gin.JSON
func (g *Gin) ResponseExt(sectionName string) {
	var err error
	if fini == nil {
		fini, err = ini.Load(responseFilePath + responseFileName)
		if err != nil {
			g.Response(http.StatusNotFound, e.INVALID_PARAMS, "Not set ext response")
			return
		}
	}
	section, er := fini.GetSection(sectionName)
	if er != nil {
		g.Response(http.StatusNotFound, e.INVALID_PARAMS, "Not find the name in ext response")
		return
	}
	code := section.Key("Code").MustInt()
	body := section.Key("Body").String()
	ct := section.Key("Header").String()
	var header map[string]string
	err = json.Unmarshal([]byte(ct), &header)
	if err != nil {
		g.Response(http.StatusInternalServerError, e.ERROR, "json unmarshl failed")
		return
	}
	logging.Debug("ResponseExt set header:", header)
	for k, v := range header {
		g.C.Header(k, v)
	}
	g.C.String(code, body)
	return
}

// Response setting gin.JSON
func ListResponseExtData() []string {
	var err error
	if fini == nil {
		fini, err = ini.Load(responseFilePath + responseFileName)
		if err != nil {
			logging.Error("ListResponseExtData Err: ", err.Error())
			return nil
		}
	}
	sections := fini.Sections()
	var bodys []string
	for _, section := range sections {
		m := section.KeysHash()
		if len(m) <= 0 {
			continue
		}
		value := fmt.Sprintf("%s: ", section.Name())
		for k, v := range m {
			value = fmt.Sprintf("%s %s:%s", value, k, v)
		}
		bodys = append(bodys, value)
	}
	return bodys
}

func DelResponseExtData(sectionName string) error {
	f, er := file.MustOpen(responseFileName, responseFilePath)
	if er != nil {
		return er
	}
	f.Close()
	var err error
	if fini == nil {
		fini, err = ini.Load(responseFilePath + responseFileName)
		if err != nil {
			return err
		}
	}
	fini.DeleteSection(sectionName)
	err = fini.SaveToIndent(responseFilePath+responseFileName, "\t")
	if err != nil {
		return err
	}
	return nil
}
