package middleware

import (
	"hotel/router"
	"hotel/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Auth is a middleware that checks if the token is valid and the user is logged in
func Auth(jwt *security.JWT, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := router.NewRouterCtx(ctx, db)
		token := c.GetTokenFromHeader()

		if token == nil {
			return
		}

		err := jwt.ValidateToken(*token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "token is invalid, expired or not provided"})
			ctx.Abort()

			return
		}
		ctx.Next()
	}
}
