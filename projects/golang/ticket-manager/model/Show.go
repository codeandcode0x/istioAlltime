package model

import (
	"ticket-manager/db"

	"gorm.io/gorm"
)

// Show entity
type Show struct {
	ID     uint64 `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Image  string `json:"image,omitempty"`
	Actors string `json:"actors,omitempty"`
	Mtype  string `json:"mtype,omitempty"`
	Minfo  string `json:"minfo,omitempty"`
	Mtime  string `json:"mtime,omitempty"`
	gorm.Model
}

// Show model
type ShowModel interface {
	Create(Show *Show) (*Show, error)
	Update(id uint64, Show *Show) (*Show, error)
	Delete(id uint64) (uint64, error)
	FindAll() ([]Show, error)
	FindByID(id uint64) (*Show, error)
	FindByName(name string) (*Show, error)
}

func (u *Show) Create(show *Show) (uint64, error) {
	result := db.DBConn.Model(&Show{}).Create(&show)
	return show.ID, result.Error
}

func (u *Show) Update(id uint64, show *Show) (int64, error) {
	result := db.DBConn.Model(&Show{}).Where("id = ?", id).Save(&show)
	return result.RowsAffected, result.Error
}

func (u *Show) Delete(id uint64) (int64, error) {
	result := db.DBConn.Model(&Show{}).Where("id = ?", id).Delete(&Show{})
	return result.RowsAffected, result.Error
}

func (u *Show) FindAll(mtype string, count int) ([]Show, error) {
	var Shows []Show
	result := db.DBConn.Model(&Show{}).Where("mtype", mtype).Order("id desc").Limit(count).Find(&Shows)
	return Shows, result.Error
}

func (u *Show) FindByID(id uint64) (*Show, error) {
	var show *Show
	result := db.DBConn.Model(&Show{}).First(&show, id)
	return show, result.Error
}
