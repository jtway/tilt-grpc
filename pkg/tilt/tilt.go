/*
 * This iBeacon portion is largely taken and inspired by https://github.com/alexhowarth/go-tilt
 */
// Package tilt provides an interface to Tilt Bluetooth devices
package tilt

import (
	"encoding/hex"
	"log"
	"math"
	"time"

	"github.com/jtway/tilt-proxy/pkg/ibeacon"
	"github.com/pkg/errors"
)

// tiltIdentifier iBeacon identifier portion (4c000215) as well as Tilt specific uuid preamble (a495)
const tiltIdentifier = "4c000215a495"

// ErrNotTilt - the BLE device does not match anything in tiltType
var ErrNotTilt = errors.New("Not a Tilt iBeacon")

// TODO(jtway) Make colors an enum

var tiltType = map[string]string{
	"a495bb10c5b14b44b5121370f02d74de": "Red",
	"a495bb20c5b14b44b5121370f02d74de": "Green",
	"a495bb30c5b14b44b5121370f02d74de": "Black",
	"a495bb40c5b14b44b5121370f02d74de": "Purple",
	"a495bb50c5b14b44b5121370f02d74de": "Orange",
	"a495bb60c5b14b44b5121370f02d74de": "Blue",
	"a495bb70c5b14b44b5121370f02d74de": "Yellow",
	"a495bb80c5b14b44b5121370f02d74de": "Pink",
}

// Tilt struct
type Tilt struct {
	Color           string    `json:"color"`
	Temp            uint16    `json:"temp"`
	SpecificGravity float64   `json:"specific_gravity"`
	ReadingTime     time.Time `json:"reading_time"`
}

// NewTilt returns a Tilt from an iBeacon
func NewTilt(b *ibeacon.IBeacon) (*Tilt, error) {
	if col, ok := tiltType[b.UUID]; ok {
		return &Tilt{
			Color:           col,
			Temp:            b.Major,
			SpecificGravity: float64(b.Minor) / 1000,
			ReadingTime:     time.Now(),
		}, nil
	}
	return nil, ErrNotTilt
}

// IsTilt tests if the data is from a Tilt
func IsTilt(d []byte) bool {
	if len(d) >= 25 && hex.EncodeToString(d)[0:12] == tiltIdentifier {
		return true
	}
	return false
}

func (t *Tilt) GetTempCelsius() float64 {
	return math.Round(float64(t.Temp-32)/1.8*100) / 100
}

func (t *Tilt) GetTempFahrenheit() uint16 {
	return t.Temp
}

func (t *Tilt) GetSpecificGravity() float64 {
	return t.SpecificGravity
}

func (t *Tilt) GetColor() string {
	return t.Color
}

func (t *Tilt) Print() {
	log.Printf("Tilt: %v", t.GetColor())
	log.Printf("Fahrenheit: %v\n", t.GetTempFahrenheit())
	log.Printf("Specific Gravity: %v\n", t.GetSpecificGravity())
	log.Printf("Celsius: %v\n", t.GetTempCelsius())
}
