package tilt

import (
	"context"

	"github.com/JuulLabs-OSS/ble"
	"github.com/JuulLabs-OSS/ble/examples/lib/dev"
	"github.com/jtway/tilt-proxy/pkg/ibeacon"
)

// Scanner for Tilt devices
type Scanner struct {
	device ble.Device
}

type TiltEventHandler func(tilt *Tilt, err error)

// NewScanner returns a Scanner
func NewScanner() (*Scanner, error) {
	device, err := dev.NewDevice("go-ble")
	if err != nil {
		return nil, err
	}
	ble.SetDefaultDevice(device)
	return &Scanner{
		device: device,
	}, nil
}

// Scans for Tilt devices called the passed in TiltEventHandler for each
// event.
func (s *Scanner) Scan(ctx context.Context, tiltEvent TiltEventHandler) error {

	ctx2, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case <-ctx.Done():
			cancel()
		case <-ctx2.Done():
		}
	}()

	advHandler := func(a ble.Advertisement) {
		// create iBeacon
		b, err := ibeacon.NewIBeacon(a.ManufacturerData())
		if err != nil {
			tiltEvent(nil, err)
			return
		}

		// create Tilt from iBeacon
		tilt, err := NewTilt(b)
		if err != nil {
			tiltEvent(nil, err)
			return
		}

		tiltEvent(tilt, nil)
	}

	advFilter := func(a ble.Advertisement) bool {
		return IsTilt(a.ManufacturerData())
	}

	// Scan until canceled, allowing duplicates to be found
	return ble.Scan(ctx, true, advHandler, advFilter)
}
