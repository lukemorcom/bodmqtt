package strategies

import (
	"github.com/lukemorcom/bodmqtt/config"
	"net"
	"time"
)

func (p PingStrategy) Execute(event config.Event) (string, error) {
	conn, err := net.DialTimeout("tcp", event.API.URL, 3*time.Second)

	if err != nil {
		return "down", err
	}

	defer conn.Close()

	return "up", nil
}
