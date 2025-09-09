package strategies

import "github.com/lukemorcom/bodmqtt/config"

type Strategy interface {
	Execute(event config.Event) (string, error)
}

type PingStrategy struct{}
