package publisher

import "log"

type Publisher interface {
	Publish(message, topic string) error
}

type MQTTPublisher struct {
	Broker string
}

func NewMQTTPublisher(broker string) *MQTTPublisher {
	return &MQTTPublisher{Broker: broker}
}

func (p *MQTTPublisher) Publish(message, topic string) error {
	log.Printf("When implemented, this will publish %s to topic %s", message, topic)

	// todo implement

	return nil
}
