import 'package:flb/api/api.dart';
import 'package:flb/models/style/style.dart';
import 'package:provider/provider.dart';
import 'package:provider/single_child_widget.dart';

List<SingleChildWidget> znkProviders = [
  ...independentServices,
  ...dependentServices,
];

List<SingleChildWidget> independentServices = [
  Provider(create: (_) => ZNKApi()),
];

List<SingleChildWidget> dependentServices = [
  //这里使用ProxyProvider来定义需要依赖其他Provider的服务
  ChangeNotifierProvider(create: (_) => ThemeStyle()),
];
