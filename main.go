package main

import (
	"alertCenter/config"
	"alertCenter/controller"
	"alertCenter/middleware"
	"fmt"
	"io"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("AlertCenter starting...")
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	logFile, err := os.OpenFile("./logs/info.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(-1)
	}
	gin.DefaultWriter = io.MultiWriter(logFile)
	gin.DefaultErrorWriter = io.MultiWriter(logFile)
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.POST("/sms", middleware.TokenRequired, controller.SMS)
	app.POST("/dingtalk", middleware.TokenRequired, controller.DingTalk)
	app.POST("/mail", middleware.TokenRequired, controller.Email)

	if err := endless.ListenAndServe(":"+fmt.Sprintf("%d", config.Server.Port), app); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
