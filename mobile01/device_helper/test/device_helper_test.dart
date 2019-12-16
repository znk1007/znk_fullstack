import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:device_helper/device_helper.dart';

void main() {
  const MethodChannel channel = MethodChannel('device_helper');

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
    expect(await DeviceHelper.platformVersion, '42');
  });
}
