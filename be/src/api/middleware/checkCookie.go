package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CheckCookie(c *gin.Context) {
	// Request for cookie
	cookie, err := c.Cookie("jwt-cookie")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"code": http.StatusUnauthorized, "message": "Unauthorized User, cookie not found"})
	}

	// ParsewithClaims
	// Parse JWT string and store results in claims
	// Pass in jwt key as well
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(LoadENV("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"code": http.StatusUnauthorized, "message": "Unauthorized User"})
		return
	}

	// Invalid token
	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"code": http.StatusUnauthorized, "message": "Unauthorized User, invalid cookie"})
		return
	}

	c.Next()
}
