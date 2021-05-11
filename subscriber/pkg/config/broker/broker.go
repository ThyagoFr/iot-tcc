package broker

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func NewMQTTClient(conf *Config) (mqtt.Client,error){
	options := mqtt.NewClientOptions()
	options.AddBroker(conf.DNS())
	options.SetUsername(conf.User)
	options.SetPassword(conf.Password)
	options.SetClientID(conf.ClientID)
	client := mqtt.NewClient(options)
	testConnection := client.Connect()
	for !testConnection.WaitTimeout(5 * time.Second) {
	}
	if err := testConnection.Error(); err != nil {
		return nil,err
	}
	return client, nil
}