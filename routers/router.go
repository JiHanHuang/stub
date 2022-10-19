package routers

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/JiHanHuang/stub/docs/swagger"
	"github.com/JiHanHuang/stub/docs/version"
	"github.com/JiHanHuang/stub/middleware/info"
	"github.com/JiHanHuang/stub/pkg/setting"
	admin "github.com/JiHanHuang/stub/routers/admin/v1"
	api "github.com/JiHanHuang/stub/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	if setting.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
	} else {
		r.Use(info.MSG())
	}

	r.Use(gin.Recovery())

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/hello"
		r.HandleContext(c)
		//c.Redirect(http.StatusMovedPermanently, "https://baidu.com")
	})
	addr, err := getServerIP()
	if err != nil {
		addr = err.Error()
	}
	r.GET("/hello", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `
		<h3>Welcome service stub.</h3>
		<script language="javascript" type="text/javascript">
		var domain = window.location.hostname;
		var httphost = "http://"+domain+":%d%s"
		document.write("<a href="+httphost+">接口介绍[HTTP]</a><br>");
		</script>
		<a href="/api/v1/ws/home">websocket</a><br><br>
		Server addr: %s<br>
		<i>Version: %s</i>
		`, setting.ServerSetting.HttpPort, "/docs/index.html", addr, version.Version)
	})

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//test
		apiv1.POST("/check/pass", api.CheckPass)
		apiv1.GET("/check/pass", api.CheckPass)

		apiv1.POST("/data/delay", api.DataDelay)
		apiv1.GET("/data/delay", api.DataDelay)
		apiv1.POST("/data/normal", api.DataNormal)
		apiv1.GET("/data/normal", api.DataNormal)
		apiv1.POST("/data/p", api.DataP)
		apiv1.GET("/data/p", api.DataP)

		apiv1.POST("/define/resp", api.DefineResp)
		apiv1.GET("/define/resp", api.DefineResp)
		apiv1.GET("/define/file", api.DefineRespFile)

		apiv1.GET("/file/download/*any", api.FileDown)
		apiv1.GET("/file/download2", api.FileDown2)
		apiv1.POST("/file/upload", api.FileUp)
		apiv1.DELETE("/file/del", api.FileDel)

		apiv1.POST("/fingerprint", api.FingerPrint)

		apiv1.POST("/show/base/*any", api.ShowBase)
		apiv1.GET("/show/base/*any", api.ShowBase)
		apiv1.POST("/show/url/*any", api.ShowUrl)
		apiv1.GET("/show/url/*any", api.ShowUrl)

		apiv1.GET("/ws/home", api.WSHome)
		apiv1.GET("/ws/echo", api.WSEcho)

	}
	apiAdmin := r.Group("/admin/v1")
	{
		apiAdmin.POST("/define/resp", admin.SetResponse)
		apiAdmin.POST("/define/file", admin.SetResponseFile)
		apiAdmin.GET("/define/resp", admin.ListResponse)
		apiAdmin.DELETE("/define/resp", admin.DelResponse)
	}

	return r
}

func getServerIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			ports := strconv.Itoa(setting.ServerSetting.HttpPort)
			if setting.ServerSetting.HttpsEn {
				ports = fmt.Sprintf("%s,%s", ports, strconv.Itoa(setting.ServerSetting.HttpsPort))
			}
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String() + ":" + ports, nil
			}

		}
	}

	return "", errors.New("Can not find the client ip address!")
}
