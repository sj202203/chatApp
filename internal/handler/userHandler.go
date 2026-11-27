package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"prob13/internal/core/entity"
)

type UserHandlerImpl struct {
	log   func(c *gin.Context, code int, message string)
	db    *gorm.DB
	chann chan chan string
}

func respondError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"message": message})
}

func NewUser(id int) entity.User {
	return entity.User{UserId: id}
}

func NewUserHandlerImpl(db *gorm.DB) UserHandlerImpl {
	return UserHandlerImpl{log: respondError, db: db}
}

var users = make(map[int]*(*entity.User), 1000)
var userGroups = make(map[int]entity.UserGroup, 1000)

func (h UserHandlerImpl) RegisterUser(c *gin.Context) {

	var user entity.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		h.log(c, http.StatusBadRequest, "User payload is malformed")
		return
	}

	var userExist entity.User
	err = h.db.Where("user_id = ?", user.UserId).First(&userExist).Error

	if err == nil {
		h.log(c, 429, "User with this user_id already there")
		return
	}

	err = h.db.Create(&user).Error
	if err != nil {
		fmt.Println("Error in creating user")
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "User registered", "User": user})
}

func (h UserHandlerImpl) GetUser(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		h.log(c, 500, "Internal server error")
		return
	}

	var user entity.User
	err = h.db.Preload("Messages").Where("user_id = ?", idInt).First(&user).Error

	if err != nil {
		h.log(c, 404, "User not found")
		return
	}

	c.JSON(200, gin.H{"User": user})
}

func (h UserHandlerImpl) RegisterUserGroup(c *gin.Context) {

	var userGroupPayload entity.UserGroupDTO

	err := c.ShouldBindJSON(&userGroupPayload)

	if err != nil {
		h.log(c, http.StatusBadRequest, "User payload is malformed")
		return
	}

	var userExist entity.UserGroup
	err = h.db.Where("id = ?", userGroupPayload.Id).First(&userExist).Error

	if err == nil {
		h.log(c, 429, "User Group with this id already there")
		return
	}

	var userGroup entity.UserGroup
	var allUsersOfUserGroup []entity.User

	for _, v := range userGroupPayload.Users {
		var tempUser entity.User
		err := h.db.Where("user_id = ?", v).First(&tempUser).Error
		if err != nil {
			respondError(c, 404, fmt.Sprintf("User %v not found", v))
			return
		}
		allUsersOfUserGroup = append(allUsersOfUserGroup, tempUser)
	}

	userGroup.Id = userGroupPayload.Id
	userGroup.Users = allUsersOfUserGroup

	if err := h.db.Create(&userGroup).Error; err != nil {
		h.log(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "User Group registered", "UserGroup": userGroup})
}

func (h UserHandlerImpl) GetUserGroup(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		h.log(c, 500, "Internal server error")
		return
	}

	var userGroup entity.UserGroup

	err = h.db.Preload("Users").Where("id = ?", idInt).First(&userGroup).Error

	if err != nil {
		h.log(c, 404, "UserGroup not found")
		return
	}

	c.JSON(200, gin.H{"UserGroup": userGroup})
}