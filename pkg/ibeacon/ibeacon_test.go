/*
 * This iBeacon portion is largely taken and inspired by https://github.com/alexhowarth/go-tilt
 */
package ibeacon_test

import (
	"errors"
	"testing"

	"github.com/jtway/tilt-proxy/pkg/ibeacon"
)

func TestNewIBeacon(t *testing.T) {
	tt := []struct {
		name  string
		data  []byte
		err   error
		uuid  string
		major uint16
		minor uint16
	}{
		{
			name:  "Valid iBeacon",
			data:  []uint8{0x4c, 0x0, 0x2, 0x15, 0xa4, 0x95, 0xbb, 0x30, 0xc5, 0xb1, 0x4b, 0x44, 0xb5, 0x12, 0x13, 0x70, 0xf0, 0x2d, 0x74, 0xde, 0x0, 0x46, 0x3, 0xfc, 0xc5},
			err:   nil,
			uuid:  "a495bb30c5b14b44b5121370f02d74de",
			major: 70,
			minor: 1020,
		},
		{
			name: "Invalid iBeacon",
			data: []uint8{0x4c, 0x0, 0x2, 0x15},
			err:  ibeacon.ErrNotBeacon,
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			got, err := ibeacon.NewIBeacon(tc.data)

			if tc.err != nil {
				// expecting an error
				if !errors.Is(err, tc.err) {
					t.Fatalf("Expected '%v' error, got '%v' error", tc.err, err)
				}
				return
			}
			if got.UUID != tc.uuid {
				t.Errorf("Expected %v, got %v", tc.uuid, got.UUID)
			}
			if got.Major != tc.major {
				t.Errorf("Expected %v, got %v", tc.major, got.Major)
			}
			if got.Minor != tc.minor {
				t.Errorf("Expected %v, got %v", tc.minor, got.Minor)
			}
		})
	}
}
