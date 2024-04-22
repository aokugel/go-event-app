package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getUsers(context *gin.Context) {
	users := db.GetUsers()
	context.JSON(http.StatusOK, users)
}

func createUser(context *gin.Context) {
	var newUser models.User
	if err := context.BindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		fmt.Println(err)
		return
	}
	var err error
	newUser.Password, err = utils.HashPassword(newUser.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.InsertUserIntoDB(&newUser)
	context.IndentedJSON(http.StatusCreated, gin.H{"message": "user account has been created successfully."})

}
