package main

import (
	"music-player-api/internal/app/config"
	"music-player-api/internal/app/repository/postgres"
	"music-player-api/pkg/logger"
)

func main() {
	log := logger.NewLogrusLogger()
	logFile, err := log.SetupLogger("logs")
	if err != nil {
		panic(err)
	}

	log.Info("logger successfully setup")

	defer logFile.Close()

	cfg := config.NewViperConfig()
	if err := cfg.LoadConfig("./config.yml"); err != nil {
		log.Error(err)

		return
	}

	log.Info("configuration successfully loaded")

	dsn, err := cfg.GetFromDB("ConnString")
	if err != nil {
		log.Error(err)

		return
	}

	_, err = postgres.NewPGPool(dsn, 100)
	if err != nil {
		log.Error(err)

		return
	}

	log.Info("postgres successfully connected")

}
