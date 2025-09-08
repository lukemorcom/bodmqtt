package main

import (
	"fmt"
	//	"github.com/lukemorcom/bodmqtt/publisher"
	"github.com/lukemorcom/bodmqtt/config"
)

func main() {
	//	pub := publisher.NewMQTTPublisher("localhost:1883")

	//err := pub.Publish("Hello from main.go", "example/topic")

	//if err != nil {
	//	fmt.Println("Uwu there was an ewwor")
	//}

	cfg, err := config.LoadConfig("config.yml")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}
