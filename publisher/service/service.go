package service

import (
	"fmt"
	"github.com/lukemorcom/bodmqtt/config"
	"github.com/lukemorcom/bodmqtt/publisher"
	"github.com/lukemorcom/bodmqtt/strategies"
	"sync"
)

type Service struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Run() {
	var wg sync.WaitGroup

	pub := publisher.NewMQTTPublisher("localhost:1883")

	for _, event := range s.cfg.Events {
		wg.Add(1)
		go workEvent(event, pub, &wg)
	}

	wg.Wait()
}

func workEvent(event config.Event, pub *publisher.MQTTPublisher, wg *sync.WaitGroup) {
	var strat strategies.Strategy

	switch event.API.Strategy {
	case config.StrategyPing:
		strat = strategies.PingStrategy{}
	default:
		fmt.Printf("Unexpected strategy name: %s\n", event.API.Strategy)
		return
	}

	status, err := strat.Execute(event)

	if err != nil {
		fmt.Printf("Error executing event %s: %v\n", event.Name, err)
		return
	}

	err = pub.Publish(status, event.Topic)

	if err != nil {
		fmt.Printf("Error publishing event %s on topic %s: %v\n", event.Name, event.Topic, err)
	}

	wg.Done()
}
