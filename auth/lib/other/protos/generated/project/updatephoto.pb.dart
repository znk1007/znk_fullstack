///
//  Generated code. Do not modify.
//  source: updatephoto.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class UpdatePhotoRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdatePhotoRequest', package: const $pb.PackageName('protos.updatephoto'), createEmptyInstance: create)
    ..aOS(1, 'userId', protoName: 'userId')
    ..aOS(2, 'sessionId', protoName: 'sessionId')
    ..aOS(3, 'photo')
    ..aOS(4, 'device')
    ..hasRequiredFields = false
  ;

  UpdatePhotoRequest._() : super();
  factory UpdatePhotoRequest() => create();
  factory UpdatePhotoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdatePhotoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UpdatePhotoRequest clone() => UpdatePhotoRequest()..mergeFromMessage(this);
  UpdatePhotoRequest copyWith(void Function(UpdatePhotoRequest) updates) => super.copyWith((message) => updates(message as UpdatePhotoRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdatePhotoRequest create() => UpdatePhotoRequest._();
  UpdatePhotoRequest createEmptyInstance() => create();
  static $pb.PbList<UpdatePhotoRequest> createRepeated() => $pb.PbList<UpdatePhotoRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdatePhotoRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdatePhotoRequest>(create);
  static UpdatePhotoRequest _defaultInstance;

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
  $core.String get photo => $_getSZ(2);
  @$pb.TagNumber(3)
  set photo($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasPhoto() => $_has(2);
  @$pb.TagNumber(3)
  void clearPhoto() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get device => $_getSZ(3);
  @$pb.TagNumber(4)
  set device($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasDevice() => $_has(3);
  @$pb.TagNumber(4)
  void clearDevice() => clearField(4);
}

class UpdatePhotoResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdatePhotoResponse', package: const $pb.PackageName('protos.updatephoto'), createEmptyInstance: create)
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(3, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  UpdatePhotoResponse._() : super();
  factory UpdatePhotoResponse() => create();
  factory UpdatePhotoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdatePhotoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UpdatePhotoResponse clone() => UpdatePhotoResponse()..mergeFromMessage(this);
  UpdatePhotoResponse copyWith(void Function(UpdatePhotoResponse) updates) => super.copyWith((message) => updates(message as UpdatePhotoResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdatePhotoResponse create() => UpdatePhotoResponse._();
  UpdatePhotoResponse createEmptyInstance() => create();
  static $pb.PbList<UpdatePhotoResponse> createRepeated() => $pb.PbList<UpdatePhotoResponse>();
  @$core.pragma('dart2js:noInline')
  static UpdatePhotoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdatePhotoResponse>(create);
  static UpdatePhotoResponse _defaultInstance;

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

