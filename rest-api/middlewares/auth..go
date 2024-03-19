package middlewares

import (
	"net/http"

	"example.com/eventbooking-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	// If the Authorization header is not present, return a 401 Unauthorized status code
	// with an error message
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})

		// Return immediately as the user is not authorized to create an event
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Token",
		})
		return
	}
	context.Set("userId", userId)
	context.Next()

}
