package global

import (
	"go-ecommerce-backend-api/packages/logger"
	"go-ecommerce-backend-api/packages/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
	Rdb *redis.Client
)

/*
Config
Redis
Mysql

*/