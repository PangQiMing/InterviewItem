package entity

import (
	"gorm.io/gorm"
)

type UploadedFile struct {
	gorm.Model
	UserID   string
	FileName string // 文件名
	FileType string // 文件类型
	FileSize int64  // 文件大小
	FileURL  string //文件地址
}
