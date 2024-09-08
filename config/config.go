package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var defaultConf = []byte(`
	http_port: "1721"
	redis_host: "127.0.0.1"
	redis_port: "6379"
	redis_cluster: true
	db_host: "127.0.0.1"
	db_port: "5432"
	trace_id: ""
	log_level: "info"
`)

var (
	// *core
	HTTP_PORT     string = "1212"
	REDIS_HOST    string = "127.0.0.1"
	REDIS_PORT    string = "6379"
	REDIS_CLUSTER bool   = false
	DB_HOST       string = "127.0.0.1"
	DB_PORT       string = "5432"

	// *log
	LOG_LEVEL string = "info"
)

func LoadConf(confPath ...string) error {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("urlshortner")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath("config")
	viper.AddConfigPath("./etc/urlshortner")
	viper.AddConfigPath("$HOME/.urlshortner")
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
	DB_HOST = viper.GetString("db_host")
	DB_PORT = viper.GetString("db_port")

	// log
	LOG_LEVEL = viper.GetString("log_level")

	return nil
}
