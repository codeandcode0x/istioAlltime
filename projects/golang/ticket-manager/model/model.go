package model

import (
	"ticket-manager/db"
	"time"

	"gorm.io/gorm"
)

// base model
type BaseModel struct {
	ID        uint64 `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// base dao
type BaseDAO interface {
	Create(entity interface{}) error
	Update(entity interface{}, uid uint64) (int64, error)
	Delete(entity interface{}, uid uint64) (int64, error)
	FindAll(entity interface{}) error
	FindByKeys(entity interface{}, keys map[string]interface{}) error
	FindByPages(entity interface{}, currentPage, pageSize int) error
}

// create
func (u *BaseModel) Create(entity interface{}) error {
	result := db.Conn.Model(entity).Create(entity)
	return result.Error
}

// update
func (u *BaseModel) Update(entity interface{}, uid uint64) (int64, error) {
	result := db.Conn.Model(entity).Where("id = ?", uid).Save(&entity)
	return result.RowsAffected, result.Error
}

// delete
func (u *BaseModel) Delete(entity interface{}, uid uint64) (int64, error) {
	result := db.Conn.Model(entity).Where("id = ?", uid).Delete(entity)
	return result.RowsAffected, result.Error
}

// find all
func (u *BaseModel) FindAll(entity interface{}) error {
	/**
	// 反射获取类型
	modelType := reflect.ValueOf(entity).Type()
	log.Println("model ........", modelType)
	// switch type 获取类型
	switch entity.(type) {
	case *Order:
		break
	}
	*/
	result := db.Conn.Model(entity).Order("id desc").Find(entity)
	return result.Error
}

// find by id
func (u *BaseModel) FindByKeys(entity interface{}, keys map[string]interface{}) error {
	result := db.Conn.Model(entity).Where(keys).Find(&entity)
	return result.Error
}

// find by pages
func (u *BaseModel) FindByPages(entity interface{}, currentPage, pageSize int) error {
	result := Paginate(currentPage, pageSize).Model(entity).Find(entity)
	return result.Error
}

// 分页 paginate
func Paginate(currentPage, pageSize int) *gorm.DB {
	offset := (currentPage - 1) * pageSize
	return db.Conn.Offset(offset).Limit(pageSize)
}
