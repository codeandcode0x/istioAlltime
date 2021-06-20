package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"ticket-manager/model"
	"ticket-manager/service"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// controller struct
type UserController struct {
	apiVersion string
	Service    *service.UserService
}

// get controller
func (uc *UserController) getCtl() *UserController {
	var svc *service.UserService
	return &UserController{"v1", svc}
}

// get all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.getCtl().Service.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": users,
	})
}

func (uc *UserController) GetUserByID() {
}

// create user
func (uc *UserController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	password, exists := c.GetPostForm("password")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "password is null",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "bcrypt password err",
		})
	}

	password = string(hash)
	email := c.PostForm("email")
	role := c.PostForm("role")
	age, _ := strconv.Atoi(c.PostForm("age"))
	user := &model.User{
		Name:         name,
		Password:     password,
		Email:        email,
		Role:         role,
		Age:          age,
		Birthday:     time.Now(),
		MemberNumber: sql.NullString{},
		BaseModel:    model.BaseModel{},
	}
	errCreate := uc.getCtl().Service.CreateUser(user)
	if errCreate != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": user,
	})
}

// update user
func (uc *UserController) UpdateUser(c *gin.Context) {
	uid, exists := c.GetPostForm("id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "id is null",
		})
	}

	uidUnit64, errConv := strconv.ParseUint(uid, 10, 64)
	if errConv != nil {
		panic(" get uid error !")
	}

	user, err := uc.getCtl().Service.FindUserById(uidUnit64)
	if err != nil {
		panic(" get user error !")
	}

	name := c.PostForm("name")
	passwd := c.PostForm("passwd")
	email := c.PostForm("email")
	age, _ := strconv.Atoi(c.PostForm("age"))

	user.ID = uidUnit64
	user.Name = name
	user.Password = passwd
	user.Email = email
	user.Age = age

	rowsAffected, updateErr := uc.getCtl().Service.UpdateUser(user)
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

func (uc *UserController) DeleteUser(c *gin.Context) {
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

	rowsAffected, delErr := uc.getCtl().Service.DeleteUser(uid_unit64)

	if delErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "delete user error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsAffected,
	})
}
