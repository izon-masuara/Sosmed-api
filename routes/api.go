package routes

import (
	"sosmed/controllers"

	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	path := router.PathPrefix("/api/v1").Subrouter()
	path.HandleFunc("/user", controllers.HandleUserRegister).Methods("POST", "GET")
}
