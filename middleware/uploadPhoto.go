package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UploadPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("image")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(header.Filename, file)
	}
}
