package main

import (
	"github.com/lukemorcom/bodmqtt/config"
	"github.com/lukemorcom/bodmqtt/service"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")

	if err != nil {
		log.Printf("Oh no something went wrong: %v\n", err)
	}

	service := service.NewService(cfg)

	service.Run()
}
