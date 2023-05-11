package controllers

import (
	"github.com/ishangoyal13/blogPost/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BaseController struct {
	DB     *gorm.DB
	Config config.Config
	Log    *logrus.Logger
}
