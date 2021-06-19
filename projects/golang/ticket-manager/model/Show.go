package model

import (
	"ticket-manager/db"

	"gorm.io/gorm"
)

// Show entity
type Show struct {
	ID     uint64 `json:"id,omitempty"`
	Name   string `json:"name,omitempty" gorm:"type:varchar(500)"`
	Image  string `json:"image,omitempty" gorm:"type:varchar(255)"`
	Actors string `json:"actors,omitempty" gorm:"type:varchar(1000)"`
	Mtype  string `json:"mtype,omitempty" gorm:"type:varchar(255)"`
	Minfo  string `json:"minfo,omitempty" gorm:"type:varchar(255)"`
	Mtime  string `json:"mtime,omitempty" gorm:"type:varchar(255)"`
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

func (u *Show) Create(instance *Show) (uint64, error) {
	result := db.Conn.Model(&Show{}).Create(&instance)
	return instance.ID, result.Error
}

func (u *Show) Update(id uint64, instance *Show) (int64, error) {
	result := db.Conn.Model(&Show{}).Where("id = ?", id).Save(&instance)
	return result.RowsAffected, result.Error
}

func (u *Show) Delete(id uint64) (int64, error) {
	result := db.Conn.Model(&Show{}).Where("id = ?", id).Delete(&Show{})
	return result.RowsAffected, result.Error
}

func (u *Show) FindAll(mtype string, count int) ([]Show, error) {
	var Shows []Show
	result := db.Conn.Model(&Show{}).Where("mtype", mtype).Order("id desc").Limit(count).Find(&Shows)
	return Shows, result.Error
}

func (u *Show) FindByID(id uint64) (*Show, error) {
	var instance *Show
	result := db.Conn.Model(&Show{}).First(&instance, id)
	return instance, result.Error
}

func (u *Show) FindByPages(currentPage, pageSize int) ([]Show, error) {
	var instances []Show
	result := Paginate(currentPage, pageSize).Model(&Show{}).Find(&instances)
	return instances, result.Error
}
