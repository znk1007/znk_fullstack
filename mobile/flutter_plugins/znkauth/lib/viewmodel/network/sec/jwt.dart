//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkxIjoidGVzdDEiLCJrZXkyIjoidGVzdDIiLCJ0aW1lc3RhbXAiOjE1ODYwMDI1MjV9.y11Z68ehIV8E7TTq3OdrlcyQgBnX9Byceabea8HifSI
import 'package:jose/jose.dart';

import 'crypt.dart';

const testStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkxIjoidGVzdDEiLCJrZXkyIjoidGVzdDIiLCJ0aW1lc3RhbXAiOiIxNTg2MDAzNzk4In0.hjnQEQ4K95bWp_AI4y5hoSlpVZ-W5Bmp5OdnutlRxmw";

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
    print('params: $_builder');
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
    print('jwt compact serialization: ${jws.toCompactSerialization()}');
  }
}

void main(List<String> args) {
  Map<String, dynamic> res =ZnkAuthJWT.parse(testStr);
  print('res: $res');
  var params = Map<String, dynamic>();
  params['key1'] = 'test1';
  params['key2'] = 'test2';
  var tk = ZnkAuthJWT.token(params, null);
}