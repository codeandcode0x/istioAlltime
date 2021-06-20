package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"ticket-manager/model"
	"ticket-manager/service"
	"ticket-manager/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ShowController struct {
	apiVersion string
	Service    *service.ShowService
	ErrorCode  error
}

// get controller
func (uc *ShowController) getCtl(c *gin.Context) *ShowController {
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
	errCreate := uc.getCtl(c).Service.CreateShow(show)
	if errCreate != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": errCreate,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": show,
	})
}

// get all shows
func (uc *ShowController) GetAllShows(c *gin.Context) {
	var currentpage, pagesize int
	pagesize = viper.GetInt("PAGE_SIZE")
	// get current page
	currentPageStr, pageExists := c.GetQuery("currentpage")
	if pageExists {
		currentpage, _ = strconv.Atoi(currentPageStr)
	}
	// get page size
	pageSizeStr, pageSizeExists := c.GetQuery("pagesize")
	if pageSizeExists {
		pagesize, _ = strconv.Atoi(pageSizeStr)
	}
	// get show type
	mtype, mtypeExists := c.GetQuery("mtype")
	if !mtypeExists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "show type is null",
		})
		return
	}

	// return error
	if uc.getCtl(c) == nil {
		return
	}
	shows, err := uc.getCtl(c).Service.FindShowByPagesWithKeys(mtype, currentpage, pagesize)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
		return
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
	if uc.getCtl(c) == nil {
		return
	}

	keys := map[string]interface{}{"id": idUint64}
	show := model.Show{}
	err := uc.getCtl(c).Service.DAO.FindByKeys(&show, keys)
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

	idUint64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	// return error
	if uc.getCtl(c) == nil {
		return
	}
	keys := map[string]interface{}{"id": idUint64}
	show := model.Show{}

	err := uc.getCtl(c).Service.DAO.FindByKeys(&show, keys)
	if err != nil {
		panic(" get Show error !")
	}

	name := c.PostForm("name")
	image := c.PostForm("image")
	actors := c.PostForm("actors")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")

	show.ID = idUint64
	show.Name = name
	show.Image = image
	show.Actors = actors
	show.Mtype = mtype
	show.Minfo = minfo
	show.Mtime = mtime

	// return error
	if uc.getCtl(c) == nil {
		return
	}

	rowsAffected, updateErr := uc.getCtl(c).Service.UpdateShow(&show)
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
	idUint64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	// return error
	if uc.getCtl(c) == nil {
		return
	}

	rowsAffected, delErr := uc.getCtl(c).Service.DeleteShow(idUint64)

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
