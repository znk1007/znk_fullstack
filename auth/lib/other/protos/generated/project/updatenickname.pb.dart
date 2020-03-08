///
//  Generated code. Do not modify.
//  source: updatenickname.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class UpdateNicknameRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdateNicknameRequest', package: const $pb.PackageName('protos.updatnickname'), createEmptyInstance: create)
    ..aOS(1, 'userId', protoName: 'userId')
    ..aOS(2, 'sessionId', protoName: 'sessionId')
    ..aOS(3, 'nickname')
    ..aOS(4, 'device')
    ..hasRequiredFields = false
  ;

  UpdateNicknameRequest._() : super();
  factory UpdateNicknameRequest() => create();
  factory UpdateNicknameRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateNicknameRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UpdateNicknameRequest clone() => UpdateNicknameRequest()..mergeFromMessage(this);
  UpdateNicknameRequest copyWith(void Function(UpdateNicknameRequest) updates) => super.copyWith((message) => updates(message as UpdateNicknameRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateNicknameRequest create() => UpdateNicknameRequest._();
  UpdateNicknameRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateNicknameRequest> createRepeated() => $pb.PbList<UpdateNicknameRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdateNicknameRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateNicknameRequest>(create);
  static UpdateNicknameRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userId => $_getSZ(0);
  @$pb.TagNumber(1)
  set userId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get sessionId => $_getSZ(1);
  @$pb.TagNumber(2)
  set sessionId($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSessionId() => $_has(1);
  @$pb.TagNumber(2)
  void clearSessionId() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get nickname => $_getSZ(2);
  @$pb.TagNumber(3)
  set nickname($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasNickname() => $_has(2);
  @$pb.TagNumber(3)
  void clearNickname() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get device => $_getSZ(3);
  @$pb.TagNumber(4)
  set device($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasDevice() => $_has(3);
  @$pb.TagNumber(4)
  void clearDevice() => clearField(4);
}

class UpdateNicknameResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdateNicknameResponse', package: const $pb.PackageName('protos.updatnickname'), createEmptyInstance: create)
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(3, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  UpdateNicknameResponse._() : super();
  factory UpdateNicknameResponse() => create();
  factory UpdateNicknameResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateNicknameResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UpdateNicknameResponse clone() => UpdateNicknameResponse()..mergeFromMessage(this);
  UpdateNicknameResponse copyWith(void Function(UpdateNicknameResponse) updates) => super.copyWith((message) => updates(message as UpdateNicknameResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateNicknameResponse create() => UpdateNicknameResponse._();
  UpdateNicknameResponse createEmptyInstance() => create();
  static $pb.PbList<UpdateNicknameResponse> createRepeated() => $pb.PbList<UpdateNicknameResponse>();
  @$core.pragma('dart2js:noInline')
  static UpdateNicknameResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateNicknameResponse>(create);
  static UpdateNicknameResponse _defaultInstance;

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
  $core.int get status => $_getIZ(2);
  @$pb.TagNumber(3)
  set status($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasStatus() => $_has(2);
  @$pb.TagNumber(3)
  void clearStatus() => clearField(3);
}

