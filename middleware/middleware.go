package middleware

import (
	"BoardGame/configs"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(configs.Cfg.JWT_SECRET_KEY)

// CORSMiddleware ...
// CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Upgrade", "Connection", "Origin"},
		AllowCredentials: true,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the cookie
		cookie, err := c.Cookie("jwt_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid JWT token cookie"})
			return
		}

		// Parse and validate the JWT token
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			// Provide the secret key used for signing the token
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		log.Println("Token: ", token)

		// Token is valid, proceed with the next middleware or handler
		c.Next()
	}
}
