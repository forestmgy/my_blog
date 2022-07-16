package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接错误，检查参数", err)
	}
	mysqlDB, err := db.DB()
	if err != nil {
		log.Println("数据库连接错误，检查参数", err)
	}
	//自动迁移模型
	_ = db.AutoMigrate(&User{}, &Article{}, &Category{})

	mysqlDB.SetConnMaxLifetime(10 * time.Second)
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)
	//mysqlDB.Close()
}
