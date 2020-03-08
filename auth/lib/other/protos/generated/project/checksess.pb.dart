///
//  Generated code. Do not modify.
//  source: checksess.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class CheckSessionRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CheckSessionRequest', package: const $pb.PackageName('proto.checksess'), createEmptyInstance: create)
    ..aOS(1, 'userId', protoName: 'userId')
    ..aOS(2, 'sessionId', protoName: 'sessionId')
    ..aOS(3, 'device')
    ..hasRequiredFields = false
  ;

  CheckSessionRequest._() : super();
  factory CheckSessionRequest() => create();
  factory CheckSessionRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CheckSessionRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  CheckSessionRequest clone() => CheckSessionRequest()..mergeFromMessage(this);
  CheckSessionRequest copyWith(void Function(CheckSessionRequest) updates) => super.copyWith((message) => updates(message as CheckSessionRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CheckSessionRequest create() => CheckSessionRequest._();
  CheckSessionRequest createEmptyInstance() => create();
  static $pb.PbList<CheckSessionRequest> createRepeated() => $pb.PbList<CheckSessionRequest>();
  @$core.pragma('dart2js:noInline')
  static CheckSessionRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CheckSessionRequest>(create);
  static CheckSessionRequest _defaultInstance;

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
  $core.String get device => $_getSZ(2);
  @$pb.TagNumber(3)
  set device($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDevice() => $_has(2);
  @$pb.TagNumber(3)
  void clearDevice() => clearField(3);
}

class CheckSessionResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CheckSessionResponse', package: const $pb.PackageName('proto.checksess'), createEmptyInstance: create)
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..aOB(3, 'isValid', protoName: 'isValid')
    ..a<$core.int>(4, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  CheckSessionResponse._() : super();
  factory CheckSessionResponse() => create();
  factory CheckSessionResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CheckSessionResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  CheckSessionResponse clone() => CheckSessionResponse()..mergeFromMessage(this);
  CheckSessionResponse copyWith(void Function(CheckSessionResponse) updates) => super.copyWith((message) => updates(message as CheckSessionResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CheckSessionResponse create() => CheckSessionResponse._();
  CheckSessionResponse createEmptyInstance() => create();
  static $pb.PbList<CheckSessionResponse> createRepeated() => $pb.PbList<CheckSessionResponse>();
  @$core.pragma('dart2js:noInline')
  static CheckSessionResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CheckSessionResponse>(create);
  static CheckSessionResponse _defaultInstance;

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
  $core.bool get isValid => $_getBF(2);
  @$pb.TagNumber(3)
  set isValid($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasIsValid() => $_has(2);
  @$pb.TagNumber(3)
  void clearIsValid() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get status => $_getIZ(3);
  @$pb.TagNumber(4)
  set status($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasStatus() => $_has(3);
  @$pb.TagNumber(4)
  void clearStatus() => clearField(4);
}

