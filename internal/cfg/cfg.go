package cfg

import (
	"log"

	"github.com/spf13/viper"
)

//Cfg is the global config
var Cfg Config

//Init initiates the Cfg
func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("Unable to decode into Config, %s", err)
	}
}

//Config is the config
type Config struct {
	App   AppConfig   `mapstructure:"app"`
	Redis RedisConfig `mapstructure:"redis"`
	Mysql MysqlConfig `mapstructure:"mysql"`
}

//AppConfig is the app config
type AppConfig struct {
	ENV        string `mapstructure:"env"`
	ServerPort string `mapstructure:"server_port"`
}

//RedisConfig is the redis config
type RedisConfig struct {
	Addr    string             `mapstructure:"addr"`
	Cluster RedisClusterConfig `mapstructure:"cluster"`
	User    string             `mapstructure:"user"`
	Pass    string             `mapstructure:"password"`
	DB      int                `mapstructure:"db"`
}

//RedisClusterConfig is the redis cluster config
type RedisClusterConfig struct {
	Addrs string `mapstructure:"addrs"`
}

//MysqlConfig is the mysql config
type MysqlConfig struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"password"`
	Schema string `mapstructure:"schema"`
}
