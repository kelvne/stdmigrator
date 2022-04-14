package main

import (
	"flag"
	"log"
	"os"

	"github.com/kelvne/stdmigrator/migrator"
)

var (
	config        *migrator.Config
	command       string
	migrationName string
)

func init() {
	config = &migrator.Config{}
	flag.StringVar(&config.MigrationsPath, "dir", "", "Migrations directory path")
	flag.StringVar(&config.Host, "host", "localhost", "PostgreSQL Host")
	flag.StringVar(&config.Port, "port", "5432", "PostgreSQL Port")
	flag.StringVar(&config.DatabaseName, "dbname", "", "PostgreSQL Database name")
	flag.StringVar(&config.Username, "user", "postgres", "PostgreSQL Database User")
	flag.StringVar(&config.Password, "password", "", "PostgreSQL Database Password")
	flag.StringVar(&migrationName, "migration-name", "", "Migration name for --command create")
	flag.StringVar(&command, "command", "", "Goose command")
	flag.Parse()
}

func createRunner() *migrator.Runner {
	runner, err := migrator.New(config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return runner
}

func main() {
	switch command {
	case "create":
		if err := migrator.CreateMigration(migrationName, config.MigrationsPath); err != nil {
			log.Fatalln(err.Error())
		}
	case "up":
		if err := createRunner().Up(); err != nil {
			log.Fatalln(err.Error())
		}
	case "down":
		if err := createRunner().Down(); err != nil {
			log.Fatalln(err.Error())
		}
	case "reset":
		if err := createRunner().Reset(); err != nil {
			log.Fatalln(err.Error())
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
