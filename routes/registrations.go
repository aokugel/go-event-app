package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func registerUserForEvent(context *gin.Context) {
	userID, err := middleware.Authenticate(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "authentication failed"})
		fmt.Println(err)
		return
	}
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
	userID, err := middleware.Authenticate(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "authentication failed"})
		fmt.Println(err)
		return
	}
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
	userID, err := middleware.Authenticate(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "authentication failed"})
		fmt.Println(err)
		return
	}

	events := db.GetEventsByRegisteredUser(int64(userID))

	context.JSON(http.StatusOK, events)

}
