package service

import (
	"fmt"
	"ticket-manager/model"
)

type UserService struct {
	Model  *model.User
}

// get user service
func (u *UserService) getUserService() *UserService {
	var entity *model.User
	return &UserService{entity}
}

// find all users
func (u *UserService) FindAllUsers() ([]model.User, error) {
	return u.getUserService().Model.FindAll()
}

// find user by id
func (u *UserService) FindUserById(uid uint64) (*model.User, error) {
	return u.getUserService().Model.FindByID(uid)
}

// find user by email
func (u *UserService) FindUserByEmail(email string) (*model.User, error) {
	return u.getUserService().Model.FindByEmail(email)
}

// create user
func (u *UserService) CreateUser(user *model.User) (uint64, error) {
	return u.getUserService().Model.Create(user)
}

// update user
func (u *UserService) UpdateUser(uid uint64, user *model.User ) (int64, error) {
	rowsAffected, err := u.getUserService().Model.Update(uid, user)
	return rowsAffected, err
}

// delete user
func (u *UserService) DeleteUser(uid uint64) (int64, error) {
	fmt.Println("uid", uid)
	return u.getUserService().Model.Delete(uid)
}

