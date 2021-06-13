package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func doMigrations(migrationsFolder, connection string) error {
	migrations, err := migrate.New(migrationsFolder, connection)
	if err != nil {
		return err
	}
	_ = migrations.Up()
	return nil
}
