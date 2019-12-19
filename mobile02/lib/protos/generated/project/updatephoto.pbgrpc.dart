///
//  Generated code. Do not modify.
//  source: updatephoto.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'updatephoto.pb.dart' as $7;
export 'updatephoto.pb.dart';

class UpdatePhotoClient extends $grpc.Client {
  static final _$photo =
      $grpc.ClientMethod<$7.UpdatePhotoRequest, $7.UpdatePhotoResponse>(
          '/protos.updatephoto.UpdatePhoto/photo',
          ($7.UpdatePhotoRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $7.UpdatePhotoResponse.fromBuffer(value));

  UpdatePhotoClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$7.UpdatePhotoResponse> photo(
      $7.UpdatePhotoRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$photo, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class UpdatePhotoServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.updatephoto.UpdatePhoto';

  UpdatePhotoServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$7.UpdatePhotoRequest, $7.UpdatePhotoResponse>(
            'photo',
            photo_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $7.UpdatePhotoRequest.fromBuffer(value),
            ($7.UpdatePhotoResponse value) => value.writeToBuffer()));
  }

  $async.Future<$7.UpdatePhotoResponse> photo_Pre($grpc.ServiceCall call,
      $async.Future<$7.UpdatePhotoRequest> request) async {
    return photo(call, await request);
  }

  $async.Future<$7.UpdatePhotoResponse> photo(
      $grpc.ServiceCall call, $7.UpdatePhotoRequest request);
}
