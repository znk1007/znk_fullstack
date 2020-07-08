import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKBaseView<T extends ZNKBaseViewModel> extends StatefulWidget {
  //初始化回调
  final Widget Function(BuildContext context, T model, Widget child) builder;
  //子树部件
  final Widget child;
  //视图模型
  final T model;
  //模型是否已初始化完成
  final Function(T) onReady;

  ZNKBaseView({Key key, this.model, this.child, this.builder, this.onReady})
      : super(key: key);

  @override
  _ZNKBaseViewState createState() => _ZNKBaseViewState();
}

class _ZNKBaseViewState<T extends ZNKBaseViewModel> extends State<ZNKBaseView> {
  T model;

  @override
  void initState() {
    model = widget.model;
    if (widget.onReady != null) {
      widget.onReady(model);
    }
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider<T>(
        create: (ctx) {
          return model;
        },
        child: Consumer<T>(
          builder: widget.builder,
          child: widget.child,
        ));
  }
}
