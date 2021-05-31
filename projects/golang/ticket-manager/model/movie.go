package model

import (
	"gorm.io/gorm"
	"ticket-manager/db"
)

// Movie entity
type Movie struct {
	ID           uint64  `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Image        string  `json:"image,omitempty"`
	Actors       string  `json:"actors,omitempty"`
	Mtype        string  `json:"mtype,omitempty"`
	Minfo        string  `json:"minfo,omitempty"`
	Mtime        string  `json:"mtime,omitempty"`
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
	result := db.DBConn.Model(&Movie{}).Limit(count).Find(&Movies)
	return Movies, result.Error
}


func (u *Movie) FindByID(id uint64) (*Movie, error) {
	var movie *Movie
	result := db.DBConn.Model(&Movie{}).First(&movie, id)
	return movie, result.Error
}












