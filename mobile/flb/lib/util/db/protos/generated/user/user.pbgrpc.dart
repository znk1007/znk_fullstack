///
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'user.pb.dart' as $2;
export 'user.pb.dart';

class UserSrvClient extends $grpc.Client {
  static final _$regist = $grpc.ClientMethod<$2.UserRgstReq, $2.UserRgstRes>(
      '/user.UserSrv/Regist',
      ($2.UserRgstReq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.UserRgstRes.fromBuffer(value));
  static final _$login = $grpc.ClientMethod<$2.UserLgnReq, $2.UserLgnRes>(
      '/user.UserSrv/Login',
      ($2.UserLgnReq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.UserLgnRes.fromBuffer(value));
  static final _$updatePassword =
      $grpc.ClientMethod<$2.UserUpdatePswReq, $2.UserUpdatePswRes>(
          '/user.UserSrv/UpdatePassword',
          ($2.UserUpdatePswReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $2.UserUpdatePswRes.fromBuffer(value));
  static final _$logout = $grpc.ClientMethod<$2.UserLgoReq, $2.UserLgoRes>(
      '/user.UserSrv/Logout',
      ($2.UserLgoReq value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.UserLgoRes.fromBuffer(value));
  static final _$status =
      $grpc.ClientMethod<$2.UserStatusReq, $2.UserStatusRes>(
          '/user.UserSrv/Status',
          ($2.UserStatusReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.UserStatusRes.fromBuffer(value));
  static final _$updatePhone =
      $grpc.ClientMethod<$2.UserUpdatePhoneReq, $2.UserUpdatePhoneRes>(
          '/user.UserSrv/UpdatePhone',
          ($2.UserUpdatePhoneReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $2.UserUpdatePhoneRes.fromBuffer(value));
  static final _$updateNickname =
      $grpc.ClientMethod<$2.UserUpdateNicknameReq, $2.UserUpdateNicknameRes>(
          '/user.UserSrv/UpdateNickname',
          ($2.UserUpdateNicknameReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $2.UserUpdateNicknameRes.fromBuffer(value));
  static final _$activeUser =
      $grpc.ClientMethod<$2.UserUpdateActiveReq, $2.UserUpdateActiveRes>(
          '/user.UserSrv/ActiveUser',
          ($2.UserUpdateActiveReq value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $2.UserUpdateActiveRes.fromBuffer(value));

  UserSrvClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$2.UserRgstRes> regist($2.UserRgstReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$regist, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.UserLgnRes> login($2.UserLgnReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$login, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.UserUpdatePswRes> updatePassword(
      $2.UserUpdatePswReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$updatePassword, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.UserLgoRes> logout($2.UserLgoReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$logout, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.UserStatusRes> status($2.UserStatusReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$status, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.UserUpdatePhoneRes> updatePhone(
      $2.UserUpdatePhoneReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$updatePhone, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.UserUpdateNicknameRes> updateNickname(
      $2.UserUpdateNicknameReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$updateNickname, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.UserUpdateActiveRes> activeUser(
      $2.UserUpdateActiveReq request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$activeUser, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class UserSrvServiceBase extends $grpc.Service {
  $core.String get $name => 'user.UserSrv';

  UserSrvServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.UserRgstReq, $2.UserRgstRes>(
        'Regist',
        regist_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.UserRgstReq.fromBuffer(value),
        ($2.UserRgstRes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.UserLgnReq, $2.UserLgnRes>(
        'Login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.UserLgnReq.fromBuffer(value),
        ($2.UserLgnRes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.UserUpdatePswReq, $2.UserUpdatePswRes>(
        'UpdatePassword',
        updatePassword_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.UserUpdatePswReq.fromBuffer(value),
        ($2.UserUpdatePswRes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.UserLgoReq, $2.UserLgoRes>(
        'Logout',
        logout_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.UserLgoReq.fromBuffer(value),
        ($2.UserLgoRes value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.UserStatusReq, $2.UserStatusRes>(
        'Status',
        status_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.UserStatusReq.fromBuffer(value),
        ($2.UserStatusRes value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$2.UserUpdatePhoneReq, $2.UserUpdatePhoneRes>(
            'UpdatePhone',
            updatePhone_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $2.UserUpdatePhoneReq.fromBuffer(value),
            ($2.UserUpdatePhoneRes value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$2.UserUpdateNicknameReq, $2.UserUpdateNicknameRes>(
            'UpdateNickname',
            updateNickname_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $2.UserUpdateNicknameReq.fromBuffer(value),
            ($2.UserUpdateNicknameRes value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$2.UserUpdateActiveReq, $2.UserUpdateActiveRes>(
            'ActiveUser',
            activeUser_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $2.UserUpdateActiveReq.fromBuffer(value),
            ($2.UserUpdateActiveRes value) => value.writeToBuffer()));
  }

  $async.Future<$2.UserRgstRes> regist_Pre(
      $grpc.ServiceCall call, $async.Future<$2.UserRgstReq> request) async {
    return regist(call, await request);
  }

  $async.Future<$2.UserLgnRes> login_Pre(
      $grpc.ServiceCall call, $async.Future<$2.UserLgnReq> request) async {
    return login(call, await request);
  }

  $async.Future<$2.UserUpdatePswRes> updatePassword_Pre($grpc.ServiceCall call,
      $async.Future<$2.UserUpdatePswReq> request) async {
    return updatePassword(call, await request);
  }

  $async.Future<$2.UserLgoRes> logout_Pre(
      $grpc.ServiceCall call, $async.Future<$2.UserLgoReq> request) async {
    return logout(call, await request);
  }

  $async.Future<$2.UserStatusRes> status_Pre(
      $grpc.ServiceCall call, $async.Future<$2.UserStatusReq> request) async {
    return status(call, await request);
  }

  $async.Future<$2.UserUpdatePhoneRes> updatePhone_Pre($grpc.ServiceCall call,
      $async.Future<$2.UserUpdatePhoneReq> request) async {
    return updatePhone(call, await request);
  }

  $async.Future<$2.UserUpdateNicknameRes> updateNickname_Pre(
      $grpc.ServiceCall call,
      $async.Future<$2.UserUpdateNicknameReq> request) async {
    return updateNickname(call, await request);
  }

  $async.Future<$2.UserUpdateActiveRes> activeUser_Pre($grpc.ServiceCall call,
      $async.Future<$2.UserUpdateActiveReq> request) async {
    return activeUser(call, await request);
  }

  $async.Future<$2.UserRgstRes> regist(
      $grpc.ServiceCall call, $2.UserRgstReq request);
  $async.Future<$2.UserLgnRes> login(
      $grpc.ServiceCall call, $2.UserLgnReq request);
  $async.Future<$2.UserUpdatePswRes> updatePassword(
      $grpc.ServiceCall call, $2.UserUpdatePswReq request);
  $async.Future<$2.UserLgoRes> logout(
      $grpc.ServiceCall call, $2.UserLgoReq request);
  $async.Future<$2.UserStatusRes> status(
      $grpc.ServiceCall call, $2.UserStatusReq request);
  $async.Future<$2.UserUpdatePhoneRes> updatePhone(
      $grpc.ServiceCall call, $2.UserUpdatePhoneReq request);
  $async.Future<$2.UserUpdateNicknameRes> updateNickname(
      $grpc.ServiceCall call, $2.UserUpdateNicknameReq request);
  $async.Future<$2.UserUpdateActiveRes> activeUser(
      $grpc.ServiceCall call, $2.UserUpdateActiveReq request);
}
