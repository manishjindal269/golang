package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manishjindal269/golang/controllers"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	// router.Use(middleware.DeserializeUser())

	router.GET("/p/products", uc.userController.GetProducts)
	router.GET("/", uc.userController.FindUsersTest)
	router.POST("/", uc.userController.CreateUser)
	router.PUT("/:userId", uc.userController.UpdateUser)

}
