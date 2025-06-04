package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"jadwalin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubjectController struct {
	subjectService services.SubjectService
}

func NewSubjectController(subjectService services.SubjectService) *SubjectController{
	return &SubjectController{
		subjectService: subjectService,
	}
}

func (c *SubjectController) Create(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)

	var subjectDTO dto.SubjectRequestDTO

	err := ctx.ShouldBindJSON(&subjectDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	err = c.subjectService.Create(userRole, subjectDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message" : "subject successfully created",
	})
}