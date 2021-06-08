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

type InfoController struct {
	apiVersion string
	Service    *service.InfoService
	ErrorCode  error
}

// get controller
func (uc *InfoController) getInfoController(c *gin.Context) *InfoController {
	errorCode, _ := c.Get("errorCode")
	if errorCode == http.StatusGatewayTimeout {
		return nil
	}
	var svc *service.InfoService
	return &InfoController{"v1", svc, nil}
}

// create Info
func (uc *InfoController) CreateInfo(c *gin.Context) {
	name := c.PostForm("name")
	title := c.PostForm("title")
	content := c.PostForm("content")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")
	//mtype, _ :=  strconv.Atoi(c.PostForm("age"))
	info := &model.Info{
		Name:    name,
		Title:   title,
		Content: content,
		Mtype:   mtype,
		Minfo:   minfo,
		Mtime:   mtime,
		Model:   gorm.Model{},
	}

	// return error
	if uc.getInfoController(c) == nil {
		return
	}
	infoId, err := uc.getInfoController(c).Service.CreateInfo(info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"id":   infoId,
		"data": info,
	})
}

// get all infos
func (uc *InfoController) GetAllInfos(c *gin.Context) {
	count := 10
	countStr, exists := c.GetQuery("count")
	if exists {
		count, _ = strconv.Atoi(countStr)
	}
	// return error
	if uc.getInfoController(c) == nil {
		return
	}
	infos, err := uc.getInfoController(c).Service.FindAllInfos(count)

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
	if uc.getInfoController(c) == nil {
		return
	}

	info, err := uc.getInfoController(c).Service.FindInfoById(idUint64)
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

	uid_unit64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	// return error
	if uc.getInfoController(c) == nil {
		return
	}
	info, err := uc.getInfoController(c).Service.FindInfoById(uid_unit64)
	if err != nil {
		panic(" get Info error !")
	}

	name := c.PostForm("name")
	title := c.PostForm("title")
	content := c.PostForm("content")
	mtype := c.PostForm("mtype")
	minfo := c.PostForm("minfo")
	mtime := c.PostForm("mtime")

	info.Name = name
	info.Title = title
	info.Content = content
	info.Mtype = mtype
	info.Minfo = minfo
	info.Mtime = mtime

	// return error
	if uc.getInfoController(c) == nil {
		return
	}

	rowsAffected, updateErr := uc.getInfoController(c).Service.UpdateInfo(uid_unit64, info)
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
	uid_unit64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	// return error
	if uc.getInfoController(c) == nil {
		return
	}

	rowsAffected, delErr := uc.getInfoController(c).Service.DeleteInfo(uid_unit64)

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

// rpc
// get all infos
// func (uc *InfoController) GetAllInfosRPC(c *gin.Context) {
// 	conn, err := grpc.Dial(":20153", grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Printf("faild to connect: %v", err)
// 	}
// 	defer conn.Close()

// 	rpcClient := info.NewInfoRPCClient(conn)
// 	rpcResponse, err := rpcClient.GetAllInfos(context.Background(), &info.InfoMsgRequest{Count: 100})
// 	if err != nil {
// 		log.Printf("could not request: %v", err)
// 	}

// 	infos := []model.Info{}
// 	json.Unmarshal([]byte(rpcResponse.Message), &infos)
// 	log.Printf("info list : %s !\n", rpcResponse.Message)
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"data": infos,
// 	})
// }
