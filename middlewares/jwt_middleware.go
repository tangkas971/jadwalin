package middlewares

import (
	"jadwalin/services"
	"jadwalin/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(blacklistService services.BlacklistTokenService) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer "){
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorized",
			})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		isBlacklisted, err := blacklistService.IsBlacklisted(token)
		if err != nil || isBlacklisted {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Token tidak valid",
			})
			ctx.Abort()
			return 
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Invalid Token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("userId", uint(claims.UserId))
		ctx.Set("userEmail", claims.Email)
		ctx.Set("userRole", claims.Role)
		ctx.Set("userProdi", uint(claims.ProdiId))
		ctx.Next()
	}
}