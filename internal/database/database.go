package database

import (
	"chaos-go/internal/config"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init(cfg *config.Config, logger *zap.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.DBConfig.Path), &gorm.Config{})
	if err != nil {
		logger.Error("db err: (Init) ", zap.Error(err))
	}
	DB = db

	if cfg.DBConfig.Migrate.Enable {
		RunMigration(cfg, logger)
	}

	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
