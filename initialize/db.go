package initialize

import (
	"fmt"
	"github.com/ohwin/core/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DB() {
	db, err := gorm.Open(mysql.Open(Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return
	}
	global.DB = db
}

func Dsn() string {
	config := global.Config.Mysql
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.DB)
}
