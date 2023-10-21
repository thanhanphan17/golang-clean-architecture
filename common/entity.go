package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// BeforeCreate is a hook that generates a UUID before creating an entity.
// It ensures that the ID field is populated with a unique identifier.
func (record *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	// Check if the ID field is empty
	if record.ID == "" {
		// Generate a new UUID string
		id := uuid.New().String()
		// Assign the generated UUID to the ID field
		record.ID = id
	}

	return
}
