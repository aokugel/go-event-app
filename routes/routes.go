package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	//events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events", postEvent)
	server.PATCH("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	//registrations
	server.POST("/events/:id/register", registerUserForEvent)
	server.DELETE("/events/:id/register", unregisterUserForEvent)

	//users
	server.GET("/users", getUsers)
	server.POST("/signup", createUser)
	server.POST("/login", userLogin)

	//jwt
	server.POST("/jwt/token")

}
