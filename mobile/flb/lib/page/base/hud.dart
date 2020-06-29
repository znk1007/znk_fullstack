import 'package:flb/page/base/tab.dart';

class Hud {
  //单例
  Hud._();
  static final Hud _instance = Hud._();  
  /* 单例方法一 */
  factory Hud() {
    return _instance;
  }
  /* 单例方法二 */
  static Hud get shared => _instance;
  //分栏页
  TabPage _tabPage;
  //设置分栏页码对象
  void wrap(TabPage tab) {
    _tabPage = tab;
  }
  //显示加载框
  void show() {
    if (_tabPage != null) {
      _tabPage.state.showLoading();
    }
  }
  //隐藏加载框
  void hide() {
    if (_tabPage != null) {
      _tabPage.state.hideLoading();
    }
  }
}