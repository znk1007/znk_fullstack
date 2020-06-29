///
//  Generated code. Do not modify.
//  source: device.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'device.pb.dart' as $0;
export 'device.pb.dart';

class DvsSrvClient extends $grpc.Client {
  static final _$updateDevice =
      $grpc.ClientMethod<$0.DvsUpdateReq, $0.DvsUpdateRes>(
          '/device.DvsSrv/UpdateDevice',
          ($0.DvsUpdateReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.DvsUpdateRes.fromBuffer(value));
  static final _$deleteDevice =
      $grpc.ClientMethod<$0.DvsDeleteReq, $0.DvsDeleteRes>(
          '/device.DvsSrv/DeleteDevice',
          ($0.DvsDeleteReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.DvsDeleteRes.fromBuffer(value));

  DvsSrvClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.DvsUpdateRes> updateDevice($0.DvsUpdateReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$updateDevice, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.DvsDeleteRes> deleteDevice($0.DvsDeleteReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$deleteDevice, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class DvsSrvServiceBase extends $grpc.Service {
  $core.String get $name => 'device.DvsSrv';

  DvsSrvServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.DvsUpdateReq, $0.DvsUpdateRes>(
        'UpdateDevice',
        updateDevice_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DvsUpdateReq.fromBuffer(value),
        ($0.DvsUpdateRes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DvsDeleteReq, $0.DvsDeleteRes>(
        'DeleteDevice',
        deleteDevice_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DvsDeleteReq.fromBuffer(value),
        ($0.DvsDeleteRes value) => value.writeToBuffer()));
  }

  $async.Future<$0.DvsUpdateRes> updateDevice_Pre(
      $grpc.ServiceCall call, $async.Future<$0.DvsUpdateReq> request) async {
    return updateDevice(call, await request);
  }

  $async.Future<$0.DvsDeleteRes> deleteDevice_Pre(
      $grpc.ServiceCall call, $async.Future<$0.DvsDeleteReq> request) async {
    return deleteDevice(call, await request);
  }

  $async.Future<$0.DvsUpdateRes> updateDevice(
      $grpc.ServiceCall call, $0.DvsUpdateReq request);
  $async.Future<$0.DvsDeleteRes> deleteDevice(
      $grpc.ServiceCall call, $0.DvsDeleteReq request);
}
