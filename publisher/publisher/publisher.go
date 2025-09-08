package publisher

import "fmt"

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
	fmt.Printf("todo publish here")

	return nil
}
