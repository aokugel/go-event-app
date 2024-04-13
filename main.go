package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("go event app")
	r := gin.Default()
	r.GET("/events", getEvents)
	r.POST("/events", postEvent)
	r.Run(":8080")
}

func getEvents(context *gin.Context) {
	// event1 := models.NewEvent(1, "Wedding", "Downtown Abbey", time.Now(), 400)
	// event2 := models.NewEvent(2, "Funeral", "St Peters Cathedral", time.Now(), 25)
	// event3 := models.NewEvent(3, "Deposition", "Downtown Courthouse", time.Now(), 400)
	// events := []models.Event{*event1, *event2, *event3}

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
