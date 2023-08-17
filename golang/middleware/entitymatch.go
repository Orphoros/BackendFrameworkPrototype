package middleware

import (
	"hotel/router"
	"hotel/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EntityIdMatch is a middleware that checks if the id in the URL matches the id in the token. Used mostly for user operations
func EntityIdMatch(jwt *security.JWT, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := router.NewRouterCtx(ctx, db)

		token := c.GetTokenFromHeader()

		if token == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "no token provided"})
			ctx.Abort()

			return
		}

		userId, err := jwt.GetUserIdClaim(*token)

		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "you are not authorized to operate on this entity"})
			ctx.Abort()

			return
		}

		if ctx.Param("id") != *userId {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "you are not authorized to operate on this entity"})
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}
