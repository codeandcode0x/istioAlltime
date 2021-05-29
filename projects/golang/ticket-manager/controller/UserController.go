package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"ticket-manager/model"
	"ticket-manager/service"
	"time"
)

type UserController struct {
	apiVersion  string
	Service    *service.UserService
}

// get controller
func (uc *UserController) getUserController() *UserController {
	var svc *service.UserService
	return &UserController{"v1", svc}
}

// create user
func (uc *UserController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	password, exists := c.GetPostForm("password")
	if ! exists {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": "password is null",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": "bcrypt password err",
		})
	}

	password = string(hash)
	email := c.PostForm("email")
	role := c.PostForm("role")
	age, _ :=  strconv.Atoi(c.PostForm("age"))
	user := &model.User{
		Name:         name,
		Password:     password,
		Email:        email,
		Role:         role,
		Age:          age,
		Birthday:     time.Now(),
		MemberNumber: sql.NullString{},
		Model:        gorm.Model{},
	}
	userId, err := uc.getUserController().Service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"id": userId,
		"data": user,
	})
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.getUserController().Service.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
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

// update user
func (uc *UserController) UpdateUser(c *gin.Context) {
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

	user, err := uc.getUserController().Service.FindUserById(uid_unit64)
	if err != nil {
		panic(" get user error !")
	}

	name := c.PostForm("name")
	passwd := c.PostForm("passwd")
	email := c.PostForm("email")
	age, _ :=  strconv.Atoi(c.PostForm("age"))

	user.Name = name
	user.Password = passwd
	user.Email = email
	user.Age = age

	rowsAffected, updateErr := uc.getUserController().Service.UpdateUser(uid_unit64, user)
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

func (uc *UserController) DeleteUser(c *gin.Context) {
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

	rowsAffected, delErr := uc.getUserController().Service.DeleteUser(uid_unit64)

	if delErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"error": "delete user error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsAffected,
	})
}