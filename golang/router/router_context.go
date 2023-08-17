package router

import (
	"encoding/json"
	"hotel/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RouterContext is a wrapper around gin.Context and GORM that adds helper functions
type RouterContext struct {
	*gin.Context
	*gorm.DB
}

// NewRouterCtx creates a new RouterContext from a gin.Context
func NewRouterCtx(ctx *gin.Context, db *gorm.DB) *RouterContext {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	return &RouterContext{ctx, db}
}

// GetTokenFromHeader gets the token from the Authorization header. If the header is not present or the token is invalid, it will abort the request and return nil
func (ctx *RouterContext) GetTokenFromHeader() *string {
	authorization := ctx.GetHeader("Authorization")
	if authorization == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "authorization required"})
		ctx.Abort()

		return nil
	}

	authKeyVal := strings.Split(authorization, " ")

	if len(authKeyVal) != 2 || authKeyVal[0] != "Bearer" || authKeyVal[1] == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "authorization required"})
		ctx.Abort()

		return nil
	}

	return &authKeyVal[1]
}

// ControlReqBodyPresent checks if the request body is present and valid json. If it is not, it will abort the request and return false. If it is, it will decode the body into the given object and return true
func (ctx *RouterContext) ControlReqBodyPresent(v any) bool {
	if ctx.Request.Body == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "request body is missing"})
		ctx.Abort()

		return false
	}

	// Check if the object has a predefined FromHttpRequest method, else use the default json decoder
	switch v := v.(type) {
	case *model.User:
		if err := v.FromHttpRequest(ctx.Request.Body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			ctx.Abort()

			return false
		}
	case *model.UserLogin:
		if err := v.FromHttpRequest(ctx.Request.Body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			ctx.Abort()

			return false
		}
	case *model.ReservationCreate:
		if err := v.FromHttpRequest(ctx.Request.Body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			ctx.Abort()

			return false
		}
	case *model.UserUpdate:
		if err := v.FromHttpRequest(ctx.Request.Body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			ctx.Abort()

			return false
		}
	default:
		jsonDecoder := json.NewDecoder(ctx.Request.Body)
		if jsonDecoder.Decode(v) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "body is not a valid json"})

			return false
		}
	}

	return true
}
