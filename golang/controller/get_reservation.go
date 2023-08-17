package controller

import (
	"errors"
	"hotel/model"
	"hotel/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetReservationByID gets a reservation from the database by the path parameter id
func GetReservationByID(ctx *router.RouterContext) {
	reservationId := ctx.Param("id")

	reservation := model.Reservation{Id: reservationId}

	// Check if reservation exists by ID
	if err := ctx.Preload("User").Preload("Room").First(&reservation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "reservation not found"})

			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	t := time.Unix(reservation.Date, 0)
	ctx.JSON(http.StatusOK, gin.H{
		"id":   reservation.Id,
		"date": t.Format("2006-01-02"), // Format date to YYYY-MM-DD for the user to be more readable
		"user": reservation.User.Name,
		"room": reservation.Room.DoorNumber,
	})

}
