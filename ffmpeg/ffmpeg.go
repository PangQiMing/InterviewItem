package ffmpeg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type ClipVideo struct {
	VideoURL  string `json:"video_url" form:"video_url"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
}

// UploadVideoHandler 上传视频
func UploadVideoHandler(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//获取文件数据
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer file.Close()

	//判断视频文件类型是否正确
	fileExt := filepath.Ext(header.Filename)
	switch fileExt {
	case ".mp4", ".avi", ".mov":
		break
	default:
		log.Println("不支持该文件类型")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "上传文件类型错误...",
		})
		return
	}

	//创建视频文件
	filename := header.Filename
	out, err := os.Create(filepath.Join("./uploads", filename))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer out.Close()

	//将视频写到磁盘中
	_, err = io.Copy(out, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s 视频上传成功!", filename),
	})
}

// ClipHandler 剪辑视频
func ClipHandler(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//做数据绑定
	var clipVideo ClipVideo
	bindErr := ctx.BindJSON(&clipVideo)
	if bindErr != nil {
		log.Println("数据绑定失败...")
		return
	}

	//调用ffmpeg剪辑视频
	savePath := "./download/output.mp4"
	cmd := exec.Command("ffmpeg", "-i", clipVideo.VideoURL, "-ss", clipVideo.StartTime, "-t", clipVideo.EndTime, "-c:a", "copy", savePath)
	err := cmd.Run()
	if err != nil {
		log.Println("调用ffmpeg错误:", err)
		return
	}

	//返回剪辑后的视频URL给用户
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "剪辑完成",
		"videoURL": savePath,
	})
}
