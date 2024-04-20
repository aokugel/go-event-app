package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getUsers(context *gin.Context) {
	users := db.GetUsers()
	context.JSON(http.StatusOK, users)
}

func postUser(context *gin.Context) {
	var newUser models.User
	if err := context.BindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		fmt.Println(err)
		return
	}
	db.InsertUserIntoDB(&newUser)
	context.IndentedJSON(http.StatusCreated, db.GetEvents())

}
