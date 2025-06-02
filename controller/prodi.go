package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
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
	var prodiDTO dto.ProdiRequestDTO

	err := ctx.ShouldBindJSON(&prodiDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	err = c.prodiService.Create(prodiDTO)
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