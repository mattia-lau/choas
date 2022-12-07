package main

import (
	"chaos-go/internal/config"
	"chaos-go/internal/database"
	"chaos-go/internal/logging"
	"log"
)

func main() {
	cfg, err := config.Load("config.yml")

	if err != nil {
		log.Fatal(err)
	}

	logger := logging.SetupLogger(cfg)

	database.RunMigration(cfg, logger)
}
