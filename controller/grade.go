package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"net/http"
	"strconv"

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
	roleAny, _ := ctx.Get("userRole")
	userRole := roleAny.(string)
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

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
	roleAny, _ := ctx.Get("userRole")
	userRole := roleAny.(string)
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

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