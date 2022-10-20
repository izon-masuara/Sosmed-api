package main

import (
	"log"
	"net/http"
	"sosmed/db"
	"sosmed/routes"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	r := mux.NewRouter()
	routes.Router(r)
	log.Fatal(http.ListenAndServe(":3000", r))
}
