package messaging

import (
  "fmt"
  "github.com/thyago/tcc/api-service/infra/env"
  	mqtt "github.com/eclipse/paho.mqtt.golang"
    "github.com/google/uuid"
    "time"
)

// Config holds information to create a mqtt connection
type Config struct {
	Host     string
	Port     int
	Password string
	User     string
	ClientID string
}

// DNS create a URI connection string
func (conf *Config) DNS() string {
	return fmt.Sprintf("tcp://%s:%d", conf.Host, conf.Port)
}

// ConfigFromEnv configure using environmnet variables
func ConfigFromEnv(conf *Config) {
	conf.Host = env.GetOrDefaultString("BROKER_MQTT_HOST","localhost")
	conf.Port = env.GetOrDefaultInt("BROKER_MQTT_PORT",1883)
	conf.Password = env.GetOrDefaultString("BROKER_MQTT_PASSWORD","test")
	conf.User = env.GetOrDefaultString("BROKER_MQTT_USER","test")
	conf.ClientID = uuid.New().String()
}

// NewMQTTClient creates a new mqtt client
func NewMQTTClient(conf *Config) (mqtt.Client, error) {
	options := mqtt.NewClientOptions()
	options.AddBroker(conf.DNS())
	options.SetUsername(conf.User)
	options.SetPassword(conf.Password)
	options.SetClientID(conf.ClientID)
	options.AutoReconnect = true
	options.ConnectRetry = true
	client := mqtt.NewClient(options)
	testConnection := client.Connect()
	for !testConnection.WaitTimeout(5 * time.Second) {
	}
	if err := testConnection.Error(); err != nil {
		return nil, err
	}
	return client, nil
}
