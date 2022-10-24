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

func UploadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("File")

		if err != nil {
			helper.ErrorHandler(ctx, "File Validation", "File required")
			return
		}

		contentType := header.Header.Get("Content-Type")

		if contentType != "video/mp4" {
			helper.ErrorHandler(ctx, "File Validation", fmt.Sprintf("Content type must be %s", contentType))
			return
		}

		if header.Size >= 15000000 {
			helper.ErrorHandler(ctx, "File Validation", "Size file must be less than 15 Mb")
			return
		}

		defer file.Close()
		buf, err := ioutil.ReadAll(file)
		if err != nil {
			helper.ErrorHandler(ctx, "Internal Server Error", nil)
			return
		}
		filename := fmt.Sprintf("%v%v-%v-%s", rand.Intn(100), rand.Intn(200)+rand.Intn(300), rand.Int(), header.Filename)
		var uploadFile = models.File{
			Filename: filename,
			Buffer:   buf,
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
