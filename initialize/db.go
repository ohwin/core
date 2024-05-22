package initialize

import (
	"core/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/yanxing?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 设置日志级别为Info，以便打印出所有SQL语句
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return
	}
	global.DB = db
}
