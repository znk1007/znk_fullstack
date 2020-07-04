#import "TableviewPlugin.h"
#if __has_include(<tableview/tableview-Swift.h>)
#import <tableview/tableview-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "tableview-Swift.h"
#endif

@implementation TableviewPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftTableviewPlugin registerWithRegistrar:registrar];
}
@end
