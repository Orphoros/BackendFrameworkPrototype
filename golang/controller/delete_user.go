package controller

import (
	"errors"
	"hotel/model"
	"hotel/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteUser deletes a user from the database by the path parameter id
func DeleteUser(ctx *router.RouterContext) {
	userId := ctx.Param("id")

	user := model.User{Id: userId}

	// Check if user exists by ID
	if err := ctx.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})

			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	result := ctx.Unscoped().Delete(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete user"})

		return
	}

	ctx.Status(http.StatusNoContent)
}
