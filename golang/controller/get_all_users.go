package controller

import (
	"hotel/model"
	"hotel/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers returns all users from the database
func GetAllUsers(ctx *router.RouterContext) {
	var users []model.User

	result := ctx.Find(&users)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not communicate with database"})

		return
	}

	// Create a new struct to return to the client to hide the password and other DB related fields
	type ReturnUser struct {
		Id   string
		Name string
	}

	returnObj := make([]ReturnUser, len(users))

	for i, user := range users {
		returnObj[i] = ReturnUser{
			Id:   user.Id,
			Name: user.Name,
		}
	}

	ctx.JSON(http.StatusOK, returnObj)
}
