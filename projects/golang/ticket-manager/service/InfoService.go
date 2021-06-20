package service

import (
	"ticket-manager/model"
)

type InfoService struct {
	DAO model.InfoDAO
}

// get Info service
func (u *InfoService) getCtl() *InfoService {
	var m model.BaseModel
	return &InfoService{
		DAO: &model.Info{BaseModel: m},
	}
}

// find all Infos
func (u *InfoService) FindAllInfos() ([]model.Info, error) {
	infos := []model.Info{}
	return infos, u.getCtl().DAO.FindAll(&infos)
}

// find Infos by pages
func (u *InfoService) FindInfoByPages(currentPage, pageSize int) ([]model.Info, error) {
	infos := []model.Info{}
	return infos, u.getCtl().DAO.FindByPages(&infos, currentPage, pageSize)
}

// find Info by id
func (u *InfoService) FindInfoById(uid uint64) (*model.Info, error) {
	keys := map[string]interface{}{"id": uid}
	info := &model.Info{}
	return info, u.getCtl().DAO.FindByKeys(info, keys)
}

// create Info
func (u *InfoService) CreateInfo(info *model.Info) error {
	return u.getCtl().DAO.Create(info)
}

// update Info
func (u *InfoService) UpdateInfo(info *model.Info) (int64, error) {
	rowsAffected, err := u.getCtl().DAO.Update(info, info.ID)
	return rowsAffected, err
}

// delete Info
func (u *InfoService) DeleteInfo(uid uint64) (int64, error) {
	return u.getCtl().DAO.Delete(&model.Info{}, uid)
}
