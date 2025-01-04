package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql)
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "success"))
	InitMySql()
	InitRedis()

	var r = InitRouter()
	r.Run(":8002")
}