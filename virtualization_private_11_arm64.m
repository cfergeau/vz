//
//  virtualization_private_11_arm64.m
//

#ifdef __arm64__
#import "apple-private/_VZPL011SerialPortConfiguration.h"
#import "virtualization_private_11_arm64.h"

/*!
 @abstract Create a new PL011 Serial Port Device configuration
 @param attachment Base class for a serial port attachment.
 @discussion
    The device creates a console which enables communication between the host and the guest through the PL011 interface.
 */
void *newVZPL011SerialPortConfiguration(void *attachment)
{
    if (@available(macOS 11, *)) {
        _VZPL011SerialPortConfiguration *config = [[_VZPL011SerialPortConfiguration alloc] init];
        [config setAttachment:(VZSerialPortAttachment *)attachment];
        return config;
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}

#endif
