package routes

import (
	"jadwalin/controller"
	"jadwalin/middlewares"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func SubjectRoute(router *gin.Engine, subjectController *controller.SubjectController, blacklistService services.BlacklistTokenService){
	v1 := router.Group("v1")
	auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))

	auth.POST("/subjects", subjectController.Create)
	auth.GET("/subjects", subjectController.GetAll)
	auth.DELETE("/subjects/:id", subjectController.Delete)
	auth.PUT("/subjects/:id", subjectController.Update)
}