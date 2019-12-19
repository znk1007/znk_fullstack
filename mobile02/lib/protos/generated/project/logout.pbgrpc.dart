///
//  Generated code. Do not modify.
//  source: logout.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'logout.pb.dart' as $3;
export 'logout.pb.dart';

class LogoutServiceClient extends $grpc.Client {
  static final _$logout =
      $grpc.ClientMethod<$3.LogoutRequest, $3.LogoutResponse>(
          '/protos.logout.LogoutService/logout',
          ($3.LogoutRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $3.LogoutResponse.fromBuffer(value));

  LogoutServiceClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$3.LogoutResponse> logout($3.LogoutRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$logout, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class LogoutServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.logout.LogoutService';

  LogoutServiceBase() {
    $addMethod($grpc.ServiceMethod<$3.LogoutRequest, $3.LogoutResponse>(
        'logout',
        logout_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.LogoutRequest.fromBuffer(value),
        ($3.LogoutResponse value) => value.writeToBuffer()));
  }

  $async.Future<$3.LogoutResponse> logout_Pre(
      $grpc.ServiceCall call, $async.Future<$3.LogoutRequest> request) async {
    return logout(call, await request);
  }

  $async.Future<$3.LogoutResponse> logout(
      $grpc.ServiceCall call, $3.LogoutRequest request);
}
