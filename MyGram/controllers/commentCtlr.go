package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PostComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Comment := models.Comment{}

	userdata := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userdata["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.User_id = strconv.Itoa(int(userId))

	err := db.Debug().Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request, comment post",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.Photo_id,
		"user_id":    Comment.User_id,
		"created_at": Comment.CreatedAt,
	})
}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	Comments, err := Comment.GetComments(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message get comment": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comments)

}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()

	Comment := models.Comment{}

	if err := db.Where("id = ?", c.Param("commentId")).First(&Comment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Comment not found!!"})
		return
	}

	var Input models.FormComment

	if err := c.ShouldBind(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	var CommentUpdate models.Comment
	CommentUpdate.Message = Input.Message
	CommentUpdate.UpdatedAt = &time.Time{}

	db.Model(&Comment).Updates(CommentUpdate)

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"updated_at": Comment.UpdatedAt,
		"photo_id":   Comment.Photo_id,
		"user_id":    Comment.User_id,
	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()

	var Comment models.Comment

	if err := db.Where("id = ?", c.Param("commentId")).First(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&Comment)

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
