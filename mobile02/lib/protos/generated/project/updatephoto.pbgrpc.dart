///
//  Generated code. Do not modify.
//  source: updatephoto.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:grpc/service_api.dart' as $grpc;

import 'dart:core' as $core show int, String, List;

import 'updatephoto.pb.dart' as $8;
import 'updatephoto.pb.dart';
export 'updatephoto.pb.dart';

class UpdatePhotoClient extends $grpc.Client {
  static final _$photo =
      $grpc.ClientMethod<UpdatePhotoRequest, UpdatePhotoResponse>(
          '/protos.updatephoto.UpdatePhoto/photo',
          (UpdatePhotoRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              UpdatePhotoResponse.fromBuffer(value));

  UpdatePhotoClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<UpdatePhotoResponse> photo(UpdatePhotoRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$photo, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class UpdatePhotoServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.updatephoto.UpdatePhoto';

  UpdatePhotoServiceBase() {
    $addMethod($grpc.ServiceMethod<UpdatePhotoRequest, UpdatePhotoResponse>(
        'photo',
        photo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => UpdatePhotoRequest.fromBuffer(value),
        (UpdatePhotoResponse value) => value.writeToBuffer()));
  }

  $async.Future<UpdatePhotoResponse> photo_Pre(
      $grpc.ServiceCall call, $async.Future request) async {
    return photo(call, await request);
  }

  $async.Future<UpdatePhotoResponse> photo(
      $grpc.ServiceCall call, UpdatePhotoRequest request);
}
