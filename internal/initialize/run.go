package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql)
	InitLogger()
	InitMySql()
	InitRedis()

	var r = InitRouter()
	r.Run(":8002")
}