package middleware

import (
	"blog/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT 验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供令牌"})
			c.Abort()
			return
		}
		userID, isValid := config.VerifyToken(token)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效令牌"})
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
