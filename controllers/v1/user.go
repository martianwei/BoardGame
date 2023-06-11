package v1

import (
	"BoardGame/models"
	"BoardGame/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

type UserController struct{}

func (u *UserController) GetUser(c *gin.Context) {
	id, err := uuid.FromString(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid user id",
			Data:    nil,
		})
		return
	}
	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User found",
		Data:    user,
	})
}

func (u *UserController) CreateUser(c *gin.Context) {
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

func (u *UserController) UpdateUser(c *gin.Context) {
	id, err := uuid.FromString(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid user id",
			Data:    nil,
		})
		return
	}

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

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
		return
	}

	user.Name = req.Name
	user.Password = req.Password

	if err := models.DB.Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User updated successfully",
		Data:    user,
	})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id, err := uuid.FromString(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid user id",
			Data:    nil,
		})
		return
	}

	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
		return
	}

	if err := models.DB.Delete(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User deleted successfully",
		Data:    nil,
	})
}
