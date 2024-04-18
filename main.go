package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	fmt.Println("go event app")
	r := gin.Default()
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEventByID)
	r.POST("/events", postEvent)
	r.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := db.GetEvents()
	context.JSON(http.StatusOK, events)
}

func postEvent(context *gin.Context) {
	var newEvent models.Event
	if err := context.BindJSON(&newEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		return
	}
	db.InsertEventIntoDB(&newEvent)
	fmt.Println(newEvent)
	context.IndentedJSON(http.StatusCreated, db.GetEvents())

}

func getEventByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	for _, event := range db.GetEvents() {
		if event.ID == int64(id) {
			context.IndentedJSON(http.StatusOK, event)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found "})
}
