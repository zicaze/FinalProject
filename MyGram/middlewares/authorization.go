package middlewares

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		userId, err := strconv.Atoi(ctx.Param("userId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}
		userdata := ctx.MustGet("userData").(jwt.MapClaims)
		iD := uint(userdata["id"].(float64))
		User := models.User{}

		err = db.Select("id").First(&User, uint(userId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if User.ID != iD {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorization",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		userdata := ctx.MustGet("userData").(jwt.MapClaims)
		email := userdata["email"]

		User := models.User{}

		err := db.Where("email=?", email).Take(&User).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not Found",
				"message": "You are illegal account ",
			})
			return
		}

		if User.Email != email {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorization",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		userdata := ctx.MustGet("userData").(jwt.MapClaims)
		iD := uint(userdata["id"].(float64))
		User := models.User{}

		err := db.Select("id").First(&User, uint(iD)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not Found",
				"message": "You are illegal account ",
			})
			return
		}

		if User.ID != iD {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorization",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}
