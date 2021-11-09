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
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/JiHanHuang/stub/routers/api/set"
	"github.com/JiHanHuang/stub/routers/api/tool"
	v1 "github.com/JiHanHuang/stub/routers/api/v1"
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
		<a href="/api/v1/websocket">websocket</a><br><br>
		Server addr: %s<br>
		<i>Version: %s</i>
		`, setting.ServerSetting.HttpPort, "/docs/index.html", addr, version.Version)
	})

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//test
		apiv1.POST("/check", v1.Tcheck)
		apiv1.GET("/check", v1.Tcheck)
		apiv1.POST("/post", v1.Tpost)
		apiv1.GET("/get", v1.Tget)
		apiv1.GET("/show/*any", v1.Show)
		apiv1.GET("/geturl/*any", v1.TgetUrl)
		apiv1.POST("/posturl/*any", v1.TpostUrl)
		apiv1.GET("/download2", v1.DownFile2)
		apiv1.GET("/download/*any", v1.DownFile)
		apiv1.POST("/upload/", v1.UpFile)
		apiv1.GET("/websocket", v1.Home)
		apiv1.GET("/websocket/echo", v1.Echo)
		apiv1.GET("/delay", v1.Delay)
		apiv1.GET("/data", v1.Data)
		apiv1.GET("/pdata", v1.PData)
	}
	apiSet := r.Group("/api/set")
	{
		apiSet.POST("/response", set.SetResponse)
		apiSet.GET("/list", set.GetResponse)
	}
	apiTool := r.Group("/api/tool")
	{
		apiTool.POST("/fingerprint", tool.FingerPrint)
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
