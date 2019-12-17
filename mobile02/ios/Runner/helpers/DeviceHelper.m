//
//  DeviceHelper.m
//  Runner
//
//  Created by Sam Huang on 2019/12/17.
//

#import "DeviceHelper.h"
#import <sys/utsname.h>

@implementation DeviceHelper

- (instancetype)init
{
    self = [super init];
    if (self) {
        _buildDict = @{};
        [self fetchDeviceInfo];
    }
    return self;
}

- (void)fetchDeviceInfo {
    UIDevice* device = [UIDevice currentDevice];
    struct utsname un;
    uname(&un);
     _buildDict = @{
           @"name" : [device name],
           @"systemName" : [device systemName],
           @"systemVersion" : [device systemVersion],
           @"model" : [device model],
           @"localizedModel" : [device localizedModel],
           @"identifierForVendor" : [[device identifierForVendor] UUIDString],
           @"isPhysicalDevice" : [self isDevicePhysical],
           @"utsname" : @{
                 @"sysname" : @(un.sysname),
                 @"nodename" : @(un.nodename),
                 @"release" : @(un.release),
                 @"version" : @(un.version),
                 @"machine" : @(un.machine),
           }
     };
}

- (NSString *)isDevicePhysical {
#if TARGET_OS_SIMULATOR
    NSString *isPhysicalDevice = @"false";
#else
    NSString *isPhysicalDevice = @"true";
#endif
    return isPhysicalDevice;
}
@end
