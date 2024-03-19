package routes

import (
	"net/http"
	"strconv"

	"example.com/eventbooking-rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := context.GetInt64("userId")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errorParsed": err})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errorGetEvent": err})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errorRegister": err})
		return
	}
}
