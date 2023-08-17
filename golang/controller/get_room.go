package controller

import (
	"errors"
	"hotel/model"
	"hotel/router"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetRoomByID returns a room by the path parameter id
func GetRoomByID(ctx *router.RouterContext) {
	roomId := ctx.Param("id")

	roomNum, err := strconv.Atoi(roomId)

	if err != nil {
		// Room is not a number, so it cannot exist
		ctx.JSON(http.StatusNotFound, gin.H{"message": "room not found"})

		return
	}

	room := model.Room{DoorNumber: roomNum}

	// Check if room exists by ID
	if err := ctx.First(&room).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "room not found"})

			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	// Return room without DB specific data
	ctx.JSON(http.StatusOK, gin.H{
		"door_number": room.DoorNumber,
		"floor":       room.Floor,
		"description": room.Description,
	})
}
