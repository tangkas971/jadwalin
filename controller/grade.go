package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"jadwalin/utils"
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
	userRole := utils.GetUserRole(ctx)
	
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

func (c *GradeController) GetAll(ctx *gin.Context){
	gradeDTOs, err := c.gradeService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gradeDTOs)
}

func (c *GradeController) Delete(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)
	id,_ := utils.GetIdParam(ctx)

	err := c.gradeService.Delete(userRole, id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "grade successfully deleted",
	})
}

func (c *GradeController) Update(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)
	id, _ := utils.GetIdParam(ctx)

	var gradeDTO dto.GradeRequestDTO
	err := ctx.ShouldBindJSON(&gradeDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	err = c.gradeService.Update(userRole, id, gradeDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(), 
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "grade successfully updated",
	})
}