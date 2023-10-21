package repository

import (
	"go-clean-architecture/internal/user/business/entity"

	"gorm.io/gorm"
)

// applyFilterScope applies the given filter to the provided database query and returns the modified query.
func applyFilterScope(filter *entity.Filter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// If the filter is nil, return the original query.
		if filter == nil {
			return db
		}

		// Apply the role filter if it is not nil.
		if filter.Role != nil {
			db = db.Where("role = ?", filter.Role)
		}

		// Apply the status filter if it is not nil.
		if filter.Status != nil {
			db = db.Where("status = ?", filter.Status)
		}

		return db
	}
}
