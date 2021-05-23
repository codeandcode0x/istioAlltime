package db

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "errors"
  "time"
)

var DBConn *gorm.DB

func init() {
	dsn := "root:root123@tcp(192.168.31.111:3306)/ticket-ops?charset=utf8mb4&parseTime=True&loc=Local"
	err := errors.New("connect database error !")
  	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

  	if err != nil {
  		panic(err.Error())
  	}

  	sqlDB, errSql := DBConn.DB()

  	if errSql != nil {
  		panic(" seting database error !")
  	}

  	// db setting
 	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
  	sqlDB.SetMaxIdleConns(10)
  	// SetMaxOpenConns 设置打开数据库连接的最大数量。
  	sqlDB.SetMaxOpenConns(100)
  	// SetConnMaxLifetime 设置了连接可复用的最大时间。
  	sqlDB.SetConnMaxLifetime(time.Hour)
}



