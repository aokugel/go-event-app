package routes

import (
	"fmt"
	"net/http"
	"os"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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

func userLogin(context *gin.Context) {
	var loginAttempt models.User
	if err := context.BindJSON(&loginAttempt); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse json object"})
		fmt.Println(err)
		return
	}
	existingUser, err := db.GetUserByEmail(loginAttempt.Email)

	if err != nil {
		fmt.Println("No user exists with that email")
		fmt.Println(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "No user with this email"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginAttempt.Password))
	if err != nil {
		fmt.Println("Password incorrectamundo")
		fmt.Println(err)
		context.JSON(http.StatusForbidden, gin.H{"message": "Incorrect password"})
		return
	}

	//this will eventually be outsourced to a function
	key := []byte(os.Getenv("TOKEN_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":  "anthonykugel.com",
			"sub":  existingUser.Email,
			"name": existingUser.FirstName + " " + existingUser.LastName,
			"foo":  2,
		})
	signedString, err := token.SignedString(key)

	if err != nil {
		fmt.Println("Error generating token")
		fmt.Println(err)
		context.JSON(http.StatusForbidden, gin.H{"message": "Error generating token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": signedString})

}
