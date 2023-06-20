package models

import (
	uuid "github.com/satori/go.uuid"
)

type GameHistory struct {
	ID       int       `gorm:"type:serial;primaryKey;" json:"id"`
	GameID   uuid.UUID `gorm:"type:uuid;not null" json:"game_id"`
	Step     int       `gorm:"not null" json:"step"`
	MoveFrom string    `gorm:"not null" json:"move_from"`
	MoveTo   string    `gorm:"not null" json:"move_to"`
}
