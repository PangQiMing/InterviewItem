package main

import (
	"github.com/PangQiMing/InterviewItem/config"
	"github.com/PangQiMing/InterviewItem/controller"
	"github.com/PangQiMing/InterviewItem/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(config.DB)

	r.Static("/download", "./download")

	authRouters := r.Group("api/auth")
	{
		authRouters.POST("register", controller.RegisterHandler)
		authRouters.POST("login", controller.LoginHandler)
	}

	userRouters := r.Group("api/user", middleware.AuthorizeJWT())
	{
		userRouters.POST("/upload", controller.UploadVideoHandler)
		userRouters.POST("/clip", controller.ClipHandler)
	}
	err := r.Run(":8080")
	if err != nil {
		panic("服务启动失败...")
	}
}
