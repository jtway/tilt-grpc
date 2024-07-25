package exporter

import "github.com/jtway/tilt-proxy/pkg/tilt"

type Exporter interface {
	Start() error
	Stop()

	EmitTiltData(tilt tilt.Tilt)
}
