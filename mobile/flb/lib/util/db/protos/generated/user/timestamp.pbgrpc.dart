///
//  Generated code. Do not modify.
//  source: timestamp.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'timestamp.pb.dart' as $1;
export 'timestamp.pb.dart';

class TimestampSrvClient extends $grpc.Client {
  static final _$userTimestamp = $grpc.ClientMethod<$1.TSReq, $1.TSReq>(
      '/timestamp.TimestampSrv/userTimestamp',
      ($1.TSReq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.TSReq.fromBuffer(value));

  TimestampSrvClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.TSReq> userTimestamp($1.TSReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$userTimestamp, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class TimestampSrvServiceBase extends $grpc.Service {
  $core.String get $name => 'timestamp.TimestampSrv';

  TimestampSrvServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.TSReq, $1.TSReq>(
        'userTimestamp',
        userTimestamp_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.TSReq.fromBuffer(value),
        ($1.TSReq value) => value.writeToBuffer()));
  }

  $async.Future<$1.TSReq> userTimestamp_Pre(
      $grpc.ServiceCall call, $async.Future<$1.TSReq> request) async {
    return userTimestamp(call, await request);
  }

  $async.Future<$1.TSReq> userTimestamp(
      $grpc.ServiceCall call, $1.TSReq request);
}