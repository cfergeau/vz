package vz

import "errors"

func NewPL011SerialPortConfiguration(attachment SerialPortAttachment) (*SerialPortConfiguration, error) {
	return nil, errors.New("PL011 serial ports unavailabe on x86_64 hardware")
}
