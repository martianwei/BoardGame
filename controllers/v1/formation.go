package v1

import (
	"BoardGame/models"
	"BoardGame/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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

	// 	if validationErr := c.ShouldBindJSON(&req); validationErr != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
	// 			Success: false,
	// 			Message: validationErr.Error(),
	// 			Data:    nil,
	// 		})
	// 		return
	// 	}

	formation := &models.Formation{
		ID:     uuid.NewV4(),
		UserID: req.UserID,
		Name:   req.Name,
	}

	// 	// Execute the create operation using GORM's Create method
	// 	if err := models.DB.Create(formation).Error; err != nil {
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{
	// 			Success: false,
	// 			Message: "Failed to create formation",
	// 			Data:    nil,
	// 		})
	// 		return
	// 	}

	// Create FormationElements
	for _, element := range req.FormationElements {
		element.ID = uuid.NewV4()
		element.FormationID = formation.ID

		if err := models.DB.Create(&element).Error; err != nil {
			// If there is an error creating a FormationElement, rollback the transaction and delete the created Formation
			models.DB.Rollback()
			models.DB.Delete(formation)
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{
				Success: false,
				Message: "Failed to create formation elements",
				Data:    nil,
			})
			return
		}
	}
	log.Println(utils.CommanderInChief)
	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Formation created",
		Data:    formation,
	})
}
