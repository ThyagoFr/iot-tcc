package database

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port int
	Database string
	User string
	Password string
}

func (conf *Config) DNS() string {
	if conf.User == "" || conf.Password == "" {
		return fmt.Sprintf("mongodb://%s:%d/%s", conf.Host, conf.Port, conf.Database)
	}
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
}

func ConfigFromEnv(conf *Config) {
	conf.Host = viper.GetString("MONGODB_HOST")
	conf.Port = viper.GetInt("MONGODB_PORT")
	conf.Database = viper.GetString("MONGODB_DB")
	conf.User = viper.GetString("MONGODB_USER")
	conf.Password = viper.GetString("MONGODB_PASSWORD")
}
