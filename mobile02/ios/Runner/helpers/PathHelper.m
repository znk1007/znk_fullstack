//
//  PathHelper.m
//  Runner
//
//  Created by Sam Huang on 2019/12/17.
//

#import "PathHelper.h"
#import <Flutter/Flutter.h>

#define kPluginPathHelperKey @"kPluginPathHelperKey"

#define kGetTemporaryDirectory @"getTemporaryDirectory"
#define kGetApplicationDocumentsDirectory @"getApplicationDocumentsDirectory"
#define kGetApplicationSupportDirectory @"getApplicationSupportDirectory"
#define kGetLibraryDirectory @"getLibraryDirectory"

NSString* GetDirectoryOfType(NSSearchPathDirectory dir) {
  NSArray* paths = NSSearchPathForDirectoriesInDomains(dir, NSUserDomainMask, YES);
  return paths.firstObject;
}

static FlutterError* getFlutterError(NSError* error) {
  if (error == nil) return nil;
  return [FlutterError errorWithCode:[NSString stringWithFormat:@"Error %ld", (long)error.code]
                             message:error.domain
                             details:error.localizedDescription];
}

@interface PathHelper ()<FlutterPlugin>

@end

@implementation PathHelper

/// 注册
/// @param registry 注册代理
+ (void)registerWithRegistry:(NSObject<FlutterPluginRegistry>*)registry {
    if (![registry hasPlugin:kPluginPathHelperKey]) {
        [self registerWithRegistrar:[registry registrarForPlugin:kPluginPathHelperKey]];
    }
}


+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
      FlutterMethodChannel* channel =
          [FlutterMethodChannel methodChannelWithName:@"path_helper_channel"
                                      binaryMessenger:registrar.messenger];
      [channel setMethodCallHandler:^(FlutterMethodCall* call, FlutterResult result) {
            if ([kGetTemporaryDirectory isEqualToString:call.method]) {
                result([self getTemporaryDirectory]);
            } else if ([kGetApplicationDocumentsDirectory isEqualToString:call.method]) {
                result([self getApplicationDocumentsDirectory]);
            } else if ([kGetApplicationSupportDirectory isEqualToString:call.method]) {
                  NSString* path = [self getApplicationSupportDirectory];
                  // Create the path if it doesn't exist
                  NSError* error;
                  NSFileManager* fileManager = [NSFileManager defaultManager];
                  BOOL success = [fileManager createDirectoryAtPath:path
                                        withIntermediateDirectories:YES
                                                         attributes:nil
                                                              error:&error];
                  if (!success) {
                      result(getFlutterError(error));
                  } else {
                      result(path);
                  }
            } else if ([kGetLibraryDirectory isEqualToString:call.method]) {
                result([self getLibraryDirectory]);
            } else {
                result(FlutterMethodNotImplemented);
            }
      }];
}

+ (NSString*)getTemporaryDirectory {
    return GetDirectoryOfType(NSCachesDirectory);
}

+ (NSString*)getApplicationDocumentsDirectory {
    return GetDirectoryOfType(NSDocumentDirectory);
}

+ (NSString*)getApplicationSupportDirectory {
    return GetDirectoryOfType(NSApplicationSupportDirectory);
}

+ (NSString*)getLibraryDirectory {
    return GetDirectoryOfType(NSLibraryDirectory);
}


@end
