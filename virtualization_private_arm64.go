package vz

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c -fno-objc-arc
#cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Virtualization -framework Cocoa
# include "virtualization_private_11_arm64.h"
*/
import "C"
import (
	"github.com/Code-Hex/vz/v3/internal/objc"
)

// PL011SerialPortConfiguration represents PL011 Serial Port Device.
//
// The device creates a console which enables communication between the host and the guest through PL011.
// This uses private APIs and could break at any time.
type PL011SerialPortConfiguration struct {
	SerialPortConfiguration
}

// NewPL011SerialPortConfiguration creates a new NewPL011SerialPortConfiguration.
//
// This is only supported on macOS 11 and newer, error will
// be returned on older versions.
func NewPL011SerialPortConfiguration(attachment SerialPortAttachment) (*SerialPortConfiguration, error) {
	if err := macOSAvailable(11); err != nil {
		return nil, err
	}

	config := &PL011SerialPortConfiguration{
		SerialPortConfiguration{
			pointer: objc.NewPointer(
				C.newVZPL011SerialPortConfiguration(
					objc.Ptr(attachment),
				),
			),
		},
	}
	objc.SetFinalizer(config, func(self *PL011SerialPortConfiguration) {
		objc.Release(self)
	})
	return &config.SerialPortConfiguration, nil
}
