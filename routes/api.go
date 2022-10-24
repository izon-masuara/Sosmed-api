package routes

import (
	"sosmed/controllers"
	"sosmed/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func Router(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	v1.POST("/", controllers.HandleUserRegister)
	v1.POST("/login", controllers.HandleUserLogin)
	v1.POST("/upload/:username", middleware.UploadFile(), controllers.UploadShortVideo)
	v1.GET("/", controllers.GetAllShortVideo)
	v1.GET("/:video-file-name", controllers.StreamShortVideo)
}
