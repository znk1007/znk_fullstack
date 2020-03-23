///
//  Generated code. Do not modify.
//  source: login.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class LoginReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LoginReq', package: const $pb.PackageName('login'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'sign')
    ..hasRequiredFields = false
  ;

  LoginReq._() : super();
  factory LoginReq() => create();
  factory LoginReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LoginReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LoginReq clone() => LoginReq()..mergeFromMessage(this);
  LoginReq copyWith(void Function(LoginReq) updates) => super.copyWith((message) => updates(message as LoginReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LoginReq create() => LoginReq._();
  LoginReq createEmptyInstance() => create();
  static $pb.PbList<LoginReq> createRepeated() => $pb.PbList<LoginReq>();
  @$core.pragma('dart2js:noInline')
  static LoginReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LoginReq>(create);
  static LoginReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get sign => $_getSZ(1);
  @$pb.TagNumber(2)
  set sign($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSign() => $_has(1);
  @$pb.TagNumber(2)
  void clearSign() => clearField(2);
}

class LoginRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LoginRes', package: const $pb.PackageName('login'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'sign')
    ..hasRequiredFields = false
  ;

  LoginRes._() : super();
  factory LoginRes() => create();
  factory LoginRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LoginRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LoginRes clone() => LoginRes()..mergeFromMessage(this);
  LoginRes copyWith(void Function(LoginRes) updates) => super.copyWith((message) => updates(message as LoginRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LoginRes create() => LoginRes._();
  LoginRes createEmptyInstance() => create();
  static $pb.PbList<LoginRes> createRepeated() => $pb.PbList<LoginRes>();
  @$core.pragma('dart2js:noInline')
  static LoginRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LoginRes>(create);
  static LoginRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get sign => $_getSZ(1);
  @$pb.TagNumber(2)
  set sign($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSign() => $_has(1);
  @$pb.TagNumber(2)
  void clearSign() => clearField(2);
}

