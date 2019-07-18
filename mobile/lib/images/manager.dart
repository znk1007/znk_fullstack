import 'package:flutter/material.dart';
import 'package:znk/utils/base/device.dart';
// 是否是iOS端
bool _isIOS = Device.isIOS;
class TabAsset {
  // 消息常态
  static Image get tabMsgNormal => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_msg_n.png') :
    Image.asset('');
  // 消息选中状态
  static Image get tabMsgSelected => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_msg_s.png') :
    Image.asset('');
    // 日程常态
  static Image get tabScheduleNormal => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_schedule_n.png') :
    Image.asset('');
  // 日程选中状态
  static Image get tabScheduleSelected => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_schedule_s.png') :
    Image.asset('');
    // 我的常态
  static Image get tabOwnerNormal => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_owner_n.png') :
    Image.asset('');
  // 消息选中状态
  static Image get tabOwnerSelected => _isIOS ? 
    Image.asset('lib/images/iOS/tabs/tab_owner_s.png') :
    Image.asset('');
}

class LaunchAsset {
    // 启动图
  static Image get launchImg => _isIOS ? 
    Image.asset('lib/images/iOS/launch/launchImg.png') :
    Image.asset('');
  // 启动图背景
  static Image get launchVersionBg => _isIOS ? 
    Image.asset('lib/images/iOS/launch/launchVersionBg.png') :
    Image.asset('');
}

class LoginAsset {
    // 登录背景图
  static Image get userBackground => _isIOS ? 
    Image.asset('lib/images/iOS/user/background.png') :
    Image.asset('');
  // 大行政
  static Image get daxingzheng => _isIOS ? 
    Image.asset('lib/images/iOS/user/daxingzheng.png') :
    Image.asset('');
  
  // 账号icon
  static Image get account => _isIOS ? 
    Image.asset('lib/images/iOS/user/account.png') :
    Image.asset('');
  // 密码icon
  static Image get password => _isIOS ? 
    Image.asset('lib/images/iOS/user/password.png') :
    Image.asset('');
}

class OwnerAsset {
    // 设置
  static Image get setting => _isIOS ? 
    Image.asset('lib/images/iOS/owner/setting.png') :
    Image.asset('');
  // 收藏
  static Image get collection => _isIOS ? 
    Image.asset('lib/images/iOS/owner/collection.png') :
    Image.asset('');
    // 网盘
  static Image get fileStore => _isIOS ? 
    Image.asset('lib/images/iOS/owner/file_store.png') :
    Image.asset('');
}

class CommonAsset {
    // 用户头像
  static Image get userHeader => _isIOS ? 
    Image.asset('lib/images/iOS/user/user_header.png') :
    Image.asset('');
  // 右箭头
  static Image get rightArrow => _isIOS ? 
    Image.asset('lib/images/iOS/common/right_arrow.png') :
    Image.asset('');
}


