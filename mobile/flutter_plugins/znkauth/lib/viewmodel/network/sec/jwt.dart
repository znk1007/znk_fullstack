//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkxIjoidGVzdDEiLCJrZXkyIjoidGVzdDIiLCJ0aW1lc3RhbXAiOjE1ODYwMDI1MjV9.y11Z68ehIV8E7TTq3OdrlcyQgBnX9Byceabea8HifSI
import 'dart:convert';
import 'dart:io';
import 'package:jose/jose.dart';
import 'package:x509/x509.dart';

const testStr = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkxIjoidGVzdDEiLCJrZXkyIjoidGVzdDIiLCJrZXkzIjoidGVzdDMiLCJ0aW1lc3RhbXAiOiIxNTg2MDcyOTQxMjY5MDk5In0.f17Aas6tw6Ou0gz2urw4R-BSQ_ZcGF1QsRMPzGW4Xv8JN2FbhTOiE1r33yc7919mp0Z5TjWDMWJX_-Ul1Qy2xwXyU0BgVMWCKY0SO1AkiXsavQXkw8rnNzvuvsv2ZN7mH-R4M2A4dm2wuZ_Y7xvyCGTuEltijDuiTZ_MO0ybJQc";

class ZnkAuthJWT {
  /* 解析jwt */
  static Map<String, dynamic> parse(String token) {
    return JsonWebToken.unverified(token).claims.toJson();
  }
  /* 生成token */
  static String token(Map<String, dynamic> params, String timestamp) {
    var ts = timestamp;
    if (ts == null || ts.length == 0) {
      ts = (DateTime.now().millisecondsSinceEpoch).toString();
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
    'kid': 'some_id'
  });
}

void main(List<String> args) {
  Map<String, dynamic> res =ZnkAuthJWT.parse(testStr);
  print('res: $res');
  var params = Map<String, dynamic>();
  params['key1'] = 'test1';
  params['key2'] = 'test2';
  var tk = ZnkAuthJWT.token(params, null);
  print('jwt compact serialization: $tk');
  Map<String, dynamic> res1 =ZnkAuthJWT.parse(tk);
  print('res1: $res1');
}

