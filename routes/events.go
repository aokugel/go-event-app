package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events := db.GetEvents()
	context.JSON(http.StatusOK, events)
}

func postEvent(context *gin.Context) {
	accessToken := context.Request.Header.Get("Authorization")
	userID, err := utils.ValidateToken(accessToken)
	if err != nil {
		errString := fmt.Sprintf("%v", err)
		context.JSON(http.StatusBadRequest, gin.H{"Message": errString})
		return
	}

	var newEvent models.Event
	if err := context.BindJSON(&newEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		return
	}

	// Clean this code up it is fucking disgusting.
	newEvent.UserID = userID

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
	accessToken := context.Request.Header.Get("Authorization")
	userID, err := utils.ValidateToken(accessToken)
	if err != nil {
		errString := fmt.Sprintf("%v", err)
		context.JSON(http.StatusBadRequest, gin.H{"Message": errString})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id"})
		fmt.Println(err)
		return
	}
	existingEvent, err := db.GetEventByID(int64(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "A record does not exist by that id"})
		return
	}
	if err := context.BindJSON(&existingEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		return
	}
	if existingEvent.UserID != userID {
		context.JSON(http.StatusForbidden, gin.H{"message": "users can only modify events that they have created"})
		return
	}
	err = db.UpdateEvent(existingEvent)
	if err != nil {
		fmt.Println(err)
		return
	}
	context.IndentedJSON(http.StatusCreated, db.GetEvents())
}

func deleteEvent(context *gin.Context) {
	accessToken := context.Request.Header.Get("Authorization")
	userID, err := utils.ValidateToken(accessToken)
	if err != nil {
		errString := fmt.Sprintf("%v", err)
		context.JSON(http.StatusBadRequest, gin.H{"Message": errString})
		return
	}
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		fmt.Println(err)
		return
	}
	existingEvent, err := db.GetEventByID(int64(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "A record does not exist by that id"})
		return
	}
	if existingEvent.UserID != userID {
		context.JSON(http.StatusForbidden, gin.H{"message": "user needs to have created event to delete said event"})
		return
	}
	err = db.DeleteEvent(int64(id))
	if err != nil {
		fmt.Println(err)
		return
	}
	context.IndentedJSON(http.StatusCreated, db.GetEvents())
}
