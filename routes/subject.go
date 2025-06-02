package routes

import (
	"jadwalin/controller"
	"jadwalin/middlewares"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func SubjectRoutes(router *gin.Engine, subjectController *controller.SubjectController, blacklistService services.BlacklistTokenService){
	v1 := router.Group("v1")
	auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))

	auth.GET("/subjects", subjectController.FindAll)
	auth.GET("/subjects/:id", subjectController.FindById)
	auth.POST("/subjects", subjectController.Create)
	auth.PUT("/subjects/:id", subjectController.Update)
	auth.DELETE("/subjects/:id", subjectController.Delete)

}