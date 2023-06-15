package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Formation struct {
	ID                uuid.UUID          `gorm:"type:uuid;primaryKey;" json:"id"`
	UserID            uuid.UUID          `gorm:"type:uuid;not null" json:"user_id"`
	Name              string             `gorm:"not null;unique" json:"name"`
	CreatedAt         time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
	FormationElements []FormationElement `gorm:"foreignKey:FormationID" json:"formation_elements"`
}
