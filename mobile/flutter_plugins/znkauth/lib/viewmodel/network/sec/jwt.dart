//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkxIjoidGVzdDEiLCJrZXkyIjoidGVzdDIiLCJ0aW1lc3RhbXAiOjE1ODYwMDI1MjV9.y11Z68ehIV8E7TTq3OdrlcyQgBnX9Byceabea8HifSI
import 'package:jose/jose.dart';

import 'crypt.dart';

const testStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkxIjoidGVzdDEiLCJrZXkyIjoidGVzdDIiLCJrZXkzIjoidGVzdDMiLCJ0aW1lc3RhbXAiOiIxNTg2MDYxMzM5NjMyMzg0In0.trm2f0n1CztebCk3NuajIISoZ_oX2W9luAMygx6NeFo";

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
    buidler.addRecipient(
      JsonWebKey.fromJson({
        'kty': 'oct',
        'k':Crypt.cryptKey,
      }),
      algorithm: 'HS256',
    );
    var jws = buidler.build();
    return jws.toCompactSerialization();
  }
}

void main(List<String> args) {
  Map<String, dynamic> res =ZnkAuthJWT.parse(testStr);
  print('res: $res');
  var params = Map<String, dynamic>();
  params['key1'] = 'test1';
  params['key2'] = 'test2';
  var tk = ZnkAuthJWT.token(params, null);
  print('jwt compact serialization: $tk');
  Map<String, dynamic> res1 =ZnkAuthJWT.parse(testStr);
  print('res1: $res1');
}