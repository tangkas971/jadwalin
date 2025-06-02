package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"net/http"
	"strconv"

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
	roleAny, exists := ctx.Get("userRole")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Unauthorized", 
		})
		return 
	}

	idAny,_ := ctx.Get("userId")

	roleUser := roleAny.(string)
	idUser := idAny.(uint)

	var scheduleDTO dto.CreateScheduleRequestDTO

	err := ctx.ShouldBindJSON(&scheduleDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return 
	}

	schedule, err := c.service.Create(idUser, roleUser, scheduleDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, schedule)

}

func (c *ScheduleController) FindAll(ctx *gin.Context){
	schedules, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : "gagal mendapatkan data schedules",
		})
		return
	}

	ctx.JSON(http.StatusOK, schedules)
}

func (c *ScheduleController) FindById(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "id tidak valid", 
		})
		return 
	}
	scheduleDTO, err := c.service.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, scheduleDTO)
}

func (c *ScheduleController) Update(ctx *gin.Context){
	role, exists := ctx.Get("userRole")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Unauthorized", 
		})
		return 
	}

	roleUser := role.(string)
	var scheduleDTO dto.UpdateScheduleRequestDTO
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "id tidak valid",
		})
		return 
	}

	err = ctx.ShouldBindJSON(&scheduleDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	schedule, err := c.service.Update(roleUser, id, scheduleDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, schedule)
}

func (c *ScheduleController) Delete(ctx *gin.Context){
	role, exists := ctx.Get("userRole")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Unauthorized", 
		})
		return 
	}

	roleUser := role.(string)

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "id tidak valid",
		})
		return 
	}

	err = c.service.Delete(roleUser, id)
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