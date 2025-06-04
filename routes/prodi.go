package routes

import (
	"jadwalin/controller"
	"jadwalin/middlewares"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func ProdiRoute(router *gin.Engine, prodiController *controller.ProdiController, blacklistService services.BlacklistTokenService){
	v1 := router.Group("v1")
	auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))

	auth.POST("/prodi", prodiController.Create)
	auth.GET("/prodi", prodiController.GetAll)
	auth.PUT("/prodi/:id", prodiController.Update)
	auth.DELETE("/prodi/:id", prodiController.Delete)
}