package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AdminAuthMiddleware(c *gin.Context) {

	accessToken := c.Request.Header.Get("Authorization")

	accessToken = strings.TrimPrefix(accessToken, "Bearer")

	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("adminaccesstokena983274uhweirbt"), nil
		//pass env secret
	})

	if err != nil {

		fmt.Println("INvalid Admin AccessToken ")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}
	c.Next()
}
