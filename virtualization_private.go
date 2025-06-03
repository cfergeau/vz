package vz

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c -fno-objc-arc
#cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Virtualization -framework Cocoa
# include "virtualization_private_11.h"
*/
import "C"
import (
	"github.com/Code-Hex/vz/v3/internal/objc"
)

// Uart16550SerialPortConfiguration represents 16550 Serial Port Device.
//
// The device creates a console which enables communication between the host and the guest through 16550.
// This uses private APIs and could break at any time.
type Uart16550SerialPortConfiguration struct {
	SerialPortConfiguration
}

// NewUart16550SerialPortConfiguration creates a new NewUart16550SerialPortConfiguration.
//
// This is only supported on macOS 11 and newer, error will
// be returned on older versions.
func NewUart16550SerialPortConfiguration(attachment SerialPortAttachment) (*SerialPortConfiguration, error) {
	if err := macOSAvailable(11); err != nil {
		return nil, err
	}

	config := &Uart16550SerialPortConfiguration{
		SerialPortConfiguration{
			pointer: objc.NewPointer(
				C.newVZ16550SerialPortConfiguration(
					objc.Ptr(attachment),
				),
			),
		},
	}
	objc.SetFinalizer(config, func(self *Uart16550SerialPortConfiguration) {
		objc.Release(self)
	})
	return &config.SerialPortConfiguration, nil
}
