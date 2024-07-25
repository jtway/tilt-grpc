/*
 * This iBeacon portion is largely taken and inspired by https://github.com/alexhowarth/go-tilt
 */
package ibeacon

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
)

// iBeaconType identifies an iBeacon
const iBeaconType uint32 = 0x4c000215

// ErrNotBeacon - the BLE device is not an iBeacon
var ErrNotBeacon = errors.New("Not an iBeacon")

// IBeacon data
type IBeacon struct {
	UUID  string
	Major uint16
	Minor uint16
}

// IsIBeacon to determine if data is an iBeacon
func IsIBeacon(data []byte) bool {
	if len(data) < 25 || binary.BigEndian.Uint32(data) != iBeaconType {
		return false
	}
	return true
}

// NewIBeacon creates an iBeacon from a valid BLE event
func NewIBeacon(data []byte) (b *IBeacon, err error) {

	if !IsIBeacon(data) {
		err = ErrNotBeacon
		return
	}
	b = &IBeacon{
		UUID:  hex.EncodeToString(data[4:20]),
		Major: binary.BigEndian.Uint16(data[20:22]),
		Minor: binary.BigEndian.Uint16(data[22:24]),
	}

	return
}
