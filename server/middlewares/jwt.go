package middlewares

import (
	"API_BIG_WORK/config"
	"API_BIG_WORK/models"
	"API_BIG_WORK/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		// 从redis中获得id
		if config.Redis {
			id, err := models.Redis.Get(token).Result()
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				c.Abort()
				return
			}
			c.Set("id", id)
		} else {
			jwt, err := utils.ParseToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				c.Abort()
				return
			}
			c.Set("id", jwt.ID)
		}
		c.Next()
	}
}
