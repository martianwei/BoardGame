package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Name      string    `gorm:"not null;unique" json:"name"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
