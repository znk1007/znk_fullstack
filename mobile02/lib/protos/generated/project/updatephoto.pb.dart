///
//  Generated code. Do not modify.
//  source: updatephoto.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

class UpdatePhotoRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdatePhotoRequest', package: const $pb.PackageName('protos.updatephoto'))
    ..aOS(1, 'userId')
    ..aOS(2, 'sessionId')
    ..aOS(3, 'photo')
    ..aOS(4, 'device')
    ..hasRequiredFields = false
  ;

  UpdatePhotoRequest() : super();
  UpdatePhotoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  UpdatePhotoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  UpdatePhotoRequest clone() => UpdatePhotoRequest()..mergeFromMessage(this);
  UpdatePhotoRequest copyWith(void Function(UpdatePhotoRequest) updates) => super.copyWith((message) => updates(message as UpdatePhotoRequest));
  $pb.BuilderInfo get info_ => _i;
  static UpdatePhotoRequest create() => UpdatePhotoRequest();
  UpdatePhotoRequest createEmptyInstance() => create();
  static $pb.PbList<UpdatePhotoRequest> createRepeated() => $pb.PbList<UpdatePhotoRequest>();
  static UpdatePhotoRequest getDefault() => _defaultInstance ??= create()..freeze();
  static UpdatePhotoRequest _defaultInstance;

  $core.String get userId => $_getS(0, '');
  set userId($core.String v) { $_setString(0, v); }
  $core.bool hasUserId() => $_has(0);
  void clearUserId() => clearField(1);

  $core.String get sessionId => $_getS(1, '');
  set sessionId($core.String v) { $_setString(1, v); }
  $core.bool hasSessionId() => $_has(1);
  void clearSessionId() => clearField(2);

  $core.String get photo => $_getS(2, '');
  set photo($core.String v) { $_setString(2, v); }
  $core.bool hasPhoto() => $_has(2);
  void clearPhoto() => clearField(3);

  $core.String get device => $_getS(3, '');
  set device($core.String v) { $_setString(3, v); }
  $core.bool hasDevice() => $_has(3);
  void clearDevice() => clearField(4);
}

class UpdatePhotoResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UpdatePhotoResponse', package: const $pb.PackageName('protos.updatephoto'))
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(3, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  UpdatePhotoResponse() : super();
  UpdatePhotoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  UpdatePhotoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  UpdatePhotoResponse clone() => UpdatePhotoResponse()..mergeFromMessage(this);
  UpdatePhotoResponse copyWith(void Function(UpdatePhotoResponse) updates) => super.copyWith((message) => updates(message as UpdatePhotoResponse));
  $pb.BuilderInfo get info_ => _i;
  static UpdatePhotoResponse create() => UpdatePhotoResponse();
  UpdatePhotoResponse createEmptyInstance() => create();
  static $pb.PbList<UpdatePhotoResponse> createRepeated() => $pb.PbList<UpdatePhotoResponse>();
  static UpdatePhotoResponse getDefault() => _defaultInstance ??= create()..freeze();
  static UpdatePhotoResponse _defaultInstance;

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

