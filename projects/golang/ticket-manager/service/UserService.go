package service

import (
	"ticket-manager/model"
)

type UserService struct {
	DAO model.UserDAO
}

// get user service
func (u *UserService) getSvc() *UserService {
	var m model.BaseModel
	return &UserService{
		DAO: &model.User{BaseModel: m},
	}
}

// find all users
func (u *UserService) FindAllUsers() (interface{}, error) {
	var users []model.User
	/**
	// use DAO Option
	err := u.getSvc().DAO.FindAll(&users, model.DAOOption{
		Order:  "id desc",
		Limit:  1,
		Where:  map[string]interface{}{"name": "itcast"},
		Select: "name,email",
	})
	*/
	err := u.getSvc().DAO.FindAll(&users)
	return users, err
}

// find user by id
func (u *UserService) FindUserById(uid uint64) (*model.User, error) {
	return u.getSvc().DAO.FindByID(uid)
}

// find user by email
func (u *UserService) FindUserByEmail(email string) (*model.User, error) {
	return u.getSvc().DAO.FindByEmail(email)
}

// create user
func (u *UserService) CreateUser(user *model.User) error {
	return u.getSvc().DAO.Create(user)
}

// update user
func (u *UserService) UpdateUser(user *model.User) (int64, error) {
	rowsAffected, err := u.getSvc().DAO.Update(user, user.ID)
	return rowsAffected, err
}

// delete user
func (u *UserService) DeleteUser(uid uint64) (int64, error) {
	return u.getSvc().DAO.Delete(&model.User{}, uid)
}

// find all users
// func (s *UserService) FindAllUserByPages(currentPage, pageSize int) ([]model.User, error) {
// 	return s.getSvc().DAO.FindByPages(currentPage, pageSize)
// }
