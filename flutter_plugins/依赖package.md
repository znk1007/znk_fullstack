使用 packages
Flutter支持使用由其他开发者贡献给Flutter和Dart生态系统的共享软件包。这使您可以快速构建应用程序，而无需从头开始开发所有应用程序。

现有的软件包支持许多使用场景，例如，网络请求（http），自定义导航/路由处理（fluro）， 集成设备API（如url_launcher＆battery） 以及使用第三方平台SDK（如 Firebase(需翻墙)））。

如果您正打算开发新的软件包，请参阅开发软件包。

如果您希望添加资源、图片或字体，无论是存储在文件还是包中，请参阅资源和图片。

使用包
搜索packages
将包依赖项添加到应用程序
开发新的packages
管理包依赖和版本
Package versions
更新依赖包
依赖未发布的packages
例子
例子: 使用 CSS Colors package
Example: 使用URL Launcher package to 启动浏览器
使用包
搜索packages
Packages会被发布到了 Pub 包仓库.

Flutter landing 页面 显示了与Flutter兼容的包（即声明依赖通常与扑兼容）。所有已发布的包都支持搜索。

将包依赖项添加到应用程序
要将包’css_colors’添加到应用中，请执行以下操作

依赖它
打开 pubspec.yaml 文件，然后在dependencies下添加css_colors:
安装它
在 terminal中: 运行 flutter packages get
或者
在 IntelliJ中: 点击pubspec.yaml文件顶部的’Packages Get’
导入它
在您的Dart代码中添加相应的import语句.
有关完整示例，请参阅下面的CSS Colors example below.

开发新的packages
如果某个软件包不适用于您的特定需求，则可以开发新的自定义package。

管理包依赖和版本
Package versions
所有软件包都有一个版本号，在他们的pubspec.yaml文件中指定。Pub会在其名称旁边显示软件包的当前版本（例如，请参阅url_launcher软件包）以及所有先前版本的列表。

当pubspec.yaml使用速记形式添加包时，plugin1: 这被解释为plugin1: any，即可以使用任何版本的包。为了确保某个包在更新后还可以正常使用，我们建议使用以下格式之一指定版本范围：

范围限制: 指定一个最小和最大的版本号,如:
dependencies:
  url_launcher: '>=0.1.2 <0.2.0'
范围限制使用 caret 语法: 与常规的范围约束类似（译者语：这和node下npm的版本管理类似）
dependencies:
  collection: '^0.1.2'
有关更多详细信息，请参阅 Pub 版本管理指南.

更新依赖包
当你在添加一个包后首次运行（IntelliJ中的’Packages Get’）flutter packages get，Flutter将找到包的版本保存在pubspec.lock。这确保了如果您或您的团队中的其他开发人员运行flutter packages get后回获取相同版本的包。

如果要升级到软件包的新版本，例如使用该软件包中的新功能，请运行flutter packages upgrade（在IntelliJ中点击Upgrade dependencies）。 这将根据您在pubspec.yaml中指定的版本约束下载所允许的最高可用版本。

依赖未发布的packages
即使未在Pub上发布，软件包也可以使用。对于不用于公开发布的专用插件，或者尚未准备好发布的软件包，可以使用其他依赖项选项：

路径 依赖: 一个Flutter应用可以依赖一个插件通过文件系统的path:依赖。路径可以是相对的，也可以是绝对的。例如，要依赖位于应用相邻目录中的插件’plugin1’，请使用以下语法

dependencies:
  plugin1:
    path: ../plugin1/
Git 依赖: 你也可以依赖存储在Git仓库中的包。如果软件包位于仓库的根目录中，请使用以下语法：

dependencies:
  plugin1:
    git:
      url: git://github.com/flutter/plugin1.git
Git 依赖于文件夹中的包: 默认情况下，Pub假定包位于Git存储库的根目录中。如果不是这种情况，您可以使用path参数指定位置，例如：

dependencies:
  package1:
    git:
      url: git://github.com/flutter/packages.git
      path: packages/package1        
最后，您可以使用ref参数将依赖关系固定到特定的git commit，branch或tag。有关更多详细信息，请参阅 Pub Dependencies article.

例子
例子: 使用 CSS Colors package
该css_colors包为CSS颜色定义颜色常量，允许您在Flutter中需要Color类型的任何位置使用它们

要使用这个包:

创建一个名为 ‘cssdemo’的新项目

打开 pubspec.yaml, 并将:
dependencies:
  flutter:
    sdk: flutter
替换为:

dependencies:
  flutter:
    sdk: flutter
  css_colors: ^1.0.0
在terminal中运行 flutter packages get, 或者在IntelliJ钟点击’Packages get’

打开 lib/main.dart 并替换其全部内容:
import 'package:flutter/material.dart';
import 'package:css_colors/css_colors.dart';

void main() {
  runApp(new MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      home: new DemoPage(),
    );
  }
}

class DemoPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      body: new Container(color: CSSColors.orange)
    );
  }
}
运行应用程序
Example: 使用URL Launcher package to 启动浏览器
URL Launcher可以使您打开移动平台上的默认浏览器显示给定的URL。 它演示了软件包如何包含特定于平台的代码（我们称这些软件包为插件）。它在Android和iOS上均受支持。

使用这个插件:

创建一个名为’launchdemo’的新项目

打开 pubspec.yaml, 并将:
dependencies:
  flutter:
    sdk: flutter
替换为:

dependencies:
  flutter:
    sdk: flutter
  url_launcher: ^0.4.1
在terminal中运行 flutter packages get, 或者在IntelliJ钟点击’Packages get’

打开 lib/main.dart 并替换其全部内容:

import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';

void main() {
  runApp(new MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return new MaterialApp(
      home: new DemoPage(),
    );
  }
}

class DemoPage extends StatelessWidget {
  launchURL() {
    launch('https://flutter.io');
  }

  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      body: new Center(
        child: new RaisedButton(
          onPressed: launchURL,
          child: new Text('Show Flutter homepage'),
        ),
      ),
    );
  }
}
运行应用程序。当您点击“Show Flutter homepage”时，您应该看到手机的默认浏览器打开，并出现Flutter主页