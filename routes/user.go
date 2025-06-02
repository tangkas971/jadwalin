package routes

import (
	"jadwalin/controller"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController *controller.UserController, blacklistService services.BlacklistTokenService) {
	// v1 := router.Group("v1")
	// auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))
	
	// auth.GET("/users", userController.FindAll)
	// auth.GET("/users/:id", userController.FindById)
	// auth.DELETE("/users/:id", userController.Delete)
}
