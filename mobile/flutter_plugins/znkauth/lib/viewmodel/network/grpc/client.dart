import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';
export 'package:znkauth/model/protos/generated/auth/regist.pbgrpc.dart';
class ZnkAuthRpc {  
  //通道选项
  ChannelOptions _options;
  //主机
  String _host;
  //端口
  int _port;
  //单例
  ZnkAuthRpc._();
  static final ZnkAuthRpc _instance = new ZnkAuthRpc._();
  //工厂模式初始化
  factory ZnkAuthRpc() {
    return _instance;
  }
  //单例调用
  static ZnkAuthRpc get shared => _instance;
  //rpc配置
  Future setRpc({
    @required bool useTls, 
    @required bool useTlsCA,
    @required String host,
    @required int port,
    @required BuildContext ctx,
  }) async {
    this._host = host;
    this._port = port;
    await _innerInit(
      useTls: useTls,
      useTlsCA: useTlsCA,
      ctx: ctx,
    );
  }
  //初始化方法
  Future _innerInit({
    bool useTls, 
    bool useTlsCA,
    BuildContext ctx,
  }) async {
    ChannelCredentials _credentials;
    if (useTls) {
      List<int> trustedRoot;
      if (useTlsCA) {
        try {
          String key = await DefaultAssetBundle.of(ctx).loadString('lib/viewmodel/network/sec/keys/client.pem');
          trustedRoot = key.codeUnits;
          _credentials = ChannelCredentials.secure(
            certificates: trustedRoot,
          );
        } catch (e) {
          print('client init err: $e');
        }
      } else {
        _credentials = ChannelCredentials.insecure();
      }
    }
    _options = ChannelOptions(
      credentials: _credentials,
    );
  }
  //启动连接服务
  Future<ClientChannel> run() async {
    if (_options == null) {
      return null;
    }
    ClientChannel channel = ClientChannel(
      this._host,
      port: this._port,
      options: _options,
    );
    return channel;
  }
}