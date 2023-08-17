package controller

import (
	"errors"
	"hotel/model"
	"hotel/router"
	"hotel/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteReservation deletes a reservation by the path parameter id
func DeleteReservation(ctx *router.RouterContext, jwt *security.JWT) {
	reservationId := ctx.Param("id")

	resInDB := model.Reservation{}

	// Check if reservation exists by ID
	if err := ctx.Preload("User").Preload("Room").First(&resInDB, "id = ?", reservationId).Error; err != nil {
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

	// Check if user is authorized to delete reservation
	if resInDB.User.Id != *currentUserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "you are not authorized to update this reservation"})

		return
	}

	reservation := model.Reservation{Id: reservationId}

	if err := ctx.Delete(&reservation).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	ctx.Status(http.StatusNoContent)
}
