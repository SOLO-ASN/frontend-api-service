package jwt

import (
	"api-service/config"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get jwt from header
		jwtString := c.GetHeader("Authorization")
		if jwtString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized."})
			c.Abort()
			return
		}

		// parse jwt
		token, err := parseJwtIntoToken(jwtString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized."})
			c.Abort()
			return
		}

		// get claims from token
		claims, err := getClaimsFromToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized."})
			c.Abort()
			return
		}

		// set claims to context
		c.Set("claims", claims)

		// next handler
		c.Next()

	}
}

func decodeSecretToBytes(secret string) ([]byte, error) {
	secretBytes, err := base64.RawStdEncoding.DecodeString(secret)
	if err != nil {
		return nil, err
	}
	return secretBytes, nil
}

func parseJwtIntoToken(jwtString string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return decodeSecretToBytes(config.Get().MiddleWare.JWTSecret)
	})
	return token, err
}

func getClaimsFromToken(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token.")
}

func ParseTokenIntoClaims(jwtString string) (jwt.MapClaims, error) {
	// parse jwt
	token, err := parseJwtIntoToken(jwtString)
	if err != nil || !token.Valid {
		return nil, err
	}
	return getClaimsFromToken(token)
}
