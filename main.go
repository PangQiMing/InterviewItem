package main

import (
	"github.com/PangQiMing/InterviewItem/config"
	"github.com/PangQiMing/InterviewItem/ffmpeg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func main() {
	r := gin.Default()
	DB = config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(DB)

	r.Static("/download", "./download")
	r.POST("/upload", ffmpeg.UploadVideoHandler)
	r.POST("/clip", ffmpeg.ClipHandler)
	err := r.Run(":8080")
	if err != nil {
		panic("服务启动失败...")
	}
}
