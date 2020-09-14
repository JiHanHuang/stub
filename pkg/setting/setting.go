package setting

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/JiHanHuang/stub/docs/version"
	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	LogStdOut   bool
	LogLevel    string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	HttpsEn      bool
	HttpsPort    int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	flag.IntVar(&(ServerSetting.HttpPort), "p", 8000, "-p	Listen port")
	flag.BoolVar(&(ServerSetting.HttpsEn), "tls", false, "-tls    Use TLS, defualt port is 8888")
	flag.IntVar(&(ServerSetting.HttpsPort), "tls-port", 8888, "--tls-port	TLS listen port")
	flag.BoolVar(&(AppSetting.LogStdOut), "log-std", false, "--log-std    Print log to terminal")
	flag.StringVar(&(AppSetting.LogLevel), "log-level", "info", "--log-level    set log level debug, info[defualt], error, warn, fatal")
	var runmod bool
	flag.BoolVar(&runmod, "debug", false, "-debug    Run debug mode")
	var v bool
	flag.BoolVar(&v, "v", false, "-v    Version")
	flag.Parse()
	if v {
		fmt.Printf("Service Stub V%s\n", version.Version)
		os.Exit(0)
	}
	if runmod {
		ServerSetting.RunMode = "debug"
	} else {
		ServerSetting.RunMode = "release"
	}
}
