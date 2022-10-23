package middleware

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sosmed/db"
	"sosmed/helper"
	"sosmed/models"

	"github.com/gin-gonic/gin"
)

func UploadPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.ParseMultipartForm(2 << 20)
		file, header, err := ctx.Request.FormFile("image")

		if err != nil {
			helper.ErrorHandler(ctx, "File Validation", "Image required")
			return
		}

		defer file.Close()
		img, err := ioutil.ReadAll(file)
		if err != nil {
			helper.ErrorHandler(ctx, "Internal Server Error", nil)
			return
		}

		filename := fmt.Sprintf("%v%v-%v-%s", rand.Intn(100), rand.Intn(200)+rand.Intn(300), rand.Int(), header.Filename)
		var uploadFile = models.FileUser{
			Filename: filename,
			Buffer:   img,
		}
		success, err := db.Db.UploadFile(uploadFile)
		if err != nil {
			helper.ErrorHandler(ctx, "Internal Server Error", nil)
			return
		}
		// oid := res.InsertedID.(primitive.ObjectID).Hex()
		ctx.Set("filename", success)
	}

	// MFA SECURITY Enable Versioning S3
}
