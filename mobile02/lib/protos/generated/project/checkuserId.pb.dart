///
//  Generated code. Do not modify.
//  source: checkuserId.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class CheckUserIdRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CheckUserIdRequest', package: const $pb.PackageName('protos.checkuserId'), createEmptyInstance: create)
    ..aOS(1, 'account')
    ..aOS(2, 'device')
    ..hasRequiredFields = false
  ;

  CheckUserIdRequest._() : super();
  factory CheckUserIdRequest() => create();
  factory CheckUserIdRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CheckUserIdRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  CheckUserIdRequest clone() => CheckUserIdRequest()..mergeFromMessage(this);
  CheckUserIdRequest copyWith(void Function(CheckUserIdRequest) updates) => super.copyWith((message) => updates(message as CheckUserIdRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CheckUserIdRequest create() => CheckUserIdRequest._();
  CheckUserIdRequest createEmptyInstance() => create();
  static $pb.PbList<CheckUserIdRequest> createRepeated() => $pb.PbList<CheckUserIdRequest>();
  @$core.pragma('dart2js:noInline')
  static CheckUserIdRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CheckUserIdRequest>(create);
  static CheckUserIdRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get account => $_getSZ(0);
  @$pb.TagNumber(1)
  set account($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get device => $_getSZ(1);
  @$pb.TagNumber(2)
  set device($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDevice() => $_has(1);
  @$pb.TagNumber(2)
  void clearDevice() => clearField(2);
}

class CheckUserIdResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CheckUserIdResponse', package: const $pb.PackageName('protos.checkuserId'), createEmptyInstance: create)
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..aOS(3, 'userId', protoName: 'userId')
    ..a<$core.int>(4, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  CheckUserIdResponse._() : super();
  factory CheckUserIdResponse() => create();
  factory CheckUserIdResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CheckUserIdResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  CheckUserIdResponse clone() => CheckUserIdResponse()..mergeFromMessage(this);
  CheckUserIdResponse copyWith(void Function(CheckUserIdResponse) updates) => super.copyWith((message) => updates(message as CheckUserIdResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CheckUserIdResponse create() => CheckUserIdResponse._();
  CheckUserIdResponse createEmptyInstance() => create();
  static $pb.PbList<CheckUserIdResponse> createRepeated() => $pb.PbList<CheckUserIdResponse>();
  @$core.pragma('dart2js:noInline')
  static CheckUserIdResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CheckUserIdResponse>(create);
  static CheckUserIdResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get message => $_getSZ(0);
  @$pb.TagNumber(1)
  set message($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasMessage() => $_has(0);
  @$pb.TagNumber(1)
  void clearMessage() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get code => $_getIZ(1);
  @$pb.TagNumber(2)
  set code($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get userId => $_getSZ(2);
  @$pb.TagNumber(3)
  set userId($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasUserId() => $_has(2);
  @$pb.TagNumber(3)
  void clearUserId() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get status => $_getIZ(3);
  @$pb.TagNumber(4)
  set status($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasStatus() => $_has(3);
  @$pb.TagNumber(4)
  void clearStatus() => clearField(4);
}

