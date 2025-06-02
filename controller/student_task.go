package controller

import (
	"jadwalin/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentTaskController struct {
	service services.StudentTaskService
}

func NewStudentTaskController(service services.StudentTaskService) *StudentTaskController{
	return &StudentTaskController{
		service: service,
	}
}

func (c *StudentTaskController) GetAll(ctx *gin.Context){
	studentIdAny, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return 
	}

	studentId,_ := studentIdAny.(uint)
	log.Println("ini id mahasiswa", studentId)
	studentTasks, err := c.service.FindByStudentTaskId(studentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, studentTasks)
}