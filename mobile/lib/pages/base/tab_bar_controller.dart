
import 'dart:collection';
import 'dart:math' as math;

import 'package:flutter/material.dart';

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

class TabBarItem {
  /// 标题组件
  final Widget title;
  /// 常态视图组件
  final Widget icon;
  /// 选中状态视图组件
  final Widget activeIcon;
  /// 背景颜色
  final Color backgroundColor;
  /// 角标组件
  final Widget badge;
  /// 角标数
  final String badgeNum;
  /// 角标颜色
  final Color badgeColor;

  /// 初始化
  TabBarItem({
    @required this.icon,
    this.title,
    Widget activeIcon,
    this.backgroundColor,
    this.badge,
    this.badgeNum,
    Color badgeColor,
  })  : activeIcon = activeIcon ?? icon,
        badgeColor = badgeColor ?? Colors.red;
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

/// 水花效果
class _Circle {
  final _TabBarWidgetState state;
  final int index;
  final Color color;
  AnimationController controller;
  CurvedAnimation animation;
  _Circle({
    @required this.state, 
    @required this.index,
    @required this.color,
    @required TickerProvider vsync, 
  })  : assert(state != null), 
        assert(index != null), 
        assert(color != null) {
          controller = AnimationController(
            duration: kThemeAnimationDuration, 
            vsync: vsync,
          );
          animation = CurvedAnimation(
            parent: controller, 
            curve: Curves.fastOutSlowIn, 
          );
          controller.forward();
        }
    double get horizontalLeadingOffset {
      double weightSum(Iterable<Animation<double>> animations) {
        return animations
                .map(state._evaluteFlex)
                .fold<double>(0.0, (double sum, double value) => sum + value);
      }

      final double allWeights = weightSum(state._animations);
      final double leadingWeights = weightSum(state._animations.sublist(0, index));

      return (leadingWeights + state._evaluteFlex(state._animations[index]) / 2.0) / allWeights;
    }

    void dispose() {
      controller.dispose();
    }
}
/// 雷达波纹效果
class _RadialPainter extends CustomPainter {
  final List<_Circle> circles;
  final TextDirection textDirection;

  _RadialPainter({
    @required this.circles, 
    @required this.textDirection, 
  })  : assert(circles != null), 
        assert(textDirection != null);

  static double _maxRadius(Offset center, Size size) {
    final double maxX = math.max(center.dx, size.width - center.dx);
    final double maxY = math.max(center.dy, size.height - center.dy);
    return math.sqrt(maxX * maxX + maxY * maxY);
  }

  @override
  void paint(Canvas canvas, Size size) {
    for (_Circle circle in circles) {
      final Paint paint = Paint()..color = circle.color;
      final Rect rect = Rect.fromLTWH(0.0, 0.0, size.width, size.height);
      canvas.clipRect(rect);
      double leftFraction;
      switch (textDirection) {
        case TextDirection.rtl:
          leftFraction = 1.0 - circle.horizontalLeadingOffset;
          break;
        case TextDirection.ltr:
          leftFraction = circle.horizontalLeadingOffset;
          break;
        default:
      }
      final Offset center = Offset(leftFraction * size.width, size.height / 2.0);
      final Tween<double> radiusTween = Tween<double>(
        begin: 0.0, 
        end: _maxRadius(center, size),
      );
      canvas.drawCircle(
        center, 
        radiusTween.transform(circle.animation.value), 
        paint,
      );
    }
  }

