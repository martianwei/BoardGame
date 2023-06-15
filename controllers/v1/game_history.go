package v1

import (
	"BoardGame/models"
	"BoardGame/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type GameHistoryController struct{}

func (g *GameHistoryController) GetGameHistories(c *gin.Context) {
	gameID, err := uuid.FromString(c.Query("game_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid game id",
			Data:    nil,
		})
		return
	}
	var gameHistories []models.GameHistory
	if err := models.DB.Where("game_id = ?", gameID).Find(&gameHistories).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get game histories",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Game histories found",
		Data:    gameHistories,
	})
}
