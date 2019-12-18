///
//  Generated code. Do not modify.
//  source: updateonline.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:grpc/service_api.dart' as $grpc;

import 'dart:core' as $core show int, String, List;

import 'updateonline.pb.dart' as $7;
import 'updateonline.pb.dart';
export 'updateonline.pb.dart';

class UpdateOnlineClient extends $grpc.Client {
  static final _$update =
      $grpc.ClientMethod<UpdateOnlineRequest, UpdateOnlineResponse>(
          '/protos.updateonline.UpdateOnline/update',
          (UpdateOnlineRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              UpdateOnlineResponse.fromBuffer(value));

  UpdateOnlineClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<UpdateOnlineResponse> update(UpdateOnlineRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$update, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class UpdateOnlineServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.updateonline.UpdateOnline';

  UpdateOnlineServiceBase() {
    $addMethod($grpc.ServiceMethod<UpdateOnlineRequest, UpdateOnlineResponse>(
        'update',
        update_Pre,
        false,
        false,
        ($core.List<$core.int> value) => UpdateOnlineRequest.fromBuffer(value),
        (UpdateOnlineResponse value) => value.writeToBuffer()));
  }

  $async.Future<UpdateOnlineResponse> update_Pre(
      $grpc.ServiceCall call, $async.Future request) async {
    return update(call, await request);
  }

  $async.Future<UpdateOnlineResponse> update(
      $grpc.ServiceCall call, UpdateOnlineRequest request);
}
