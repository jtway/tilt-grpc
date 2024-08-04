package httpexporter

// Going to keep this overly simple for the moment. Eventually we could add
// TLS support, configuration of the path/route, formatting, timeouts, etc.
// However, most of this is largely over kill.
type Config struct {
	Port    int  `mapstructure:"port"`
	Enabled bool `mapstructure:"enabled"`
}
