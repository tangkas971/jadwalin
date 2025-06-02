package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectController struct {
	service services.SubjectServices
}

func NewSubjectController(service services.SubjectServices) *SubjectController {
	return &SubjectController{
		service: service,
	}
}

func (c *SubjectController) Create(ctx *gin.Context){
	roleAny, exists := ctx.Get("userRole")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return
	}
	roleUser := roleAny.(string)
	var subjectDTO dto.CreateSubjectRequest

	err := ctx.ShouldBindJSON(&subjectDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	subject, err := c.service.Create(roleUser, subjectDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, subject)
}

func (c *SubjectController) FindAll(ctx *gin.Context){
	subjects, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : "gagal mendapatkan data subjects",
		})
		return 
	}

	ctx.JSON(http.StatusOK, subjects)
}

func (c *SubjectController) FindById(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "id tidak valid", 
		})
		return 
	}

	subject, err := c.service.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(), 
		}) 
		return 
	}

	ctx.JSON(http.StatusOK, subject)
}

func (c *SubjectController) Update(ctx *gin.Context){
	roleAny, exists := ctx.Get("roleUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return 
	}

	roleUser := roleAny.(string)

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "id tidak valid", 
		})
		return 
	}

	var subjectDTO dto.CreateSubjectRequest
	err = ctx.ShouldBindJSON(&subjectDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		}) 
		return 
	}

	subject, err := c.service.Update(roleUser, id, subjectDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, subject)
}

func (c *SubjectController) Delete(ctx *gin.Context){
	roleAny, exists := ctx.Get("roleUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return 
	}

	roleUser := roleAny.(string)

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "id tidak valid", 
		})
		return 
	}

	err = c.service.Delete(roleUser, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : "gagal menghapus subject", 
		})
		return 
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "subject berhasil dihapus",
	})
}