package model

import (
	"ticket-manager/db"

	"gorm.io/gorm"
)

// Info entity
type Info struct {
	ID      uint64 `json:"id,omitempty"`
	Image   string `json:"image,omitempty" gorm:"type:varchar(255)"`
	Title   string `json:"title,omitempty" gorm:"type:varchar(255)"`
	Content string `json:"content,omitempty" gorm:"type:text"`
	Mtype   string `json:"mtype,omitempty" gorm:"type:varchar(255)"`
	Minfo   string `json:"minfo,omitempty" gorm:"type:varchar(255)"`
	Mtime   string `json:"mtime,omitempty" gorm:"type:varchar(255)"`
	gorm.Model
}

// Info model
type InfoModel interface {
	Create(Info *Info) (*Info, error)
	Update(id uint64, Info *Info) (*Info, error)
	Delete(id uint64) (uint64, error)
	FindAll() ([]Info, error)
	FindByID(id uint64) (*Info, error)
	FindByName(name string) (*Info, error)
}

func (u *Info) Create(instance *Info) (uint64, error) {
	result := db.Conn.Model(&Info{}).Create(&instance)
	return instance.ID, result.Error
}

func (u *Info) Update(id uint64, instance *Info) (int64, error) {
	result := db.Conn.Model(&Info{}).Where("id = ?", id).Save(&instance)
	return result.RowsAffected, result.Error
}

func (u *Info) Delete(id uint64) (int64, error) {
	result := db.Conn.Model(&Info{}).Where("id = ?", id).Delete(&Info{})
	return result.RowsAffected, result.Error
}

func (u *Info) FindAll(count int) ([]Info, error) {
	var Infos []Info
	result := db.Conn.Model(&Info{}).Order("id desc").Limit(count).Find(&Infos)
	return Infos, result.Error
}

func (u *Info) FindByID(id uint64) (*Info, error) {
	var instance *Info
	result := db.Conn.Model(&Info{}).First(&instance, id)
	return instance, result.Error
}

func (u *Info) FindByPages(currentPage, pageSize int) ([]Info, error) {
	var users []Info
	result := Paginate(currentPage, pageSize).Model(&Info{}).Find(&users)
	return users, result.Error
}
