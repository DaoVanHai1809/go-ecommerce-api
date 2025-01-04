package setting

type Config struct {
	Mysql MySqlSetting `mapstructure:"mysql"`
}

type MySqlSetting struct {
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	DbName          string `mapstructure:"dbname"`
	Port            int    `mapstructure:"port"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}