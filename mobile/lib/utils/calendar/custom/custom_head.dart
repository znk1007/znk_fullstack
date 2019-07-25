import 'package:flutter/material.dart';

abstract class CustomHead {
  // 头部背景视图
  Widget get backgroundView;
  // 左侧视图
  Widget get leftView;
  // 左侧视图位置
  EdgeInsets get leftViewPosition;
  // 标题视图
  Widget get titleView;
  // 标题内容
  String get title;
  // 设置标题
  set title(String txt);
  // 标题位置
  EdgeInsets get titleViewPosition;
  // 状态视图
  Widget get statusView;
  // 状态标题内容
  String get statusText;
  set statuText(String txt);
  // 状态视图位置
  EdgeInsets get statusViewPostion;
  // 右侧视图
  Widget get rightView;
  // 右侧视图位置
  EdgeInsets get rightViewPosition;
  
  
}