//
//  DeviceHelper.h
//  Runner
//
//  Created by Sam Huang on 2019/12/17.
//

#import <UIKit/UIKit.h>

NS_ASSUME_NONNULL_BEGIN

@interface DeviceHelper : NSObject

/// 设备信息
@property (nonatomic, strong) NSDictionary<NSString *, NSObject *> *buildDict;

@end

NS_ASSUME_NONNULL_END
