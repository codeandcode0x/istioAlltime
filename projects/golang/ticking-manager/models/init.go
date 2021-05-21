package models

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql" //加载mysql
    "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Setup() {
    var err error
    // var user userModel.User
    DB, err = gorm.Open("localhost:3306", "root:12345678@/iu56?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }

    if DB.Error != nil {
        fmt.Printf("database error %v", DB.Error)
    }
    AutoMigrateAll()
}

func AutoMigrateAll() {
    DB.AutoMigrate(&User{})
}
