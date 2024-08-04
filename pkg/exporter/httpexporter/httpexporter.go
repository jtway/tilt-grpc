package httpexporter

import "github.com/jtway/tilt-proxy/pkg/tilt"

type HttpExporter struct {
	
}

func NewHttpExporter(config *Config) (*HttpExporter, error) {

	return nil, nil
}

func (e *HttpExporter) Start() error {

	return nil
}

func (e *HttpExporter) Stop() {

}

func (e *HttpExporter) TiltDataEvent(tilt tilt.Tilt) {

}
