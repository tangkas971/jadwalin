package routes

import (
	"jadwalin/controller"
	"jadwalin/middlewares"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func GradeRoute(router *gin.Engine, gradeController *controller.GradeController, blacklistService services.BlacklistTokenService){
	v1 := router.Group("v1")
	auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))

	auth.POST("/grades", gradeController.Create)
	auth.GET("/grades", gradeController.GetAll)
	auth.PUT("/grades/:id", gradeController.Update)
	auth.DELETE("/grades/:id", gradeController.Delete)
}