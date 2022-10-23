package routes

import (
	"sosmed/controllers"
	"sosmed/middleware"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	v1 := router.Group("/api/v1/user")
	v1.POST("/", middleware.UploadPhoto(), controllers.HandleUserRegister)
	v1.POST("/login", controllers.HandleUserLogin)
}
