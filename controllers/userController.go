package controllers

import (
	databases "SchoolTest/database"
	"SchoolTest/helpers"
	"SchoolTest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func HasData(u *models.User) bool {
	if u.ID != 0 || u.Username != "" {
		return true
	}
	return false
}

func UserRegister(c *gin.Context) {
	db := databases.GetDB()
	contentType := c.ContentType()

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"username": User.Username,
		"role": User.Role,
	})
}

func UserLogin(c *gin.Context) {
	db := databases.GetDB()
	contentType := c.ContentType()

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Where("username = ?", User.Username).First(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Username is not found!",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Password incorrect!",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Username, User.Role)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
