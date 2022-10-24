package models

type File struct {
	Filename string  `json:"filename" bson:"filename" binding:"required"`
	Buffer   []uint8 `json:"buffer" bson:"buffer" binding:"required"`
}

type ShortVideo struct {
	Filename   string `json:"filename" bson:"filename" binding:"required"`
	Caption    string `json:"caption" bson:"caption" binding:"required"`
	Category   string `json:"category" bson:"category" binding:"required"`
	UploadedBy string `json:"uploaded_by" bson:"uploaded_by" binding:"required"`
}
