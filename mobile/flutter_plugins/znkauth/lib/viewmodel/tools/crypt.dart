
import 'dart:convert';

import 'package:encrypt/encrypt.dart';
//应用标识
final String _appKey = 'fullstack^@-znk';
class CryptManager {
  
  static final _key = Key.fromUtf8(_appKey);
  static final _iv = IV.fromUtf8(_appKey);
  static final encrypter = Encrypter(AES(_key));
  /* 加密 */
  static String encrypt(String src){
    final encrypted = encrypter.encrypt(src, iv: _iv);
    return encrypted.base64.toString();
  }
  /* 解密 */
  static String decrypt(String src) {
    final encrypted = Encrypted(base64.decode(src));
    final decrypted = encrypter.decrypt(encrypted, iv:_iv);
    return decrypted;
  }
}