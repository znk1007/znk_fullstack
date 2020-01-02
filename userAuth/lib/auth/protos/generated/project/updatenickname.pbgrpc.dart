///
//  Generated code. Do not modify.
//  source: updatenickname.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'updatenickname.pb.dart' as $5;
export 'updatenickname.pb.dart';

class UpdateNicknameClient extends $grpc.Client {
  static final _$nickname =
      $grpc.ClientMethod<$5.UpdateNicknameRequest, $5.UpdateNicknameResponse>(
          '/protos.updatnickname.UpdateNickname/nickname',
          ($5.UpdateNicknameRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $5.UpdateNicknameResponse.fromBuffer(value));

  UpdateNicknameClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$5.UpdateNicknameResponse> nickname(
      $5.UpdateNicknameRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$nickname, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class UpdateNicknameServiceBase extends $grpc.Service {
  $core.String get $name => 'protos.updatnickname.UpdateNickname';

  UpdateNicknameServiceBase() {
    $addMethod($grpc.ServiceMethod<$5.UpdateNicknameRequest,
            $5.UpdateNicknameResponse>(
        'nickname',
        nickname_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $5.UpdateNicknameRequest.fromBuffer(value),
        ($5.UpdateNicknameResponse value) => value.writeToBuffer()));
  }

  $async.Future<$5.UpdateNicknameResponse> nickname_Pre($grpc.ServiceCall call,
      $async.Future<$5.UpdateNicknameRequest> request) async {
    return nickname(call, await request);
  }

  $async.Future<$5.UpdateNicknameResponse> nickname(
      $grpc.ServiceCall call, $5.UpdateNicknameRequest request);
}
