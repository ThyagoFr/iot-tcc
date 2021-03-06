// package subscribers

// import (
// 	"context"
// 	"encoding/json"
// 	mqtt "github.com/eclipse/paho.mqtt.golang"
// 	"github.com/spf13/viper"
// 	"github.com/thyago/tcc/api-service/domain/model"
// 	"github.com/thyago/tcc/api-service/domain/repository"
// 	"log"
// 	"time"
// )

// type DeviceSubscriber struct {
// 	client     mqtt.Client
// 	topic      string
// 	repository repository.DeviceRepository
// }

// func NewDeviceSubscriber(client mqtt.Client, repository repository.DeviceRepository) Subscriber {
// 	sub := &DeviceSubscriber{client: client, repository: repository}
// 	topic := viper.GetString("MQTT_TOPIC_DEVICES")
// 	sub.topic = topic
// 	return sub
// }

// func (s *DeviceSubscriber) Subscribe() {
// 	log.Printf("subscribing topic %s \n", s.topic)
// 	s.client.Subscribe(s.topic, 0, s.callback)
// }

// func (s *DeviceSubscriber) callback(_ mqtt.Client, message mqtt.Message) {
// 	data := &model.Device{}
// 	err := json.Unmarshal(message.Payload(), data)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	err = s.repository.Insert(ctx, data)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
