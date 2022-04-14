package migrator

import "github.com/jinzhu/gorm"

// Config describes Runner settings
//
// It holds information for database authentication and migrations location
type Config struct {
	// Host Postgres host
	Host string

	// Port Postgres host
	Port string

	// DatabaseName name of the Postgres database
	DatabaseName string

	// Username database user's username with access to the given DatabaseName
	Username string

	// Password database user's password with access to the given DatabaseName
	Password string

	// MigrationsPath directory path
	MigrationsPath string
}

// Runner runs migration actions
type Runner struct {
	config *Config
	rootDB *gorm.DB
	db     *gorm.DB
}
