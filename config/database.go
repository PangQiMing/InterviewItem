package config

import (
	"fmt"
	"github.com/PangQiMing/InterviewItem/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// SetupDatabaseConnection 连接数据库,并且返回数据库接口
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("加载配置文件出错...")
	}
	//获取配置文件.env里的信息
	dbUSER := os.Getenv("DB_USER")
	dbPASS := os.Getenv("DB_PASS")
	dbHOST := os.Getenv("DB_HOST")
	dbNAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUSER, dbPASS, dbHOST, dbNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	err = db.AutoMigrate(&entity.User{}, &entity.UploadedFile{})
	if err != nil {
		panic(err)
	}
	return db
}

// CloseDatabaseConnection 关闭数据库连接
func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("关闭数据库失败")
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
