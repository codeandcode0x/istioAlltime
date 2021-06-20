package service

import (
	"ticket-manager/model"
)

// struct movie service
type MovieService struct {
	DAO model.MovieDAO
}

// get Movie service
func (u *MovieService) getCtl() *MovieService {
	var m model.BaseModel
	return &MovieService{
		DAO: &model.Movie{BaseModel: m},
	}
}

// find all Movies
func (u *MovieService) FindAllMovies() ([]model.Movie, error) {
	movies := []model.Movie{}
	return movies, u.getCtl().DAO.FindAll(&movies)
}

// find Movies by pages
func (u *MovieService) FindMovieByPages(currentPage, pageSize int) ([]model.Movie, error) {
	movies := []model.Movie{}
	return movies, u.getCtl().DAO.FindByPages(&movies, currentPage, pageSize)
}

// find Movie by id
func (u *MovieService) FindMovieById(uid uint64) (*model.Movie, error) {
	keys := map[string]interface{}{"id": uid}
	movie := &model.Movie{}
	return movie, u.getCtl().DAO.FindByKeys(movie, keys)
}

// create Movie
func (u *MovieService) CreateMovie(movie *model.Movie) error {
	return u.getCtl().DAO.Create(movie)
}

// update Movie
func (u *MovieService) UpdateMovie(movie *model.Movie) (int64, error) {
	rowsAffected, err := u.getCtl().DAO.Update(movie, movie.ID)
	return rowsAffected, err
}

// delete Movie
func (u *MovieService) DeleteMovie(uid uint64) (int64, error) {
	return u.getCtl().DAO.Delete(&model.Movie{}, uid)
}
