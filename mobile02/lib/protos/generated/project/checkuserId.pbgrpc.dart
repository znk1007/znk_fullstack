///
//  Generated code. Do not modify.
//  source: checkuserId.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'checkuserId.pb.dart' as $1;
export 'checkuserId.pb.dart';

class CheckUserIdClient extends $grpc.Client {
  static final _$check =
      $grpc.ClientMethod<$1.CheckUserIdRequest, $1.CheckUserIdResponse>(
          '/protos.checkuserId.CheckUserId/check',
          ($1.CheckUserIdRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.CheckUserIdResponse.fromBuffer(value));

  CheckUserIdClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.CheckUserIdResponse> check(
      $1.CheckUserIdRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$check, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class CheckUserIdServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.checkuserId.CheckUserId';

  CheckUserIdServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$1.CheckUserIdRequest, $1.CheckUserIdResponse>(
            'check',
            check_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.CheckUserIdRequest.fromBuffer(value),
            ($1.CheckUserIdResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.CheckUserIdResponse> check_Pre($grpc.ServiceCall call,
      $async.Future<$1.CheckUserIdRequest> request) async {
    return check(call, await request);
  }

  $async.Future<$1.CheckUserIdResponse> check(
      $grpc.ServiceCall call, $1.CheckUserIdRequest request);
}
