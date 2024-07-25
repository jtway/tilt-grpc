package main

import (
	"context"

	"github.com/JuulLabs-OSS/ble"
	"github.com/jtway/tilt-proxy/pkg/tilt"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	tiltScanner, err := tilt.NewScanner()
	if err != nil {
		logger.Errorf("Unable to create new scanner, %s", err.Error())
	}

	// This will give use the background context with signal handling.
	// @TODO(jtway) Is it possible to maybe just pass in a context. I would like
	// to keep the ble library out of direct use here. Also would be nice to differ
	// the cancel.
	ctx := ble.WithSigHandler(context.WithCancel(context.Background()))

	tiltChan := make(chan tilt.Tilt)

	tiltEventHandler := func(tilt *tilt.Tilt, err error) {
		if err != nil {
			logger.Errorf("error with getting tilt events. %s\n", err.Error())
			return
		}
		if tilt == nil {
			logger.Errorf("error getting tilt readings, but not error returned\n")
			return
		}
		tiltChan <- *tilt
	}

	// @TODO(jtway) Think of how I can go about just doing this in the background.
	// I feel like I'm making more contexts than I should need to.
	go func() {
		tiltScanner.Scan(ctx, tiltEventHandler)
	}()

ScanFor:
	for {
		select {
		case <-ctx.Done():
			break ScanFor
		case tilt := <-tiltChan:
			tilt.Print()
		}
	}

	logger.Infof("Finished Scanning")

}
