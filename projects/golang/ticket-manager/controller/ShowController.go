package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"ticket-manager/model"
	"ticket-manager/service"
	"ticket-manager/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShowController struct {
	apiVersion string
	Service    *service.ShowService
	ErrorCode  error
}

// get controller
func (uc *ShowController) getShowController(c *gin.Context) *ShowController {
	errorCode, _ := c.Get("errorCode")
	if errorCode == http.StatusGatewayTimeout {
		return nil
	}
	var svc *service.ShowService
	return &ShowController{"v1", svc, nil}
}

// create Show
func (uc *ShowController) CreateShow(c *gin.Context) {
	name := c.PostForm("name")
	image := c.PostForm("image")
	actors := c.PostForm("actors")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")
	//mtype, _ :=  strconv.Atoi(c.PostForm("age"))
	show := &model.Show{
		Name:   name,
		Image:  image,
		Actors: actors,
		Mtype:  mtype,
		Minfo:  minfo,
		Mtime:  mtime,
		Model:  gorm.Model{},
	}

	// return error
	if uc.getShowController(c) == nil {
		return
	}
	showId, err := uc.getShowController(c).Service.CreateShow(show)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"id":   showId,
		"data": show,
	})
}

// get all shows
func (uc *ShowController) GetAllShows(c *gin.Context) {
	count := 10
	countStr, countExists := c.GetQuery("count")
	mtype, mtypeExists := c.GetQuery("mtype")
	if countExists && mtypeExists {
		count, _ = strconv.Atoi(countStr)
	}
	// return error
	if uc.getShowController(c) == nil {
		return
	}
	shows, err := uc.getShowController(c).Service.FindAllShows(mtype, count)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": shows,
	})
}

// get show by id
func (uc *ShowController) GetShowByID(c *gin.Context) {
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
	if uc.getShowController(c) == nil {
		return
	}

	show, err := uc.getShowController(c).Service.FindShowById(idUint64)
	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": show,
	})
}

// update Show
func (uc *ShowController) UpdateShow(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "id is null",
		})
	}

	uid_unit64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	// return error
	if uc.getShowController(c) == nil {
		return
	}
	show, err := uc.getShowController(c).Service.FindShowById(uid_unit64)
	if err != nil {
		panic(" get Show error !")
	}

	name := c.PostForm("name")
	image := c.PostForm("image")
	actors := c.PostForm("actors")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")

	show.Name = name
	show.Image = image
	show.Actors = actors
	show.Mtype = mtype
	show.Minfo = minfo
	show.Mtime = mtime

	// return error
	if uc.getShowController(c) == nil {
		return
	}

	rowsAffected, updateErr := uc.getShowController(c).Service.UpdateShow(uid_unit64, show)
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

// delete show
func (uc *ShowController) DeleteShow(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "id is null",
		})
	}
	fmt.Println("uid", uid)
	uid_unit64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	// return error
	if uc.getShowController(c) == nil {
		return
	}

	rowsAffected, delErr := uc.getShowController(c).Service.DeleteShow(uid_unit64)

	if delErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "delete show error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsAffected,
	})
}

// rpc
// get all shows
// func (uc *ShowController) GetAllShowsRPC(c *gin.Context) {
// 	conn, err := grpc.Dial(":20153", grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Printf("faild to connect: %v", err)
// 	}
// 	defer conn.Close()

// 	rpcClient := show.NewShowRPCClient(conn)
// 	rpcResponse, err := rpcClient.GetAllShows(context.Background(), &show.ShowMsgRequest{Count: 100})
// 	if err != nil {
// 		log.Printf("could not request: %v", err)
// 	}

// 	shows := []model.Show{}
// 	json.Unmarshal([]byte(rpcResponse.Message), &shows)
// 	log.Printf("show list : %s !\n", rpcResponse.Message)
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"data": shows,
// 	})
// }
