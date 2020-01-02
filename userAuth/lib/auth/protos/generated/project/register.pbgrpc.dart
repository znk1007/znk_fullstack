///
//  Generated code. Do not modify.
//  source: register.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'register.pb.dart' as $4;
export 'register.pb.dart';

class RegisterClient extends $grpc.Client {
  static final _$regist =
      $grpc.ClientMethod<$4.RegistRequest, $4.RegistResponse>(
          '/protos.register.Register/Regist',
          ($4.RegistRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $4.RegistResponse.fromBuffer(value));

  RegisterClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$4.RegistResponse> regist($4.RegistRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$regist, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class RegisterServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.register.Register';

  RegisterServiceBase() {
    $addMethod($grpc.ServiceMethod<$4.RegistRequest, $4.RegistResponse>(
        'Regist',
        regist_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.RegistRequest.fromBuffer(value),
        ($4.RegistResponse value) => value.writeToBuffer()));
  }

  $async.Future<$4.RegistResponse> regist_Pre(
      $grpc.ServiceCall call, $async.Future<$4.RegistRequest> request) async {
    return regist(call, await request);
  }

  $async.Future<$4.RegistResponse> regist(
      $grpc.ServiceCall call, $4.RegistRequest request);
}
