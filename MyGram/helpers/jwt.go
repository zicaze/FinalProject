package helpers

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = "THE LORD OF SHADOW"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	parsetoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parsetoken.SignedString([]byte(secretKey))
	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
