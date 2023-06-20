package models

import (
	"BoardGame/utils"

	uuid "github.com/satori/go.uuid"
)

type FormationElement struct {
	ID          int                `gorm:"type:serial;primaryKey;" json:"id"`
	FormationID uuid.UUID          `gorm:"type:uuid;not null" json:"formation_id"`
	Commission  utils.MilitaryRank `gorm:"not null" json:"commission"`
	Position    int                `gorm:"not null" json:"position"`
}
