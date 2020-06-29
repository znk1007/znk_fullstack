///
//  Generated code. Do not modify.
//  source: device.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class DvsUpdateReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DvsUpdateReq', package: const $pb.PackageName('device'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  DvsUpdateReq._() : super();
  factory DvsUpdateReq() => create();
  factory DvsUpdateReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DvsUpdateReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DvsUpdateReq clone() => DvsUpdateReq()..mergeFromMessage(this);
  DvsUpdateReq copyWith(void Function(DvsUpdateReq) updates) => super.copyWith((message) => updates(message as DvsUpdateReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DvsUpdateReq create() => DvsUpdateReq._();
  DvsUpdateReq createEmptyInstance() => create();
  static $pb.PbList<DvsUpdateReq> createRepeated() => $pb.PbList<DvsUpdateReq>();
  @$core.pragma('dart2js:noInline')
  static DvsUpdateReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DvsUpdateReq>(create);
  static DvsUpdateReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class DvsUpdateRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DvsUpdateRes', package: const $pb.PackageName('device'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  DvsUpdateRes._() : super();
  factory DvsUpdateRes() => create();
  factory DvsUpdateRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DvsUpdateRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DvsUpdateRes clone() => DvsUpdateRes()..mergeFromMessage(this);
  DvsUpdateRes copyWith(void Function(DvsUpdateRes) updates) => super.copyWith((message) => updates(message as DvsUpdateRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DvsUpdateRes create() => DvsUpdateRes._();
  DvsUpdateRes createEmptyInstance() => create();
  static $pb.PbList<DvsUpdateRes> createRepeated() => $pb.PbList<DvsUpdateRes>();
  @$core.pragma('dart2js:noInline')
  static DvsUpdateRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DvsUpdateRes>(create);
  static DvsUpdateRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class DvsDeleteReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DvsDeleteReq', package: const $pb.PackageName('device'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  DvsDeleteReq._() : super();
  factory DvsDeleteReq() => create();
  factory DvsDeleteReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DvsDeleteReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DvsDeleteReq clone() => DvsDeleteReq()..mergeFromMessage(this);
  DvsDeleteReq copyWith(void Function(DvsDeleteReq) updates) => super.copyWith((message) => updates(message as DvsDeleteReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DvsDeleteReq create() => DvsDeleteReq._();
  DvsDeleteReq createEmptyInstance() => create();
  static $pb.PbList<DvsDeleteReq> createRepeated() => $pb.PbList<DvsDeleteReq>();
  @$core.pragma('dart2js:noInline')
  static DvsDeleteReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DvsDeleteReq>(create);
  static DvsDeleteReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

class DvsDeleteRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DvsDeleteRes', package: const $pb.PackageName('device'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'data')
    ..hasRequiredFields = false
  ;

  DvsDeleteRes._() : super();
  factory DvsDeleteRes() => create();
  factory DvsDeleteRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DvsDeleteRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  DvsDeleteRes clone() => DvsDeleteRes()..mergeFromMessage(this);
  DvsDeleteRes copyWith(void Function(DvsDeleteRes) updates) => super.copyWith((message) => updates(message as DvsDeleteRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DvsDeleteRes create() => DvsDeleteRes._();
  DvsDeleteRes createEmptyInstance() => create();
  static $pb.PbList<DvsDeleteRes> createRepeated() => $pb.PbList<DvsDeleteRes>();
  @$core.pragma('dart2js:noInline')
  static DvsDeleteRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DvsDeleteRes>(create);
  static DvsDeleteRes _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get data => $_getSZ(1);
  @$pb.TagNumber(2)
  set data($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

