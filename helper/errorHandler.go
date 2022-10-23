package helper

import (
	"fmt"
	"net/http"
	"sosmed/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	errorMessage models.ErrorsMsg
)

func ErrorHandler(ctx *gin.Context, validation string, err interface{}) {
	switch validation {
	case "Field Validation":
		var errMsg []string
		for _, e := range err.(validator.ValidationErrors) {
			errMsg = append(errMsg, fmt.Sprintf("Error on field %s, Condintion %s", e.Field(), e.ActualTag()))
		}
		errorMessage.Code = http.StatusBadRequest
		errorMessage.Message = errMsg
		ctx.JSON(errorMessage.Code, errorMessage)
	case "Unauthorized":
		errorMessage.Code = http.StatusUnauthorized
		errorMessage.Message = err
		ctx.JSON(errorMessage.Code, errorMessage)
	case "File Validation":
		errorMessage.Code = http.StatusBadRequest
		errorMessage.Message = err
		ctx.JSON(errorMessage.Code, errorMessage)
	default:
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
}
