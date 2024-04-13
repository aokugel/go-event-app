package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("go event app")
	r := gin.Default()
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEventByID)
	r.POST("/events", postEvent)
	r.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, models.Events)
}

func postEvent(context *gin.Context) {
	var newEvent models.Event
	if err := context.BindJSON(&newEvent); err != nil {
		return
	}
	models.Events = append(models.Events, newEvent)
	context.IndentedJSON(http.StatusCreated, models.Events)

}

func getEventByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	for _, event := range models.Events {
		if event.ID == id {
			context.IndentedJSON(http.StatusOK, event)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found "})
}
