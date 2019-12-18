///
//  Generated code. Do not modify.
//  source: register.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:grpc/service_api.dart' as $grpc;

import 'dart:core' as $core show int, String, List;

import 'register.pb.dart' as $5;
import 'register.pb.dart';
export 'register.pb.dart';

class RegisterClient extends $grpc.Client {
  static final _$regist = $grpc.ClientMethod<RegistRequest, RegistResponse>(
      '/protos.register.Register/Regist',
      (RegistRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => RegistResponse.fromBuffer(value));

  RegisterClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<RegistResponse> regist(RegistRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$regist, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class RegisterServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.register.Register';

  RegisterServiceBase() {
    $addMethod($grpc.ServiceMethod<RegistRequest, RegistResponse>(
        'Regist',
        regist_Pre,
        false,
        false,
        ($core.List<$core.int> value) => RegistRequest.fromBuffer(value),
        (RegistResponse value) => value.writeToBuffer()));
  }

  $async.Future<RegistResponse> regist_Pre(
      $grpc.ServiceCall call, $async.Future request) async {
    return regist(call, await request);
  }

  $async.Future<RegistResponse> regist(
      $grpc.ServiceCall call, RegistRequest request);
}
