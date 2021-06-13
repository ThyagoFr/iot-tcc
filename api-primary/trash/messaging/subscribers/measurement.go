// package subscribers

// import (
// 	"context"
// 	"encoding/json"
// 	mqtt "github.com/eclipse/paho.mqtt.golang"
// 	"github.com/spf13/viper"
// 	model2 "github.com/thyago/tcc/api-service/domain/model"
// 	"github.com/thyago/tcc/api-service/domain/repository"
// 	"log"
// 	"time"
// )

// type MeasurementSubscriber struct {
// 	client     mqtt.Client
// 	topic      string
// 	repository repository.MeasurementRepository
// }

// func NewMeasurementSubscriber(client mqtt.Client, repository repository.MeasurementRepository) Subscriber {
// 	sub := &MeasurementSubscriber{client: client, repository: repository}
// 	topic := viper.GetString("MQTT_TOPIC_MEASUREMENTS")
// 	sub.topic = topic
// 	return sub
// }

// func (s *MeasurementSubscriber) Subscribe() {
// 	log.Printf("subscribing topic %s \n", s.topic)
// 	s.client.Subscribe(s.topic, 0, s.callback)
// }

// func (s *MeasurementSubscriber) callback(_ mqtt.Client, message mqtt.Message) {
// 	measure := &model2.Measurement{}
// 	err := json.Unmarshal(message.Payload(), measure)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	err = s.repository.Insert(ctx, measure)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
