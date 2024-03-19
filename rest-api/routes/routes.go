package routes

import (
	"example.com/eventbooking-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// EVENTS
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	//EVENT REGISTRATION with USER
	authenticated.POST("/events/:id/register", registerEvent)
	// authenticated.DELETE("/events/:id/cancel", cancelEvent)

	// USERS
	server.POST("/users", signup)
	server.POST("/login", login)
}
