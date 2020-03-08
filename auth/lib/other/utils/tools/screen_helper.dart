import 'dart:ui';

import 'package:flutter/material.dart';

class ScreenHelper {
  ScreenHelper._();
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
  static ScreenHelper _instance;
  /// 在程序启动时执行，或
  /// @override
  /// Widget build(BuildContext context) {}
  /// 前执行
  static ScreenHelper get _shared {
    if (_instance == null) {
      _instance = ScreenHelper._();
    }
    return _instance;
  }
  /// 设置设计参数
  static void setDesignParams(double width, double height, [bool allowFontScaling = false]) {
    if (width > 0) {
      ScreenHelper._shared._width = width;
    }
    if (height > 0) {
      ScreenHelper._shared._height = height;
    }
    ScreenHelper._shared._allowFontScaling = allowFontScaling;
    _config();
  }
  /// 设置上下文
  static void setContext(BuildContext context) {
    ScreenHelper._shared._context = context;
    _config();
  }
  /// 配置
  static void _config() {
    BuildContext ctx = ScreenHelper._shared._context;
    if (ctx != null) {
      MediaQueryData data = MediaQuery.of(ctx);
      ScreenHelper._shared._screenWidth = data.size.width;
      ScreenHelper._shared._screenHeight = data.size.height;
      ScreenHelper._shared._pixelRadio = data.devicePixelRatio;
      ScreenHelper._shared._textScaleFactor = data.textScaleFactor;
      ScreenHelper._shared._safeTopArea = data.padding.top;
      ScreenHelper._shared._safeBottomArea = data.padding.bottom;
    } else {
      ScreenHelper._shared._screenWidth = window.physicalSize.width;
      ScreenHelper._shared._screenHeight = window.physicalSize.height;
      ScreenHelper._shared._pixelRadio = window.devicePixelRatio;
      ScreenHelper._shared._textScaleFactor = window.textScaleFactor;
      ScreenHelper._shared._safeTopArea = window.padding.top;
      ScreenHelper._shared._safeBottomArea = window.padding.bottom;
    }
    ScreenHelper._shared._scaleWidth = ScreenHelper._shared._screenWidth / ScreenHelper._shared._width;
    ScreenHelper._shared._scaleHeight = ScreenHelper._shared._screenHeight / ScreenHelper._shared._height;

  }
  /// 宽度缩放比例
  static double get scaleWidth => ScreenHelper._shared._scaleWidth;

  /// 高度缩放比
  static double get scaleHeight => ScreenHelper._shared._scaleHeight;

  /// 像素比
  static double get pixelRadio => ScreenHelper._shared._pixelRadio;

  /// 屏幕宽度
  static double get screenWidth => ScreenHelper._shared._screenWidth;

  /// 屏幕高度
  static double get screenHeight => ScreenHelper._shared._screenHeight;

  /// 顶部安全区域，状态栏高度
  static double get safeTopArea => ScreenHelper._shared._safeTopArea;

  /// 底部安全区域，分栏高度
  static double get safeBottomArea => ScreenHelper._shared._safeBottomArea;

  /// 缩放后的宽度
  static num setWidth(num w) => w * scaleWidth;

  /// 缩放后的高度
  static num setHeight(num h) => h * scaleHeight;

  /// 缩放后字体大小
  static num setFont(num fontSize) => ScreenHelper._shared._allowFontScaling ? 
                                        setWidth(fontSize) : 
                                        setWidth(fontSize) / ScreenHelper._shared._textScaleFactor;
}