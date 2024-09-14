package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	// *core
	HTTP_PORT       = "1212"
	REDIS_HOST      = "127.0.0.1"
	REDIS_PORT      = "6379"
	REDIS_CLUSTER   = false
	REDIS_POOL_SIZE = 10
	DB_HOST         = "127.0.0.1"
	DB_PORT         = "5432"
	DB_USERNAME     = "postgres"
	DB_PASSWORD     = "postgres"
	DB_NAME         = "shrinklink"
	DOMAIN          = "https://shrinklink.com"

	// *log
	LOG_LEVEL = "info"
)

func LoadConf(confPath ...string) error {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("shrinklink")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath("config")
	viper.AddConfigPath("./etc/shrinklink")
	viper.AddConfigPath("$HOME/.shrinklink")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		return err
	}

	// core
	HTTP_PORT = viper.GetString("http_port")
	REDIS_HOST = viper.GetString("redis_host")
	REDIS_PORT = viper.GetString("redis_port")
	REDIS_CLUSTER = viper.GetBool("redis_cluster")
	REDIS_POOL_SIZE = viper.GetInt("redis_pool_size")
	DB_HOST = viper.GetString("db_host")
	DB_PORT = viper.GetString("db_port")
	DB_USERNAME = viper.GetString("db_username")
	DB_PASSWORD = viper.GetString("db_password")
	DB_NAME = viper.GetString("db_name")
	DOMAIN = viper.GetString("domain")

	// log
	LOG_LEVEL = viper.GetString("log_level")

	return nil
}
