
package middleware

import (
	"homework4/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
return func(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" || !strings.HasPrefix(header, "Bearer "){
		utils.Error(c, http.StatusUnauthorized, "未提供认证令牌")
		c.Abort()
		return
	}
	tokenStr := strings.TrimPrefix(header, "Bearer ")
	claims, err := utils.ParseToken(tokenStr)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, "令牌无效或已过期")
		c.Abort()
		return
	}
	c.Set("userID", claims.UserID)
	c.Set("username", claims.Username)
	c.Next()
  }
}