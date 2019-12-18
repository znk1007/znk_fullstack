///
//  Generated code. Do not modify.
//  source: checksess.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:grpc/service_api.dart' as $grpc;

import 'dart:core' as $core show int, String, List;

import 'checksess.pb.dart' as $0;
import 'checksess.pb.dart';
export 'checksess.pb.dart';

class CheckSessionClient extends $grpc.Client {
  static final _$checkSession =
      $grpc.ClientMethod<CheckSessionRequest, CheckSessionResponse>(
          '/proto.checksess.CheckSession/checkSession',
          (CheckSessionRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              CheckSessionResponse.fromBuffer(value));

  CheckSessionClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<CheckSessionResponse> checkSession(
      CheckSessionRequest request,
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
    $addMethod($grpc.ServiceMethod<CheckSessionRequest, CheckSessionResponse>(
        'checkSession',
        checkSession_Pre,
        false,
        false,
        ($core.List<$core.int> value) => CheckSessionRequest.fromBuffer(value),
        (CheckSessionResponse value) => value.writeToBuffer()));
  }

  $async.Future<CheckSessionResponse> checkSession_Pre(
      $grpc.ServiceCall call, $async.Future request) async {
    return checkSession(call, await request);
  }

  $async.Future<CheckSessionResponse> checkSession(
      $grpc.ServiceCall call, CheckSessionRequest request);
}
