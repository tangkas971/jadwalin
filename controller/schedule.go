package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"jadwalin/utils"
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
	// utils
	userRole := utils.GetUserRole(ctx)
	userId := utils.GetUserId(ctx)
	userProdi := utils.GetUserProdi(ctx)

	var scheduleDTO dto.ScheduleRequestDTO

	err := ctx.ShouldBindJSON(&scheduleDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return 
	}

	scheduleDTO.UserId = userId
	scheduleDTO.UserRole = userRole
	scheduleDTO.ProdiId = userProdi
	
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

func (c *ScheduleController) GetAll(ctx *gin.Context){
	scheduleDTOs, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, scheduleDTOs)
}

func (c *ScheduleController) Delete(ctx *gin.Context){
	id, _ := utils.GetIdParam(ctx)

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "berhasil menghapus schedule", 
	})
}

func (c *ScheduleController) Update(ctx *gin.Context){
	id, _ := utils.GetIdParam(ctx)
	
	var schedule dto.ScheduleRequestDTO
	err := ctx.ShouldBindJSON(&schedule)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(), 
		})
		return 
	}

	err = c.service.Update(id, schedule)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "schedule successfully updated",
	})
}