///
//  Generated code. Do not modify.
//  source: login.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:grpc/service_api.dart' as $grpc;

import 'dart:core' as $core show int, String, List;

import 'login.pb.dart' as $3;
import 'login.pb.dart';
export 'login.pb.dart';

class LoginClient extends $grpc.Client {
  static final _$login = $grpc.ClientMethod<LoginRequest, LoginResponse>(
      '/protos.login.Login/login',
      (LoginRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => LoginResponse.fromBuffer(value));

  LoginClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<LoginResponse> login(LoginRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$login, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class LoginServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.login.Login';

  LoginServiceBase() {
    $addMethod($grpc.ServiceMethod<LoginRequest, LoginResponse>(
        'login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => LoginRequest.fromBuffer(value),
        (LoginResponse value) => value.writeToBuffer()));
  }

  $async.Future<LoginResponse> login_Pre(
      $grpc.ServiceCall call, $async.Future request) async {
    return login(call, await request);
  }

  $async.Future<LoginResponse> login(
      $grpc.ServiceCall call, LoginRequest request);
}
