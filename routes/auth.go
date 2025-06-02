package routes

import (
	"jadwalin/controller"
	"jadwalin/middlewares"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authController *controller.AuthController, blacklistService services.BlacklistTokenService){
	v1 := router.Group("v1")
	auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))

	auth.POST("/register/student", authController.RegisterStudent)
	auth.POST("/register/lecturer", authController.RegisterLecturer)
	auth.POST("/register/admin", authController.RegisterAdmin)
	v1.POST("/login", authController.Login)
	auth.POST("/logout", authController.Logout)

}
