package global

import (
	"go-ecommerce-backend-api/packages/logger"
	"go-ecommerce-backend-api/packages/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
)

/*
Config
Redis
Mysql

*/