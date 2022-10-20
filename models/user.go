package models

type FileUser struct {
	Filename string
	Buffer   []uint8
}

type UserRegister struct {
	Username string
	Password string
	File     *FileUser
}
