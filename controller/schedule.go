package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	service services.ScheduleService
}

func NewScheduleController(service services.ScheduleService) *ScheduleController {
	return &ScheduleController{
		service: service,
	}
}

func (c *ScheduleController) Create(ctx *gin.Context){
	roleAny, _ := ctx.Get("userRole")
	idAny,_ := ctx.Get("userId")
	prodiAny,_ := ctx.Get("userProdi")
	log.Println("ini prodiId ctrl", prodiAny)
	log.Println("ini user Role", roleAny)

	userRole := roleAny.(string)
	userId := idAny.(uint)
	userProdi := prodiAny.(uint)
	log.Println("ini user prodi uint", userProdi)

	var scheduleDTO dto.ScheduleRequestDTO

	err := ctx.ShouldBindJSON(&scheduleDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return 
	}

	scheduleDTO.UserId = int(userId)
	scheduleDTO.UserRole = userRole
	scheduleDTO.ProdiId = int(userProdi)
	
	err = c.service.Create(scheduleDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message" : "schedules successfully created",
	})

}
