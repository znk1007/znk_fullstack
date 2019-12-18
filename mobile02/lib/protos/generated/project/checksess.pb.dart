///
//  Generated code. Do not modify.
//  source: checksess.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

class CheckSessionRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CheckSessionRequest', package: const $pb.PackageName('proto.checksess'))
    ..aOS(1, 'userId')
    ..aOS(2, 'sessionId')
    ..aOS(3, 'device')
    ..hasRequiredFields = false
  ;

  CheckSessionRequest() : super();
  CheckSessionRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  CheckSessionRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  CheckSessionRequest clone() => CheckSessionRequest()..mergeFromMessage(this);
  CheckSessionRequest copyWith(void Function(CheckSessionRequest) updates) => super.copyWith((message) => updates(message as CheckSessionRequest));
  $pb.BuilderInfo get info_ => _i;
  static CheckSessionRequest create() => CheckSessionRequest();
  CheckSessionRequest createEmptyInstance() => create();
  static $pb.PbList<CheckSessionRequest> createRepeated() => $pb.PbList<CheckSessionRequest>();
  static CheckSessionRequest getDefault() => _defaultInstance ??= create()..freeze();
  static CheckSessionRequest _defaultInstance;

  $core.String get userId => $_getS(0, '');
  set userId($core.String v) { $_setString(0, v); }
  $core.bool hasUserId() => $_has(0);
  void clearUserId() => clearField(1);

  $core.String get sessionId => $_getS(1, '');
  set sessionId($core.String v) { $_setString(1, v); }
  $core.bool hasSessionId() => $_has(1);
  void clearSessionId() => clearField(2);

  $core.String get device => $_getS(2, '');
  set device($core.String v) { $_setString(2, v); }
  $core.bool hasDevice() => $_has(2);
  void clearDevice() => clearField(3);
}

class CheckSessionResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CheckSessionResponse', package: const $pb.PackageName('proto.checksess'))
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..aOB(3, 'isValid')
    ..a<$core.int>(4, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  CheckSessionResponse() : super();
  CheckSessionResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  CheckSessionResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  CheckSessionResponse clone() => CheckSessionResponse()..mergeFromMessage(this);
  CheckSessionResponse copyWith(void Function(CheckSessionResponse) updates) => super.copyWith((message) => updates(message as CheckSessionResponse));
  $pb.BuilderInfo get info_ => _i;
  static CheckSessionResponse create() => CheckSessionResponse();
  CheckSessionResponse createEmptyInstance() => create();
  static $pb.PbList<CheckSessionResponse> createRepeated() => $pb.PbList<CheckSessionResponse>();
  static CheckSessionResponse getDefault() => _defaultInstance ??= create()..freeze();
  static CheckSessionResponse _defaultInstance;

  $core.String get message => $_getS(0, '');
  set message($core.String v) { $_setString(0, v); }
  $core.bool hasMessage() => $_has(0);
  void clearMessage() => clearField(1);

  $core.int get code => $_get(1, 0);
  set code($core.int v) { $_setSignedInt32(1, v); }
  $core.bool hasCode() => $_has(1);
  void clearCode() => clearField(2);

  $core.bool get isValid => $_get(2, false);
  set isValid($core.bool v) { $_setBool(2, v); }
  $core.bool hasIsValid() => $_has(2);
  void clearIsValid() => clearField(3);

  $core.int get status => $_get(3, 0);
  set status($core.int v) { $_setSignedInt32(3, v); }
  $core.bool hasStatus() => $_has(3);
  void clearStatus() => clearField(4);
}

