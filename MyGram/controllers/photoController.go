package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PostPhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}

	userdata := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userdata["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.User_id = userId

	err := db.Debug().Create(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request, photo post",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"photo_url":  Photo.Photo_url,
		"caption":    Photo.Caption,
		"user_id":    Photo.User_id,
		"created_at": Photo.CreatedAt,
	})

}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photos, err := Photo.GetPhotos(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, photos)

}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()

	Photo := models.Photo{}

	if err := db.Where("id = ?", c.Param("photoId")).First(&Photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Photo not found!!"})
		return
	}

	var Input models.FormPhoto

	if err := c.ShouldBind(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	var PhotoUpdate models.Photo
	PhotoUpdate.Title = Input.Title
	PhotoUpdate.Caption = Input.Caption
	PhotoUpdate.Photo_url = Input.Photo_url
	PhotoUpdate.UpdatedAt = &time.Time{}

	db.Model(&Photo).Updates(PhotoUpdate)

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"updated_at": Photo.UpdatedAt,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_url,
		"user_id":    Photo.User_id,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()

	var Photo models.Photo

	if err := db.Where("id = ?", c.Param("photoId")).First(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&Photo)

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
