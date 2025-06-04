package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"jadwalin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProdiController struct {
	prodiService services.ProdiService
}

func NewProdiController(prodiService services.ProdiService) *ProdiController{
	return &ProdiController{
		prodiService: prodiService,
	}
}

func (c *ProdiController) Create(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)
	
	var prodiDTO dto.ProdiRequestDTO

	err := ctx.ShouldBindJSON(&prodiDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	err = c.prodiService.Create(userRole, prodiDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(), 
		})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message" : "prodi succesfully created",
	})
}

func (c *ProdiController) GetAll(ctx *gin.Context){
	prodiDTOs, err := c.prodiService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, prodiDTOs)
}

func (c *ProdiController) Update(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)
	id, _ := utils.GetIdParam(ctx)

	var prodiDTO dto.ProdiRequestDTO
	err := ctx.ShouldBindJSON(&prodiDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	err = c.prodiService.Update(userRole, id, prodiDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(), 
		})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message" : "prodi successfully updated",
	})
}

func (c *ProdiController) Delete(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)
	id, _ := utils.GetIdParam(ctx)

	err := c.prodiService.Delete(userRole, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "berhasil menghapus prodi",
	})
}