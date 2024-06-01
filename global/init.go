package global

import (
	"admin/configs"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *configs.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
