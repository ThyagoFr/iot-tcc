package main

import (
	"githu"
	"github.com/spf13/viper"
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
	fmt.Println("in construction...")

	// err := Initialize()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// dbConf := &database.Config{}
	// database.ConfigFromEnv(dbConf)
	// db, err := database.NewMongoDBClient(dbConf)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// deviceRepository := device.NewMongoDB(db)
	// measurementRepository := measurement.NewMongoDB(db)

	// brokerConf := &broker.Config{}
	// broker.ConfigFromEnv(brokerConf)
	// mqttClient, err := broker.NewMQTTClient(brokerConf)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// deviceListener := mqtt.NewDeviceSubscriber(mqttClient,deviceRepository)
	// measurementListener := measurement.NewSubscriber(mqttClient, measurementRepository)

	// go deviceListener.Listen()
	// go measurementListener.Listen()

	// apiService := service.NewService(deviceRepository, measurementRepository)
	// server := http.NewAPI(apiService)

	// err = server.ListenAndServe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
