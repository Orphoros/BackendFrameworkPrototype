package controller

import (
	"errors"
	"hotel/model"
	"hotel/router"
	"hotel/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login logs a user in by returning a JWT token
func Login(ctx *router.RouterContext, jwt *security.JWT) {
	user := model.UserLogin{}

	if !ctx.ControlReqBodyPresent(&user) {
		return
	}

	userInDB := model.User{Email: user.Email}

	// Check if user exists by email as email is unique and the user does not login with the id, but by email and pw
	result := ctx.First(&userInDB, "email = ?", user.Email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "username or password is incorrect"})

			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not communicate with database"})

		return
	}

	if bcrypt.CompareHashAndPassword([]byte(userInDB.Password), []byte(user.Password)) != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "username or password is incorrect"})

		return
	}

	token, err := jwt.GenerateJWT(security.TokenDuration, userInDB.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred while processing your request"})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
