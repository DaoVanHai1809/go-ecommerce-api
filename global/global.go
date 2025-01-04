package global

import (
	"go-ecommerce-backend-api/packages/logger"
	"go-ecommerce-backend-api/packages/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)

/*
Config
Redis
Mysql

*/