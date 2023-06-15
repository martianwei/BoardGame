package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type Game struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Player1ID uuid.UUID `gorm:"type:uuid;not null" json:"player1_id"`
	Player2ID uuid.UUID `gorm:"type:uuid;not null" json:"player2_id"`
	Player3ID uuid.UUID `gorm:"type:uuid;not null" json:"player3_id"`
	Player4ID uuid.UUID `gorm:"type:uuid;not null" json:"player4_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
