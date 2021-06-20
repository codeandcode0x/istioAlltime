package model

// Show entity
type Show struct {
	ID     uint64 `json:"id,omitempty"`
	Name   string `json:"name,omitempty" gorm:"type:varchar(500)"`
	Image  string `json:"image,omitempty" gorm:"type:varchar(255)"`
	Actors string `json:"actors,omitempty" gorm:"type:varchar(1000)"`
	Mtype  string `json:"mtype,omitempty" gorm:"type:varchar(255)"`
	Minfo  string `json:"minfo,omitempty" gorm:"type:varchar(255)"`
	Mtime  string `json:"mtime,omitempty" gorm:"type:varchar(255)"`
	BaseModel
}

// instance model
type ShowDAO interface {
	BaseDAO
}
