package controller

import (
	"hotel/model"
	"hotel/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllReservations returns all reservations from the database
func GetAllReservations(ctx *router.RouterContext) {
	var reservations []model.Reservation

	// We need to load in the related user and room data to get the objects
	result := ctx.Preload("User").Preload("Room").Find(&reservations)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	// Create a new struct to return to the client to hide the password and other DB related fields
	type ReturnReservation struct {
		Id   string
		Date string
		User string
		Room int
	}

	returnObj := make([]ReturnReservation, len(reservations))

	for i, reservation := range reservations {
		t := time.Unix(reservation.Date, 0)
		returnObj[i] = ReturnReservation{
			Id:   reservation.Id,
			Date: t.Format("2006-01-02"), // format unix epoch for user to be more readable
			User: reservation.User.Name,
			Room: reservation.Room.DoorNumber,
		}
	}

	ctx.JSON(http.StatusOK, returnObj)
}
