package controller

import (
	"errors"
	"hotel/model"
	"hotel/router"
	"hotel/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateReservation creates a reservation for a user
func CreateReservation(ctx *router.RouterContext, jwt *security.JWT) {
	token := ctx.GetTokenFromHeader()

	if token == nil {
		return
	}

	userId, err := jwt.GetUserIdClaim(*token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	// Get reservation data from request body
	reservationReq := model.ReservationCreate{}
	if err := reservationReq.FromHttpRequest(ctx.Request.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	date, err := time.Parse("2006-01-02", reservationReq.Date)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid date format, should be yyyy-mm-dd"})

		return
	}

	// Check if reservation already exists
	reservation := model.Reservation{
		Date:           date.Unix(),
		UserID:         *userId,
		RoomDoorNumber: reservationReq.RoomNumber,
	}

	var exists bool
	if err := ctx.Model(reservation).
		Select("count(*) > 0").
		Where("user_id = ?", reservation.UserID).
		Where("date = ?", reservation.Date).
		Where("room_door_number = ?", reservation.RoomDoorNumber).
		Find(&exists).
		Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not control reservation data"})

		return
	}

	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "reservation already exists"})

		return
	}

	reservation.Id = uuid.New().String()

	// Check if room exists for the reservation
	var roomExists bool
	if err := ctx.Model(model.Room{
		DoorNumber: reservation.RoomDoorNumber}).
		Select("count(*) > 0").
		Where("door_number = ?", reservation.RoomDoorNumber).
		Find(&roomExists).
		Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}
	if !roomExists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "room does not exist"})

		return
	}

	// Create reservation after the checks
	if err := ctx.Create(&reservation).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	// Query back the created reservation to send it back to the user as confirmation
	var createdReservation model.Reservation

	result := ctx.Model(&reservation).Preload("User").Preload("Room").First(&createdReservation)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "could not check if reservation was created"})

			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create reservation in database"})

		return
	}

	// Create response object as we don't want to send the user's password back and other DB fields
	t := time.Unix(reservation.Date, 0)
	ctx.JSON(http.StatusCreated, gin.H{
		"reservation_id": createdReservation.Id,
		"room_number":    createdReservation.Room.DoorNumber,
		"floor":          createdReservation.Room.Floor,
		"user_name":      createdReservation.User.Name,
		"date":           t.Format("2006-01-02"),
	})
}
