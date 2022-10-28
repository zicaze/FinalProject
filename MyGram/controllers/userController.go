package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
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
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email=?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()

	User := models.User{}

	if err := db.Where("id = ?", c.Param("userId")).First(&User).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	var Input models.UpdateUserInput

	if err := c.ShouldBind(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	var UserUpdate models.User
	UserUpdate.Username = Input.Username
	UserUpdate.Email = Input.Email
	UserUpdate.UpdatedAt = &time.Time{}

	db.Model(&User).Updates(UserUpdate)

	c.JSON(http.StatusOK, gin.H{
		"id":         User.ID,
		"updated_at": User.UpdatedAt,
		"username":   User.Username,
		"email":      User.Email,
		"age":        User.Age,
	})
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()

	var User models.User

	if err := db.Where("id = ?", c.Param("userId")).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&User)

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
