package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"ticket-manager/model"
	"ticket-manager/service"
)

type MovieController struct {
	apiVersion  string
	Service    *service.MovieService
}

// get controller
func (uc *MovieController) getMovieController() *MovieController {
	var svc *service.MovieService
	return &MovieController{"v1", svc}
}

// create Movie
func (uc *MovieController) CreateMovie(c *gin.Context) {
	name := c.PostForm("name")
	image := c.PostForm("image")
	actors := c.PostForm("actors")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")
	//mtype, _ :=  strconv.Atoi(c.PostForm("age"))
	movie := &model.Movie{
		Name:       name,
		Image:      image,
		Actors:     actors,
		Mtype:      mtype,
		Minfo:      minfo,
		Mtime:      mtime,
		Model:      gorm.Model{},
	}
	movieId, err := uc.getMovieController().Service.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"id": movieId,
		"data": movie,
	})
}

func (uc *MovieController) GetAllMovies(c *gin.Context) {
	// var movies []model.Movie{}
	// var err error
	count := 10
	countStr, exists := c.GetQuery("count")
	if exists {
		count, _ = strconv.Atoi(countStr)
	}
	movies, err := uc.getMovieController().Service.FindAllMovies(count)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": movies,
	})
}


func (uc *MovieController) GetMovieByID(c *gin.Context) {
	id, exists := c.GetQuery("count")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": "id is null",
		})
		return
	}

	idUint64, _ := strconv.ParseUint(id, 10, 64)
	movie, err := uc.getMovieController().Service.FindMovieById(idUint64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": movie,
	})
}

// update Movie
func (uc *MovieController) UpdateMovie(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": "id is null",
		})
	}

	uid_unit64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	movie, err := uc.getMovieController().Service.FindMovieById(uid_unit64)
	if err != nil {
		panic(" get Movie error !")
	}

	name := c.PostForm("name")
	image := c.PostForm("image")
	actors := c.PostForm("actors")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")

	movie.Name = name
	movie.Image = image
	movie.Actors = actors
	movie.Mtype = mtype
	movie.Minfo = minfo
	movie.Mtime = mtime

	rowsAffected, updateErr := uc.getMovieController().Service.UpdateMovie(uid_unit64, movie)
	if updateErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": updateErr,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsAffected,
	})
}

func (uc *MovieController) DeleteMovie(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": "id is null",
		})
	}
	fmt.Println("uid", uid)
	uid_unit64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	rowsAffected, delErr := uc.getMovieController().Service.DeleteMovie(uid_unit64)

	if delErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": "delete movie error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsAffected,
	})
}