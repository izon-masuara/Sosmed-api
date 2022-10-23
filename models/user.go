package models

type FileUser struct {
	Filename string  `json:"filename" bson:"filename" binding:"required"`
	Buffer   []uint8 `json:"buffer" bson:"buffer" binding:"required"`
}

type UserRegister struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Filename string `json:"filename" bson:"filename" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
