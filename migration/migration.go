package migration

import (
	userEntity "go-clean-architecture/internal/user/business/entity"

	"gorm.io/gorm"
)

// Migration is a function that performs database migration for the User entity.
// It takes a *gorm.DB as a parameter and returns an error if the migration fails.
func Migration(db *gorm.DB) error {
	// db.Migrator().DropTable(
	// 	userEntity.User{},
	// )

	// Perform automatic migration for entities
	err := db.AutoMigrate(
		userEntity.User{},
	)

	// Return the error if migration fails
	return err
}
