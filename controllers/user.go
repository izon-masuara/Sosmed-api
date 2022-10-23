package controllers

import (
	"net/http"
	"sosmed/db"
	"sosmed/helper"
	"sosmed/models"

	"github.com/gin-gonic/gin"
)

var responseSuccess models.SuccessResponse

func HandleUserRegister(ctx *gin.Context) {
	var user models.UserRegister
	user.Filename = ctx.MustGet("filename").(string)
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
	db.Db.UserRegister(user)
	if err != nil {
		db.Db.FindAndDeletePhoto(user.Filename)
		helper.ErrorHandler(ctx, "Internal Server Error", nil)
		return
	}
	responseSuccess.Code = http.StatusOK
	responseSuccess.Message = "Success Register"
	ctx.JSON(responseSuccess.Code, responseSuccess)
}

func HandleUserLogin(ctx *gin.Context) {
	var user models.UserLogin

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

// func HandleUserRegister(ctx *gin.Context) {
// 	database := db.DB
// 	data, err := database.Collection("user").Find(context.Background(), bson.D{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	var users []models.UserRegister
// 	if err = data.All(context.Background(), &users); err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx.String(http.StatusOK, string(users[0].File.Buffer))
// 	// ctx.Data(http.StatusOK, "application/octet-stream", users[0].File.Buffer)
// }
