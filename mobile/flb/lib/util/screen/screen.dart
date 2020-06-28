import 'dart:ui';

import 'package:flutter/material.dart';

class Screen {
  
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
  Screen._();
  static final Screen _instance = Screen._();
  static Screen get shared => _instance;
  /* Screen 工厂模式 */
  factory Screen() {
    return _instance;
  }

  /// 设置设计参数
  static void setDesignParams(double width, double height, [bool allowFontScaling = false]) {
    if (width > 0) {
      Screen.shared._width = width;
    }
    if (height > 0) {
      Screen.shared._height = height;
    }
    Screen.shared._allowFontScaling = allowFontScaling;
    _config();
  }
  /// 设置上下文
  static void setContext(BuildContext context) {
    Screen.shared._context = context;
    _config();
  }
  /// 配置
  static void _config() {
    BuildContext ctx = Screen.shared._context;
    if (ctx != null) {
      MediaQueryData data = MediaQuery.of(ctx);
      Screen.shared._screenWidth = data.size.width;
      Screen.shared._screenHeight = data.size.height;
      Screen.shared._pixelRadio = data.devicePixelRatio;
      Screen.shared._textScaleFactor = data.textScaleFactor;
      Screen.shared._safeTopArea = data.padding.top;
      Screen.shared._safeBottomArea = data.padding.bottom;
    } else {
      Screen.shared._screenWidth = window.physicalSize.width;
      Screen.shared._screenHeight = window.physicalSize.height;
      Screen.shared._pixelRadio = window.devicePixelRatio;
      Screen.shared._textScaleFactor = window.textScaleFactor;
      Screen.shared._safeTopArea = window.padding.top;
      Screen.shared._safeBottomArea = window.padding.bottom;
    }
    Screen.shared._scaleWidth = Screen.shared._screenWidth / Screen.shared._width;
    Screen.shared._scaleHeight = Screen.shared._screenHeight / Screen.shared._height;

  }
  /// 宽度缩放比例
  static double get scaleWidth => Screen.shared._scaleWidth;

  /// 高度缩放比
  static double get scaleHeight => Screen.shared._scaleHeight;

  /// 像素比
  static double get pixelRadio => Screen.shared._pixelRadio;

  /// 屏幕宽度
  static double get screenWidth => Screen.shared._screenWidth;

  /// 屏幕高度
  static double get screenHeight => Screen.shared._screenHeight;

  /// 顶部安全区域，状态栏高度
  static double get safeTopArea => Screen.shared._safeTopArea;

  /// 底部安全区域，分栏高度
  static double get safeBottomArea => Screen.shared._safeBottomArea;

  /// 缩放后的宽度
  static num setWidth(num w) => w * scaleWidth;

  /// 缩放后的高度
  static num setHeight(num h) => h * scaleHeight;

  /// 缩放后字体大小
  static num setFont(num fontSize) => Screen.shared._allowFontScaling ? 
                                        setWidth(fontSize) : 
                                        setWidth(fontSize) / Screen.shared._textScaleFactor;
}