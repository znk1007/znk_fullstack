///
//  Generated code. Do not modify.
//  source: login.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'login.pb.dart' as $0;
export 'login.pb.dart';

class LoginSrvClient extends $grpc.Client {
  static final _$userLogin = $grpc.ClientMethod<$0.LoginReq, $0.LoginReq>(
      '/login.LoginSrv/userLogin',
      ($0.LoginReq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.LoginReq.fromBuffer(value));

  LoginSrvClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.LoginReq> userLogin($0.LoginReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$userLogin, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class LoginSrvServiceBase extends $grpc.Service {
  $core.String get $name => 'login.LoginSrv';

  LoginSrvServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.LoginReq, $0.LoginReq>(
        'userLogin',
        userLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.LoginReq.fromBuffer(value),
        ($0.LoginReq value) => value.writeToBuffer()));
  }

  $async.Future<$0.LoginReq> userLogin_Pre(
      $grpc.ServiceCall call, $async.Future<$0.LoginReq> request) async {
    return userLogin(call, await request);
  }

  $async.Future<$0.LoginReq> userLogin(
      $grpc.ServiceCall call, $0.LoginReq request);
}
