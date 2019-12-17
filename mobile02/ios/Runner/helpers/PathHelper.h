//
//  PathHelper.h
//  Runner
//
//  Created by Sam Huang on 2019/12/17.
//

#import <Foundation/Foundation.h>
#define kGetTemporaryDirectory @"getTemporaryDirectory"
#define kGetApplicationDocumentsDirectory @"getApplicationDocumentsDirectory"
#define kGetApplicationSupportDirectory @"getApplicationSupportDirectory"
#define kGetLibraryDirectory @"getLibraryDirectory"

NS_ASSUME_NONNULL_BEGIN

@interface PathHelper : NSObject

/// 获取指定方法路径
/// @param method 方法名
+ (NSObject *)getDevicePathWithMethod:(NSString *)method;

@end

NS_ASSUME_NONNULL_END
