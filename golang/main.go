package main

import (
	"fmt"
	"hotel/controller"
	"hotel/middleware"
	"hotel/model"
	"hotel/router"
	"hotel/security"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database credentials for local development
const (
	DB_SCHEMA   = "reservation"
	DB_USERNAME = "api_user"
	DB_PW       = "Test12345!"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PW, DB_SCHEMA)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.Room{}, &model.Reservation{})

	jwt := security.NewJWT("root-pw")

	ginRouter := gin.Default()
	ginRouter.HandleMethodNotAllowed = true

	api := ginRouter.Group("/api/v1") // root path for all api endpoints
	secured := api.Group("/")         // path for all secured endpoints that require a valid token
	userApi := api.Group("/")         // path for all endpoints that require a valid token and the user id in the path to match the id in the token

	secured.Use(middleware.Auth(jwt, db))
	userApi.Use(middleware.Auth(jwt, db)).Use(middleware.EntityIdMatch(jwt, db))

	api.POST("/login", func(c *gin.Context) { controller.Login(router.NewRouterCtx(c, db), jwt) })
	api.GET("/users/:id", func(c *gin.Context) { controller.GetUserByID(router.NewRouterCtx(c, db)) })
	api.POST("/users", func(c *gin.Context) { controller.CreateUser(router.NewRouterCtx(c, db)) })
	api.GET("/users", func(c *gin.Context) { controller.GetAllUsers(router.NewRouterCtx(c, db)) })
	api.GET("/rooms", func(c *gin.Context) { controller.GetAllRooms(router.NewRouterCtx(c, db)) })
	api.GET("/rooms/:id", func(c *gin.Context) { controller.GetRoomByID(router.NewRouterCtx(c, db)) })
	api.GET("/reservations", func(c *gin.Context) { controller.GetAllReservations(router.NewRouterCtx(c, db)) })
	api.GET("/reservations/:id", func(c *gin.Context) { controller.GetReservationByID(router.NewRouterCtx(c, db)) })
	userApi.DELETE("/users/:id", func(c *gin.Context) { controller.DeleteUser(router.NewRouterCtx(c, db)) })
	userApi.PATCH("/users/:id", func(c *gin.Context) { controller.PatchUser(router.NewRouterCtx(c, db)) })
	secured.POST("/reservations", func(c *gin.Context) { controller.CreateReservation(router.NewRouterCtx(c, db), jwt) })
	secured.PATCH("/reservations/:id", func(c *gin.Context) { controller.PatchReservation(router.NewRouterCtx(c, db), jwt) })
	secured.DELETE("/reservations/:id", func(c *gin.Context) { controller.DeleteReservation(router.NewRouterCtx(c, db), jwt) })
	secured.GET("/rooms/status", func(c *gin.Context) { controller.GetRoomsStatus(router.NewRouterCtx(c, db)) })

	ginRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})

	ginRouter.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "method not allowed"})
	})

	ginRouter.Run(":8080") // listen and serve on localhost:8080
}
