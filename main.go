package main

import (
	"github.com/PangQiMing/InterviewItem/config"
	"github.com/PangQiMing/InterviewItem/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(config.DB)

	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)
	r.Static("/download", "./download")
	r.POST("/upload", controller.UploadVideoHandler)
	r.POST("/clip", controller.ClipHandler)
	err := r.Run(":8080")
	if err != nil {
		panic("服务启动失败...")
	}
}
