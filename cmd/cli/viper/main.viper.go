package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host string `mapstructure:"host"`
		DbName string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
	Security struct {
		Jwt struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	var viper = viper.New()
	viper.AddConfigPath("./config/")	// path to config
	viper.SetConfigName("local") // name file config
	viper.SetConfigType("yaml")

	// read configuration
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Security jwt key::", viper.GetString("security.jwt.key"))

	// get configure structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Println("config port::", config.Server.Port)
	for count, db := range config.Databases {
		fmt.Printf("Stt: %d, User: %s, Password: %s, Host: %s, DbName: %s \n", count + 1, db.User, db.Password, db.Host, db.DbName)
	}
}