package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	// *core
	HTTP_PORT       string = "1212"
	REDIS_HOST      string = "127.0.0.1"
	REDIS_PORT      string = "6379"
	REDIS_CLUSTER   bool   = false
	REDIS_POOL_SIZE int    = 10
	DB_HOST         string = "127.0.0.1"
	DB_PORT         string = "5432"
	DB_USERNAME     string = "postgres"
	DB_PASSWORD     string = "postgres"
	DB_NAME         string = "shrink-link"

	// *log
	LOG_LEVEL string = "info"
)

func LoadConf(confPath ...string) error {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("shrink-link")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath("config")
	viper.AddConfigPath("./etc/shrink-link")
	viper.AddConfigPath("$HOME/.shrink-link")
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

	// log
	LOG_LEVEL = viper.GetString("log_level")

	return nil
}
