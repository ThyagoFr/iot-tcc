package main

import (
	"github.com/spf13/viper"
	"github.com/thyago/tcc/subscriber/internal/device"
	"github.com/thyago/tcc/subscriber/internal/measurement"
	"github.com/thyago/tcc/subscriber/pkg/api"
	"github.com/thyago/tcc/subscriber/pkg/api/service"
	"github.com/thyago/tcc/subscriber/pkg/config/broker"
	"github.com/thyago/tcc/subscriber/pkg/config/database"
	deviceSub "github.com/thyago/tcc/subscriber/pkg/listener/device"
	measurementSub "github.com/thyago/tcc/subscriber/pkg/listener/measurement"
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

	deviceListener := deviceSub.NewSubscriber(mqttClient, deviceRepository)
	measurementListener := measurementSub.NewSubscriber(mqttClient, measurementRepository)

	go deviceListener.Listen()
	go measurementListener.Listen()

	apiService := service.NewService(deviceRepository,measurementRepository)
	server := api.NewAPI(apiService)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
