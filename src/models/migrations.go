package models

import (
	"fmt"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Skill{},
	)

	if err != nil {
		return fmt.Errorf("unable to run migrations: %w", err)
	}

	return nil
}
