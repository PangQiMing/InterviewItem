package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Name         string `gorm:"type:varchar(20);not null" json:"name"`               // 用户名
	Account      string `gorm:"primaryKey;type:varchar(20);not null" json:"account"` // 邮箱
	Password     string `gorm:"not null" json:"-"`                                   // 密码
	PhoneNumber  string `gorm:"type:varchar(20)" json:"phone_number"`                //手机号码
	Token        string `gorm:"-" json:"token"`
	UploadedFile []UploadedFile
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
}
