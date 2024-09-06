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

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
    return []byte(cfg.ADMINSECRET), nil
	})

	if err != nil {

		fmt.Println("INvalid Admin AccessToken ")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

claims,ok  :=token.Claims.(jwt.MapClaims)
  if !ok{
    c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid authorization"})
		c.AbortWithStatus(http.StatusUnauthorized)
    return
  } 
 
  id,ok:=claims["id"].(float64)
  fmt.Println("map claim printing to check",claims)
  c.Set("id",int(id))
	c.Next()

}
