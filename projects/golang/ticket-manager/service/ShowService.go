package service

import (
	"ticket-manager/model"
)

type ShowService struct {
	Model *model.Show
}

// get Show service
func (u *ShowService) getShowService() *ShowService {
	var entity *model.Show
	return &ShowService{entity}
}

// find all Shows
func (u *ShowService) FindAllShows(mtype string, count int) ([]model.Show, error) {
	return u.getShowService().Model.FindAll(mtype, count)
}

// find Show by id
func (u *ShowService) FindShowById(uid uint64) (*model.Show, error) {
	return u.getShowService().Model.FindByID(uid)
}

// create Show
func (u *ShowService) CreateShow(Show *model.Show) (uint64, error) {
	return u.getShowService().Model.Create(Show)
}

// update Show
func (u *ShowService) UpdateShow(uid uint64, Show *model.Show) (int64, error) {
	rowsAffected, err := u.getShowService().Model.Update(uid, Show)
	return rowsAffected, err
}

// delete Show
func (u *ShowService) DeleteShow(uid uint64) (int64, error) {
	return u.getShowService().Model.Delete(uid)
}
