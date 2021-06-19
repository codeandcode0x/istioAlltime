package model

import (
	"ticket-manager/db"

	"gorm.io/gorm"
)

// Movie entity
type Movie struct {
	ID     uint64 `json:"id,omitempty"`
	Name   string `json:"name,omitempty" gorm:"type:varchar(255)"`
	Image  string `json:"image,omitempty" gorm:"type:varchar(255)"`
	Actors string `json:"actors,omitempty" gorm:"type:varchar(1000)"`
	Mtype  string `json:"mtype,omitempty" gorm:"type:varchar(255)"`
	Minfo  string `json:"minfo,omitempty" gorm:"type:varchar(255)"`
	Mtime  string `json:"mtime,omitempty" gorm:"type:varchar(255)"`
	gorm.Model
}

// Movie model
type MovieModel interface {
	Create(Movie *Movie) (*Movie, error)
	Update(id uint64, Movie *Movie) (*Movie, error)
	Delete(id uint64) (uint64, error)
	FindAll() ([]Movie, error)
	FindByID(id uint64) (*Movie, error)
	FindByName(name string) (*Movie, error)
}

func (u *Movie) Create(instance *Movie) (uint64, error) {
	result := db.Conn.Model(&Movie{}).Create(&instance)
	return instance.ID, result.Error
}

func (u *Movie) Update(id uint64, instance *Movie) (int64, error) {
	result := db.Conn.Model(&Movie{}).Where("id = ?", id).Save(&instance)
	return result.RowsAffected, result.Error
}

func (u *Movie) Delete(id uint64) (int64, error) {
	result := db.Conn.Model(&Movie{}).Where("id = ?", id).Delete(&Movie{})
	return result.RowsAffected, result.Error
}

func (u *Movie) FindAll(count int) ([]Movie, error) {
	var Movies []Movie
	result := db.Conn.Model(&Movie{}).Order("id desc").Limit(count).Find(&Movies)
	return Movies, result.Error
}

func (u *Movie) FindByID(id uint64) (*Movie, error) {
	var instance *Movie
	result := db.Conn.Model(&Movie{}).First(&instance, id)
	return instance, result.Error
}

func (u *Movie) FindByPages(currentPage, pageSize int) ([]Movie, error) {
	var instances []Movie
	result := Paginate(currentPage, pageSize).Model(&Movie{}).Find(&instances)
	return instances, result.Error
}
