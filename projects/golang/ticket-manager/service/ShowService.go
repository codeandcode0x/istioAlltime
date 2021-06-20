package service

import (
	"ticket-manager/model"
)

// struct show service
type ShowService struct {
	DAO model.ShowDAO
}

// get Show service
func (u *ShowService) getCtl() *ShowService {
	var m model.BaseModel
	return &ShowService{
		DAO: &model.Show{BaseModel: m},
	}
}

// find all Shows
func (u *ShowService) FindAllShows() ([]model.Show, error) {
	shows := []model.Show{}
	return shows, u.getCtl().DAO.FindAll(&shows)
}

// find shows by page
func (u *ShowService) FindShowByPagesWithKeys(mtype string, currentPage, pageSize int) ([]model.Show, error) {
	keys := map[string]interface{}{"mtype": mtype}
	shows := []model.Show{}
	return shows, u.getCtl().DAO.FindByPagesWithKeys(&shows, keys, currentPage, pageSize)
}

// create Show
func (u *ShowService) CreateShow(show *model.Show) error {
	return u.getCtl().DAO.Create(show)
}

// update Show
func (u *ShowService) UpdateShow(show *model.Show) (int64, error) {
	rowsAffected, err := u.getCtl().DAO.Update(show, show.ID)
	return rowsAffected, err
}

// delete Show
func (u *ShowService) DeleteShow(uid uint64) (int64, error) {
	return u.getCtl().DAO.Delete(&model.Show{}, uid)
}
