package entity

import (
	"gorm.io/gorm"
)

type UploadedFile struct {
	gorm.Model
	UserAccount string
	FileName    string `gorm:"type:varchar(40);not null" json:"file_name"`
	FileType    string `gorm:"type:varchar(10);not null" json:"file_type"`
	FileSize    int64  `gorm:"not null" json:"file_size"`
	FileURL     string `gorm:"type:varchar(255)" json:"file_url"`
}
