package entity

import (
	"gorm.io/gorm"
)

type UploadedFile struct {
	gorm.Model
	UserID   uint
	FileName string // 文件名
	FileType string // 文件类型
	FileSize int64  // 文件大小
}
