package promexporter

// May want to allow restricting to listening on an interface, timeouts, etc.
type Config struct {
	Port    int  `mapstructure:"port"`
	Enabled bool `mapstructure:"enabled"`
}
