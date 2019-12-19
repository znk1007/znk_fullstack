///
//  Generated code. Do not modify.
//  source: updateonline.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'updateonline.pb.dart' as $6;
export 'updateonline.pb.dart';

class UpdateOnlineClient extends $grpc.Client {
  static final _$update =
      $grpc.ClientMethod<$6.UpdateOnlineRequest, $6.UpdateOnlineResponse>(
          '/protos.updateonline.UpdateOnline/update',
          ($6.UpdateOnlineRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $6.UpdateOnlineResponse.fromBuffer(value));

  UpdateOnlineClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$6.UpdateOnlineResponse> update(
      $6.UpdateOnlineRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$update, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class UpdateOnlineServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.updateonline.UpdateOnline';

  UpdateOnlineServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$6.UpdateOnlineRequest, $6.UpdateOnlineResponse>(
            'update',
            update_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $6.UpdateOnlineRequest.fromBuffer(value),
            ($6.UpdateOnlineResponse value) => value.writeToBuffer()));
  }

  $async.Future<$6.UpdateOnlineResponse> update_Pre($grpc.ServiceCall call,
      $async.Future<$6.UpdateOnlineRequest> request) async {
    return update(call, await request);
  }

  $async.Future<$6.UpdateOnlineResponse> update(
      $grpc.ServiceCall call, $6.UpdateOnlineRequest request);
}
