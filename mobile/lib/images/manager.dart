import 'package:flutter/material.dart';
import 'package:flutter/painting.dart';
import 'package:znk/utils/base/device.dart';
// 是否是iOS端
bool _isIOS = Device.isIOS;
class TabAsset {
  // 消息常态
  static String get tabMsgNormal => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_msg_n.png') :
    '';
  // 消息选中状态
  static String get tabMsgSelected => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_msg_s.png') :
    '';
    // 日程常态
  static String get tabScheduleNormal => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_schedule_n.png') :
    '';
  // 日程选中状态
  static String get tabScheduleSelected => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_schedule_s.png') :
    '';
    // 我的常态
  static String get tabOwnerNormal => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_owner_n.png') :
    '';
  // 消息选中状态
  static String get tabOwnerSelected => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_owner_s.png') :
    '';
}

class LaunchAsset {
    // 启动图
  static String get launchImg => _isIOS ? 
    Image.asset('lib/images/iOS/launch/launchImg.png') :
    '';
  // 启动图背景
  static String get launchVersionBg => _isIOS ? 
    Image.asset('lib/images/iOS/launch/launchVersionBg.png') :
    '';
}

class LoginAsset {
    // 登录背景图
  static String get userBackground => _isIOS ? 
    Image.asset('lib/images/iOS/user/background.png') :
    '';
  // 大行政
  static String get daxingzheng => _isIOS ? 
    Image.asset('lib/images/iOS/user/daxingzheng.png') :
    '';
  
  // 账号icon
  static String get account => _isIOS ? 
    Image.asset('lib/images/iOS/user/account.png') :
    '';
  // 密码icon
  static String get password => _isIOS ? 
    Image.asset('lib/images/iOS/user/password.png') :
    '';
}

class OnwerAsset {
    // 设置
  static String get setting => _isIOS ? 
    Image.asset('lib/images/iOS/owner/setting.png') :
    '';
  // 收藏
  static String get collection => _isIOS ? 
    Image.asset('lib/images/iOS/owner/collection.png') :
    '';
    // 网盘
  static String get fileStore => _isIOS ? 
    Image.asset('lib/images/iOS/owner/file_store.png') :
    '';
}

class CommonAsset {
    // 用户头像
  static String get userHeader => _isIOS ? 
    Image.asset('lib/images/iOS/user/user_header.png') :
    '';
  // 右箭头
  static String get rightArrow => _isIOS ? 
    Image.asset('lib/images/iOS/common/right_arrow.png') :
    '';
}


