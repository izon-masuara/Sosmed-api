package main

import (
	"sosmed/db"
	"sosmed/routes"
)

func main() {
	db.Connect()
	r := routes.SetupRouter()
	routes.Router(r)
	r.Run(":3000")
}
