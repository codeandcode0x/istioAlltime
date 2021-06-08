package service

import (
	"fmt"
	"ticket-manager/model"
)

type InfoService struct {
	Model *model.Info
}

// get Info service
func (u *InfoService) getInfoService() *InfoService {
	var entity *model.Info
	return &InfoService{entity}
}

// find all Infos
func (u *InfoService) FindAllInfos(count int) ([]model.Info, error) {
	return u.getInfoService().Model.FindAll(count)
}

// find Info by id
func (u *InfoService) FindInfoById(uid uint64) (*model.Info, error) {
	return u.getInfoService().Model.FindByID(uid)
}

// create Info
func (u *InfoService) CreateInfo(Info *model.Info) (uint64, error) {
	return u.getInfoService().Model.Create(Info)
}

// update Info
func (u *InfoService) UpdateInfo(uid uint64, Info *model.Info) (int64, error) {
	rowsAffected, err := u.getInfoService().Model.Update(uid, Info)
	return rowsAffected, err
}

// delete Info
func (u *InfoService) DeleteInfo(uid uint64) (int64, error) {
	fmt.Println("uid", uid)
	return u.getInfoService().Model.Delete(uid)
}
