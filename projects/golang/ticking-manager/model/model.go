package model

type Model interface {
	Create(entity *interface{}) (*interface{}, error)
	FindAll() ([]*interface{}, error)
	FindByID(uid uint64) (*interface{}, error)
	Update(uid uint64, entity *interface{}) (*interface{}, error)
	Delete(uid uint64) (uint64, error)
}