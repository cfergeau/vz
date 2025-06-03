//
//  virtualization_11_private_arm64.h
//

#pragma once

#import "virtualization_helper.h"
#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

#ifdef __arm64__

void *newVZPL011SerialPortConfiguration(void *attachment);

#endif
