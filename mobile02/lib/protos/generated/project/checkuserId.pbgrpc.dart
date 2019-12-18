///
//  Generated code. Do not modify.
//  source: checkuserId.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:grpc/service_api.dart' as $grpc;

import 'dart:core' as $core show int, String, List;

import 'checkuserId.pb.dart' as $1;
import 'checkuserId.pb.dart';
export 'checkuserId.pb.dart';

class CheckUserIdClient extends $grpc.Client {
  static final _$check =
      $grpc.ClientMethod<CheckUserIdRequest, CheckUserIdResponse>(
          '/protos.checkuserId.CheckUserId/check',
          (CheckUserIdRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              CheckUserIdResponse.fromBuffer(value));

  CheckUserIdClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<CheckUserIdResponse> check(CheckUserIdRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$check, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class CheckUserIdServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.checkuserId.CheckUserId';

  CheckUserIdServiceBase() {
    $addMethod($grpc.ServiceMethod<CheckUserIdRequest, CheckUserIdResponse>(
        'check',
        check_Pre,
        false,
        false,
        ($core.List<$core.int> value) => CheckUserIdRequest.fromBuffer(value),
        (CheckUserIdResponse value) => value.writeToBuffer()));
  }

  $async.Future<CheckUserIdResponse> check_Pre(
      $grpc.ServiceCall call, $async.Future request) async {
    return check(call, await request);
  }

  $async.Future<CheckUserIdResponse> check(
      $grpc.ServiceCall call, CheckUserIdRequest request);
}
