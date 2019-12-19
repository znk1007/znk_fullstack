///
//  Generated code. Do not modify.
//  source: login.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'login.pb.dart' as $2;
export 'login.pb.dart';

class LoginClient extends $grpc.Client {
  static final _$login = $grpc.ClientMethod<$2.LoginRequest, $2.LoginResponse>(
      '/protos.login.Login/login',
      ($2.LoginRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.LoginResponse.fromBuffer(value));

  LoginClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$2.LoginResponse> login($2.LoginRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$login, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class LoginServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.login.Login';

  LoginServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.LoginRequest, $2.LoginResponse>(
        'login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.LoginRequest.fromBuffer(value),
        ($2.LoginResponse value) => value.writeToBuffer()));
  }

  $async.Future<$2.LoginResponse> login_Pre(
      $grpc.ServiceCall call, $async.Future<$2.LoginRequest> request) async {
    return login(call, await request);
  }

  $async.Future<$2.LoginResponse> login(
      $grpc.ServiceCall call, $2.LoginRequest request);
}
