package middleware

import (
	"fmt"
	"net/http"

	helper "github.com/Franso/golang-jwt/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authorization headers provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
	}
}

func ValidateToken() {}
