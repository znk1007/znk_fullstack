//
//  PathHelper.h
//  Runner
//
//  Created by Sam Huang on 2019/12/17.
//

#import <Flutter/Flutter.h>

NS_ASSUME_NONNULL_BEGIN

@interface PathHelper : NSObject

/// 注册
/// @param registry 注册代理
+ (void)registerWithRegistry:(NSObject<FlutterPluginRegistry>*)registry;

@end

NS_ASSUME_NONNULL_END
