import 'package:test/test.dart';
import 'package:znkauth/viewmodel/tools/crypt.dart';
void main() {
  test('encrypt test', () {
    final plainText = 'Lorem ipsum dolor sit amet, consectetur adipiscing elit';
    final encryptStr = CryptManager.encrypt(plainText);
    final decryptStr = CryptManager.decrypt(encryptStr);
    expect(decryptStr, plainText);
  });
}