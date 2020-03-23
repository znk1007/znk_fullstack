//
//  Generated file. Do not edit.
//

#import "GeneratedPluginRegistrant.h"

#if __has_include(<sqflite/SqflitePlugin.h>)
#import <sqflite/SqflitePlugin.h>
#else
@import sqflite;
#endif

#if __has_include(<znk_auth/ZnkAuthPlugin.h>)
#import <znk_auth/ZnkAuthPlugin.h>
#else
@import znk_auth;
#endif

@implementation GeneratedPluginRegistrant

+ (void)registerWithRegistry:(NSObject<FlutterPluginRegistry>*)registry {
  [SqflitePlugin registerWithRegistrar:[registry registrarForPlugin:@"SqflitePlugin"]];
  [ZnkAuthPlugin registerWithRegistrar:[registry registrarForPlugin:@"ZnkAuthPlugin"]];
}

@end
