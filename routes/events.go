package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

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
	context.IndentedJSON(http.StatusCreated, db.GetEvents())

}

func getEventByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}

	event, err := db.GetEventByID(int64(id))

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"Message": "Record not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	//var updatedEvent models.Event

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		fmt.Println(err)
		return
	}
	existingEvent, err := db.GetEventByID(int64(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		return
	}
	if err := context.BindJSON(&existingEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		return
	}
	//updatedEvent.ID = int64(id)
	err = db.UpdateEvent(existingEvent)
	if err != nil {
		fmt.Println(err)
		return
	}
	context.IndentedJSON(http.StatusCreated, db.GetEvents())
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		fmt.Println(err)
		return
	}
	err = db.DeleteEvent(int64(id))
	if err != nil {
		fmt.Println(err)
		return
	}
	context.IndentedJSON(http.StatusCreated, db.GetEvents())
}
