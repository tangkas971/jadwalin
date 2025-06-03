package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GradeController struct {
	gradeService services.GradeService
}

func NewGradeController(gradeService services.GradeService) *GradeController{
	return &GradeController{
		gradeService:gradeService ,
	}
}

func (c *GradeController) Create(ctx *gin.Context){
	roleAny, exists := ctx.Get("userRole")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return 
	}
	userRole := roleAny.(string)
	var gradeDTO dto.GradeRequestDTO

	err := ctx.ShouldBindJSON(&gradeDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	err = c.gradeService.Create(userRole, gradeDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message" : "grade berhasil dibuat",
	})
}