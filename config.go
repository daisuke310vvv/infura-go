package infura

type Config struct {
	ApiKey  string
	Network Network
}

func NewConfig(apiKey string, network Network) *Config {
	return &Config{
		ApiKey:  apiKey,
		Network: network,
	}
}