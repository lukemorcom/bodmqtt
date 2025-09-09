package main

import (
	"github.com/lukemorcom/bodmqtt/config"
	"github.com/lukemorcom/bodmqtt/service"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")

	if err != nil {
		panic(err)
	}

	service := service.NewService(cfg)

	service.Run()
}
