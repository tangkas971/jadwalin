package routes

import (
	"jadwalin/controller"
	"jadwalin/middlewares"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine, taskController *controller.TaskController, blacklistService services.BlacklistTokenService){
	v1 := router.Group("v1")
	auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))

	auth.POST("/tasks", taskController.Create)
	auth.GET("/tasks", taskController.GetAll)
	auth.DELETE("/tasks/:id", taskController.Delete)
	auth.PUT("/tasks/:id", taskController.Update)
}