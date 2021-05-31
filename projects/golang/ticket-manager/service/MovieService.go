package service

import (
	"fmt"
	"ticket-manager/model"
)

type MovieService struct {
	Model  *model.Movie
}

// get Movie service
func (u *MovieService) getMovieService() *MovieService {
	var entity *model.Movie
	return &MovieService{entity}
}

// find all Movies
func (u *MovieService) FindAllMovies(count int) ([]model.Movie, error) {
	return u.getMovieService().Model.FindAll(count)
}

// find Movie by id
func (u *MovieService) FindMovieById(uid uint64) (*model.Movie, error) {
	return u.getMovieService().Model.FindByID(uid)
}

// create Movie
func (u *MovieService) CreateMovie(Movie *model.Movie) (uint64, error) {
	return u.getMovieService().Model.Create(Movie)
}

// update Movie
func (u *MovieService) UpdateMovie(uid uint64, Movie *model.Movie ) (int64, error) {
	rowsAffected, err := u.getMovieService().Model.Update(uid, Movie)
	return rowsAffected, err
}

// delete Movie
func (u *MovieService) DeleteMovie(uid uint64) (int64, error) {
	fmt.Println("uid", uid)
	return u.getMovieService().Model.Delete(uid)
}

