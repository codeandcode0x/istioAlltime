package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"ticket-manager/model"
	"ticket-manager/service"
	"ticket-manager/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	movie "ticket-manager/rpc/grpc/protos/movie"
)

type MovieController struct {
	apiVersion string
	Service    *service.MovieService
	ErrorCode  error
}

// get controller
func (uc *MovieController) getCtl(c *gin.Context) *MovieController {
	errorCode, _ := c.Get("errorCode")
	if errorCode == http.StatusGatewayTimeout {
		return nil
	}
	var svc *service.MovieService
	return &MovieController{"v1", svc, nil}
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
		Name:      name,
		Image:     image,
		Actors:    actors,
		Mtype:     mtype,
		Minfo:     minfo,
		Mtime:     mtime,
		BaseModel: model.BaseModel{},
	}

	// return error
	if uc.getCtl(c) == nil {
		return
	}
	err := uc.getCtl(c).Service.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": movie,
	})
}

// get all movies
func (uc *MovieController) GetAllMovies(c *gin.Context) {
	var currentpage, pagesize int
	pagesize = viper.GetInt("PAGE_SIZE")
	currentPageStr, pageExists := c.GetQuery("currentpage")
	if pageExists {
		currentpage, _ = strconv.Atoi(currentPageStr)
	}
	// get page size
	pageSizeStr, pageSizeExists := c.GetQuery("pagesize")
	if pageSizeExists {
		pagesize, _ = strconv.Atoi(pageSizeStr)
	}

	// return error
	if uc.getCtl(c) == nil {
		return
	}
	movies, err := uc.getCtl(c).Service.FindMovieByPages(currentpage, pagesize)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": movies,
	})
}

// get movie by id
func (uc *MovieController) GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "id is null",
		})
		return
	}

	idUint64, _ := strconv.ParseUint(id, 10, 64)

	// return error
	if uc.getCtl(c) == nil {
		return
	}

	movie, err := uc.getCtl(c).Service.FindMovieById(idUint64)
	if err != nil {
		util.SendError(c, err.Error())
		return
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
			"code":  -1,
			"error": "id is null",
		})
	}

	uidUint64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	// return error
	if uc.getCtl(c) == nil {
		return
	}
	movie, err := uc.getCtl(c).Service.FindMovieById(uidUint64)
	if err != nil {
		panic(" get Movie error !")
	}

	name := c.PostForm("name")
	image := c.PostForm("image")
	actors := c.PostForm("actors")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")

	movie.ID = uidUint64
	movie.Name = name
	movie.Image = image
	movie.Actors = actors
	movie.Mtype = mtype
	movie.Minfo = minfo
	movie.Mtime = mtime

	// return error
	if uc.getCtl(c) == nil {
		return
	}

	rowsAffected, updateErr := uc.getCtl(c).Service.UpdateMovie(movie)
	if updateErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": updateErr,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsAffected,
	})
}

// delete movie
func (uc *MovieController) DeleteMovie(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "id is null",
		})
	}
	fmt.Println("uid", uid)
	uidUint64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	// return error
	if uc.getCtl(c) == nil {
		return
	}

	rowsAffected, delErr := uc.getCtl(c).Service.DeleteMovie(uidUint64)

	if delErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "delete movie error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsAffected,
	})
}

// rpc
// get all movies
func (uc *MovieController) GetAllMoviesRPC(c *gin.Context) {
	conn, err := grpc.Dial(":20153", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	rpcClient := movie.NewMovieRPCClient(conn)
	rpcResponse, err := rpcClient.GetAllMovies(context.Background(), &movie.MovieMsgRequest{Count: 100})
	if err != nil {
		log.Printf("could not request: %v", err)
	}

	movies := []model.Movie{}
	json.Unmarshal([]byte(rpcResponse.Message), &movies)
	log.Printf("movie list : %s !\n", rpcResponse.Message)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": movies,
	})
}
