///
//  Generated code. Do not modify.
//  source: checksess.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'checksess.pb.dart' as $0;
export 'checksess.pb.dart';

class CheckSessionClient extends $grpc.Client {
  static final _$checkSession =
      $grpc.ClientMethod<$0.CheckSessionRequest, $0.CheckSessionResponse>(
          '/proto.checksess.CheckSession/checkSession',
          ($0.CheckSessionRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.CheckSessionResponse.fromBuffer(value));

  CheckSessionClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.CheckSessionResponse> checkSession(
      $0.CheckSessionRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$checkSession, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class CheckSessionServiceBase extends $grpc.Service {
  $core.String get $name => 'proto.checksess.CheckSession';

  CheckSessionServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$0.CheckSessionRequest, $0.CheckSessionResponse>(
            'checkSession',
            checkSession_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.CheckSessionRequest.fromBuffer(value),
            ($0.CheckSessionResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.CheckSessionResponse> checkSession_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.CheckSessionRequest> request) async {
    return checkSession(call, await request);
  }

  $async.Future<$0.CheckSessionResponse> checkSession(
      $grpc.ServiceCall call, $0.CheckSessionRequest request);
}
