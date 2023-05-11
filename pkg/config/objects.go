package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
	Config Config
}

type Config struct {
	ServiceDatabase MainDatabaseConfiguration `mapstructure:",squash"`
	Server          ServerConfiguration       `mapstructure:",squash"`
}

type ServerConfiguration struct {
	Debug bool   `mapstructure:"DEBUG"`
	Port  string `mapstructure:"SERVER_PORT"`
}

type MainDatabaseConfiguration struct {
	Name     string `mapstructure:"MAIN_DB_NAME"`
	Username string `mapstructure:"MAIN_DB_USER"`
	Password string `mapstructure:"MAIN_DB_PASSWORD"`
	Host     string `mapstructure:"MAIN_DB_HOST"`
	Port     string `mapstructure:"MAIN_DB_PORT"`
	LogMode  bool   `mapstructure:"MAIN_DB_LOG_MODE"`
	SSLMode  string `mapstructure:"MAIN_DB_SSL_MODE"`
}
