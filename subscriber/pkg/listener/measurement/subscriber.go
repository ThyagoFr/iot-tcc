package measurement

import (
	"context"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"github.com/thyago/tcc/subscriber/internal/measurement"
	"github.com/thyago/tcc/subscriber/pkg/listener"
	"log"
	"time"
)

type Subscriber struct {
	client     mqtt.Client
	topic      string
	repository measurement.Repository
}

func NewSubscriber(client mqtt.Client,repository measurement.Repository) listener.Listener {
	sub :=  &Subscriber{client: client, repository: repository}
	topic := viper.GetString("MQTT_TOPIC_MEASUREMENTS")
	sub.topic = topic
	return sub
}

func (s *Subscriber) Listen() {
	log.Printf("subscribing topic %s \n", s.topic)
	s.client.Subscribe(s.topic, 0, s.callback)
}

func (s *Subscriber) callback(_ mqtt.Client, message mqtt.Message) {
	measure := &measurement.Measurement{}
	err := json.Unmarshal(message.Payload(), measure)
	if err != nil {
		log.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = s.repository.Insert(ctx, measure)
	if err != nil {
		log.Println(err)
	}
}
