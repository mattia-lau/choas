package database

import (
	"chaos-go/internal/config"
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
	"go.uber.org/zap"
)

func RunMigration(cfg *config.Config, logger *zap.Logger) {
	db, err := sql.Open("sqlite3", cfg.DBConfig.Path)
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}

	fSrc, err := (&file.File{}).Open(cfg.DBConfig.Migrate.Dir)
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}

	// modify for Down
	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Warn("Migration Error", zap.Error(err))
		} else {
			logger.Fatal("", zap.Error(err))
		}
	}
	m.Close()

	logger.Info("Done Migration")
}
