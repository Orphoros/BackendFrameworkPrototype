package controller

import (
	"errors"
	"hotel/model"
	"hotel/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserByID returns a user by the path parameter id
func GetUserByID(ctx *router.RouterContext) {
	userId := ctx.Param("id")

	user := model.User{}

	// Check if user exists by ID
	if err := ctx.First(&user, "id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			ctx.Abort()

			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	// Return user without DB specific data and password
	ctx.JSON(http.StatusOK, gin.H{"id": user.Id, "name": user.Name})
}
