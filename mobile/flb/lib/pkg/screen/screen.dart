import 'dart:ui';

import 'package:flutter/material.dart';

/*

XR/11	            6.1 inch	326 ppi	414*896 pt	828*1792 px	@2x
XS Max/11 Pro Max	6.5 inch	458 ppi	414*896 pt	1242*2688 px	@3x
X/XS/11 Pro	      5.8 inch	458 ppi	375*812 pt	1125*2436 px	@3x
6P/6SP/7P/8P	    5.5 inch	401 ppi	414*736 pt	1242*2208 px	@3x
6/6S/7/8	        4.7 inch	326 ppi	375*667 pt	750*1334 px	@2x
5/5S/5c/SE  	    4.0 inch	326 ppi	320*568 pt	640*1136 px	@2x
4/4s	            3.5 inch	326 ppi	320*480 pt	640*960 px	@2x

*/

class ZNKScreen {
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
  ZNKScreen._();
  static final ZNKScreen _instance = ZNKScreen._();
  static ZNKScreen get shared => _instance;
  /* ZNKScreen 工厂模式 */
  factory ZNKScreen() {
    return _instance;
  }

  /// 设置设计参数
  static void setDesignParams(double width, double height,
      [bool allowFontScaling = false]) {
    if (width > 0) {
      ZNKScreen.shared._width = width;
    }
    if (height > 0) {
      ZNKScreen.shared._height = height;
    }
    ZNKScreen.shared._allowFontScaling = allowFontScaling;
    _config();
  }

  /// 设置上下文
  static void setContext(BuildContext context) {
    ZNKScreen.shared._context = context;
    _config();
  }

  /// 配置
  static void _config() {
    BuildContext ctx = ZNKScreen.shared._context;
    if (ctx != null) {
      MediaQueryData data = MediaQuery.of(ctx);
      ZNKScreen.shared._screenWidth = data.size.width;
      ZNKScreen.shared._screenHeight = data.size.height;
      ZNKScreen.shared._pixelRadio = data.devicePixelRatio;
      ZNKScreen.shared._textScaleFactor = data.textScaleFactor;
      ZNKScreen.shared._safeTopArea = data.padding.top;
      ZNKScreen.shared._safeBottomArea = data.padding.bottom;
    } else {
      ZNKScreen.shared._screenWidth = window.physicalSize.width;
      ZNKScreen.shared._screenHeight = window.physicalSize.height;
      ZNKScreen.shared._pixelRadio = window.devicePixelRatio;
      ZNKScreen.shared._textScaleFactor = window.textScaleFactor;
      ZNKScreen.shared._safeTopArea = window.padding.top;
      ZNKScreen.shared._safeBottomArea = window.padding.bottom;
    }
    ZNKScreen.shared._scaleWidth =
        ZNKScreen.shared._screenWidth / ZNKScreen.shared._width;
    ZNKScreen.shared._scaleHeight =
        ZNKScreen.shared._screenHeight / ZNKScreen.shared._height;
  }

  /// 宽度缩放比例
  static double get scaleWidth => ZNKScreen.shared._scaleWidth;

  /// 高度缩放比
  static double get scaleHeight => ZNKScreen.shared._scaleHeight;

  /// 像素比
  static double get pixelRadio => ZNKScreen.shared._pixelRadio;

  /// 屏幕宽度
  static double get screenWidth => ZNKScreen.shared._screenWidth;

  /// 屏幕高度
  static double get screenHeight => ZNKScreen.shared._screenHeight;

  /// 顶部安全区域，状态栏高度
  static double get safeTopArea => ZNKScreen.shared._safeTopArea;

  /// 底部安全区域，分栏高度
  static double get safeBottomArea => ZNKScreen.shared._safeBottomArea;

  /// 缩放后的宽度
  static num setWidth(num w) => w * scaleWidth;

  /// 缩放后的高度
  static num setHeight(num h) => h * scaleHeight;

  /// 缩放后字体大小
  static num setFont(num fontSize) => ZNKScreen.shared._allowFontScaling
      ? setWidth(fontSize)
      : setWidth(fontSize) / ZNKScreen.shared._textScaleFactor;
}
