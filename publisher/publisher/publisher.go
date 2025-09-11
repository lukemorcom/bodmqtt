package publisher

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type Publisher interface {
	Publish(message, topic string) error
}

type MQTTPublisher struct {
	Client mqtt.Client
}

func NewMQTTPublisher(broker string) *MQTTPublisher {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("bodmqtt-go-pub")

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Couldn't connect to broker: %v\n", token.Error())
	}

	log.Printf("Connected to broker on: %s", broker)

	return &MQTTPublisher{Client: client}
}

func (p *MQTTPublisher) Publish(message, topic string) error {
	token := p.Client.Publish(topic, 0, false, message)

	log.Printf("Published %s to %s", message, topic)

	token.Wait()

	return token.Error()
}
