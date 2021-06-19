package model

import (
	"database/sql"
	"ticket-manager/db"
	"time"

	"gorm.io/gorm"
)

// instance entity
type User struct {
	ID           uint64         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty" gorm:"type:varchar(255)"`
	Password     string         `json:"password,omitempty" gorm:"type:varchar(1000)"`
	Email        string         `json:"email,omitempty" gorm:"type:varchar(255)"`
	Age          int            `json:"age,omitempty"`
	Birthday     time.Time      `json:"birthday,omitempty"`
	MemberNumber sql.NullString `json:"member_number,omitempty" gorm:"type:varchar(255)"`
	Role         string         `json:"Role,omitempty" gorm:"type:varchar(100)"`
	gorm.Model
}

// instance model
type UserModel interface {
	Create(instance *User) (*User, error)
	Update(uid uint64, instance *User) (*User, error)
	Delete(uid uint64) (uint64, error)
	FindAll() ([]User, error)
	FindByID(uid uint64) (*User, error)
	FindByName(name string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByPages(currentPage, pageSize int) ([]User, error)
}

func (u *User) Create(instance *User) (uint64, error) {
	result := db.Conn.Model(&User{}).Create(&instance)
	return instance.ID, result.Error
}

func (u *User) Update(uid uint64, instance *User) (int64, error) {
	result := db.Conn.Model(&User{}).Where("id = ?", uid).Save(&instance)
	return result.RowsAffected, result.Error
}

func (u *User) Delete(uid uint64) (int64, error) {
	result := db.Conn.Model(&User{}).Where("id = ?", uid).Delete(&User{})
	return result.RowsAffected, result.Error
}

func (u *User) FindAll() ([]User, error) {
	var instances []User
	result := db.Conn.Model(&User{}).Order("id desc").Find(&instances)
	return instances, result.Error
}

func (u *User) FindByID(uid uint64) (*User, error) {
	var instance *User
	result := db.Conn.Model(&User{}).First(&instance, uid)
	return instance, result.Error
}

func (u *User) FindByEmail(email string) (*User, error) {
	var instance *User
	result := db.Conn.Model(&User{}).Where("email", email).First(&instance)
	return instance, result.Error
}

func (u *User) FindByName(name string) (*User, error) {
	return &User{}, nil
}

func (u *User) FindByPages(currentPage, pageSize int) ([]User, error) {
	var instances []User
	result := Paginate(currentPage, pageSize).Model(&User{}).Find(&instances)
	return instances, result.Error
}
