//
//  DeviceHelper.m
//  Runner
//
//  Created by Sam Huang on 2019/12/17.
//

#import "DeviceHelper.h"
#import <sys/utsname.h>
#define kPluginDeviceHelperKey @"kPluginDeviceHelperKey"

@interface DeviceHelper ()<FlutterPlugin>

@end

@implementation DeviceHelper

/// 注册
/// @param registry 注册代理
+ (void)registerWithRegistry:(NSObject<FlutterPluginRegistry>*)registry {
    if (![registry hasPlugin:kPluginDeviceHelperKey]) {
        [self registerWithRegistrar:[registry registrarForPlugin:kPluginDeviceHelperKey]];
    }
}

+ (void)registerWithRegistrar:(nonnull NSObject<FlutterPluginRegistrar> *)registrar {
    FlutterMethodChannel *channel = [FlutterMethodChannel methodChannelWithName:@"device_helper_channel" binaryMessenger:[registrar messenger]];
    DeviceHelper *instance = [[DeviceHelper alloc] init];
    [registrar addMethodCallDelegate:instance channel:channel];
}

- (void)handleMethodCall:(FlutterMethodCall *)call result:(FlutterResult)result {
    if ([call.method isEqualToString:@"getIOSDeviceInfo"]) {
        result([self fetchDeviceInfo]);
    } else {
        result(FlutterMethodNotImplemented);
    }
}


- (NSDictionary *)fetchDeviceInfo {
    UIDevice* device = [UIDevice currentDevice];
    struct utsname un;
    uname(&un);
     return @{
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
