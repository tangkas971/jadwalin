package routes

import (
	"jadwalin/controller"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func StudentTask(router *gin.Engine, studentTaskController *controller.StudentTaskController, blacklistService services.BlacklistTokenService){
	// v1 := router.Group("v1")
	// auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))

	// auth.GET("/studentTasks", studentTaskController.GetAll)
}