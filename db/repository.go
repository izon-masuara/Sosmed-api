package db

import (
	"context"
	"sosmed/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	UploadPhoto(file models.FileUser) (models.FileUser, error)
	FindAndDeletePhoto(filename string) (string, error)
	UserRegister(user models.UserRegister)
	UserLogin(username string) models.UserLogin
}

type Collections struct {
	user  *mongo.Collection
	files *mongo.Collection
}

func NewRepository(db *mongo.Database) *Collections {
	user := db.Collection("user")
	files := db.Collection("files")
	return &Collections{user: user, files: files}
}

// Files

func (r *Collections) UploadFile(file models.FileUser) (string, error) {
	_, err := r.files.InsertOne(context.Background(), file)
	return file.Filename, err
}

func (r *Collections) FindAndDeletePhoto(filename string) {
	r.files.FindOneAndDelete(context.Background(), bson.M{"filename": filename})
}

// User

func (r *Collections) UserRegister(user models.UserRegister) (string, error) {
	_, err := r.user.InsertOne(context.Background(), user)
	return "Success Register", err
}

func (r *Collections) UserLogin(username string) models.UserLogin {
	var result models.UserLogin
	r.user.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	return result
}
