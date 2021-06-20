package model

import (
	"database/sql"
	"ticket-manager/db"
	"time"
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
	BaseModel
}

// instance model
type UserDAO interface {
	BaseDAO
	FindByID(uid uint64) (*User, error)
	FindByName(name string) (*User, error)
	FindByEmail(email string) (*User, error)
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
