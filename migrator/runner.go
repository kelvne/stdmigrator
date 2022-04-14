package migrator

import (
	"fmt"

	"github.com/pressly/goose"
)

// New returns a new Runner
func New(c *Config) (*Runner, error) {
	return (&Runner{
		config: c,
	}).init()
}

func (r *Runner) init() (*Runner, error) {
	if r.rootDB != nil {
		r.rootDB.DB().Close()
	}

	rootDB, err := openGorm(r.config.rootConnectionString(), 1, 0, 1)
	if err != nil {
		return nil, err
	}

	db, err := openGorm(r.config.connectionString(), 1, 0, 1)
	if err != nil {
		return nil, err
	}

	r.rootDB = rootDB
	r.db = db

	return r, nil
}

func (r *Runner) isDatabaseCreated() (bool, error) {
	response, err := r.rootDB.DB().Exec(fmt.Sprintf(
		"SELECT 1 FROM pg_database WHERE datname='%s'",
		r.config.DatabaseName,
	))
	if err != nil {
		return false, nil
	}

	rows, err := response.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}

func (r *Runner) createDatabase() error {
	exists, err := r.isDatabaseCreated()
	if err != nil {
		return err
	}

	if !exists {
		if _, err := r.rootDB.DB().Exec(fmt.Sprintf("CREATE DATABASE %s", r.config.DatabaseName)); err != nil {
			return err
		}
	}

	return nil
}

func (r *Runner) Up() error {
	if err := r.createDatabase(); err != nil {
		return err
	}
	return goose.Up(r.db.DB(), r.config.MigrationsPath)
}

func (r *Runner) Down() error {
	return goose.Down(r.db.DB(), r.config.MigrationsPath)
}

func (r *Runner) Reset() error {
	return goose.Reset(r.db.DB(), r.config.MigrationsPath)
}
