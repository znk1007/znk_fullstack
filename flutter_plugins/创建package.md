开发Packages和插件
Package 介绍
Package 类型
Developing Dart packages
Step 1: 开发Dart包
Step 2: 实现package
开发插件包
Step 1: 创建 package
Step 2: 实现包 package
Step 2a: 定义包API（.dart）
Step 2b: 添加Android平台代码（.java / .kt）
Step 2c: 添加iOS平台代码 (.h+.m/.swift)
Step 2d: 连接API和平台代码
添加文档
API documentation
发布 packages
处理包的相互依赖
Android
iOS
解决冲突
Package 介绍
使用package可以创建可轻松共享的模块化代码。一个最小的package包括

一个pubspec.yaml文件：声明了package的名称、版本、作者等的元数据文件。

一个 lib 文件夹：包括包中公开的(public)代码，最少应有一个<package-name>.dart文件

Package 类型
Packages可以包含多种内容：

Dart包：其中一些可能包含Flutter的特定功能，因此对Flutter框架具有依赖性，仅将其用于Flutter，例如fluro包。

插件包：一种专用的Dart包，其中包含用Dart代码编写的API，以及针对Android（使用Java或Kotlin）和/或针对iOS（使用ObjC或Swift）平台的特定实现。一个具体的例子是battery插件包。

Developing Dart packages
Step 1: 开发Dart包
要创建Dart包，请使用--template=package 来执行 flutter create

flutter create --template=package hello
这将在hello/文件夹下创建一个具有以下专用内容的package工程：

lib/hello.dart:
Package的Dart代码
test/hello_test.dart:
Package的单元测试代码.
Step 2: 实现package
对于纯Dart包，只需在主lib/<package name>.dart文件内或lib目录中的文件中添加功能 。

要测试软件包，请在test目录中添加unit tests。

有关如何组织包内容的更多详细信息，请参阅Dart library package文档。

开发插件包
如果你想开发一个调用特定平台API的包，你需要开发一个插件包，插件包是Dart包的专用版本。 插件包包含针对Android（Java或Kotlin代码）或iOS（Objective-C或Swift代码）编写的特定于平台的实现（可以同时包含Android和Ios原生的代码）。 API使用platform channels连接到特定平台（Android或IOS）。

Step 1: 创建 package
要创建插件包，请使用--template=plugin参数执行flutter create

使用--org选项指定您的组织，并使用反向域名表示法。该值用于生成的Android和iOS代码中的各种包和包标识符。

flutter create --org com.example --template=plugin hello
这将在hello/文件夹下创建一个具有以下专用内容的插件工程：

lib/hello.dart:
插件包的Dart API.
android/src/main/java/com/yourcompany/​hello/HelloPlugin.java:
插件包API的Android实现.
ios/Classes/HelloPlugin.m:
插件包API的ios实现.
example/:
一个依赖于该插件的Flutter应用程序，来说明如何使用它
默认情况下，插件项目针对iOS代码使用Objective-C，Android代码使用Java。如果您更喜欢Swift或Kotlin，则可以使用-i 或 -a 为iOS或Android指定语言。例如：

flutter create --template=plugin -i swift -a kotlin hello
Step 2: 实现包 package
由于插件包中包含用多种编程语言编写的多个平台的代码，因此需要一些特定的步骤来确保顺畅的体验。

Step 2a: 定义包API（.dart）
插件包的API在Dart代码中定义。打开主文件夹hello/ 。找到lib/hello.dart

Step 2b: 添加Android平台代码（.java / .kt）
我们建议您使用Android Studio编辑Android代码。

在Android Studio中编辑Android平台代码之前，首先确保代码至少已经构建过一次（例如，从IntelliJ运行示例应用程序或在终端执行cd hello/example; flutter build apk）

接下来

启动Android Studio
在’Welcome to Android Studio’对话框选择 ‘Import project’, 或者在菜单栏 ‘File > New > Import Project…‘,然后选择hello/example/android/build.gradle文件.
在’Gradle Sync’ 对话框, 选择 ‘OK’.
在’Android Gradle Plugin Update’ 对话框, 选择 ‘Don’t remind me again for this project’.
您插件的Android平台代码位于 hello/java/com.yourcompany.hello/​HelloPlugin.

您可以通过按下 ▶ 按钮从Android Studio运行示例应用程序.

Step 2c: 添加iOS平台代码 (.h+.m/.swift)
我们建议您使用Xcode编辑iOS代码 。

在编辑Xcode中的iOS平台代码之前，首先确保代码至少已经构建过一次（例如，从Xcode中运行示例应用程序或终端执行cd hello/example; flutter build ios --no-codesign）。

接下来

