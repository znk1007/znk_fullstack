import 'package:flb/api/api.dart';
import 'package:flutter/material.dart';

class ZNKBaseViewModel extends ChangeNotifier{
  //API对象
  final ZNKApi api;
  //是否已被销毁
  bool _disposed = false;

  ZNKBaseViewModel({@required this.api});

  @override
  void dispose() {
    super.dispose();
    _disposed = true;
  }

  @override
  void notifyListeners() {
    if (!_disposed) {
      super.notifyListeners();
    }
  }
}
  