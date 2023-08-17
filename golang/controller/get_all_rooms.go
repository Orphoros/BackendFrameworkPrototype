package controller

import (
	"hotel/model"
	"hotel/router"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRooms(ctx *router.RouterContext) {
	freeAtEpoch := ctx.Query("free")

	rooms := []model.Room{}

	// This is the struct that will be returned to the client as we want to hide the many-to-many reservation relationship object prop of room
	type ReturnRoom struct {
		DoorNumber  int
		Floor       int
		Description string
	}

	var returnObj []ReturnRoom

	// if we want to filter the rooms by the date they are free
	if freeAtEpoch != "" {
		freeAt, err := strconv.ParseInt(freeAtEpoch, 10, 64)

		if err == nil {
			result := ctx.Where("door_number NOT IN (SELECT room_door_number FROM reservations WHERE date = ?)", freeAt).Find(&rooms)

			if result.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

				return
			}

			for _, room := range rooms {
				returnObj = append(returnObj, ReturnRoom{
					DoorNumber:  room.DoorNumber,
					Floor:       room.Floor,
					Description: room.Description,
				})
			}

			ctx.JSON(http.StatusOK, returnObj)

			return

		}
	}

	// No filter, just return all rooms

	result := ctx.Find(&rooms)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	for _, room := range rooms {
		returnObj = append(returnObj, ReturnRoom{
			DoorNumber:  room.DoorNumber,
			Floor:       room.Floor,
			Description: room.Description,
		})
	}

	ctx.JSON(http.StatusOK, returnObj)
}
