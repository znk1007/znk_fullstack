///
//  Generated code. Do not modify.
//  source: regist.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'regist.pb.dart' as $1;
export 'regist.pb.dart';

class RegistSrvClient extends $grpc.Client {
  static final _$userReigst = $grpc.ClientMethod<$1.RegistReq, $1.RegistRes>(
      '/regist.RegistSrv/userReigst',
      ($1.RegistReq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.RegistRes.fromBuffer(value));

  RegistSrvClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.RegistRes> userReigst($1.RegistReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$userReigst, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class RegistSrvServiceBase extends $grpc.Service {
  $core.String get $name => 'regist.RegistSrv';

  RegistSrvServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.RegistReq, $1.RegistRes>(
        'userReigst',
        userReigst_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.RegistReq.fromBuffer(value),
        ($1.RegistRes value) => value.writeToBuffer()));
  }

  $async.Future<$1.RegistRes> userReigst_Pre(
      $grpc.ServiceCall call, $async.Future<$1.RegistReq> request) async {
    return userReigst(call, await request);
  }

  $async.Future<$1.RegistRes> userReigst(
      $grpc.ServiceCall call, $1.RegistReq request);
}
