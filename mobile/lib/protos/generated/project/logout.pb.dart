///
//  Generated code. Do not modify.
//  source: logout.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

class LogoutRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LogoutRequest', package: const $pb.PackageName('protos.logout'))
    ..aOS(1, 'userId')
    ..aOS(2, 'sessionId')
    ..aOS(3, 'device')
    ..hasRequiredFields = false
  ;

  LogoutRequest() : super();
  LogoutRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  LogoutRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  LogoutRequest clone() => LogoutRequest()..mergeFromMessage(this);
  LogoutRequest copyWith(void Function(LogoutRequest) updates) => super.copyWith((message) => updates(message as LogoutRequest));
  $pb.BuilderInfo get info_ => _i;
  static LogoutRequest create() => LogoutRequest();
  LogoutRequest createEmptyInstance() => create();
  static $pb.PbList<LogoutRequest> createRepeated() => $pb.PbList<LogoutRequest>();
  static LogoutRequest getDefault() => _defaultInstance ??= create()..freeze();
  static LogoutRequest _defaultInstance;

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

class LogoutResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LogoutResponse', package: const $pb.PackageName('protos.logout'))
    ..aOS(1, 'message')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(3, 'status', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  LogoutResponse() : super();
  LogoutResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  LogoutResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  LogoutResponse clone() => LogoutResponse()..mergeFromMessage(this);
  LogoutResponse copyWith(void Function(LogoutResponse) updates) => super.copyWith((message) => updates(message as LogoutResponse));
  $pb.BuilderInfo get info_ => _i;
  static LogoutResponse create() => LogoutResponse();
  LogoutResponse createEmptyInstance() => create();
  static $pb.PbList<LogoutResponse> createRepeated() => $pb.PbList<LogoutResponse>();
  static LogoutResponse getDefault() => _defaultInstance ??= create()..freeze();
  static LogoutResponse _defaultInstance;

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

