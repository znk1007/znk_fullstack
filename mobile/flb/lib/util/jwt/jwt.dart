//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkxIjoidGVzdDEiLCJrZXkyIjoidGVzdDIiLCJ0aW1lc3RhbXAiOjE1ODYwMDI1MjV9.y11Z68ehIV8E7TTq3OdrlcyQgBnX9Byceabea8HifSI
import 'dart:convert';
import 'dart:io';
import 'package:jose/jose.dart';
import 'package:x509/x509.dart';

class ZnkAuthJWT {
  /* 解析jwt */
  static Future<Map<String, dynamic>>parse(String token) async {
    var jwt = JsonWebToken.unverified(token);
    return jwt.claims.toJson();
  }
  /* 生成token */
  static String token(Map<String, dynamic> params, String timestamp) {
    var ts = timestamp;
    if (ts == null || ts.length == 0) {
      var mse = DateTime.now().millisecondsSinceEpoch / 1000;
      ts = (mse.toInt()).toString();
    }
    if (params == null) {
      params = Map<String, dynamic>();
    }
    params['timestamp'] = ts;
    var clms = JsonWebTokenClaims.fromJson(params);
    var buidler = JsonWebSignatureBuilder();
    buidler.jsonContent = clms.toJson();
    var key = _readPrivateKeyFromFile('lib/viewmodel/network/sec/keys/jwt.rsa');
    buidler.addRecipient(
      key,
      algorithm: 'RS512',
    );
    var jws = buidler.build();
    return jws.toCompactSerialization();
  }
}
/* 读取私钥 */
JsonWebKey _readPrivateKeyFromFile(String path) {
  var v = parsePem(File(path).readAsStringSync()).first;
  var keyPair = (v is PrivateKeyInfo) ? v.keyPair : v as KeyPair;
  var pKey = keyPair.privateKey as RsaPrivateKey;
  String _bytesToBase64(List<int> bytes) {
    return base64Url.encode(bytes).replaceAll('=', '');
  }

  String _intToBase64(BigInt v) {
    return _bytesToBase64(v
        .toRadixString(16)
        .replaceAllMapped(RegExp('[0-9a-f]{2}'), (m) => '${m.group(0)},')
        .split(',')
        .where((v) => v.isNotEmpty)
        .map((v) => int.parse(v, radix: 16))
        .toList());
  }
  return JsonWebKey.fromJson({
    'kty': 'RSA',
    'n': _intToBase64(pKey.modulus),
    'd': _intToBase64(pKey.privateExponent),
    'p': _intToBase64(pKey.firstPrimeFactor),
    'q': _intToBase64(pKey.secondPrimeFactor),
    'alg': 'RS512',
    'kid': 'fullstack-manznk'
  });
}

void main(List<String> args) async {
  var testStr = "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJrZXkxIjoidGVzdDEiLCJrZXkyIjoidGVzdDIiLCJrZXkzIjoidGVzdDMiLCJ0aW1lc3RhbXAiOiIxNTg2MDg0OTIxIn0.fmVr591aBbnP6gkyedBu1bgCI2ZbadmzRXNjLY1XLB4-eVO1qOS1U9xEJgCBhyof5J_dc5Mmmq-aNb0YfCMPHIYw0C8PHKT5NYLGz1A80U-1-vp-bSDPwoitjJbGk7QsUbhmdCgYam5z92_gFm1aHXi1nI9mtXRYJVufgVg5nmo";
  Map<String, dynamic> res = await ZnkAuthJWT.parse(testStr);
  print('res: $res');
  var params = Map<String, dynamic>();
  params['key1'] = 'test1';
  params['key2'] = 'test2';
  var tk = ZnkAuthJWT.token(params, null);
  print('jwt compact serialization: $tk');
  Map<String, dynamic> res1 = await ZnkAuthJWT.parse(tk);
  print('res1: $res1');
}
