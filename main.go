package main

import (
	"github.com/ishangoyal13/blogPost/models"
	logger "github.com/ishangoyal13/blogPost/pkg/log"
	"github.com/ishangoyal13/blogPost/routers"

	"github.com/ishangoyal13/blogPost/pkg/config"
	"github.com/ishangoyal13/blogPost/pkg/database"
)

func main() {

	// Get Config from env
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("config GetConfig() error: %s", err)
	}

	// Get DB connection
	db, err := database.GetDB(cfg)
	if err != nil {
		logger.Fatalf("database GetDB() error: %s", err)
	}

	app := config.App{
		Config: cfg,
		DB:     db,
	}

	err = db.AutoMigrate(models.GetMigrationModels()...)
	if err != nil {
		logger.Fatalf("database migration error: %s", err)
	}

	routers.SetupAndRunServer(&app)
}
