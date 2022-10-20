package controllers

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"sosmed/db"
	"sosmed/models"

	"go.mongodb.org/mongo-driver/bson"
)

func HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	database := db.DB
	if r.Method == "GET" {
		data, err := database.Collection("user").Find(context.Background(), bson.D{})
		if err != nil {
			panic(err.Error())
		}
		var users []models.UserRegister
		if err = data.All(context.Background(), &users); err != nil {
			log.Fatal(err)
		}

		w.Write(users[0].File.Buffer)
	} else {
		err := r.ParseMultipartForm(2 << 20)
		if err != nil {
			panic(err.Error())
		}

		file, header, err := r.FormFile("image")
		if err != nil {
			panic(err.Error())
		}

		defer file.Close()

		img, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err.Error())
		}
		var user = models.UserRegister{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
			File: &models.FileUser{
				Filename: header.Filename,
				Buffer:   img,
			},
		}
		_, err = database.Collection("user").InsertOne(context.TODO(), user)
		if err != nil {
			panic(err.Error())
		}

		w.Write(user.File.Buffer)
	}
}
