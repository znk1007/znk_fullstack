#import "DeviceHelperPlugin.h"
#if __has_include(<device_helper/device_helper-Swift.h>)
#import <device_helper/device_helper-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "device_helper-Swift.h"
#endif

@implementation DeviceHelperPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftDeviceHelperPlugin registerWithRegistrar:registrar];
}
@end
