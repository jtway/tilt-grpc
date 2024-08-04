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
var ErrIvalidUUID = errors.New("Invalid iBeacon UUID")

// IBeacon data
type IBeacon struct {
	UUID    string
	Major   uint16
	Minor   uint16
	TxPower int8
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
		// @TODO(jtway) Figure out how to turn this in to distance.
		TxPower: int8(data[24]),
	}
	return
}

func (i *IBeacon) EncodeBLEEvent() ([]byte, error) {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, iBeaconType)
	uuid, err := hex.DecodeString(i.UUID)
	if err != nil {
		return nil, err
	}
	if len(uuid) != 16 {
		return nil, ErrIvalidUUID
	}
	data = append(data, uuid...)
	data = binary.BigEndian.AppendUint16(data, i.Major)
	data = binary.BigEndian.AppendUint16(data, i.Minor)
	data = append(data, byte(i.TxPower))

	return data, nil
}
