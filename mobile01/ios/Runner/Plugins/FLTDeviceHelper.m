//
//  FLTDeviceHelper.m
//  Runner
//
//  Created by Sam Huang on 2019/12/16.
//

#import "FLTDeviceHelper.h"
#import <sys/utsname.h>
@implementation FLTDeviceHelper

+ (void)registerWithRegistrar:(nonnull NSObject<FlutterPluginRegistrar> *)registrar {
    FlutterMethodChannel *channel = [FlutterMethodChannel methodChannelWithName:@"device_helper" binaryMessenger:[registrar messenger]];
    FLTDeviceHelper *instance = [[FLTDeviceHelper alloc] init];
    [registrar addMethodCallDelegate:instance channel:channel];
}

- (void)handleMethodCall:(FlutterMethodCall *)call result:(FlutterResult)result {
    if ([@"getIOSDeviceInfo" isEqualToString:call.method]) {
           UIDevice* device = [UIDevice currentDevice];
           struct utsname un;
           uname(&un);

           result(@{
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
           });
     } else {
         result(FlutterMethodNotImplemented);
     }
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
