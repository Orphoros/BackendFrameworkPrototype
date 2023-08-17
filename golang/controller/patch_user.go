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

// PatchUser updates a user by the path parameter id
func PatchUser(ctx *router.RouterContext) {
	user := model.UserUpdate{}

	if !ctx.ControlReqBodyPresent(&user) {
		return
	}

	var hashedPw, name, email string

	// Using if statements due to optional data
	if user.Name != "" {
		name = user.Name
	}

	if user.Email != "" {
		email = user.Email
	}

	if user.Password != "" {
		var err error
		hashedPw, err = security.HashPassword(user.Password)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate user data"})

			return
		}
	}

	saveUser := model.User{
		Id:       ctx.Param("id"),
		Name:     name,
		Email:    email,
		Password: hashedPw,
	}

	// Update the user

	result := ctx.Model(&saveUser).Where("id = ?", saveUser.Id).Updates(saveUser)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})

			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	// Query back the user to return the updated data as a confirmation

	updatedUser := model.User{}

	if err := ctx.First(&updatedUser, "id = ?", saveUser.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			ctx.Abort()

			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": updatedUser.Id, "name": updatedUser.Name, "email": updatedUser.Email})
}
