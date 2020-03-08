#import "ZnkAuthPlugin.h"
#if __has_include(<znk_auth/znk_auth-Swift.h>)
#import <znk_auth/znk_auth-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "znk_auth-Swift.h"
#endif

@implementation ZnkAuthPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftZnkAuthPlugin registerWithRegistrar:registrar];
}
@end
