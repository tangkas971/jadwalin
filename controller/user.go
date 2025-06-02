package controller

import (
	"jadwalin/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// func (c *UserController) FindAll(ctx *gin.Context){
// 	users, err := c.service.FindAll()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error" : err.Error(),
// 		})
// 		return 
// 	}

// 	ctx.JSON(http.StatusOK, users)
// }

// func (c *UserController) FindById(ctx *gin.Context){
// 	idParam := ctx.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error" : "id tidak valid",
// 		})
// 		return
// 	}

// 	userDTO, err := c.service.FindById(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error" : err.Error(),
// 		})
// 		return 
// 	}

// 	ctx.JSON(http.StatusOK, userDTO)
// }

// func (c *UserController) Delete(ctx *gin.Context){
// 	role, exists := ctx.Get("userRole")
// 	if !exists {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{
// 			"error" : "unauthorized",
// 		})
// 		return
// 	}

// 	roleUser := role.(string)

// 	idParam := ctx.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error" : "id tidak valid", 
// 		})
// 		return 
// 	}
// 	err = c.service.Delete(roleUser, id)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error" : err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message" : "user telah berhasl dihapus",
// 	})
// }