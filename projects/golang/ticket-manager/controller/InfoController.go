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

type InfoController struct {
	apiVersion string
	Service    *service.InfoService
	ErrorCode  error
}

// get controller
func (uc *InfoController) getCtl(c *gin.Context) *InfoController {
	errorCode, _ := c.Get("errorCode")
	if errorCode == http.StatusGatewayTimeout {
		return nil
	}
	var svc *service.InfoService
	return &InfoController{"v1", svc, nil}
}

// create Info
func (uc *InfoController) CreateInfo(c *gin.Context) {
	image := c.PostForm("image")
	title := c.PostForm("title")
	content := c.PostForm("content")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")
	//mtype, _ :=  strconv.Atoi(c.PostForm("age"))
	info := &model.Info{
		Image:     image,
		Title:     title,
		Content:   content,
		Mtype:     mtype,
		Minfo:     minfo,
		Mtime:     mtime,
		BaseModel: model.BaseModel{},
	}

	// return error
	if uc.getCtl(c) == nil {
		return
	}
	err := uc.getCtl(c).Service.CreateInfo(info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": info,
	})
}

// get all infos
func (uc *InfoController) GetAllInfos(c *gin.Context) {
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
	infos, err := uc.getCtl(c).Service.FindInfoByPages(currentpage, pagesize)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": infos,
	})
}

// get info by id
func (uc *InfoController) GetInfoByID(c *gin.Context) {
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

	info, err := uc.getCtl(c).Service.FindInfoById(idUint64)
	if err != nil {
		util.SendError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": info,
	})
}

// update Info
func (uc *InfoController) UpdateInfo(c *gin.Context) {
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
	info, err := uc.getCtl(c).Service.FindInfoById(uidUint64)
	if err != nil {
		panic(" get Info error !")
	}

	image := c.PostForm("image")
	title := c.PostForm("title")
	content := c.PostForm("content")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")

	info.ID = uidUint64
	info.Image = image
	info.Title = title
	info.Content = content
	info.Mtype = mtype
	info.Minfo = minfo
	info.Mtime = mtime

	// return error
	if uc.getCtl(c) == nil {
		return
	}

	rowsAffected, updateErr := uc.getCtl(c).Service.UpdateInfo(info)
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

// delete info
func (uc *InfoController) DeleteInfo(c *gin.Context) {
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

	rowsAffected, delErr := uc.getCtl(c).Service.DeleteInfo(uidUint64)

	if delErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "delete info error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsAffected,
	})
}
