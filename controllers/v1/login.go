package v1

import (
	"BoardGame/configs"
	"BoardGame/models"
	"BoardGame/utils"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(configs.Cfg.JWT_SECRET_KEY)

type LoginController struct{}

func (l *LoginController) Login(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if validationErr := c.ShouldBindJSON(&req); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: validationErr.Error(),
			Data:    nil,
		})
		return
	}

	var user models.User
	if err := models.DB.Where("name = ?", req.Name).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if user.Password != req.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Invalid credentials",
			Data:    nil,
		})
		return
	}
	// Assuming the validation is successful, create a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (e.g., 24 hours)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	log.Println("Token: ", tokenString)

	// Set the token in a cookie
	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true, // Set it to true if using HTTPS
		Path:     "/",
	}
	http.SetCookie(c.Writer, &cookie)

	// Return a success response or redirect to the desired page
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}
