package controller

import (
	"hotel/model"
	"hotel/router"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetRoomsStatus returns the status of all rooms
func GetRoomsStatus(ctx *router.RouterContext) {
	period := ctx.Query("period") // epoch format

	if period == "" {
		period = "0" // Default to 0 for converting the string to a number
	}

	rooms := []model.Room{}

	periodDurationEpochSeconds, err := strconv.ParseInt(period, 10, 64)
	periodStartEpoch := time.Now().Unix()

	if err != nil {
		ctx.JSON(http.StatusOK, rooms)

		return
	}

	result := ctx.Preload("Reservations").Find(&rooms)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	// Create a new struct to return to the client to hide the password and other DB related fields
	type ReturnReservation struct {
		IsReserved bool
		ReservedOn string
		Room       int
		ReservedBy string
	}

	// Create a new empty array to return to the client
	returnObj := []ReturnReservation{}

	// Loop through all rooms and check if they are reserved
	for _, room := range rooms {
		// Check if room has any reservations
		if room.Reservations == nil || len(room.Reservations) == 0 { // No reservations
			returnObj = append(returnObj, ReturnReservation{
				IsReserved: false,
				ReservedOn: "", // No date reserved
				Room:       room.DoorNumber,
				ReservedBy: "", // No one
			})
			continue
		}

		// Room has reservation(s). Loop through them and check if they are in the period
		for _, reservation := range room.Reservations {
			// Check if a period is not defined. If not, return all reservation data for the rooms
			if periodDurationEpochSeconds == 0 {
				returnObj = append(returnObj, ReturnReservation{
					IsReserved: true,
					ReservedOn: time.Unix(reservation.Date, 0).Format("2006-01-02"),
					Room:       room.DoorNumber,
					ReservedBy: reservation.UserID,
				})
				continue
			}
			// If period is defined, check if reservation is in the period
			if reservation.Date >= periodStartEpoch && reservation.Date < periodStartEpoch+periodDurationEpochSeconds {
				returnObj = append(returnObj, ReturnReservation{
					IsReserved: true,
					ReservedOn: time.Unix(reservation.Date, 0).Format("2006-01-02"),
					Room:       room.DoorNumber,
					ReservedBy: reservation.UserID,
				})
				continue
			}

			returnObj = append(returnObj, ReturnReservation{
				IsReserved: false,
				ReservedOn: "",
				Room:       room.DoorNumber,
				ReservedBy: "",
			})
		}
	}

	ctx.JSON(http.StatusOK, returnObj)
}
