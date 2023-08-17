package controller

import (
	"encoding/json"
	"errors"
	"hotel/model"
	"hotel/router"
	"hotel/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PatchReservation updates a reservation by the path parameter id
func PatchReservation(ctx *router.RouterContext, jwt *security.JWT) {
	reservationId := ctx.Param("id")

	resInDB := model.Reservation{Id: reservationId}

	// Check if reservation exists by ID
	if err := ctx.Preload("User").Preload("Room").First(&resInDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "reservation not found"})

			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	token := ctx.GetTokenFromHeader()

	currentUserId, err := jwt.GetUserIdClaim(*token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})

		return
	}

	// Check if user is authorized to update reservation
	if resInDB.User.Id != *currentUserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "you are not authorized to update this reservation"})

		return
	}

	reservation := model.ReservationCreate{}

	// Check if body is defined
	if ctx.Request.Body == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "body is not defined"})

		return
	}
	jsonDecoder := json.NewDecoder(ctx.Request.Body)
	if jsonDecoder.Decode(&reservation) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "body is not a valid json"})

		return
	}

	var date int64
	var roomNumber int

	// Check if optional date is defined
	if reservation.Date != "" {
		t, err := time.Parse("2006-01-02", reservation.Date)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "date is not of the format yyyy-mm-dd"})

			return
		}

		if t.Unix() < time.Now().Unix() {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "date cannot be in the past or today"})

			return
		}

		date = t.Unix()
	}

	// Check if optional room number is defined. Negative numbers are not allowed, and thus will be ignored
	if reservation.RoomNumber > 0 {
		var roomExists bool
		if err := ctx.Model(model.Room{
			DoorNumber: reservation.RoomNumber}).
			Select("count(*) > 0").
			Where("door_number = ?", reservation.RoomNumber).
			Find(&roomExists).
			Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

			return
		}

		if !roomExists {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "room does not exist"})

			return
		}

		roomNumber = reservation.RoomNumber
	}

	// create reservation to update
	reservationToUpdate := model.Reservation{
		Id:             reservationId,
		Date:           date,
		RoomDoorNumber: roomNumber,
	}

	// update reservation
	if err := ctx.Model(&reservationToUpdate).Updates(reservationToUpdate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}
}
