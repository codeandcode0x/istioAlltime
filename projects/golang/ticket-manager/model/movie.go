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

func (u *Movie) Create(movie *Movie) (uint64, error) {
	result := db.DBConn.Model(&Movie{}).Create(&movie)
	return movie.ID, result.Error
}

func (u *Movie) Update(id uint64, movie *Movie) (int64, error) {
	result := db.DBConn.Model(&Movie{}).Where("id = ?", id).Save(&movie)
	return result.RowsAffected, result.Error
}

func (u *Movie) Delete(id uint64) (int64, error) {
	result := db.DBConn.Model(&Movie{}).Where("id = ?", id).Delete(&Movie{})
	return result.RowsAffected, result.Error
}

func (u *Movie) FindAll(count int) ([]Movie, error) {
	var Movies []Movie
	result := db.DBConn.Model(&Movie{}).Order("id desc").Limit(count).Find(&Movies)
	return Movies, result.Error
}

func (u *Movie) FindByID(id uint64) (*Movie, error) {
	var movie *Movie
	result := db.DBConn.Model(&Movie{}).First(&movie, id)
	return movie, result.Error
}
