import 'package:flutter/material.dart';

abstract class CustomHead {
  // 头部背景视图
  Widget get backgroundView;
  // 左侧视图
  Widget get leftView;
  // 标题视图
  Widget get titleView;
  // 设置标题内容
  set title(String txt);
  // 状态视图
  Widget get statusView;
  // 状态标题内容
  set statuText(String txt);
  // 右侧视图
  Widget get rightView;
  
}