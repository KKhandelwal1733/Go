package middleware

import (
	"net/http"
	"example.com/myapp/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context){
	token := c.Request.Header.Get("Authorization")
	if token =="" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: No token provided"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: Invalid token"})
	}
	c.Set("userId", userId)
	c.Next()
}