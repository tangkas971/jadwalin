package routes

import (
	"jadwalin/controller"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func ProdiRoute(router *gin.Engine, prodiController *controller.ProdiController, blacklistService services.BlacklistTokenService){
	v1 := router.Group("v1")
	// auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))

	v1.POST("/prodi", prodiController.Create)

}