package routes

import (
	"jadwalin/controller"
	"jadwalin/middlewares"
	"jadwalin/services"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(router *gin.Engine, scheduleController *controller.ScheduleController, blacklistService services.BlacklistTokenService){
	v1 := router.Group("v1")
	auth := v1.Group("/", middlewares.AuthMiddleware(blacklistService))

	auth.GET("/schedules", scheduleController.FindAll)
	auth.GET("/schedules/:id", scheduleController.FindById)
	auth.POST("/schedules", scheduleController.Create)
	auth.PUT("/schedules/:id", scheduleController.Update)
	auth.DELETE("/schedules/:id", scheduleController.Delete)

}