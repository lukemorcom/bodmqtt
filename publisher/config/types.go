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
	Name     string `yaml:"name"`
	APIname  string `yaml:"api"`
	Topic    string `yaml:"topic"`
	Interval int    `yaml:"interval"`
	API      API    `yaml:"-"`
}

type Config struct {
	APIs   []API
	Events []Event
}
