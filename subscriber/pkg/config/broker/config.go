package broker

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port int
	Password string
	User string
	ClientID string
	Topic string
}

func (conf *Config) DNS() string {
	return fmt.Sprintf("tcp://%s:%d",conf.Host,conf.Port)
}

func ConfigFromEnv(conf *Config) {
	conf.Host = viper.GetString("BROKER_MQTT_HOST")
	conf.Port = viper.GetInt("BROKER_HOST_PORT")
	conf.Password = viper.GetString("BROKER_MQTT_PASSWORD")
	conf.User = viper.GetString("BROKER_MQTT_USER")
	conf.ClientID = viper.GetString("MQTT_CLIENT_ID")
	conf.Topic = viper.GetString("MQTT_TOPIC")
}