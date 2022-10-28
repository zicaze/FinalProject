package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), middlewares.UserAuthorization(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.Authentication(), middlewares.UserAuthorization(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.POST("/", middlewares.Authentication(), middlewares.PhotoAuthorization(), controllers.PostPhoto)
		photoRouter.GET("/", middlewares.Authentication(), middlewares.PhotoAuthorization(), controllers.GetPhoto)
		photoRouter.PUT("/:photoId", middlewares.Authentication(), middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.Authentication(), middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.POST("/", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.PostComment)
		commentRouter.GET("/", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.GetComment)
		commentRouter.PUT("/:commentId", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.DeleteComment)
	}
	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.POST("/", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.PostSosmed)
		sosmedRouter.GET("/", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.GetSosmed)
		sosmedRouter.PUT("/:socialMediaId", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.UpdateSosmed)
		sosmedRouter.DELETE("/:socialMediaId", middlewares.Authentication(), middlewares.CommentAuthorization(), controllers.DeleteSosmed)
	}

	return r
}
