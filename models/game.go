package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Game struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey;" json:"id"`
	CreatedAt   time.Time     `gorm:"autoCreateTime" json:"created_at"`
	GameHistory []GameHistory `gorm:"foreignKey:GameID" json:"game_history"`
}
