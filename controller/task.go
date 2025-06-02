package controller

import (
	"jadwalin/services"
)

type TaskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *TaskController{
	return &TaskController{
		taskService: taskService,
	}
}

// func (c *TaskController) Create(ctx *gin.Context){
// 	var task dto.CreateTaskRequest
// 	err := ctx.ShouldBindJSON(&task)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error" : err.Error(),
// 		})
// 		return 
// 	}

// 	taskDTO, err := c.taskService.Create(task)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error" : err.Error(),
// 		})
// 		return 
// 	}

// 	ctx.JSON(http.StatusCreated, taskDTO)
// }