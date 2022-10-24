package db

import (
	"context"
	"errors"
	"sosmed/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	UploadFile(file models.File) (string, error)
	FindAndDeleteFile(filename string) (string, error)
	UploadShortVideo(videoInfo models.ShortVideo) (string, error)
	StreamShortVideo(filename string) uint8
	FindAllShortVideo() ([]models.ShortVideo, error)

	UserRegister(user models.User)
	UserLogin(username string) models.User
}

type Collections struct {
	user       *mongo.Collection
	files      *mongo.Collection
	shortVideo *mongo.Collection
}

func NewRepository(db *mongo.Database) *Collections {
	user := db.Collection("user")
	files := db.Collection("files")
	shortVideo := db.Collection("short_video")
	return &Collections{user: user, files: files, shortVideo: shortVideo}
}

// Files

func (r *Collections) UploadFile(file models.File) (string, error) {
	_, err := r.files.InsertOne(context.Background(), file)
	return file.Filename, err
}

func (r *Collections) FindAndDeleteFile(filename string) {
	r.files.FindOneAndDelete(context.Background(), bson.M{"filename": filename})
}

func (r *Collections) UploadShortVideo(infoVideo models.ShortVideo) (string, error) {
	_, err := r.shortVideo.InsertOne(context.Background(), infoVideo)
	return "Success Upload Short Video", err
}

func (r *Collections) StreamShortVideo(filename string) *models.File {
	var result models.File
	r.files.FindOne(context.Background(), bson.M{"filename": filename}).Decode(&result)
	return &result
}

func (r *Collections) GetAllShortVideo() ([]*models.ShortVideo, error) {
	csr, err := r.shortVideo.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer csr.Close(context.Background())

	var result []*models.ShortVideo
	for csr.Next(context.Background()) {
		var row models.ShortVideo
		err := csr.Decode(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, &row)
	}
	return result, err
}

// User

func (r *Collections) UserRegister(user models.User) (string, error) {
	var found models.User
	r.user.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&found)
	if len(found.Username) != 0 {
		err := errors.New("username already exits")
		return "", err
	}
	_, err := r.user.InsertOne(context.Background(), user)
	return "Success Register", err
}

func (r *Collections) UserLogin(username string) models.User {
	var result models.User
	r.user.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	return result
}
