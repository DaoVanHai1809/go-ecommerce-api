package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	var viper = viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // name file config
	viper.SetConfigType("yaml")

	// read configuration
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Security jwt key::", viper.GetString("security.jwt.key"))

	// get configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}