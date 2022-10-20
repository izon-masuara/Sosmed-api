package routes

import (
	"sosmed/controllers"
	"sosmed/middleware"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	url := "/api/v1/user"
	router.POST(url+"/", middleware.UploadPhoto(), controllers.HandleUserRegister)
}
