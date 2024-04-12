package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	ID       int
	Name     string
	Location string
	Date     time.Time
	Invitees int
}

func NewEvent(id int, name string, location string, date time.Time, invitees int) *Event {
	return &Event{
		ID:       id,
		Name:     name,
		Location: location,
		Date:     date,
		Invitees: invitees,
	}

}

func main() {

	fmt.Println("go event app")
	r := gin.Default()
	r.GET("/events", getEvents)
	r.Run(":8080")
}

func getEvents(context *gin.Context) {
	event1 := NewEvent(1, "Wedding", "Downtown Abbey", time.Now(), 400)
	event2 := NewEvent(2, "Funeral", "St Peters Cathedral", time.Now(), 25)
	event3 := NewEvent(3, "Deposition", "Downtown Courthouse", time.Now(), 400)
	events := []Event{*event1, *event2, *event3}

	context.JSON(http.StatusOK, events)
}

//gin.H{}
