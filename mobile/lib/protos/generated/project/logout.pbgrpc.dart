///
//  Generated code. Do not modify.
//  source: logout.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:grpc/service_api.dart' as $grpc;

import 'dart:core' as $core show int, String, List;

import 'logout.pb.dart' as $4;
import 'logout.pb.dart';
export 'logout.pb.dart';

class LogoutServiceClient extends $grpc.Client {
  static final _$logout = $grpc.ClientMethod<LogoutRequest, LogoutResponse>(
      '/protos.logout.LogoutService/logout',
      (LogoutRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => LogoutResponse.fromBuffer(value));

  LogoutServiceClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<LogoutResponse> logout(LogoutRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$logout, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class LogoutServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.logout.LogoutService';

  LogoutServiceBase() {
    $addMethod($grpc.ServiceMethod<LogoutRequest, LogoutResponse>(
        'logout',
        logout_Pre,
        false,
        false,
        ($core.List<$core.int> value) => LogoutRequest.fromBuffer(value),
        (LogoutResponse value) => value.writeToBuffer()));
  }

  $async.Future<LogoutResponse> logout_Pre(
      $grpc.ServiceCall call, $async.Future request) async {
    return logout(call, await request);
  }

  $async.Future<LogoutResponse> logout(
      $grpc.ServiceCall call, LogoutRequest request);
}
