package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"jadwalin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *TaskController{
	return &TaskController{
		taskService: taskService,
	}
}

func (c *TaskController) Create(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)
	userId := utils.GetUserId(ctx)
	
	var taskDTO dto.TaskRequestDTO
	err := ctx.ShouldBindJSON(&taskDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	taskDTO.LecturerId = int(userId)

	err = c.taskService.Create(userRole, taskDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message" : "task successfully created",
	})
}