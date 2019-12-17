//
//  PathHelper.m
//  Runner
//
//  Created by Sam Huang on 2019/12/17.
//

#import "PathHelper.h"
#import <Flutter/Flutter.h>

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

@implementation PathHelper

- (instancetype)init
{
    self = [super init];
    if (self) {
    }
    return self;
}

/// 获取指定方法路径
/// @param method 方法名
+ (NSObject *)getDevicePathWithMethod:(NSString *)method {
    if ([kGetTemporaryDirectory isEqualToString:method]) {
      return [self getTemporaryDirectory];
    } else if ([kGetApplicationDocumentsDirectory isEqualToString:method]) {
      return [self getApplicationDocumentsDirectory];
    } else if ([kGetApplicationSupportDirectory isEqualToString:method]) {
      NSString* path = [self getApplicationSupportDirectory];

      // Create the path if it doesn't exist
      NSError* error;
      NSFileManager* fileManager = [NSFileManager defaultManager];
      BOOL success = [fileManager createDirectoryAtPath:path
                            withIntermediateDirectories:YES
                                             attributes:nil
                                                  error:&error];
      if (!success) {
        return getFlutterError(error);
      } else {
        return path;
      }
    } else if ([kGetLibraryDirectory isEqualToString:method]) {
      return [self getLibraryDirectory];
    } else {
      return (NSObject *)FlutterMethodNotImplemented;
    }
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
