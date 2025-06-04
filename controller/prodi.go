package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"net/http"
	"strconv"

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
	roleAny, exists := ctx.Get("userRole")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return 
	}
	userRole := roleAny.(string)
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
	roleAny, _ := ctx.Get("userRole")
	userRole := roleAny.(string)
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

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
	roleAny, _ := ctx.Get("userRole")
	userRole := roleAny.(string)
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

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