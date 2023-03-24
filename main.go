package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/manishjindal269/golang/config"
	"github.com/manishjindal269/golang/controllers"
	"github.com/manishjindal269/golang/routes"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	server *gin.Engine

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	db = config.ConnectDB()
	UserController = controllers.NewUserController(db)
	UserRouteController = routes.NewRouteUserController(UserController)
	server = gin.Default()
}
func main() {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/", func(ctx *gin.Context) {
		message := "connected"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	UserRouteController.UserRoute(router)
	server.Run("localhost:5000")
}