  @override
  bool shouldRepaint(_RadialPainter oldDelegate) {
    if (textDirection != oldDelegate.textDirection) {
      return true;
    }
    if (circles == oldDelegate.circles) {
      return false;
    }

    if (circles.length != oldDelegate.circles.length) {
      return true;
    }
    for (var i = 0; i < circles.length; i++) {
      if (circles[i] != oldDelegate.circles[i]) {
        return true;
      }
    }
    return false;
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
/// splash 飞溅效果分栏
class _TabBarWidgetState extends State<TabBarWidget> 
  with TickerProviderStateMixin {

  
  List<AnimationController> _controllers = <AnimationController>[];
  List<CurvedAnimation> _animations;
  /// 水花纹 
  final Queue<_Circle> _circles = Queue<_Circle>();
  /// 背景颜色
  Color _backgroundColor;

  static final Animatable<double> _flexTween = 
    Tween<double>(begin: 1.0, end: 1.5);

  void _resetState() {
    for (AnimationController controller in _controllers) {
      controller.dispose();
    }
    for (var circle in _circles) {
      circle.dispose();
    }
    _circles.clear();

    _controllers = List<AnimationController>.generate(widget.items.length, (int index) {
      return AnimationController(
        duration: kThemeAnimationDuration, 
        vsync: this,
      )..addListener(_rebuild);
    });
    _animations = List<CurvedAnimation>.generate(widget.items.length, (int index) {
      return CurvedAnimation(
        parent: _controllers[index], 
        curve: Curves.fastOutSlowIn, 
        reverseCurve: Curves.fastOutSlowIn.flipped, 
      );
    });
    _controllers[widget.currentIndex].value = 1.0;
    _backgroundColor = widget.items[widget.currentIndex].backgroundColor;
  }

  @override
  void initState() {
    super.initState();
    _resetState();
  }

  @override
  Widget build(BuildContext context) {
    assert(debugCheckHasDirectionality(context));
    assert(debugCheckHasMaterialLocalizations(context));
    final double additionalBottomPadding = 
      math.max(MediaQuery.of(context).padding.bottom - _kBottomMargin, 0.0);
    Color backgroundColor;
    switch (widget.type) {
      case TabBarTapType.fixed:
        
        break;
      case TabBarTapType.shifting:
        backgroundColor = _backgroundColor;
        break;
      default:
    }
    return Semantics(
      container: true,
      explicitChildNodes: true,
      child: Stack(
        children: <Widget>[
          Positioned.fill(
            child: Material(
              elevation: 8.0,//shadow
              color: backgroundColor,
            ),
          ),
          ConstrainedBox(
            constraints: BoxConstraints(
              minHeight: kBottomNavigationBarHeight + additionalBottomPadding, 
            ),
            child: Stack(
              children: <Widget>[
                Positioned.fill(
                  child: CustomPaint(
                    painter: _RadialPainter(
                      circles: _circles.toList(),
                      textDirection: Directionality.of(context), 
                    ),
                  ),
                ),
                Material(
                  type: MaterialType.transparency,
                  child: Padding(
                    padding: EdgeInsets.only(
                      bottom: additionalBottomPadding, 
                    ),
                    child: MediaQuery.removePadding(
                      context: context,
                      removeBottom: true,
                      child: _createContainer(_createTiles()),
                    ),
                  ),
                )
              ],
            ),
          ), 
        ],
      ),
    );
  }

  @override
  void dispose() {
    for (var controller in _controllers) {
      controller.dispose();
    }
    for (var circle in _circles) {
      circle.dispose();
    }
    super.dispose();
  }

  void _rebuild() {
    setState(() {
      
    });
  }

  double _evaluteFlex(Animation<double> animation) => 
      _flexTween.evaluate(animation);

  void _pushCircle(int index) {
    if (widget.items.length <= index || widget.items[index].backgroundColor == null) {
      return;
    }
    _circles.add(
      _Circle(
        state: this, 
        index: index, 
        color: widget.items[index].backgroundColor, 
        vsync: this, 
      )..controller.addStatusListener( (AnimationStatus status) {
        switch (status) {
          case AnimationStatus.completed:
            setState(() {
              final _Circle circle = _circles.removeFirst();
              _backgroundColor = circle.color;
              circle.dispose();
            });
            break;
          case AnimationStatus.dismissed:
          case AnimationStatus.forward:
          case AnimationStatus.reverse:
            break;
          default:
        }
      }), 
    );
  }

  Widget _createContainer(List<Widget> tiles) {
      return DefaultTextStyle.merge(
        overflow: TextOverflow.ellipsis, 
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: tiles,
        ),
      );
    }

    List<Widget> _createTiles() {
      final MaterialLocalizations localizations = MaterialLocalizations.of(context);
      assert(localizations != null);
      final List<Widget> children = <Widget>[];
      switch (widget.type) {
        case TabBarTapType.fixed: 
          final ThemeData themeData = Theme.of(context);
          final TextTheme textTheme = themeData.textTheme;
          Color themeColor;
          switch (themeData.brightness) {
            case Brightness.dark:
              themeColor = themeData.accentColor;
              break;
            case Brightness.light:
              themeColor = themeData.primaryColor;
              break;
            default:
          }
          final ColorTween colorTween = ColorTween(
            begin: textTheme.caption.color, 
            end: widget.fixedColor ?? themeColor, 
          );
          for (var i = 0; i < widget.items.length; i++) {
            children.add(
              _TabBarTitle(
                widget.type, 
                widget.items[i], 
                _animations[i], 
                widget.iconSize,
                colorTween: colorTween,
                selected: i == widget.currentIndex,
                indexLabel: localizations.tabLabel(
                  tabIndex: i + 1, 
                  tabCount: widget.items.length, 
                ),
                isAnimation: widget.isAnimation,
                isInkResponse: widget.isInkResponse,
                badgeColor: widget.badgeColor == null ?
                            widget.fixedColor :
                            widget.badgeColor,
                onTap: () {
                  if (widget.onTap != null) {
                    widget.onTap(i);
                  }
                },
              ),
            );
          }
          break;
        case TabBarTapType.shifting:
          for (var i = 0; i < widget.items.length; i++) {
            children.add(
              _TabBarTitle(
                widget.type, 
                widget.items[i], 
                _animations[i], 
                widget.iconSize, 
                flex: _evaluteFlex(_animations[i]), 
                selected: i == widget.currentIndex,
                indexLabel: localizations.tabLabel(
                  tabIndex: i + 1, 
                  tabCount: widget.items.length, 
                ),
                isAnimation: widget.isAnimation,
                isInkResponse: widget.isInkResponse,
                badgeColor: widget.badgeColor == null ?
                            widget.fixedColor :
                            widget.badgeColor,
                onTap: () {
                  if (widget.onTap != null) widget.onTap(i);
                },
              )
            );
          }
          break;
        default:
      }
      return children;
    }

  @override
  void didUpdateWidget(TabBarWidget oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (widget.items.length != oldWidget.items.length) {
      _resetState();
      return;
    }
    if (widget.currentIndex != oldWidget.currentIndex) {
      switch (widget.type) {
        case TabBarTapType.fixed:
        break;
        case TabBarTapType.shifting:
          _pushCircle(widget.currentIndex);
        break;
      }
      _controllers[oldWidget.currentIndex].reverse();
      _controllers[widget.currentIndex].forward();
    } else {
      if (_backgroundColor != widget.items[widget.currentIndex].backgroundColor) {
        _backgroundColor = widget.items[widget.currentIndex].backgroundColor;
      }
    }
  }
}


