package controllers

import (
	"io"
	"io/ioutil"
	"sosmed/models"

	"github.com/gin-gonic/gin"
)

func HandleUserRegister(ctx *gin.Context) {
	// database := db.DB
	ctx.Request.ParseMultipartForm(2 << 20)
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	img, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	var user = models.UserRegister{
		Username: ctx.Request.FormValue("username"),
		Password: ctx.Request.FormValue("password"),
		File: &models.FileUser{
			Filename: header.Filename,
			Buffer:   img,
		},
	}
	// _, err = database.Collection("user").InsertOne(context.TODO(), user)
	// if err != nil {
	// 	panic(err.Error())
	// }

	ctx.Stream(func(w io.Writer) bool {
		w.Write(user.File.Buffer)
		return false
	})
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
