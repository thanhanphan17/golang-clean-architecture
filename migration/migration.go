package migration

import (
	usrentity "go-clean-architecture/internal/user/business/entity"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		usrentity.User{},
	)

	if err != nil {
		return
	}
}
