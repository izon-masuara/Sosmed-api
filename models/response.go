package models

type ErrorsMsg struct {
	Message interface{}
	Code    int
}

type SuccessResponse struct {
	Message interface{}
	Code    int
}
