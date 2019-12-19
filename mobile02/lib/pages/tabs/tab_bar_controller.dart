
import 'package:flutter/material.dart';
import 'tab_bar_item.dart';

const double _kActiveFontSize = 14.0;
const double _kInactiveFontSize = 12.0;
const double _kTopMargin = 6.0;
const double _kBottomMargin = 8.0;
/// 点击效果类型
enum TabBarTapType {
  /// 点击时固定，不可移动
  fixed,
  /// 点击时可移动
  shifting,
}

class _TabBarTitle extends StatelessWidget {

  final TabBarTapType type;
  final TabBarItem item;
  final Animation<double> animation;
  final double iconSize;
  final VoidCallback onTap;
  final ColorTween colorTween;
  final double flex;
  final bool selected;
  final String indexLabel;
  final bool isAnimation;
  final bool isInkResponse;
  final Color badgeColor;

  const _TabBarTitle(
    this.type,
    this.item,
    this.animation,
    this.iconSize, {
    this.onTap,
    this.colorTween,
    this.flex,
    this.selected = false,
    this.indexLabel,
    this.isAnimation = true,
    this.isInkResponse = true,
    this.badgeColor,
  }) : assert(selected != null);

  Widget _buildIcon() {
    double tweenBegin;
    Color iconColor;
    switch (type) {
      case TabBarTapType.fixed:
        tweenBegin = 8.0;
        iconColor = colorTween.evaluate(animation);
        break;
      case TabBarTapType.shifting:
        tweenBegin = 16.0;
        iconColor = Colors.white;
        break;
      default:
    }
    return Align(
      alignment: Alignment.topCenter,
      heightFactor: 1.0,
      child: Container(
        margin: EdgeInsets.only(
          top: isAnimation ?
                Tween<double>(
                  begin: tweenBegin, 
                  end: _kTopMargin,
                ).evaluate(animation) :
                _kTopMargin,
        ),
        child: IconTheme(
          data: IconThemeData(
            color: iconColor, 
            size: iconSize, 
          ),
          child: selected ? item.activeIcon : item.icon,
        ),
      ),
    );
  }

  Widget _buildFixedLabel() {
    double scale = isAnimation ?
            Tween<double>(
              begin: _kInactiveFontSize / _kActiveFontSize, 
              end: 1.0,
            ).evaluate(animation) : 1.0;
    return Align(
      alignment: Alignment.bottomCenter,
      heightFactor: 1.0,
      child: Container(
        margin: const EdgeInsets.only(
          bottom: _kBottomMargin, 
        ),
        child: DefaultTextStyle.merge(
          style: TextStyle(
            fontSize: _kActiveFontSize,
            color: colorTween.evaluate(animation),
          ),
          child: Transform(
            transform: Matrix4.diagonal3Values(
              scale, 
              scale, 
              scale
            ),
            alignment: Alignment.bottomCenter,
            child: item.title,
          ),
        ),
      ),
    );
  }

  Widget _buildShiftingLabel() {
    return Align(
      alignment: Alignment.bottomCenter,
      heightFactor: 1.0,
      child: Container(
        margin: EdgeInsets.only(
          bottom: Tween<double>(
            begin: 2.0, 
            end: _kBottomMargin
          ).evaluate(animation),
        ),
        child: FadeTransition(
          alwaysIncludeSemantics: true,
          opacity: animation,
          child: DefaultTextStyle(
            style: const TextStyle(
              fontSize: _kActiveFontSize,
              color: Colors.white,
            ), 
            child: item.title,
          ),
        ),
      ),
    );
  }

  Widget _buildBadge() {
    String badgeNum = item.badgeNum;
    if (item.badge == null && badgeNum == null) {
      return Container();
    }
    if (item.badge != null) {
      return item.badge;
    }
    badgeNum = badgeNum ?? '●';
    int num = int.tryParse(badgeNum);
    if (num == null) {
      badgeNum = '●';
    } else {
      if (num > 99) {
        badgeNum = '99+';
      }
    }
    
    return Container(
      width: 24,
      padding: EdgeInsets.fromLTRB(0, 2, 0, 2),
      alignment: Alignment.center,
      decoration: BoxDecoration(
        color: badgeColor, 
        shape: BoxShape.rectangle, 
        borderRadius: BorderRadius.all(Radius.circular(10)), 
      ),
      child: Text(
        badgeNum, 
        style: TextStyle(
          fontSize: 10, 
          color: item.badgeColor,
        ),
      ),
    );
  }

  Widget _buildInkWidget(Widget label) {
    if (isInkResponse) {
      return InkResponse(
        onTap: onTap, 
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          mainAxisSize: MainAxisSize.min,
          children: <Widget>[
            _buildIcon(), 
            label,
          ],
        ),
      );
    }
    return GestureDetector(
      onTap: onTap,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        mainAxisSize: MainAxisSize.max,
        children: <Widget>[
          _buildIcon(), 
          label, 
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    int tempFlex;
    Widget label;
    switch (type) {
      case TabBarTapType.fixed:
        tempFlex = 1;
        label = _buildFixedLabel();
        break;
      case TabBarTapType.shifting:
      tempFlex = (flex * 1000.0).round();
      label = _buildShiftingLabel();
      break;
      default:
    }
    return Expanded(
      flex: tempFlex,
      child: Semantics(
        container: true,
        header: true,
        selected: selected,
        child: Stack(
          children: <Widget>[
            Positioned(
              right: 4,
              top: 4,
              child: _buildBadge(),
            ),
            _buildInkWidget(label), 
            Semantics(
              label: indexLabel,
            )
          ],
        ),
      ),
    );
  }
}

class TabBarWidget extends StatefulWidget {
  /// 是否可显示动画，默认true
  final bool isAnimation;
  /// 角标颜色
  final Color badgeColor;
  /// 是否可墨汁轮廓响应, 默认true
  final bool isInkResponse;
  /// 标题数组
  final List<TabBarItem> items;
  /// 点击事件回调
  final ValueChanged<int> onTap;

  /// The index into [items] of the current active item.
  final int currentIndex;

  /// 点击效果类型
  final TabBarTapType type;

  /// 固定颜色
  final Color fixedColor;

  /// item标题字体大小
  final double iconSize;

  TabBarWidget({
    Key key,
    @required this.items,
    this.onTap,
    this.currentIndex = 0,
    TabBarTapType type,
    this.fixedColor,
    this.iconSize = 24.0,
    this.isAnimation = true,
    this.badgeColor,
    this.isInkResponse = true,
  }) :  assert(items != null), 
        assert(items.length >= 2), 
        assert(
          items.every((TabBarItem item) => item.title != null) == true,
          'All item title cannot be null'
        ),
        assert(0 <= currentIndex && currentIndex < items.length),
        assert(iconSize != null), 
        type = type ?? (items.length <= 3 ? TabBarTapType.fixed : TabBarTapType.shifting),
        super(key: key);

  @override
  _TabBarWidgetState createState() => _TabBarWidgetState();
}

class _TabBarWidgetState extends State<TabBarWidget> {
  @override
  Widget build(BuildContext context) {
    return Container(
       child: child,
    );
  }
}


