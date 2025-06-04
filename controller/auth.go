package controller

import (
	"jadwalin/dto"
	"jadwalin/services"
	"jadwalin/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
	blackListTokenService services.BlacklistTokenService
}

func NewAuthController(authService services.AuthService, blackListTokenService services.BlacklistTokenService) *AuthController{
	return &AuthController{
		authService: authService,
		blackListTokenService: blackListTokenService,
	}
}

func (c *AuthController) RegisterStudent(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)

	var userDTO dto.StudentRegisterRequest

	err := ctx.ShouldBindJSON(&userDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	}

	err = c.authService.CreateStudent(userRole, userDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message" : "student berhasil di buat",
	})
}

func (c *AuthController) RegisterLecturer(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)

	var userDTO dto.LecturerRegisterRequest

	err := ctx.ShouldBindJSON(&userDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	err = c.authService.CreateLecturer(userRole, userDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message" : "created lecturer successfully",
	})
}

func (c *AuthController) RegisterAdmin(ctx *gin.Context){
	userRole := utils.GetUserRole(ctx)
	
	var userDTO dto.AdminRegisterRequest
	err := ctx.ShouldBindJSON(&userDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	err = c.authService.CreateAdmin(userRole, userDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(), 
		})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message" : "admin successfully created",
	})
}

func (c *AuthController) Login(ctx *gin.Context){
	var userLogin dto.LoginUserRequest

	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	user, err := c.authService.Login(userLogin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(), 
		})
		return 
	}

	 ctx.JSON(http.StatusOK, user)
}

func (c *AuthController) Logout(ctx *gin.Context){
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "auth header kosong",
		})
		return 
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error" : "token tidak valid",
		})
		return
	}

	err = c.blackListTokenService.BlacklistToken(tokenString, claims.ExpiresAt.Time)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : "gagal logout", 
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "berhasil logout",
	})
}

