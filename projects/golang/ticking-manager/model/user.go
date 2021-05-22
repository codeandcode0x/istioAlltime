package model

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"ticket-manager/db"
	"time"
)

// user entity
type User struct {
	ID           uint64  `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Password     string  `json:"password,omitempty"`
  	Email        string  `json:"email,omitempty"`
  	Age          int   `json:"age,omitempty"`
  	Birthday     time.Time  `json:"birthday,omitempty"`
  	MemberNumber sql.NullString  `json:"member_number,omitempty"`
	gorm.Model
}

// user model
type UserModel interface {
	Create(user *User) (*User, error)
	Update(uid uint64, user *User) (*User, error)
	Delete(uid uint64) (uint64, error)
	FindAll() ([]User, error)
	FindByID(uid uint64) (*User, error)
	FindByName(name string) (*User, error)
	FindByEmail(email string) (*User, error)
}



func (u *User) Create(user *User) (uint64, error) {
	result := db.DBConn.Model(&User{}).Create(&user)
	return user.ID, result.Error
}

func (u *User) Update(uid uint64, user *User) (int64, error) {
	result := db.DBConn.Model(&User{}).Where("id = ?", uid).Save(&user)
	return result.RowsAffected, result.Error
}


func (u *User) Delete(uid uint64) (int64, error) {
	fmt.Println("uid", uid)
	result := db.DBConn.Model(&User{}).Where("id = ?", uid).Delete(&User{})
	return result.RowsAffected, result.Error
}


func (u *User) FindAll() ([]User, error) {
	var users []User
	result := db.DBConn.Model(&User{}).Find(&users)
	return users, result.Error
}


func (u *User) FindByID(uid uint64) (*User, error) {
	var user *User
	result := db.DBConn.Model(&User{}).First(&user, uid)
	return user, result.Error
}


func (u *User) FindByName(name string) (*User, error) {
	return &User{}, nil
}


func (u *User) FindByEmail(email string) (*User, error) {
	return &User{}, nil
}











