/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/3 11:15 下午
 * @Description: 主程序启动入口
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/togettoyou/go-miaosha/docs"
	"github.com/togettoyou/go-miaosha/models"
	"github.com/togettoyou/go-miaosha/pkg/setting"
	"github.com/togettoyou/go-miaosha/routers"
	"log"
	"net/http"
	"time"
)

func init() {
	setting.Setup()
	models.Setup()
}

// @title go实现的秒杀系统
// @version 1.0
// @description go实现的秒杀系统api文档
// @contact.name 夜央 Oh oh oh oh oh oh
// @contact.email zoujh99@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info]  Set Time Zone to Asia/Chongqing")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		log.Printf("[error] Set Time Zone to Asia/Chongqing failed %s", err)
	}
	time.Local = timeLocal
	log.Printf("[info] start http server listening %s", endPoint)

	if err := server.ListenAndServe(); err != nil {
		log.Printf("start http server failed %s", err)
	}
}
