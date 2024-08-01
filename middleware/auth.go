package middleware

import (
	"fmt"
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) (userID int64) {
	accessToken := context.Request.Header.Get("Authorization")
	userID, err := utils.ValidateToken(accessToken)
	if err != nil {
		errString := fmt.Sprintf("%v", err)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": errString})
		return
	}
	context.Next()

	//I could also run a context.set() here to set the userid
	return userID
}
