package controller

import (
	"hotel/model"
	"hotel/router"
	"hotel/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateUser creates a new user
func CreateUser(ctx *router.RouterContext) {
	user := model.User{}

	// Check if the request body is present and is valid
	if !ctx.ControlReqBodyPresent(&user) {
		return
	}

	var exists bool // false by default
	if err := ctx.Model(user).
		Select("count(*) > 0").
		Where("email = ?", user.Email). // Check if the email is already in use, as we don't know the user's id yet
		Find(&exists).
		Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})
		ctx.Abort()

		return
	}

	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user already exists"})
		ctx.Abort()

		return
	}

	user.Id = uuid.New().String()
	hashedPw, err := security.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate user data"})

		return
	}

	user.Password = hashedPw

	result := ctx.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not create user in database"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": user.Id, "name": user.Name, "email": user.Email})
}
