package model

// Info entity
type Info struct {
	ID      uint64 `json:"id,omitempty"`
	Image   string `json:"image,omitempty" gorm:"type:varchar(255)"`
	Title   string `json:"title,omitempty" gorm:"type:varchar(255)"`
	Content string `json:"content,omitempty" gorm:"type:text"`
	Mtype   string `json:"mtype,omitempty" gorm:"type:varchar(255)"`
	Minfo   string `json:"minfo,omitempty" gorm:"type:varchar(255)"`
	Mtime   string `json:"mtime,omitempty" gorm:"type:varchar(255)"`
	BaseModel
}

// instance model
type InfoDAO interface {
	BaseDAO
}
