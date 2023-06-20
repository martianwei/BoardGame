package models

import (
	uuid "github.com/satori/go.uuid"
)

type GameFormation struct {
	ID          int       `gorm:"type:serial;primaryKey;" json:"id"`
	GameID      uuid.UUID `gorm:"type:uuid" json:"game_id"`
	FormationID uuid.UUID `gorm:"type:uuid" json:"formation_id"`
	MoveOrder   int       `gorm:"type:int" json:"move_order"`
}
