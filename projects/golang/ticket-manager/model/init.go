package model

import "ticket-manager/db"

// init
func init() {
	AutoMigrateAll()
}

// Migrate Model
func AutoMigrateAll() {
	_ = db.Conn.Table("users").AutoMigrate(&User{})
	_ = db.Conn.Table("movies").AutoMigrate(&Movie{})
	_ = db.Conn.Table("shows").AutoMigrate(&Show{})
	_ = db.Conn.Table("infos").AutoMigrate(&Info{})
	_ = db.Conn.Table("orders").AutoMigrate(&Order{})
}
