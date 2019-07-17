///
//  Generated code. Do not modify.
//  source: login.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

import 'user.pb.dart' as $2;

class LoginRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LoginRequest', package: const $pb.PackageName('protos.login'))
    ..aOS(1, 'userId')
    ..aOS(2, 'account')
    ..aOS(3, 'password')
    ..aOS(4, 'device')
    ..hasRequiredFields = false
  ;

  LoginRequest() : super();
  LoginRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  LoginRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  LoginRequest clone() => LoginRequest()..mergeFromMessage(this);
  LoginRequest copyWith(void Function(LoginRequest) updates) => super.copyWith((message) => updates(message as LoginRequest));
  $pb.BuilderInfo get info_ => _i;
  static LoginRequest create() => LoginRequest();
  LoginRequest createEmptyInstance() => create();
  static $pb.PbList<LoginRequest> createRepeated() => $pb.PbList<LoginRequest>();
  static LoginRequest getDefault() => _defaultInstance ??= create()..freeze();
  static LoginRequest _defaultInstance;

  $core.String get userId => $_getS(0, '');
  set userId($core.String v) { $_setString(0, v); }
  $core.bool hasUserId() => $_has(0);
  void clearUserId() => clearField(1);

  $core.String get account => $_getS(1, '');
  set account($core.String v) { $_setString(1, v); }
  $core.bool hasAccount() => $_has(1);
  void clearAccount() => clearField(2);

  $core.String get password => $_getS(2, '');
  set password($core.String v) { $_setString(2, v); }
  $core.bool hasPassword() => $_has(2);
  void clearPassword() => clearField(3);

  $core.String get device => $_getS(3, '');
  set device($core.String v) { $_setString(3, v); }
  $core.bool hasDevice() => $_has(3);
  void clearDevice() => clearField(4);
}

class LoginResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LoginResponse', package: const $pb.PackageName('protos.login'))
    ..aOS(2, 'message')
    ..a<$core.int>(3, 'code', $pb.PbFieldType.O3)
    ..a<$2.User>(4, 'user', $pb.PbFieldType.OM, $2.User.getDefault, $2.User.create)
    ..a<$core.int>(5, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  LoginResponse() : super();
  LoginResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  LoginResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  LoginResponse clone() => LoginResponse()..mergeFromMessage(this);
  LoginResponse copyWith(void Function(LoginResponse) updates) => super.copyWith((message) => updates(message as LoginResponse));
  $pb.BuilderInfo get info_ => _i;
  static LoginResponse create() => LoginResponse();
  LoginResponse createEmptyInstance() => create();
  static $pb.PbList<LoginResponse> createRepeated() => $pb.PbList<LoginResponse>();
  static LoginResponse getDefault() => _defaultInstance ??= create()..freeze();
  static LoginResponse _defaultInstance;

  $core.String get message => $_getS(0, '');
  set message($core.String v) { $_setString(0, v); }
  $core.bool hasMessage() => $_has(0);
  void clearMessage() => clearField(2);

  $core.int get code => $_get(1, 0);
  set code($core.int v) { $_setSignedInt32(1, v); }
  $core.bool hasCode() => $_has(1);
  void clearCode() => clearField(3);

  $2.User get user => $_getN(2);
  set user($2.User v) { setField(4, v); }
  $core.bool hasUser() => $_has(2);
  void clearUser() => clearField(4);

  $core.int get status => $_get(3, 0);
  set status($core.int v) { $_setSignedInt32(3, v); }
  $core.bool hasStatus() => $_has(3);
  void clearStatus() => clearField(5);
}

