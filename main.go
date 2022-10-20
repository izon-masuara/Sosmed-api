package main

import (
	"sosmed/db"
	"sosmed/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()
	routes.Router(r)
	r.Run(":3000")
}
