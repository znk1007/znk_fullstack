import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:znk_auth/znk_auth.dart';

void main() {
  const MethodChannel channel = MethodChannel('znk_auth');

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
    expect(await ZnkAuth.platformVersion, '42');
  });
}
