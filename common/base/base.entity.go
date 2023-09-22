package base

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        *string        `json:"id"`
	Status    string         `json:"status" gorm:"column:status;"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// Hook to generate uuid before create entity
func (record *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if record.ID == nil {
		id := uuid.New().String()
		record.ID = &id
	}

	return
}
