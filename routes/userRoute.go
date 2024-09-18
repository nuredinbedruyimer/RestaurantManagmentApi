package routes

import (
	controller "restaurant_manegment_api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(userRoutes *gin.Engine) {
	//  Get All User
	userRoutes.GET("/users", controller.GetUsers())
	//  Get Single User
	userRoutes.GET("/users/:user_id", controller.GetUser())
	//  Register User Using Post
	userRoutes.POST("/users/signup", controller.SignUp())
	//  Login User Using Post Method
	userRoutes.POST("/users/signin", controller.SignIn())

}
