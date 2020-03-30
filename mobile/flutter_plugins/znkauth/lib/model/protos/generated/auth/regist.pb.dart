///
//  Generated code. Do not modify.
//  source: regist.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class RegistReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegistReq', package: const $pb.PackageName('regist'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'token')
    ..hasRequiredFields = false
  ;

  RegistReq._() : super();
  factory RegistReq() => create();
  factory RegistReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegistReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegistReq clone() => RegistReq()..mergeFromMessage(this);
  RegistReq copyWith(void Function(RegistReq) updates) => super.copyWith((message) => updates(message as RegistReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegistReq create() => RegistReq._();
  RegistReq createEmptyInstance() => create();
  static $pb.PbList<RegistReq> createRepeated() => $pb.PbList<RegistReq>();
  @$core.pragma('dart2js:noInline')
  static RegistReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegistReq>(create);
  static RegistReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);
}

class RegistRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RegistRes', package: const $pb.PackageName('regist'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'token')
    ..hasRequiredFields = false
  ;

  RegistRes._() : super();
  factory RegistRes() => create();
  factory RegistRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegistRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RegistRes clone() => RegistRes()..mergeFromMessage(this);
  RegistRes copyWith(void Function(RegistRes) updates) => super.copyWith((message) => updates(message as RegistRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegistRes create() => RegistRes._();
  RegistRes createEmptyInstance() => create();
  static $pb.PbList<RegistRes> createRepeated() => $pb.PbList<RegistRes>();
  @$core.pragma('dart2js:noInline')
  static RegistRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegistRes>(create);
  static RegistRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);
}

