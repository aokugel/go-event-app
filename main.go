package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
