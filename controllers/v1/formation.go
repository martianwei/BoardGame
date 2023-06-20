package v1

import (
	"BoardGame/models"
	"BoardGame/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type FormationController struct{}

func (f *FormationController) GetFormation(c *gin.Context) {
	formationID, err := uuid.FromString(c.Query("formation_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid formation id",
			Data:    nil,
		})
		return
	}
	var formation models.Formation

	if err := models.DB.Preload("FormationElements").First(&formation, formationID).Error; err != nil {
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
		UserID            uuid.UUID                 `json:"user_id"`
		Name              string                    `json:"name"`
		FormationElements []models.FormationElement `json:"formation_elements"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request data",
			Data:    err.Error(),
		})
		return
	}

	formation := &models.Formation{
		ID:        uuid.NewV4(),
		UserID:    req.UserID,
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 创建 Formation
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(formation).Error; err != nil {
			return err
		}

		// 创建 FormationElements
		for i := range req.FormationElements {
			req.FormationElements[i].FormationID = formation.ID
		}
		if err := tx.Create(&req.FormationElements).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create formation",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Formation created",
		Data:    formation,
	})
}
