package main

import (
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	fmt.Println("go event app")
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
