///
//  Generated code. Do not modify.
//  source: updatenickname.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

class UpdateNicknameRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdateNicknameRequest', package: const $pb.PackageName('protos.updatnickname'))
    ..aOS(1, 'userId')
    ..aOS(2, 'sessionId')
    ..aOS(3, 'nickname')
    ..aOS(4, 'device')
    ..hasRequiredFields = false
  ;

  UpdateNicknameRequest() : super();
  UpdateNicknameRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  UpdateNicknameRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  UpdateNicknameRequest clone() => UpdateNicknameRequest()..mergeFromMessage(this);
  UpdateNicknameRequest copyWith(void Function(UpdateNicknameRequest) updates) => super.copyWith((message) => updates(message as UpdateNicknameRequest));
  $pb.BuilderInfo get info_ => _i;
  static UpdateNicknameRequest create() => UpdateNicknameRequest();
  UpdateNicknameRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateNicknameRequest> createRepeated() => $pb.PbList<UpdateNicknameRequest>();
  static UpdateNicknameRequest getDefault() => _defaultInstance ??= create()..freeze();
  static UpdateNicknameRequest _defaultInstance;

  $core.String get userId => $_getS(0, '');
  set userId($core.String v) { $_setString(0, v); }
  $core.bool hasUserId() => $_has(0);
  void clearUserId() => clearField(1);

  $core.String get sessionId => $_getS(1, '');
  set sessionId($core.String v) { $_setString(1, v); }
  $core.bool hasSessionId() => $_has(1);
  void clearSessionId() => clearField(2);

  $core.String get nickname => $_getS(2, '');
  set nickname($core.String v) { $_setString(2, v); }
  $core.bool hasNickname() => $_has(2);
  void clearNickname() => clearField(3);

  $core.String get device => $_getS(3, '');
  set device($core.String v) { $_setString(3, v); }
  $core.bool hasDevice() => $_has(3);
  void clearDevice() => clearField(4);
}

class UpdateNicknameResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdateNicknameResponse', package: const $pb.PackageName('protos.updatnickname'))
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(3, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  UpdateNicknameResponse() : super();
  UpdateNicknameResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  UpdateNicknameResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  UpdateNicknameResponse clone() => UpdateNicknameResponse()..mergeFromMessage(this);
  UpdateNicknameResponse copyWith(void Function(UpdateNicknameResponse) updates) => super.copyWith((message) => updates(message as UpdateNicknameResponse));
  $pb.BuilderInfo get info_ => _i;
  static UpdateNicknameResponse create() => UpdateNicknameResponse();
  UpdateNicknameResponse createEmptyInstance() => create();
  static $pb.PbList<UpdateNicknameResponse> createRepeated() => $pb.PbList<UpdateNicknameResponse>();
  static UpdateNicknameResponse getDefault() => _defaultInstance ??= create()..freeze();
  static UpdateNicknameResponse _defaultInstance;

  $core.String get message => $_getS(0, '');
  set message($core.String v) { $_setString(0, v); }
  $core.bool hasMessage() => $_has(0);
  void clearMessage() => clearField(1);

  $core.int get code => $_get(1, 0);
  set code($core.int v) { $_setSignedInt32(1, v); }
  $core.bool hasCode() => $_has(1);
  void clearCode() => clearField(2);

  $core.int get status => $_get(2, 0);
  set status($core.int v) { $_setSignedInt32(2, v); }
  $core.bool hasStatus() => $_has(2);
  void clearStatus() => clearField(3);
}

