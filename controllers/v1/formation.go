package v1

import (
	"BoardGame/models"
	"BoardGame/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type FormationController struct{}

func (f *FormationController) GetFormation(c *gin.Context) {
	user_id, err := uuid.FromString(c.Query("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid user id",
			Data:    nil,
		})
		return
	}
	var formation models.Formation

	if err := models.DB.Where("user_id = ?", user_id).First(&formation).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "Formation not found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Formation found",
		Data:    formation,
	})
}

func (f *FormationController) CreateFormation(c *gin.Context) {
	var req struct {
		UserID    uuid.UUID        `json:"user_id"`
		Formation models.Formation `json:"formation"`
	}

	if validationErr := c.ShouldBindJSON(&req); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: validationErr.Error(),
			Data:    nil,
		})
		return
	}

	formation := &models.Formation{
		ID:       uuid.NewV4(),
		UserID:   req.UserID,
		Name:     req.Formation.Name,
		Strategy: pq.Int32Array(req.Formation.Strategy),
	}

	if err := models.DB.Create(formation).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create formation",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Formation created",
		Data:    formation,
	})
}
