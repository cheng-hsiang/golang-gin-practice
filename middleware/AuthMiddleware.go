package middleware

import (
	"gin_api/common"
	"gin_api/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "權限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "權限不足"})
			ctx.Abort()
			return
		}

		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "權限不足"})
			ctx.Abort()
			return
		}

		// 用戶存在 將user 寫入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
