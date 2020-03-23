import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:znkauth/znkauth.dart';

void main() {
  const MethodChannel channel = MethodChannel('znkauth');

  TestWidgetsFlutterBinding.ensureInitialized();

  setUp(() {
    channel.setMockMethodCallHandler((MethodCall methodCall) async {
      return '42';
    });
  });

  tearDown(() {
    channel.setMockMethodCallHandler(null);
  });

  test('getPlatformVersion', () async {
    expect(await Znkauth.platformVersion, '42');
  });
}
