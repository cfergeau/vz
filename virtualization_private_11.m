//
//  virtualization_private_11.m
//

#import "apple-private/_VZ16550SerialPortConfiguration.h"
#import "virtualization_private_11.h"

/*!
 @abstract Create a new 16550 Serial Port Device configuration
 @param attachment Base class for a serial port attachment.
 @discussion
    The device creates a console which enables communication between the host and the guest through the 16550 interface.
 */
void *newVZ16550SerialPortConfiguration(void *attachment)
{
    if (@available(macOS 11, *)) {
        _VZ16550SerialPortConfiguration *config = [[_VZ16550SerialPortConfiguration alloc] init];
        [config setAttachment:(VZSerialPortAttachment *)attachment];
        return config;
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
