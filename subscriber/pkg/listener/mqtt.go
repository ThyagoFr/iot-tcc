package listener

import (
	"context"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thyago/tcc/subscriber/internal/device"
	"log"
	"time"
)

type Subscriber struct {
	client mqtt.Client
	topic string
	repository device.Repository
}

func NewSubscriber(client mqtt.Client, topic string, repository device.Repository) Listener {
	return &Subscriber{client: client, topic: topic, repository: repository}
}

func (s *Subscriber) Listen() {
	s.client.Subscribe(s.topic, 0, s.callback)
}

func (s *Subscriber) callback(_ mqtt.Client, message mqtt.Message) {
	data := &device.DeviceData{}
	err := json.Unmarshal(message.Payload(),data)
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

