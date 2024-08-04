package promexporter

import "github.com/jtway/tilt-proxy/pkg/tilt"

type PromExporter struct {
}

func NewPromExporter() (*PromExporter, error) {

	return nil, nil
}

func (e *PromExporter) Start() error {

	return nil
}

func (e *PromExporter) Stop() {

}

func (e *PromExporter) TiltDataEvent(tilt tilt.Tilt) {

}
