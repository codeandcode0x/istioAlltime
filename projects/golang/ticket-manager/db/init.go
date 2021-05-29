package db

import (
  "errors"
  "github.com/spf13/viper"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "log"
  "os"
  "ticket-manager/util"
  "time"
  "fmt"
)

// DBConn
var DBConn *gorm.DB

// init
func init() {
  // init config
  util.InitConfig()
  // get db config
  dbHost := os.Getenv("DB_HOST")
  if dbHost == "" {
    dbHost = viper.GetString("DB_HOST")
  }

  // get db port
  dbPort := os.Getenv("DB_PORT")
  if dbPort == "" {
    dbPort = viper.GetString("DB_PORT")
  }

  // get db name
  dbName := os.Getenv("DB_DATABASE")
  if dbName == "" {
    dbName = viper.GetString("DB_DATABASE")
  }

  // get db user
  dbUser := os.Getenv("DB_USER")
  if dbUser == "" {
    dbUser = viper.GetString("DB_USER")
  }

  // get db passwd
  dbPasswd := os.Getenv("DB_PASSWD")
  if dbPasswd == "" {
    dbPasswd = viper.GetString("DB_PASSWD")
  }

  // set db log level
  logLevel := logger.Info
  dbLogMode := os.Getenv("DB_LOGMODE")
  if dbLogMode == "" {
    dbLogMode = viper.GetString("DB_LOGMODE")
    if dbLogMode == "Warn" {
      logLevel = logger.Warn
    }
  }

  fmt.Println("get env :", dbHost, dbPort, dbUser, dbName, dbLogMode)

  DBLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold:              3 * time.Second,   // Slow SQL threshold
      LogLevel:                   logLevel,     // Log level
      IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
      Colorful:                  false,          // Disable color
    },
    )

  // db dsn
  dsn := dbUser+":"+dbPasswd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"
  err := errors.New("connect database error !")
  DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: DBLogger,
  })
  // connect err
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



