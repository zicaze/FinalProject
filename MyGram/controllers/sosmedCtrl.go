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

func PostSosmed(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Sosmed := models.Sosmed{}

	userdata := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userdata["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.User_id = userId

	err := db.Debug().Create(&Sosmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request, Sosmed post",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               Sosmed.ID,
		"name":             Sosmed.Name,
		"social_media_url": Sosmed.Social_media_url,
		"user_id":          Sosmed.User_id,
		"created_at":       Sosmed.CreatedAt,
	})
}

func GetSosmed(c *gin.Context) {
	db := database.GetDB()
	Sosmed := models.Sosmed{}

	Sosmeds, err := Sosmed.GetSosmeds(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message get comment": err.Error(),
		})
		return
	}

	response := make(map[string]interface{})
	response["social_medias"] = Sosmeds

	c.JSON(http.StatusOK, response)

}

func UpdateSosmed(c *gin.Context) {
	db := database.GetDB()

	Sosmed := models.Sosmed{}

	if err := db.Where("id = ?", c.Param("socialMediaId")).First(&Sosmed).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Social Media not found!!"})
		return
	}

	var Input models.FormSosmed

	if err := c.ShouldBind(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	var SosmedUpdate models.Sosmed
	SosmedUpdate.Name = Input.Name
	SosmedUpdate.Social_media_url = Input.Social_media_url
	SosmedUpdate.UpdatedAt = &time.Time{}

	db.Model(&Sosmed).Updates(SosmedUpdate)

	c.JSON(http.StatusOK, gin.H{
		"id":               Sosmed.ID,
		"name":             Sosmed.Name,
		"social_media_url": Sosmed.Social_media_url,
		"user_id":          Sosmed.User_id,
		"updated_at":       Sosmed.UpdatedAt,
	})
}

func DeleteSosmed(c *gin.Context) {
	db := database.GetDB()

	var Sosmed models.Sosmed

	if err := db.Where("id = ?", c.Param("socialMediaId")).First(&Sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&Sosmed)

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Social Meida has been successfully deleted",
	})
}
