package database

import (
	"fmt"

	log "github.com/ishangoyal13/blogPost/pkg/log"

	"github.com/ishangoyal13/blogPost/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var (
	db  *gorm.DB
	err error
)

// Connection create database connection
func GetDB(cfg config.Config) (*gorm.DB, error) {
	mainDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.ServiceDatabase.Host, cfg.ServiceDatabase.Username, cfg.ServiceDatabase.Password, cfg.ServiceDatabase.Name, cfg.ServiceDatabase.Port, cfg.ServiceDatabase.SSLMode,
	)
	replicaDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.ServiceDatabase.Host, cfg.ServiceDatabase.Username, cfg.ServiceDatabase.Password, cfg.ServiceDatabase.Name, cfg.ServiceDatabase.Port, cfg.ServiceDatabase.SSLMode,
	)

	logMode := cfg.ServiceDatabase.LogMode
	debug := cfg.Server.Debug

	loglevel := logger.Silent
	if logMode {
		loglevel = logger.Info
	}

	db, err = gorm.Open(postgres.Open(mainDSN), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})
	if !debug {
		db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{
				postgres.Open(replicaDSN),
			},
			Policy: dbresolver.RandomPolicy{},
		}))
	}

	if err != nil {
		log.Fatal("Db connection error")
		return nil, err
	}
	return db, nil
}
