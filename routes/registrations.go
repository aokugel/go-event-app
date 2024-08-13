package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"github.com/gin-gonic/gin"
)

func registerUserForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")

	fmt.Println("registerforeventid", userID)

	eventID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		fmt.Println(err)
		return
	}
	_, err = db.GetEventByID(int64(eventID))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "An event does not exist by that id"})
		return
	}

	db.CreateUserRegistration(userID, int64(eventID))
	context.JSON(http.StatusOK, gin.H{"message": "registration successful"})

}

func unregisterUserForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")

	eventID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		fmt.Println(err)
		return
	}
	_, err = db.GetEventByID(int64(eventID))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "An event does not exist by that id"})
		return
	}
	fmt.Printf("userID: %v, eventID: %v\n", userID, eventID)

	db.DeleteUserRegistration(userID, int64(eventID))

	context.JSON(http.StatusOK, gin.H{"message": "unregistration successful"})
}

func getEventsByRegisteredUser(context *gin.Context) {
	userID := context.GetInt64("userID")

	events := db.GetEventsByRegisteredUser(userID)

	fmt.Println(userID, events)

	context.JSON(http.StatusOK, events)

}
