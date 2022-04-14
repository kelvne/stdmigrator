package migrator

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func openGorm(connStr string, idleCount, openCount, lifetimeSecs int) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connStr)
	if err == nil {
		db.DB().SetMaxIdleConns(idleCount)
		db.DB().SetMaxOpenConns(openCount)
		db.DB().SetConnMaxLifetime(time.Duration(lifetimeSecs) * time.Second)

		if err := db.DB().Ping(); err != nil {
			return nil, err
		}
	}
	return db, err
}

// CreateMigration creates a new migration
func CreateMigration(migrationName, migrationsDirectory string) error {
	return goose.Create(nil, migrationsDirectory, migrationName, "sql")
}
