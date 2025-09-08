package config

type APIStrategy string

const (
	StrategyPing  APIStrategy = "ping"
	StrategyFetch APIStrategy = "fetch"
)

type API struct {
	Name     string      `yaml:"name"`
	URL      string      `yaml:"url"`
	Strategy APIStrategy `yaml:"strategy"`
}

type Event struct {
	Name  string `yaml:"name"`
	API   string `yaml:"api"`
	Topic string `yaml:"topic"`
}

type Config struct {
	APIs   []API
	Events []Event
}
