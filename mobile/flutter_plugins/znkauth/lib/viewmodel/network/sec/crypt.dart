// import 'package:encrypt/encrypt.dart' as Encrypt;
// class Crypt {
//   //加密密钥
//   static const _cryptKey = 'fullstack-manznk';
//   static const _cryptIV = '0000000000000000';
//   //encode 加密
//   static String encode(String src) {
//     String encodeStr = '';
//     try {
//       final key = Encrypt.Key.fromUtf8(_cryptKey);
//       final iv = Encrypt.IV.fromUtf8(_cryptIV);
//       final encrypter = Encrypt.Encrypter(Encrypt.AES(key, mode: Encrypt.AESMode.cbc));
//       final encrypted = encrypter.encrypt(src, iv: iv);
//       return encrypted.base16;
//     } catch (e) {
      
//     }
//     return encodeStr;
//   }
// }