package controller

import (
	"github.com/PangQiMing/InterviewItem/dto"
	"github.com/PangQiMing/InterviewItem/entity"
	"github.com/PangQiMing/InterviewItem/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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
	filename := filepath.Join("./uploads", header.Filename)
	out, err := os.Create(filename)
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
		"message":   "文件上传成功",
		"video_url": filename,
	})
}

// ClipHandler 剪辑视频
func ClipHandler(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//做数据绑定
	var video dto.VideoDTO
	bindErr := ctx.BindJSON(&video)
	if bindErr != nil {
		log.Println("数据绑定失败...")
		return
	}

	//生成唯一的文件名称
	outFilename := filepath.Join("download", uuid.New().String()+".mp4")
	//调用ffmpeg剪辑视频
	cmd := exec.Command("ffmpeg", "-i", video.VideoURL, "-ss", video.StartTime, "-t", video.EndTime, "-c:a", "copy", outFilename)
	err := cmd.Run()
	if err != nil {
		log.Println("调用ffmpeg错误:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//把\\双斜杠替换成/斜杠
	videoURL := strings.Replace(outFilename, "\\", "/", -1)

	//获取文件详细信息
	var file entity.UploadedFile
	filename, fileType, fileSize, err := getFileInfo(videoURL)
	if err != nil {
		log.Println(err)
		return
	}

	account := utils.VerificationToken(ctx)

	log.Println(account)
	file.UserID = account
	file.FileName = filename
	file.FileType = fileType
	file.FileSize = fileSize
	file.FileURL = videoURL

	//返回剪辑后的视频URL给用户
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "剪辑完成",
		"video_url": videoURL,
	})
}

func getFileInfo(filename string) (string, string, int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", "", 0, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return "", "", 0, err
	}
	log.Println(stat.Name(), stat.ModTime(), stat.Mode().Type())
	return stat.Name(), filepath.Ext(filename), stat.Size() / 1024, nil
}
