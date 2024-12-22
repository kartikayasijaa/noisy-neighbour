package commands

import (
	"log"
	"noisy_neighbour/common/database"
)

type MigrateCommand struct{}

func (migrateCommand MigrateCommand) GetName() string {
	return "migrate"
}

func (migrateCommand MigrateCommand) Run() error {
	log.Println("Migrations started.")
	err := database.Migrate()

	if err != nil {
		return err
	}

	log.Println("Migrations completed.")
	return nil
}
