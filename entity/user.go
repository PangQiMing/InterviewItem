package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `gorm:"type:varchar(20);not null" json:"name"`                // 用户名
	Account      string `gorm:"uniqueIndex;type:varchar(20);not null" json:"Account"` // 邮箱
	Password     string `gorm:"not null" json:"-"`                                    // 密码
	PhoneNumber  string `gorm:"type:varchar(20)" json:"phone_number"`                 //手机号码
	Token        string `gorm:"-" json:"token"`
	UploadedFile []UploadedFile
}
