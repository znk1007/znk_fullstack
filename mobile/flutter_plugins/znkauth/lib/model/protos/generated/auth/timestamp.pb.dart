///
//  Generated code. Do not modify.
//  source: timestamp.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class TSReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('TSReq', package: const $pb.PackageName('timestamp'), createEmptyInstance: create)
    ..aOS(1, 'platform')
    ..hasRequiredFields = false
  ;

  TSReq._() : super();
  factory TSReq() => create();
  factory TSReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TSReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  TSReq clone() => TSReq()..mergeFromMessage(this);
  TSReq copyWith(void Function(TSReq) updates) => super.copyWith((message) => updates(message as TSReq));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TSReq create() => TSReq._();
  TSReq createEmptyInstance() => create();
  static $pb.PbList<TSReq> createRepeated() => $pb.PbList<TSReq>();
  @$core.pragma('dart2js:noInline')
  static TSReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TSReq>(create);
  static TSReq _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get platform => $_getSZ(0);
  @$pb.TagNumber(1)
  set platform($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasPlatform() => $_has(0);
  @$pb.TagNumber(1)
  void clearPlatform() => clearField(1);
}

class TSRes extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('TSRes', package: const $pb.PackageName('timestamp'), createEmptyInstance: create)
    ..aInt64(1, 'timestamp')
    ..hasRequiredFields = false
  ;

  TSRes._() : super();
  factory TSRes() => create();
  factory TSRes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TSRes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  TSRes clone() => TSRes()..mergeFromMessage(this);
  TSRes copyWith(void Function(TSRes) updates) => super.copyWith((message) => updates(message as TSRes));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TSRes create() => TSRes._();
  TSRes createEmptyInstance() => create();
  static $pb.PbList<TSRes> createRepeated() => $pb.PbList<TSRes>();
  @$core.pragma('dart2js:noInline')
  static TSRes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TSRes>(create);
  static TSRes _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get timestamp => $_getI64(0);
  @$pb.TagNumber(1)
  set timestamp($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTimestamp() => $_has(0);
  @$pb.TagNumber(1)
  void clearTimestamp() => clearField(1);
}

