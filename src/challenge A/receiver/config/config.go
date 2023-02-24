package config

// AppConfig holds the application config
type AppConfig struct {
	InProduction bool
	BrokerDSN    string
	MySQLDSN     string
	PortServer   string
}
