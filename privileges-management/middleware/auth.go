package middleware

import (
	"github.com/gin-gonic/gin"
	auth "github.com/korylprince/go-ad-auth"
	"privileges-management/database"
)

func LDAPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, _ := c.Request.BasicAuth()
		config := database.GetActiveDirectoryAuthConfig()
		status, err := auth.Authenticate(config, username, password)

		if err != nil {
			c.JSON(401, gin.H{"error": "Authentication failed"})
			c.Abort()
			return
		}

		if !status {
			c.JSON(401, gin.H{"error": "Authentication failed"})
			c.Abort()
			return
		}

		c.Next()
	}
}
