package vz

/*
#cgo darwin CFLAGS: -x objective-c -fno-objc-arc -mmacosx-version-min=11.0
#cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Virtualization
# include "virtualization.h"
*/
import "C"
import "runtime"

// VirtioEntropyDeviceConfiguration is used to expose a source of entropy for the guest operating system’s random-number generator.
// When you create this object and add it to your virtual machine’s configuration, the virtual machine configures a Virtio-compliant
// entropy device. The guest operating system uses this device as a seed to generate random numbers.
//
// see: https://developer.apple.com/documentation/virtualization/vzvirtioentropydeviceconfiguration?language=objc
type VirtioEntropyDeviceConfiguration struct {
	pointer
}

// NewVirtioEntropyDeviceConfiguration creates a new Virtio Entropy Device confiuration.
//
// This is only supported on macOS 11 and newer, ErrUnsupportedOSVersion will
// be returned on older versions.
func NewVirtioEntropyDeviceConfiguration() (*VirtioEntropyDeviceConfiguration, error) {
	if macosMajorVersionLessThan(11) {
		return nil, ErrUnsupportedOSVersion
	}

	config := &VirtioEntropyDeviceConfiguration{
		pointer: pointer{
			ptr: C.newVZVirtioEntropyDeviceConfiguration(),
		},
	}
	runtime.SetFinalizer(config, func(self *VirtioEntropyDeviceConfiguration) {
		self.Release()
	})
	return config, nil
}
