package infura

type Config struct {
	ProjectID string
	Network   Network
}

func NewConfig(projectID string, network Network) *Config {
	return &Config{
		ProjectID: projectID,
		Network:   network,
	}
}
