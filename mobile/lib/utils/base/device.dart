import 'dart:io';
import 'dart:ui' as ui;
import 'package:device_info/device_info.dart';
import 'package:flutter/widgets.dart';
import 'package:package_info/package_info.dart';

class Device {
  // 是否为iOS系统
  static bool get isIOS => Platform.isIOS;

  // 包信息
  static PackageInfo _packageInfo;
  // 获取包信息
  static getPackageInfo() async {
    _packageInfo = await PackageInfo.fromPlatform();
  }
  // 版本
  static String get version {
    return _packageInfo?.version ?? '1.0';
  }



  // 当前系统名称
  static String get systemName {
    return Platform.operatingSystem;
  }
  // 证书地址
  static String get pemPath {
    return 'lib/utils/security/keys/ca.pem';
  }
  // 获取window对象
  static MediaQueryData get window {
    return MediaQueryData.fromWindow(ui.window);
  }
  // 屏幕宽度
  static double get width {
    return window == null ? 0 : window.size.width;
  }
  // 屏幕宽度
  static double get height {
    return window == null ? 0 : window.size.height;
  }
  // 像素比
  static double get scale {
    return window == null ? 0 : window.devicePixelRatio;
  }
  // 状态栏高度
  static double get statusBarHeight {
    return window == null ? 0 : window.padding.top;
  }
  // 宽度比
  static double get iOSWidthScale {
    return width / 414.0;
  }
  // 高度比
  static double get iOSHeightScale {
    return height / 736.0;
  }
  // 相对宽度 iOS
  static double iOSRelativeWidth(double other) {
    return other * iOSWidthScale;
  }
  // 相对高度 iOS
  static double iOSRelativeHeight(double other) {
    return other * iOSHeightScale;
  }
  // 相对宽度 Android
  static double androidRelativeWidth(double other) {
    return other;
  }
  // 相对高度 Android
  static double androidRelativeHeight(double other) {
    return other;
  }

  // 相对宽度 通用
  static double relativeWidth(double other) {
    return isIOS ? iOSRelativeWidth(other) : androidRelativeWidth(other);
  }
  // 相对高度 通用
  static double relativeHeight(double other) {
    return isIOS ? iOSRelativeHeight(other) : androidRelativeHeight(other);
  }

}