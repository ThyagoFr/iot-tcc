package device

import (
	"context"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"github.com/thyago/tcc/subscriber/internal/device"
	"github.com/thyago/tcc/subscriber/pkg/listener"
	"log"
	"time"
)

type Subscriber struct {
	client     mqtt.Client
	topic      string
	repository device.Repository
}

func NewSubscriber(client mqtt.Client, repository device.Repository) listener.Listener {
	sub := &Subscriber{client: client,repository: repository}
	topic := viper.GetString("MQTT_TOPIC_DEVICES")
	sub.topic = topic
	return sub
}

func (s *Subscriber) Listen() {
	log.Printf("subscribing topic %s \n", s.topic)
	s.client.Subscribe(s.topic, 0, s.callback)
}

func (s *Subscriber) callback(_ mqtt.Client, message mqtt.Message) {
	data := &device.Device{}
	err := json.Unmarshal(message.Payload(), data)
	if err != nil {
		log.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = s.repository.Insert(ctx, data)
	if err != nil {
		log.Println(err)
	}
}
