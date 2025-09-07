package main

import (
	"fmt"
	"github.com/lukemorcom/bodmqtt/publisher"
)

func main() {
	pub := publisher.NewMQTTPublisher("localhost:1883")

	err := pub.Publish("Hello from main.go", "example/topic")

	if err != nil {
		fmt.Println("Uwu there was an ewwor")
	}
}
