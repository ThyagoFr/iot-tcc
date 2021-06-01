package main

import (
	"github.com/spf13/viper"
	"github.com/thyago/tcc/api-service/application/service"
	broker "github.com/thyago/tcc/api-service/config/broker"
	database "github.com/thyago/tcc/api-service/config/database"
	"github.com/thyago/tcc/api-service/infra/http"
	"github.com/thyago/tcc/api-service/infra/mqtt"
	"github.com/thyago/tcc/api-service/infra/mqtt"
	"github.com/thyago/tcc/api-service/internal/device"
	"github.com/thyago/tcc/api-service/internal/measurement"
	"log"
)

// Initialize loads the application configuration
func Initialize() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./cmd")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.AutomaticEnv()
	return nil
}

func main() {

	err := Initialize()
	if err != nil {
		log.Fatal(err)
	}

	dbConf := &database.Config{}
	database.ConfigFromEnv(dbConf)
	db, err := database.NewMongoDBClient(dbConf)
	if err != nil {
		log.Fatal(err)
	}

	deviceRepository := device.NewMongoDB(db)
	measurementRepository := measurement.NewMongoDB(db)

	brokerConf := &broker.Config{}
	broker.ConfigFromEnv(brokerConf)
	mqttClient, err := broker.NewMQTTClient(brokerConf)
	if err != nil {
		log.Fatal(err)
	}

	deviceListener := mqtt.NewDeviceSubscriber(mqttClient,deviceRepository)
	measurementListener := measurement.NewSubscriber(mqttClient, measurementRepository)

	go deviceListener.Listen()
	go measurementListener.Listen()

	apiService := service.NewService(deviceRepository, measurementRepository)
	server := http.NewAPI(apiService)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
