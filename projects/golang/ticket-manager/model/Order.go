package model

// instance entity
type Order struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty" gorm:"type:varchar(255)"`
	// Birthday time.Time `json:"birthday,omitempty"`
	// MemberNumber sql.NullString `json:"member_number,omitempty" gorm:"type:varchar(255)"`
	Money uint64 `json:"money,omitempty"`
	BaseModel
}

// instance model
type OrderDAO interface {
	BaseDAO
	FindByUserId(entity interface{}) error
}

// custom order dto
func (u *Order) FindByUserId(entity interface{}) error {
	return nil
}
