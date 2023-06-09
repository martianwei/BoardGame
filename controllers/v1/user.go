package v1

import (
	"golangPrac/models"
	"golangPrac/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

type UserController struct{}

// func (u UserController) GetUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User

// 	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
// 		utils.Response(c, 404, "User not found")
// 		return
// 	}

// 	c.JSON(200, user)
// }

func (u *UserController) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if validationErr := c.ShouldBindJSON(&req); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: validationErr.Error(),
			Data:    nil,
		})
		return
	}

	user := &models.User{
		ID:       uuid.NewV4(),
		Name:     req.Name,
		Password: req.Password,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}
