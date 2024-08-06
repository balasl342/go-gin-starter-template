package routes

import (
	"github.com/balasl342/go-gin-starter-template/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.POST("/", controllers.CreateUser)
	}
}
