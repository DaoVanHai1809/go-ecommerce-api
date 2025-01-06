package setting

type Config struct {
	Mysql  MySqlSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type MySqlSetting struct {
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	DbName          string `mapstructure:"dbname"`
	Port            int    `mapstructure:"port"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Max_size      int    `mapstructure:"max_size"`
	Compress      bool   `mapstructure:"compress"`
}