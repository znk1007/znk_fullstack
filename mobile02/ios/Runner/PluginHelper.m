//
//  PluginHelper.m
//  Runner
//
//  Created by Sam Huang on 2019/12/17.
//

#import "PluginHelper.h"
#import "DeviceHelper.h"
#import "PathHelper.h"
#define kPluginHelperKey @"kPluginHelperKey"


@interface PluginHelper() <FlutterPlugin>

@end

@implementation PluginHelper

/// 注册
/// @param registry 注册代理
+ (void)registerWithRegistry:(NSObject<FlutterPluginRegistry>*)registry {
    if (![registry hasPlugin:kPluginHelperKey]) {
        [self registerWithRegistrar:[registry registrarForPlugin:kPluginHelperKey]];
    }
}

+ (void)registerWithRegistrar:(nonnull NSObject<FlutterPluginRegistrar> *)registrar {
    FlutterMethodChannel *channel = [FlutterMethodChannel methodChannelWithName:@"method_channel_helper" binaryMessenger:[registrar messenger]];
    PluginHelper *instance = [[PluginHelper alloc] init];
    [registrar addMethodCallDelegate:instance channel:channel];
}

- (void)handleMethodCall:(FlutterMethodCall *)call result:(FlutterResult)result {
    if ([call.method isEqualToString:@"getIOSDeviceInfo"]) {
        DeviceHelper *helper = [[DeviceHelper alloc] init];
        result(helper.buildDict);
    } else if ([call.method isEqualToString:kGetTemporaryDirectory] ||
                [call.method isEqualToString:kGetApplicationDocumentsDirectory] ||
               [call.method isEqualToString:kGetApplicationSupportDirectory] ||
               [call.method isEqualToString:kGetLibraryDirectory]
               ) {
        [PathHelper getDevicePathWithMethod:call.method];
    } else {
        result(FlutterMethodNotImplemented);
    }
}

@end
