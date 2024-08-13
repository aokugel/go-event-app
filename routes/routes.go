package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/login", userLogin)
	server.GET("/users", getUsers)
	server.POST("/signup", createUser)

	authorized := server.Group("/")
	authorized.Use(middleware.Authenticate)

	authorized.POST("/events", postEvent)
	authorized.PATCH("/events/:id", updateEvent)
	authorized.DELETE("/events/:id", deleteEvent)
	authorized.POST("/events/:id/register", registerUserForEvent)
	authorized.DELETE("/events/:id/register", unregisterUserForEvent)
	authorized.GET("/events/registered", getEventsByRegisteredUser)

	//jwt
	server.POST("/jwt/token")

}
