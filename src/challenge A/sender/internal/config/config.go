package config

// AppConfig holds the application config
type AppConfig struct {
	InProduction bool
	BrokerDSN    string
	PortServer   string
}
