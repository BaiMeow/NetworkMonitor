package db

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	ldb, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open db fail:%v", err)
	}
	db = ldb
	if err := db.AutoMigrate(&ASUp{}); err != nil {
		return fmt.Errorf("auto migrate fail:%v", err)
	}
	return nil
}

var ErrDatabase = fmt.Errorf("database error")
