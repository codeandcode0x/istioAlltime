package model

import (
	. "ticket-manager/db"
)

// init
func init() {
	AutoMigrateAll()
}

// Migrate Model
func AutoMigrateAll() {
    DBConn.Table("users").AutoMigrate(&User{})
    DBConn.Table("movies").AutoMigrate(&Movie{})
}