package initialize

import (
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/packages/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}