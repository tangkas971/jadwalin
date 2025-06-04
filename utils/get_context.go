package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdParam(ctx *gin.Context)(id int, error error){
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}

	return id, nil 
}

func GetUserRole(ctx *gin.Context)(role string){
	roleAny, _ := ctx.Get("userRole")
	userRole := roleAny.(string)
	return userRole
}	

func GetUserEmail(ctx *gin.Context)(email string){
	emailAny, _ := ctx.Get("userEmail")
	userEmail := emailAny.(string)
	return userEmail
}

func GetUserId(ctx *gin.Context)(id int){
	idAny,_ := ctx.Get("userId")
	userId := idAny.(uint)
	return int(userId)
}

func GetUserProdi(ctx *gin.Context)(prodiId int){
	prodiAny,_ := ctx.Get("userProdi")
	userProdi := prodiAny.(uint)
	return int(userProdi)
}
