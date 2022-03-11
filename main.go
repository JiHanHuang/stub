package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/stub/pkg/file"
	"github.com/JiHanHuang/stub/pkg/gredis"
	"github.com/JiHanHuang/stub/pkg/logging"
	"github.com/JiHanHuang/stub/pkg/setting"
	"github.com/JiHanHuang/stub/pkg/util"
	"github.com/JiHanHuang/stub/routers"
)

func init() {
	setting.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Stub
// @version 1.0
// @description A server using to test
// @termsOfService https://github.com/JiHanHuang/stub

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()

	if setting.ServerSetting.HttpsEn {
		portTLS := fmt.Sprintf(":%d", setting.ServerSetting.HttpsPort)
		serverTLS := &http.Server{
			Addr:    portTLS,
			Handler: routersInit,
		}
		if file.CheckNotExist("server.key") {
			// cmd := "openssl genrsa -out server.key 2048"
			// if _, err := util.Cmder(cmd); err != nil && err.Error() != "Generating RSA private key, 2048 bit long modulus" {
			// 	log.Fatal("[ERRO] ", err)
			// 	return
			// }
			log.Fatalf("[%s] Please run [openssl genrsa -out server.key 2048] to create server.key.",
				logging.LevelFlags[logging.INFO])
		}
		if file.CheckNotExist("server.crt") {
			log.Fatalf("[%s] Please run [openssl req -new -x509 -key server.key -out server.crt -days 365] to create server.crt.",
				logging.LevelFlags[logging.INFO])
		}
		log.Printf("[%s] start https server listening %s", logging.LevelFlags[logging.INFO], portTLS)
		go serverTLS.ListenAndServeTLS("server.crt", "server.key")
	}

	port := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	server := &http.Server{
		Addr:    port,
		Handler: routersInit,
	}

	log.Printf("[%s] Start http server...	Port[%d]", logging.LevelFlags[logging.INFO], setting.ServerSetting.HttpPort)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("[%s] Start http server failed. ERR:%s", logging.LevelFlags[logging.ERROR], err.Error())
	}
}
