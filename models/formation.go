package models

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Formation struct {
	ID        uuid.UUID     `gorm:"type:uuid;primaryKey;" json:"id"`
	UserID    uuid.UUID     `gorm:"type:uuid;not null" json:"user_id"`
	Name      string        `gorm:"not null;unique" json:"name"`
	Strategy  pq.Int32Array `gorm:"type:int[]" json:"strategy"`
	CreatedAt time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	User      User          `gorm:"foreignKey:UserID" json:"user"`
}
