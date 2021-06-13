package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(config *Config) (*gorm.DB, error) {
	dns := config.DNS()
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}
	return db, nil
}
