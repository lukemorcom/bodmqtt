package service

import (
	"github.com/lukemorcom/bodmqtt/config"
	"github.com/lukemorcom/bodmqtt/publisher"
	"github.com/lukemorcom/bodmqtt/strategies"
	"log"
	"time"
)

type Service struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Run() {
	pub := publisher.NewMQTTPublisher("localhost:1883")

	for _, event := range s.cfg.Events {
		go workEvent(event, pub)
	}

	select {}
}

func workEvent(event config.Event, pub *publisher.MQTTPublisher) {
	for {
		var strat strategies.Strategy

		switch event.API.Strategy {
		case config.StrategyPing:
			strat = strategies.PingStrategy{}
		default:
			log.Printf("Unexpected strategy name: %s\n", event.API.Strategy)
			return
		}

		status, err := strat.Execute(event)

		if err != nil {
			log.Printf("Error executing event %s: %v\n", event.Name, err)
			return
		}

		err = pub.Publish(status, event.Topic)

		if err != nil {
			log.Printf("Error publishing event %s on topic %s: %v\n", event.Name, event.Topic, err)
		}

		time.Sleep(time.Duration(event.Interval) * time.Second)
	}
}
