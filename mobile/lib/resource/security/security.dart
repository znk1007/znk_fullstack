import 'dart:convert';
import 'dart:io';
import 'package:encrypt/encrypt.dart' as Encrpyt;
import 'package:flutter/cupertino.dart';
import 'package:grpc/grpc.dart';

const encrpytKey = 'aes-znk1007man20';
const encryptIV = '0000000000000000';
/// 安全类型
class Security {  
  //安全类型
  bool _useTls;
  //证书
  String _certFile;
  //地址
  String _address;
  //端口
  int _port;
  //构造函数
  Security({String address = "localhost", int port = 9001, bool useTls = false, String certFile = ""}) {
    this._address = address;
    this._port = port;
    this._useTls = useTls;
    this._certFile = certFile;
  }
  ///元数据选项
  CallOptions configurateCallOptions() {
    return CallOptions(metadata: {'app_key': 'znk_project-item=20', 'app_secret': '19911007'});
  }

  ///配置通道
  Future<ClientChannel> configurateChannel(BuildContext ctx) async {
    ChannelCredentials credentials = await this._configurateCredentials(ctx);
    ChannelOptions channelOptions = ChannelOptions(
      credentials: credentials,
    );
    ClientChannel channel = ClientChannel(
      this._address,
      port: this._port,
      options: channelOptions
    );
    return channel;
  }

  /// 配置安全验证
  Future<ChannelCredentials> _configurateCredentials(BuildContext ctx) async {
    ChannelCredentials credentials;
    if (this._useTls) {
      List<int> trusted;
      if (this._certFile.isEmpty == false) {
        try {
          String key = await DefaultAssetBundle.of(ctx).loadString(this._certFile);
          trusted = key.codeUnits;
        } catch (e) {
          print("credentials err: $e");
        }
      }
      
      credentials = ChannelCredentials.secure(
          certificates: trusted,
          onBadCertificate: (X509Certificate certificate, String host) {
            return true;
          }
      );
    } else {
      credentials = ChannelCredentials.insecure();
    }
    return credentials;
  }

  // aes加密
  static String aesEncode(String src) {
    try {
      final key = Encrpyt.Key.fromUtf8(encrpytKey);
      final iv = Encrpyt.IV.fromUtf8(encryptIV);
      final encrpyter = Encrpyt.Encrypter(Encrpyt.AES(key, mode: Encrpyt.AESMode.cbc));
      final encrypted = encrpyter.encrypt(src, iv: iv);
      return encrypted.base16;
    } catch (_) {
      return '';
    }
  }

}

