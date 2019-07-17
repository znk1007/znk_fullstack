///
//  Generated code. Do not modify.
//  source: updatenickname.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:grpc/service_api.dart' as $grpc;

import 'dart:core' as $core show int, String, List;

import 'updatenickname.pb.dart' as $6;
import 'updatenickname.pb.dart';
export 'updatenickname.pb.dart';

class UpdateNicknameClient extends $grpc.Client {
  static final _$nickname =
      $grpc.ClientMethod<UpdateNicknameRequest, UpdateNicknameResponse>(
          '/protos.updatnickname.UpdateNickname/nickname',
          (UpdateNicknameRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              UpdateNicknameResponse.fromBuffer(value));

  UpdateNicknameClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<UpdateNicknameResponse> nickname(
      UpdateNicknameRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$nickname, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class UpdateNicknameServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.updatnickname.UpdateNickname';

  UpdateNicknameServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<UpdateNicknameRequest, UpdateNicknameResponse>(
            'nickname',
            nickname_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                UpdateNicknameRequest.fromBuffer(value),
            (UpdateNicknameResponse value) => value.writeToBuffer()));
  }

  $async.Future<UpdateNicknameResponse> nickname_Pre(
      $grpc.ServiceCall call, $async.Future request) async {
    return nickname(call, await request);
  }

  $async.Future<UpdateNicknameResponse> nickname(
      $grpc.ServiceCall call, UpdateNicknameRequest request);
}
