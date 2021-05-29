package model

import . "ticket-manager/db"

// init
func init() {
	AutoMigrateAll()
}

// Migrate Model
func AutoMigrateAll() {
	_ = DBConn.Table("users").AutoMigrate(&User{})
	_ = DBConn.Table("movies").AutoMigrate(&Movie{})
}