启动 Xcode
选择 ‘File > Open’, 然后选择 hello/example/ios/Runner.xcworkspace 文件.
您插件的iOS平台代码位于Pods/DevelopmentPods/hello/Classes/中。

您可以通过按下&#9654按钮来运行示例应用程序

Step 2d: 连接API和平台代码
最后，您需要将用Dart代码编写的API与平台特定的实现连接起来。这是通过platform channels完成的。

添加文档
建议将以下文档添加到所有软件包：

README.md:介绍包的文件
CHANGELOG.md 记录每个版本中的更改
LICENSE 包含软件包许可条款的文件
所有公共API的API文档 (详情见下文)
API documentation
在发布软件包时，API文档会自动生成并发布到dartdocs.org，示例请参阅device_info docs

如果您希望在本地生成API文档，请使用以下命令：

将目录更改为您的软件包的位置:

cd ~/dev/mypackage

告诉文档工具Flutter SDK的位置:

export FLUTTER_ROOT=~/dev/flutter (on macOS or Linux)

set FLUTTER_ROOT=~/dev/flutter (on Windows)

运行 dartdoc 工具 (它是Flutter SDK的一部分):

$FLUTTER_ROOT/bin/cache/dart-sdk/bin/dartdoc (on macOS or Linux)

%FLUTTER_ROOT%\bin\cache\dart-sdk\bin\dartdoc (on Windows)

有关如何编写API文档的提示，请参阅Effective Dart: Documentation

发布 packages
一旦你实现了一个包，你可以在Pub上发布它 ，这样其他开发人员就可以轻松使用它

在发布之前，检查pubspec.yaml、README.md以及CHANGELOG.md文件，以确保其内容的完整性和正确性。

然后, 运行 dry-run 命令以查看是否都准备OK了:

flutter packages pub publish --dry-run
最后, 运行发布命令:

flutter packages pub publish
有关发布的详细信息，请参阅Pub publishing docs

处理包的相互依赖
如果您正在开发一个hello包，它依赖于另一个包，则需要将该依赖包添加到pubspec.yaml文件的dependencies部分。 下面的代码使url_launcher插件的Dart API ，这在hello包中是可用的：

In hello/pubspec.yaml:

dependencies:
  url_launcher: ^0.4.2
现在你可以在hello中import 'package:url_launcher/url_launcher.dart' 然后 launch(someUrl) 了。

这与在Flutter应用程序或任何其他Dart项目中引用软件包没有什么不同

但是，如果hello碰巧是一个插件包，其平台特定的代码需要访问url_launcher公开的特定于平台的API，那么您还需要为特定于平台的构建文件添加合适的依赖声明，如下所示。

Android
在 hello/android/build.gradle:

android {
    // lines skipped
    dependencies {
        provided rootProject.findProject(":url_launcher")
    }
}
您现在可以在hello/android/src源码中import io.flutter.plugins.urllauncher.UrlLauncherPlugin访问UrlLauncherPlugin类

iOS
在hello/ios/hello.podspec:

Pod::Spec.new do |s|
  # lines skipped
  s.dependency 'url_launcher'
您现在可以在hello/ios/Classes源码中 #import "UrlLauncherPlugin.h" 然后访问 UrlLauncherPlugin 类

解决冲突
假设你想在你的hello包中使用some_package和other_package，并且这两个包都依赖url_launcher，但是依赖的是url_launcher的不同的版本。 那我们就有潜在的冲突。避免这种情况的最好方法是在指定依赖关系时，程序包作者使用版本范围而不是特定版本。

dependencies:
  url_launcher: ^0.4.2    # Good, any 0.4.x with x >= 2 will do.
  image_picker: '0.1.1'   # Not so good, only 0.1.1 will do.
如果some_package声明了上面的依赖关系,other_package声明了url_launcher版本像’0.4.5’或’^0.4.0’，pub将能够自动解决问题。 类似的注释适用于插件包对Gradle模块和Cocoa pods的平台特定的依赖关系。

即使some_package和other_package声明了不兼容的url_launcher版本，它仍然可能会和url_launcher以兼容的方式正常工作。 你可以通过向hello包的pubspec.yaml文件中添加依赖性覆盖声明来处理冲突，从而强制使用特定版本：

强制使用 0.4.3版本的url_launcher，在 hello/pubspec.yaml中:

dependencies:
  some_package:
  other_package:
dependency_overrides:
  url_launcher: '0.4.3'
如果冲突的依赖不是一个包，而是一个特定于Android的库，比如guava，那么必须将依赖重写声明添加到Gradle构建逻辑中。

强制使用23.0版本的guava库，在hello/android/build.gradle中：

configurations.all {
    resolutionStrategy {
        force 'com.google.guava:guava:23.0-android'
    }
}
Cocoapods目前不提供依赖覆盖功能。