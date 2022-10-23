package models

type File struct {
	Filename string  `form:"filename" bson:"filename" binding:"required"`
	Buffer   []uint8 `form:"buffer" bson:"buffer" binding:"required"`
}

type ShortVideo struct {
	Filename   string `form:"filename" bson:"filename" binding:"required"`
	Caption    string `form:"caption" bson:"caption" binding:"required"`
	Category   string `form:"category" bson:"category" binding:"required"`
	UploadedBy string `form:"uploaded_by" bson:"uploaded_by" binding:"required"`
}
