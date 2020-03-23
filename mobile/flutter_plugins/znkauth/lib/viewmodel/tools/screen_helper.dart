import 'dart:ui';

import 'package:flutter/material.dart';

class ScreenHelper {
  
  /// 设计图宽度
  double _width = 375;
  /// 设计图高度
  double _height = 667;
  /// 上下文
  BuildContext _context;
  /// 宽缩放比例
  double _scaleWidth;
  /// 高缩放比例
  double _scaleHeight;
  /// 屏幕宽度
  double _screenWidth;
  /// 屏幕高度
  double _screenHeight;
  /// 像素比
  double _pixelRadio;
  /// 文本缩放因子
  double _textScaleFactor;
  /// 是否允许根据系统字体大小缩放字体
  bool _allowFontScaling;
  /// 顶部安全区域
  double _safeTopArea;
  /// 底部安全区域
  double _safeBottomArea;
  
  /// 单例
  ScreenHelper._();
  static final ScreenHelper _instance = ScreenHelper._();
  static ScreenHelper get shared => _instance;
  /* ScreenHelper 工厂模式 */
  factory ScreenHelper() {
    return _instance;
  }

  /// 设置设计参数
  static void setDesignParams(double width, double height, [bool allowFontScaling = false]) {
    if (width > 0) {
      ScreenHelper.shared._width = width;
    }
    if (height > 0) {
      ScreenHelper.shared._height = height;
    }
    ScreenHelper.shared._allowFontScaling = allowFontScaling;
    _config();
  }
  /// 设置上下文
  static void setContext(BuildContext context) {
    ScreenHelper.shared._context = context;
    _config();
  }
  /// 配置
  static void _config() {
    BuildContext ctx = ScreenHelper.shared._context;
    if (ctx != null) {
      MediaQueryData data = MediaQuery.of(ctx);
      ScreenHelper.shared._screenWidth = data.size.width;
      ScreenHelper.shared._screenHeight = data.size.height;
      ScreenHelper.shared._pixelRadio = data.devicePixelRatio;
      ScreenHelper.shared._textScaleFactor = data.textScaleFactor;
      ScreenHelper.shared._safeTopArea = data.padding.top;
      ScreenHelper.shared._safeBottomArea = data.padding.bottom;
    } else {
      ScreenHelper.shared._screenWidth = window.physicalSize.width;
      ScreenHelper.shared._screenHeight = window.physicalSize.height;
      ScreenHelper.shared._pixelRadio = window.devicePixelRatio;
      ScreenHelper.shared._textScaleFactor = window.textScaleFactor;
      ScreenHelper.shared._safeTopArea = window.padding.top;
      ScreenHelper.shared._safeBottomArea = window.padding.bottom;
    }
    ScreenHelper.shared._scaleWidth = ScreenHelper.shared._screenWidth / ScreenHelper.shared._width;
    ScreenHelper.shared._scaleHeight = ScreenHelper.shared._screenHeight / ScreenHelper.shared._height;

  }
  /// 宽度缩放比例
  static double get scaleWidth => ScreenHelper.shared._scaleWidth;

  /// 高度缩放比
  static double get scaleHeight => ScreenHelper.shared._scaleHeight;

  /// 像素比
  static double get pixelRadio => ScreenHelper.shared._pixelRadio;

  /// 屏幕宽度
  static double get screenWidth => ScreenHelper.shared._screenWidth;

  /// 屏幕高度
  static double get screenHeight => ScreenHelper.shared._screenHeight;

  /// 顶部安全区域，状态栏高度
  static double get safeTopArea => ScreenHelper.shared._safeTopArea;

  /// 底部安全区域，分栏高度
  static double get safeBottomArea => ScreenHelper.shared._safeBottomArea;

  /// 缩放后的宽度
  static num setWidth(num w) => w * scaleWidth;

  /// 缩放后的高度
  static num setHeight(num h) => h * scaleHeight;

  /// 缩放后字体大小
  static num setFont(num fontSize) => ScreenHelper.shared._allowFontScaling ? 
                                        setWidth(fontSize) : 
                                        setWidth(fontSize) / ScreenHelper.shared._textScaleFactor;
}