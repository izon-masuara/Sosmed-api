package controllers

import (
	"io"
	"net/http"
	"sosmed/db"
	"sosmed/helper"
	"sosmed/models"

	"github.com/gin-gonic/gin"
)

var responseSuccess models.SuccessResponse

func HandleUserRegister(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		helper.ErrorHandler(ctx, "Field Validation", err)
		return
	}

	password, err := helper.HashPassword(user.Password)

	if err != nil {
		helper.ErrorHandler(ctx, "Internal Server Error", nil)
		return
	}

	// Change password to password that has been hasing
	user.Password = password
	_, err = db.Db.UserRegister(user)
	if err != nil {
		helper.ErrorHandler(ctx, "Unauthorized", err.Error())
		return
	}
	responseSuccess.Code = http.StatusOK
	responseSuccess.Message = "Success Register"
	ctx.JSON(responseSuccess.Code, responseSuccess)
}

func HandleUserLogin(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBind(&user)
	if err != nil {
		helper.ErrorHandler(ctx, "Field Validation", err)
		return
	}

	result := db.Db.UserLogin(user.Username)

	correctPass := helper.CheckPasswordHash(user.Password, result.Password)
	if !correctPass {
		helper.ErrorHandler(ctx, "Unauthorized", "Invalid Username or Password")
		return
	}
	responseSuccess.Code = http.StatusOK
	responseSuccess.Message = "Token"
	ctx.JSON(responseSuccess.Code, responseSuccess.Message)
}

func UploadShortVideo(ctx *gin.Context) {
	var payload models.ShortVideo
	payload.Filename = ctx.MustGet("filename").(string)
	payload.UploadedBy = ctx.Param("username")
	err := ctx.ShouldBind(&payload)
	if err != nil {
		db.Db.FindAndDeleteFile(payload.Filename)
		helper.ErrorHandler(ctx, "Field Validation", err)
		return
	}

	result := db.Db.UserLogin(payload.UploadedBy)
	if len(result.Username) == 0 {
		db.Db.FindAndDeleteFile(payload.Filename)
		helper.ErrorHandler(ctx, "Unauthorized", "Invalid User To Upload Video")
		return
	}

	resp, err := db.Db.UploadShortVideo(payload)
	if err != nil {
		db.Db.FindAndDeleteFile(payload.Filename)
		helper.ErrorHandler(ctx, "Internal Server Error", nil)
		return
	}

	responseSuccess.Code = http.StatusOK
	responseSuccess.Message = resp
	ctx.JSON(responseSuccess.Code, responseSuccess)
}

func StreamShortVideo(ctx *gin.Context) {
	videoFileName := ctx.Param("video-file-name")
	result := db.Db.StreamShortVideo(videoFileName)
	ctx.Stream(func(w io.Writer) bool {
		w.Write(result.Buffer)
		return false
	})
}

func GetAllShortVideo(ctx *gin.Context) {
	result, err := db.Db.GetAllShortVideo()
	if err != nil {
		helper.ErrorHandler(ctx, "Internal Server Error", nil)
	}
	responseSuccess.Code = http.StatusOK
	responseSuccess.Message = result
	ctx.JSON(responseSuccess.Code, responseSuccess.Message)
}
