package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type GameHistory struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	GameID    uuid.UUID `gorm:"type:uuid;not null" json:"game_id"`
	Step      int       `gorm:"not null" json:"step"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	MoveFrom  string    `gorm:"not null" json:"move_from"`
	MoveTo    string    `gorm:"not null" json:"move_to"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
