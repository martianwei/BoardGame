package v1

import (
	"BoardGame/models"
	"BoardGame/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type GameController struct{}

func (g *GameController) CreateGame(c *gin.Context) {
	// 获取请求中的用户信息
	type FormationInfo struct {
		FormationID uuid.UUID `json:"formation_id"`
		MoveOrder   int       `json:"move_order"`
	}
	var req struct {
		Formations []FormationInfo `json:"formations"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request data",
			Data:    err.Error(),
		})
		return
	}

	// 创建游戏
	game := models.Game{
		ID:        uuid.NewV4(), // 使用uuid生成唯一的游戏ID
		CreatedAt: time.Now(),
	}

	// 在此处执行数据库插入操作将游戏保存到数据库中
	err = models.DB.Create(&game).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create game",
			Data:    err.Error(),
		})
		return
	}

	// 创建game_users记录
	for _, userRequest := range req.Formations {
		gameUser := models.GameFormation{
			GameID:      game.ID,
			FormationID: userRequest.FormationID,
			MoveOrder:   userRequest.MoveOrder,
		}
		err := models.DB.Create(&gameUser).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.Response{
				Success: false,
				Message: "Failed to create game user",
				Data:    err.Error(),
			})
			return
		}
	}

	// 返回创建成功的响应
	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Game create success",
		Data:    game,
	})
}
