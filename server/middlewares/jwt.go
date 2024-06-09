package middlewares

import (
	"API_BIG_WORK/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		// printf("")
		jwt, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("id", jwt.ID)
		c.Next()
	}
}
